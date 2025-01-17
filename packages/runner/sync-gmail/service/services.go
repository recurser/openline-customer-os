package service

import (
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"github.com/openline-ai/openline-customer-os/packages/runner/sync-gmail/config"
	"github.com/openline-ai/openline-customer-os/packages/runner/sync-gmail/repository"
	"gorm.io/gorm"
)

type Services struct {
	cfg          *config.Config
	Repositories *repository.Repositories

	TenantService TenantService
	UserService   UserService
	OpenAiService OpenAiService

	SyncService    SyncService
	EmailService   EmailService
	MeetingService MeetingService
}

func InitServices(driver *neo4j.DriverWithContext, gormDb *gorm.DB, cfg *config.Config) *Services {
	repositories := repository.InitRepos(driver, gormDb)

	services := new(Services)
	services.cfg = cfg
	services.Repositories = repositories

	services.TenantService = NewTenantService(repositories)
	services.UserService = NewUserService(repositories)
	services.OpenAiService = NewOpenAiService(cfg, repositories)

	services.SyncService = NewSyncService(cfg, repositories, services)
	services.EmailService = NewEmailService(cfg, repositories, services)
	services.MeetingService = NewMeetingService(cfg, repositories, services)

	return services
}
