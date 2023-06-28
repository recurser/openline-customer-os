package service

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j/db"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j/dbtype"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/common"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/entity"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/repository"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/tracing"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-common-module/logger"
	common_module "github.com/openline-ai/openline-customer-os/packages/server/customer-os-common-module/service"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-common-module/utils"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/log"
	"net/http"
	"strings"
)

type EmailService interface {
	GetAllFor(ctx context.Context, entityType entity.EntityType, entityId string) (*entity.EmailEntities, error)
	GetAllForEntityTypeByIds(ctx context.Context, entityType entity.EntityType, entityIds []string) (*entity.EmailEntities, error)
	MergeEmailTo(ctx context.Context, entityType entity.EntityType, entityId string, entity *entity.EmailEntity) (*entity.EmailEntity, error)
	UpdateEmailFor(ctx context.Context, entityType entity.EntityType, entityId string, entity *entity.EmailEntity) (*entity.EmailEntity, error)
	DetachFromEntity(ctx context.Context, entityType entity.EntityType, entityId, email string) (bool, error)
	DetachFromEntityById(ctx context.Context, entityType entity.EntityType, entityId, emailId string) (bool, error)
	DeleteById(ctx context.Context, emailId string) (bool, error)
	GetById(ctx context.Context, emailId string) (*entity.EmailEntity, error)
	CheckValidation(ctx context.Context, email string) (*EmailValidationResponse, error)

	mapDbNodeToEmailEntity(node dbtype.Node) *entity.EmailEntity
}

type EmailValidateRequest struct {
	Email string `json:"email" validate:"required,email"`
}

type EmailValidationResponse struct {
	Error           string `json:"error"`
	IsReachable     string `json:"isReachable"`
	Email           string `json:"email"`
	AcceptsMail     bool   `json:"acceptsMail"`
	CanConnectSmtp  bool   `json:"canConnectSmtp"`
	HasFullInbox    bool   `json:"hasFullInbox"`
	IsCatchAll      bool   `json:"isCatchAll"`
	IsDeliverable   bool   `json:"isDeliverable"`
	IsDisabled      bool   `json:"isDisabled"`
	Address         string `json:"address"`
	Domain          string `json:"domain"`
	IsValidSyntax   bool   `json:"isValidSyntax"`
	Username        string `json:"username"`
	NormalizedEmail string `json:"normalizedEmail"`
}

type emailService struct {
	log          logger.Logger
	repositories *repository.Repositories
	services     *Services
}

func NewEmailService(log logger.Logger, repositories *repository.Repositories, services *Services) EmailService {
	return &emailService{
		log:          log,
		repositories: repositories,
		services:     services,
	}
}

func (s *emailService) getDriver() neo4j.DriverWithContext {
	return *s.repositories.Drivers.Neo4jDriver
}

func (s *emailService) GetAllFor(ctx context.Context, entityType entity.EntityType, entityId string) (*entity.EmailEntities, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "EmailService.GetAllFor")
	defer span.Finish()
	tracing.SetDefaultServiceSpanTags(ctx, span)
	span.LogFields(log.String("entityType", entityType.String()), log.String("entityId", entityId))

	queryResult, err := s.repositories.EmailRepository.GetAllFor(ctx, common.GetContext(ctx).Tenant, entityType, entityId)
	if err != nil {
		return nil, err
	}

	emailEntities := make(entity.EmailEntities, 0, len(queryResult.([]*db.Record)))
	for _, dbRecord := range queryResult.([]*db.Record) {
		emailEntity := s.mapDbNodeToEmailEntity(dbRecord.Values[0].(dbtype.Node))
		s.addDbRelationshipToEmailEntity(dbRecord.Values[1].(dbtype.Relationship), emailEntity)
		emailEntities = append(emailEntities, *emailEntity)
	}

	return &emailEntities, nil
}

