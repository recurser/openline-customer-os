package service

import (
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/config"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/grpc_client"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/repository"
	commonAuthService "github.com/openline-ai/openline-customer-os/packages/server/customer-os-common-auth/service"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-common-module/logger"
	commonService "github.com/openline-ai/openline-customer-os/packages/server/customer-os-common-module/service"
)

type Services struct {
	cfg   *config.Config
	Cache CacheService

	CommonServices     *commonService.Services
	CommonAuthServices *commonAuthService.Services

	ContactService                  ContactService
	OrganizationService             OrganizationService
	CustomFieldService              CustomFieldService
	PhoneNumberService              PhoneNumberService
	EmailService                    EmailService
	UserService                     UserService
	FieldSetService                 FieldSetService
	EntityTemplateService           EntityTemplateService
	FieldSetTemplateService         FieldSetTemplateService
	CustomFieldTemplateService      CustomFieldTemplateService
	TimelineEventService            TimelineEventService
	NoteService                     NoteService
	JobRoleService                  JobRoleService
	CalendarService                 CalendarService
	LocationService                 LocationService
	TagService                      TagService
	SearchService                   SearchService
	QueryService                    DashboardService
	DomainService                   DomainService
	IssueService                    IssueService
	InteractionSessionService       InteractionSessionService
	InteractionEventService         InteractionEventService
	PageViewService                 PageViewService
	AnalysisService                 AnalysisService
	AttachmentService               AttachmentService
	MeetingService                  MeetingService
	TenantService                   TenantService
	WorkspaceService                WorkspaceService
	SocialService                   SocialService
	PlayerService                   PlayerService
	OrganizationRelationshipService OrganizationRelationshipService
	ExternalSystemService           ExternalSystemService
	ActionService                   ActionService
	CountryService                  CountryService
	ActionItemService               ActionItemService
	BillableService                 BillableService
	LogEntryService                 LogEntryService
}

func InitServices(log logger.Logger, driver *neo4j.DriverWithContext, cfg *config.Config, commonServices *commonService.Services, commonAuthServices *commonAuthService.Services, grpcClients *grpc_client.Clients) *Services {
	repositories := repository.InitRepos(driver, cfg.Neo4j.Database)

	services := Services{
		CommonServices:                  commonServices,
		CommonAuthServices:              commonAuthServices,
		OrganizationService:             NewOrganizationService(log, repositories, grpcClients),
		CustomFieldService:              NewCustomFieldService(log, repositories),
		UserService:                     NewUserService(log, repositories, grpcClients),
		FieldSetService:                 NewFieldSetService(log, repositories),
		EntityTemplateService:           NewEntityTemplateService(log, repositories),
		FieldSetTemplateService:         NewFieldSetTemplateService(log, repositories),
		CustomFieldTemplateService:      NewCustomFieldTemplateService(log, repositories),
		LocationService:                 NewLocationService(log, repositories),
		TagService:                      NewTagService(log, repositories),
		DomainService:                   NewDomainService(log, repositories),
		PageViewService:                 NewPageViewService(log, repositories),
		AttachmentService:               NewAttachmentService(log, repositories),
		TenantService:                   NewTenantService(log, repositories),
		WorkspaceService:                NewWorkspaceService(log, repositories),
		SocialService:                   NewSocialService(log, repositories),
		OrganizationRelationshipService: NewOrganizationRelationshipService(log, repositories),
		ExternalSystemService:           NewExternalSystemService(log, repositories),
		ActionService:                   NewActionService(log, repositories),
		CountryService:                  NewCountryService(log, repositories),
		ActionItemService:               NewActionItemService(log, repositories),
		BillableService:                 NewBillableService(log, repositories),
		LogEntryService:                 NewLogEntryService(log, repositories),
	}
	services.IssueService = NewIssueService(log, repositories, &services)
	services.PhoneNumberService = NewPhoneNumberService(log, repositories, grpcClients, &services)
	services.JobRoleService = NewJobRoleService(log, repositories, &services)
	services.CalendarService = NewCalendarService(log, repositories, &services)
	services.EmailService = NewEmailService(log, repositories, &services, grpcClients)
	services.ContactService = NewContactService(log, repositories, grpcClients, &services)
	services.NoteService = NewNoteService(log, repositories, &services)
	services.TimelineEventService = NewTimelineEventService(log, repositories, &services)
	services.SearchService = NewSearchService(log, repositories, &services)
	services.QueryService = NewDashboardService(log, repositories, &services)
	services.InteractionEventService = NewInteractionEventService(log, repositories, &services)
	services.InteractionSessionService = NewInteractionSessionService(log, repositories, &services)
	services.AnalysisService = NewAnalysisService(log, repositories, &services)
	services.MeetingService = NewMeetingService(log, repositories, &services)
	services.PlayerService = NewPlayerService(repositories, &services)

	log.Info("Init cache service")
	services.Cache = NewCacheService(&services)
	services.Cache.InitCache()
	log.Info("Init cache service done")

	services.cfg = cfg
	return &services
}
