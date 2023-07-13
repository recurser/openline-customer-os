package repository

import (
	"fmt"
	"github.com/openline-ai/openline-customer-os/packages/runner/sync-customer-os-data/source/hubspot/entity"
	"gorm.io/gorm"
	"time"
)

const (
	CompanyEntity      = "companies"
	OwnerEntity        = "owners"
	NoteEntity         = "engagements_notes"
	MeetingEntity      = "engagements_meetings"
	EmailMessageEntity = "engagements_emails"
)

func GetContacts(db *gorm.DB, limit int, runId string) (entity.Contacts, error) {
	var contacts entity.Contacts

	cte := `
		WITH UpToDateData AS (
    		SELECT row_number() OVER (PARTITION BY id ORDER BY updatedat DESC) AS row_num, *
    		FROM contacts
		)`
	err := db.
		Raw(cte+" SELECT u.* FROM UpToDateData u left join openline_sync_status_contacts s "+
			" on u.id = s.id and u._airbyte_ab_id = s._airbyte_ab_id and u._airbyte_contacts_hashid = s._airbyte_contacts_hashid "+
			" WHERE u.row_num = ? "+
			" and (s.synced_to_customer_os is null or s.synced_to_customer_os = ?) "+
			" and (s.synced_to_customer_os_attempt is null or s.synced_to_customer_os_attempt < ?) "+
			" and (s.run_id is null or s.run_id <> ?) "+
			" limit ?", 1, false, 10, runId, limit).
		Find(&contacts).Error

	if err != nil {
		return nil, err
	}
	return contacts, nil
}

func GetContactProperties(db *gorm.DB, airbyteAbId, airbyteContactsHashId string) (entity.ContactProperties, error) {
	contactProperties := entity.ContactProperties{}
	err := db.Table(entity.ContactProperties{}.TableName()).
		Where(&entity.ContactProperties{AirbyteAbId: airbyteAbId, AirbyteContactsHashid: airbyteContactsHashId}).
		First(&contactProperties).Error
	return contactProperties, err
}

func MarkContactProcessed(db *gorm.DB, contact entity.Contact, synced bool, runId string) error {
	syncStatusContact := entity.SyncStatusContact{
		Id:                    contact.Id,
		AirbyteAbId:           contact.AirbyteAbId,
		AirbyteContactsHashid: contact.AirbyteContactsHashid,
	}
	db.FirstOrCreate(&syncStatusContact, syncStatusContact)

	return db.Model(&syncStatusContact).
		Where(&entity.SyncStatusContact{Id: contact.Id, AirbyteAbId: contact.AirbyteAbId, AirbyteContactsHashid: contact.AirbyteContactsHashid}).
		Updates(entity.SyncStatusContact{
			SyncedToCustomerOs: synced,
			SyncedAt:           time.Now(),
			SyncAttempt:        syncStatusContact.SyncAttempt + 1,
			RunId:              runId,
		}).Error
}

func GetAirbyteUnprocessedRecords(db *gorm.DB, limit int, runId, tableSuffix string) (entity.AirbyteRaws, error) {
	var airbyteRecords entity.AirbyteRaws

	err := db.
		Raw(fmt.Sprintf(`SELECT a.*
FROM _airbyte_raw_%s a
LEFT JOIN openline_sync_status s ON a._airbyte_ab_id = s._airbyte_ab_id and s.entity = ?
WHERE (s.synced_to_customer_os IS NULL OR s.synced_to_customer_os = FALSE)
  AND (s.synced_to_customer_os_attempt IS NULL OR s.synced_to_customer_os_attempt < ?)
  AND (s.run_id IS NULL OR s.run_id <> ?)
ORDER BY a._airbyte_emitted_at ASC
LIMIT ?`, tableSuffix), tableSuffix, 10, runId, limit).
		Find(&airbyteRecords).Error

	if err != nil {
		return nil, err
	}
	return airbyteRecords, nil
}

func MarkProcessed(db *gorm.DB, syncedEntity, airbyteAbId string, synced bool, runId, externalSyncId string) error {
	syncStatus := entity.SyncStatus{
		Entity:         syncedEntity,
		AirbyteAbId:    airbyteAbId,
		ExternalSyncId: externalSyncId,
	}
	db.FirstOrCreate(&syncStatus, syncStatus)

	return db.Model(&syncStatus).
		Where(&entity.SyncStatus{AirbyteAbId: airbyteAbId, Entity: syncedEntity}).
		Updates(entity.SyncStatus{
			SyncedToCustomerOs: synced,
			SyncedAt:           time.Now(),
			SyncAttempt:        syncStatus.SyncAttempt + 1,
			RunId:              runId,
		}).Error
}