func (s *emailService) GetAllForEntityTypeByIds(ctx context.Context, entityType entity.EntityType, entityIds []string) (*entity.EmailEntities, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "EmailService.GetAllForEntityTypeByIds")
	defer span.Finish()
	tracing.SetDefaultServiceSpanTags(ctx, span)
	span.LogFields(log.String("entityType", entityType.String()), log.Object("entityIds", entityIds))

	emails, err := s.repositories.EmailRepository.GetAllForIds(ctx, common.GetContext(ctx).Tenant, entityType, entityIds)
	if err != nil {
		return nil, err
	}

	emailEntities := make(entity.EmailEntities, 0, len(emails))
	for _, v := range emails {
		emailEntity := s.mapDbNodeToEmailEntity(*v.Node)
		s.addDbRelationshipToEmailEntity(*v.Relationship, emailEntity)
		emailEntity.DataloaderKey = v.LinkedNodeId
		emailEntities = append(emailEntities, *emailEntity)
	}
	return &emailEntities, nil
}

func (s *emailService) MergeEmailTo(ctx context.Context, entityType entity.EntityType, entityId string, inputEntity *entity.EmailEntity) (*entity.EmailEntity, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "EmailService.MergeEmailTo")
	defer span.Finish()
	tracing.SetDefaultServiceSpanTags(ctx, span)
	span.LogFields(log.String("entityType", entityType.String()), log.String("entityId", entityId))

	session := utils.NewNeo4jWriteSession(ctx, s.getDriver())
	defer session.Close(ctx)

	var err error
	var emailNode *dbtype.Node
	var emailRelationship *dbtype.Relationship

	_, err = session.ExecuteWrite(ctx, func(tx neo4j.ManagedTransaction) (interface{}, error) {
		emailNode, emailRelationship, err = s.repositories.EmailRepository.MergeEmailToInTx(ctx, tx, common.GetContext(ctx).Tenant, entityType, entityId, *inputEntity)
		if err != nil {
			return nil, err
		}
		emailId := utils.GetPropsFromNode(*emailNode)["id"].(string)
		if inputEntity.Primary == true {
			err = s.repositories.EmailRepository.SetOtherEmailsNonPrimaryInTx(ctx, tx, common.GetContext(ctx).Tenant, entityType, entityId, emailId)
			if err != nil {
				return nil, err
			}
		}
		return nil, err
	})
	if err != nil {
		return nil, err
	}

	if entityType == entity.ORGANIZATION {
		s.services.OrganizationService.UpdateLastTouchpointSync(ctx, entityId)
	} else if entityType == entity.CONTACT {
		s.services.OrganizationService.UpdateLastTouchpointSyncByContactId(ctx, entityId)
	}

	var emailEntity = s.mapDbNodeToEmailEntity(*emailNode)
	s.addDbRelationshipToEmailEntity(*emailRelationship, emailEntity)
	return emailEntity, nil
}

