package service

import (
	"github.com/google/uuid"
	"github.com/openline-ai/openline-customer-os/packages/runner/sync-customer-os-data/common"
	"github.com/openline-ai/openline-customer-os/packages/runner/sync-customer-os-data/entity"
	"github.com/openline-ai/openline-customer-os/packages/runner/sync-customer-os-data/repository"
	"github.com/sirupsen/logrus"
	"golang.org/x/net/context"
	"time"
)

type ContactSyncService interface {
	SyncContacts(ctx context.Context, dataService common.SourceDataService, syncDate time.Time, tenant, runId string) (int, int)
}

type contactSyncService struct {
	repositories *repository.Repositories
}

func NewContactSyncService(repositories *repository.Repositories) ContactSyncService {
	return &contactSyncService{
		repositories: repositories,
	}
}

func (s *contactSyncService) SyncContacts(ctx context.Context, dataService common.SourceDataService, syncDate time.Time, tenant, runId string) (int, int) {
	completed, failed := 0, 0
	for {
		contacts := dataService.GetContactsForSync(batchSize, runId)
		if len(contacts) == 0 {
			logrus.Debugf("no contacts found for sync from %s for tenant %s", dataService.SourceId(), tenant)
			break
		}
		logrus.Infof("syncing %d contacts from %s for tenant %s", len(contacts), dataService.SourceId(), tenant)

		for _, v := range contacts {
			var failedSync = false

			contactId, err := s.repositories.ContactRepository.GetMatchedContactId(ctx, tenant, v)
			if err != nil {
				failedSync = true
				logrus.Errorf("failed finding existing matched contact with external reference %v for tenant %v :%v", v.ExternalId, tenant, err)
			}

			// Create new organization id if not found
			if len(contactId) == 0 {
				orgUuid, _ := uuid.NewRandom()
				contactId = orgUuid.String()
			}
			v.Id = contactId

			if !failedSync {
				err = s.repositories.ContactRepository.MergeContact(ctx, tenant, syncDate, v)
				if err != nil {
					failedSync = true
					logrus.Errorf("failed merge contact with external reference %v for tenant %v :%v", v.ExternalId, tenant, err)
				}
			}

			if len(v.PrimaryEmail) > 0 && !failedSync {
				if err = s.repositories.ContactRepository.MergePrimaryEmail(ctx, tenant, contactId, v.PrimaryEmail, v.ExternalSystem, v.CreatedAt); err != nil {
					failedSync = true
					logrus.Errorf("failed merge primary email for contact with external reference %v , tenant %v :%v", v.ExternalId, tenant, err)
				}
			}

			if !failedSync {
				for _, additionalEmail := range v.AdditionalEmails {
					if len(additionalEmail) > 0 {
						if err = s.repositories.ContactRepository.MergeAdditionalEmail(ctx, tenant, contactId, additionalEmail, v.ExternalSystem, v.CreatedAt); err != nil {
							failedSync = true
							logrus.Errorf("failed merge additional email for contact with external reference %v , tenant %v :%v", v.ExternalId, tenant, err)
						}
					}
				}
			}

			if v.HasPhoneNumber() && !failedSync {
				if err = s.repositories.ContactRepository.MergePrimaryPhoneNumber(ctx, tenant, contactId, v.PhoneNumber, v.ExternalSystem, v.CreatedAt); err != nil {
					failedSync = true
					logrus.Errorf("failed merge primary phone number for contact with external reference %v , tenant %v :%v", v.ExternalId, tenant, err)
				}
			}

			if v.HasOrganizations() && !failedSync {
				for _, organizationExternalId := range v.OrganizationsExternalIds {
					if err = s.repositories.ContactRepository.LinkContactWithOrganization(ctx, tenant, contactId, organizationExternalId, dataService.SourceId()); err != nil {
						failedSync = true
						logrus.Errorf("failed link contact %v to organization with external id %v, tenant %v :%v", contactId, organizationExternalId, tenant, err)
					}
				}
			}

			if !failedSync {
				if err = s.repositories.RoleRepository.RemoveOutdatedJobRoles(ctx, tenant, contactId, dataService.SourceId(), v.PrimaryOrganizationExternalId); err != nil {
					failedSync = true
					logrus.Errorf("failed removing outdated roles for contact %v, tenant %v :%v", contactId, tenant, err)
				}
			}

			if len(v.PrimaryOrganizationExternalId) > 0 && !failedSync {
				if err = s.repositories.RoleRepository.MergeJobRole(ctx, tenant, contactId, v.JobTitle, v.PrimaryOrganizationExternalId, dataService.SourceId()); err != nil {
					failedSync = true
					logrus.Errorf("failed merge primary role for contact %v, tenant %v :%v", contactId, tenant, err)
				}
			}

			if len(v.UserExternalOwnerId) > 0 && !failedSync {
				if err = s.repositories.ContactRepository.SetOwnerRelationship(ctx, tenant, contactId, v.UserExternalOwnerId, dataService.SourceId()); err != nil {
					// Do not mark sync as failed in case owner relationship is not set
					logrus.Errorf("failed set owner user for contact %v, tenant %v :%v", contactId, tenant, err)
				}
			}

			if v.HasNotes() && !failedSync {
				for _, note := range v.Notes {
					noteId, err := s.repositories.NoteRepository.MergeNote(ctx, tenant, syncDate, entity.NoteData{
						Html:           note.Note,
						CreatedAt:      v.CreatedAt,
						ExternalId:     string(note.FieldSource) + "-" + v.ExternalId,
						ExternalSystem: v.ExternalSystem,
					})
					if err != nil {
						failedSync = true
						logrus.Errorf("failed merge note for contact %v, tenant %v :%v", contactId, tenant, err)
					}
					err = s.repositories.NoteRepository.NoteLinkWithContactByExternalId(ctx, tenant, noteId, v.ExternalId, v.ExternalSystem)
					if err != nil {
						failedSync = true
						logrus.Errorf("failed link note with contact %v, tenant %v :%v", contactId, tenant, err)
					}
				}
			}

			if !failedSync {
				for _, organization_id := range v.TextCustomFields {
					if err = s.repositories.ContactRepository.MergeTextCustomField(ctx, tenant, contactId, organization_id); err != nil {
						failedSync = true
						logrus.Errorf("failed merge custom field %v for contact %v, tenant %v :%v", organization_id.Name, contactId, tenant, err)
					}
				}
			}

			if v.HasLocation() && !failedSync {
				err = s.repositories.ContactRepository.MergeContactDefaultPlace(ctx, tenant, contactId, v)
				if err != nil {
					failedSync = true
					logrus.Errorf("failed merge place for contact %v, tenant %v :%v", contactId, tenant, err)
				}
			}

			if len(v.TagName) > 0 {
				err = s.repositories.ContactRepository.MergeTagForContact(ctx, tenant, contactId, v.TagName, v.ExternalSystem)
				if err != nil {
					failedSync = true
					logrus.Errorf("failed to merge tag for contact %v, tenant %v :%v", contactId, tenant, err)
				}
			}

			logrus.Debugf("successfully merged contact with id %v for tenant %v from %v", contactId, tenant, dataService.SourceId())
			if err := dataService.MarkContactProcessed(v.ExternalSyncId, runId, failedSync == false); err != nil {
				failed++
				continue
			}
			if failedSync == true {
				failed++
			} else {
				completed++
			}
		}
		if len(contacts) < batchSize {
			break
		}
	}
	return completed, failed
}