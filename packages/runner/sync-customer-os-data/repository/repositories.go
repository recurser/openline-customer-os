package repository

import (
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"github.com/openline-ai/openline-customer-os/packages/runner/sync-customer-os-data/config"
	"github.com/openline-ai/openline-customer-os/packages/runner/sync-customer-os-data/logger"
	commonRepository "github.com/openline-ai/openline-customer-os/packages/server/customer-os-common-module/repository"
	"gorm.io/gorm"
)

type Dbs struct {
	GormDB         *gorm.DB
	Neo4jDriver    *neo4j.DriverWithContext
	RawDataStoreDB *config.RawDataStoreDB
}

type Repositories struct {
	Dbs Dbs

	CommonRepositories *commonRepository.Repositories

	TenantSyncSettingsRepository TenantSyncSettingsRepository
	TenantSettingsRepository     TenantSettingsRepository
	SyncRunRepository            SyncRunRepository

	ContactRepository          ContactRepository
	EmailRepository            EmailRepository
	PhoneNumberRepository      PhoneNumberRepository
	LocationRepository         LocationRepository
	ExternalSystemRepository   ExternalSystemRepository
	OrganizationRepository     OrganizationRepository
	UserRepository             UserRepository
	LogEntryRepository         LogEntryRepository
	InteractionEventRepository InteractionEventRepository
	IssueRepository            IssueRepository
	MeetingRepository          MeetingRepository
	ActionRepository           ActionRepository
}

func InitRepos(driver *neo4j.DriverWithContext, gormDB *gorm.DB, airbyteStoreDb *config.RawDataStoreDB, log logger.Logger) *Repositories {
	repositories := Repositories{
		Dbs: Dbs{
			Neo4jDriver:    driver,
			GormDB:         gormDB,
			RawDataStoreDB: airbyteStoreDb,
		},
		CommonRepositories:           commonRepository.InitRepositories(gormDB, driver),
		TenantSyncSettingsRepository: NewTenantSyncSettingsRepository(gormDB),
		TenantSettingsRepository:     NewTenantSettingsRepository(gormDB),
		SyncRunRepository:            NewSyncRunRepository(gormDB),
		ContactRepository:            NewContactRepository(driver),
		EmailRepository:              NewEmailRepository(driver),
		PhoneNumberRepository:        NewPhoneNumberRepository(driver),
		LocationRepository:           NewLocationRepository(driver),
		ExternalSystemRepository:     NewExternalSystemRepository(driver),
		OrganizationRepository:       NewOrganizationRepository(driver, log),
		UserRepository:               NewUserRepository(driver),
		InteractionEventRepository:   NewInteractionEventRepository(driver),
		IssueRepository:              NewIssueRepository(driver),
		MeetingRepository:            NewMeetingRepository(driver),
		ActionRepository:             NewActionRepository(driver),
		LogEntryRepository:           NewLogEntryRepository(driver),
	}
	return &repositories
}