func (s *emailService) UpdateEmailFor(ctx context.Context, entityType entity.EntityType, entityId string, inputEntity *entity.EmailEntity) (*entity.EmailEntity, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "EmailService.UpdateEmailFor")
	defer span.Finish()
	tracing.SetDefaultServiceSpanTags(ctx, span)
	span.LogFields(log.String("entityType", string(entityType)), log.String("entityId", entityId))

	session := utils.NewNeo4jWriteSession(ctx, s.getDriver())
	defer session.Close(ctx)

	var err error
	var emailNode *dbtype.Node
	var emailRelationship *dbtype.Relationship
	var detachCurrentEmail = false

	_, err = session.ExecuteWrite(ctx, func(tx neo4j.ManagedTransaction) (interface{}, error) {
		currentEmailNode, err := s.repositories.EmailRepository.GetByIdAndRelatedEntity(ctx, entityType, common.GetTenantFromContext(ctx), inputEntity.Id, entityId)
		if err != nil {
			return nil, err
		}

		currentEmail := utils.GetStringPropOrEmpty(utils.GetPropsFromNode(*currentEmailNode), "email")
		currentRawEmail := utils.GetStringPropOrEmpty(utils.GetPropsFromNode(*currentEmailNode), "rawEmail")

		var emailExists = false
		if currentRawEmail == "" {
			emailExists, err = s.repositories.EmailRepository.Exists(ctx, common.GetContext(ctx).Tenant, inputEntity.RawEmail)
			if err != nil {
				return nil, err
			}
		}

		if len(inputEntity.RawEmail) == 0 || inputEntity.RawEmail == currentEmail || inputEntity.RawEmail == currentRawEmail ||
			(currentRawEmail == "" && !emailExists) {
			// proceed with update
			emailNode, emailRelationship, err = s.repositories.EmailRepository.UpdateEmailForInTx(ctx, tx, common.GetContext(ctx).Tenant, entityType, entityId, *inputEntity)
			if err != nil {
				return nil, err
			}
			emailId := utils.GetPropsFromNode(*emailNode)["id"].(string)
			if inputEntity.Primary == true {
				err := s.repositories.EmailRepository.SetOtherEmailsNonPrimaryInTx(ctx, tx, common.GetContext(ctx).Tenant, entityType, entityId, emailId)
				if err != nil {
					return nil, err
				}
			}
		} else {
			// proceed with email address replace
			// merge new email address
			emailNode, emailRelationship, err = s.repositories.EmailRepository.MergeEmailToInTx(ctx, tx, common.GetContext(ctx).Tenant, entityType, entityId, *inputEntity)
			if err != nil {
				return nil, err
			}
			emailId := utils.GetPropsFromNode(*emailNode)["id"].(string)
			if inputEntity.Primary == true {
				err := s.repositories.EmailRepository.SetOtherEmailsNonPrimaryInTx(ctx, tx, common.GetContext(ctx).Tenant, entityType, entityId, emailId)
				if err != nil {
					return nil, err
				}
			}
			detachCurrentEmail = true
		}
		return nil, nil
	})
	if err != nil {
		return nil, err
	}

	if detachCurrentEmail {
		_, err = s.DetachFromEntityById(ctx, entityType, entityId, inputEntity.Id)
	}

	if entityType == entity.ORGANIZATION {
		s.services.OrganizationService.UpdateLastTouchpointSync(ctx, entityId)
	} else if entityType == entity.CONTACT {
		s.services.OrganizationService.UpdateLastTouchpointSyncByContactId(ctx, entityId)
	}

	var emailEntity = s.mapDbNodeToEmailEntity(*emailNode)
	s.addDbRelationshipToEmailEntity(*emailRelationship, emailEntity)
	return emailEntity, nil
}

func (s *emailService) DetachFromEntity(ctx context.Context, entityType entity.EntityType, entityId, email string) (bool, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "EmailService.DetachFromEntity")
	defer span.Finish()
	tracing.SetDefaultServiceSpanTags(ctx, span)
	span.LogFields(log.String("email", email), log.String("entityId", entityId), log.String("entityType", string(entityType)))

	err := s.repositories.EmailRepository.RemoveRelationship(ctx, entityType, common.GetTenantFromContext(ctx), entityId, email)

	if entityType == entity.ORGANIZATION {
		s.services.OrganizationService.UpdateLastTouchpointSync(ctx, entityId)
	} else if entityType == entity.CONTACT {
		s.services.OrganizationService.UpdateLastTouchpointSyncByContactId(ctx, entityId)
	}

	return err == nil, err
}

func (s *emailService) DetachFromEntityById(ctx context.Context, entityType entity.EntityType, entityId, emailId string) (bool, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "EmailService.DetachFromEntityById")
	defer span.Finish()
	tracing.SetDefaultServiceSpanTags(ctx, span)
	span.LogFields(log.String("emailId", emailId), log.String("entityId", entityId), log.String("entityType", string(entityType)))

	err := s.repositories.EmailRepository.RemoveRelationshipById(ctx, entityType, common.GetTenantFromContext(ctx), entityId, emailId)

	if entityType == entity.ORGANIZATION {
		s.services.OrganizationService.UpdateLastTouchpointSync(ctx, entityId)
	} else if entityType == entity.CONTACT {
		s.services.OrganizationService.UpdateLastTouchpointSyncByContactId(ctx, entityId)
	}

	return err == nil, err
}

func (s *emailService) DeleteById(ctx context.Context, emailId string) (bool, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "EmailService.DeleteById")
	defer span.Finish()
	tracing.SetDefaultServiceSpanTags(ctx, span)
	span.LogFields(log.String("emailId", emailId))

	err := s.repositories.EmailRepository.DeleteById(ctx, common.GetTenantFromContext(ctx), emailId)
	return err == nil, err
}

func (s *emailService) GetById(ctx context.Context, emailId string) (*entity.EmailEntity, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "EmailService.GetById")
	defer span.Finish()
	tracing.SetDefaultServiceSpanTags(ctx, span)
	span.LogFields(log.String("emailId", emailId))

	emailNode, err := s.repositories.EmailRepository.GetById(ctx, emailId)
	if err != nil {
		return nil, err
	}
	var emailEntity = s.mapDbNodeToEmailEntity(*emailNode)
	return emailEntity, nil
}

func (s *emailService) CheckValidation(ctx context.Context, email string) (*EmailValidationResponse, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "EmailService.CheckValidation")
	defer span.Finish()
	tracing.SetDefaultServiceSpanTags(ctx, span)
	span.LogFields(log.String("email", email))

	emailValidationRequest := EmailValidateRequest{
		Email: strings.TrimSpace(email),
	}
	evJSON, err := json.Marshal(emailValidationRequest)
	if err != nil {
		tracing.TraceErr(span, err)
		return nil, err
	}
	requestBody := []byte(string(evJSON))
	req, err := http.NewRequest("POST", s.services.cfg.Service.ValidationApi+"/validateEmail", bytes.NewBuffer(requestBody))
	if err != nil {
		tracing.TraceErr(span, err)
		return nil, err
	}
	// Set the request headers
	req.Header.Set(common_module.ApiKeyHeader, s.services.cfg.Service.ValidationApiKey)
	req.Header.Set(common_module.TenantHeader, common.GetTenantFromContext(ctx))

	// Make the HTTP request
	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		tracing.TraceErr(span, err)
		return nil, err
	}
	defer response.Body.Close()
	var emailValidationResponse EmailValidationResponse
	err = json.NewDecoder(response.Body).Decode(&emailValidationResponse)
	if err != nil {
		tracing.TraceErr(span, err)
		return nil, err
	}
	return &emailValidationResponse, nil
}

func (s *emailService) mapDbNodeToEmailEntity(node dbtype.Node) *entity.EmailEntity {
	props := utils.GetPropsFromNode(node)
	result := entity.EmailEntity{
		Id:            utils.GetStringPropOrEmpty(props, "id"),
		Email:         utils.GetStringPropOrEmpty(props, "email"),
		RawEmail:      utils.GetStringPropOrEmpty(props, "rawEmail"),
		Source:        entity.GetDataSource(utils.GetStringPropOrEmpty(props, "source")),
		SourceOfTruth: entity.GetDataSource(utils.GetStringPropOrEmpty(props, "sourceOfTruth")),
		AppSource:     utils.GetStringPropOrEmpty(props, "appSource"),
		CreatedAt:     utils.GetTimePropOrEpochStart(props, "createdAt"),
		UpdatedAt:     utils.GetTimePropOrEpochStart(props, "updatedAt"),

		Validated:      utils.GetBoolPropOrNil(props, "validated"),
		IsReachable:    utils.GetStringPropOrNil(props, "isReachable"),
		IsValidSyntax:  utils.GetBoolPropOrNil(props, "isValidSyntax"),
		CanConnectSMTP: utils.GetBoolPropOrNil(props, "canConnectSmtp"),
		AcceptsMail:    utils.GetBoolPropOrNil(props, "acceptsMail"),
		HasFullInbox:   utils.GetBoolPropOrNil(props, "hasFullInbox"),
		IsCatchAll:     utils.GetBoolPropOrNil(props, "isCatchAll"),
		IsDeliverable:  utils.GetBoolPropOrNil(props, "isDeliverable"),
		IsDisabled:     utils.GetBoolPropOrNil(props, "isDisabled"),
		Error:          utils.GetStringPropOrNil(props, "validationError"),
	}
	return &result
}

func (s *emailService) addDbRelationshipToEmailEntity(relationship dbtype.Relationship, emailEntity *entity.EmailEntity) {
	props := utils.GetPropsFromRelationship(relationship)
	emailEntity.Primary = utils.GetBoolPropOrFalse(props, "primary")
	emailEntity.Label = utils.GetStringPropOrEmpty(props, "label")
}
