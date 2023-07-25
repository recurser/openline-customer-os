package events

import (
	"github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform/domain/organization/models"
	"github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform/eventstore"
	"github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform/validator"
	"time"
)

const (
	OrganizationCreateV1          = "V1_ORGANIZATION_CREATE"
	OrganizationUpdateV1          = "V1_ORGANIZATION_UPDATE"
	OrganizationPhoneNumberLinkV1 = "V1_ORGANIZATION_PHONE_NUMBER_LINK"
	OrganizationEmailLinkV1       = "V1_ORGANIZATION_EMAIL_LINK"
	OrganizationLinkDomainV1      = "V1_ORGANIZATION_LINK_DOMAIN"
	OrganizationAddSocialV1       = "V1_ORGANIZATION_ADD_SOCIAL"
)

type OrganizationCreateEvent struct {
	Tenant            string    `json:"tenant" validate:"required"`
	Name              string    `json:"name"`
	Description       string    `json:"description"`
	Website           string    `json:"website"`
	Industry          string    `json:"industry"`
	SubIndustry       string    `json:"subIndustry"`
	IndustryGroup     string    `json:"industryGroup"`
	TargetAudience    string    `json:"targetAudience"`
	ValueProposition  string    `json:"valueProposition"`
	IsPublic          bool      `json:"isPublic"`
	Employees         int64     `json:"employees"`
	Market            string    `json:"market"`
	LastFundingRound  string    `json:"lastFundingRound"`
	LastFundingAmount string    `json:"lastFundingAmount"`
	Source            string    `json:"source"`
	SourceOfTruth     string    `json:"sourceOfTruth"`
	AppSource         string    `json:"appSource"`
	CreatedAt         time.Time `json:"createdAt"`
	UpdatedAt         time.Time `json:"updatedAt"`
}

func NewOrganizationCreateEvent(aggregate eventstore.Aggregate, organizationFields *models.OrganizationFields, createdAt, updatedAt time.Time) (eventstore.Event, error) {
	eventData := OrganizationCreateEvent{
		Tenant:            organizationFields.Tenant,
		Name:              organizationFields.OrganizationDataFields.Name,
		Description:       organizationFields.OrganizationDataFields.Description,
		Website:           organizationFields.OrganizationDataFields.Website,
		Industry:          organizationFields.OrganizationDataFields.Industry,
		SubIndustry:       organizationFields.OrganizationDataFields.SubIndustry,
		IndustryGroup:     organizationFields.OrganizationDataFields.IndustryGroup,
		TargetAudience:    organizationFields.OrganizationDataFields.TargetAudience,
		ValueProposition:  organizationFields.OrganizationDataFields.ValueProposition,
		IsPublic:          organizationFields.OrganizationDataFields.IsPublic,
		Employees:         organizationFields.OrganizationDataFields.Employees,
		Market:            organizationFields.OrganizationDataFields.Market,
		LastFundingRound:  organizationFields.OrganizationDataFields.LastFundingRound,
		LastFundingAmount: organizationFields.OrganizationDataFields.LastFundingAmount,
		Source:            organizationFields.Source.Source,
		SourceOfTruth:     organizationFields.Source.SourceOfTruth,
		AppSource:         organizationFields.Source.AppSource,
		CreatedAt:         createdAt,
		UpdatedAt:         updatedAt,
	}

	if err := validator.GetValidator().Struct(eventData); err != nil {
		return eventstore.Event{}, err
	}

	event := eventstore.NewBaseEvent(aggregate, OrganizationCreateV1)
	if err := event.SetJsonData(&eventData); err != nil {
		return eventstore.Event{}, err
	}
	return event, nil
}

type OrganizationUpdateEvent struct {
	IgnoreEmptyFields bool      `json:"ignoreEmptyFields"`
	Tenant            string    `json:"tenant" validate:"required"`
	SourceOfTruth     string    `json:"sourceOfTruth"`
	UpdatedAt         time.Time `json:"updatedAt"`
	Name              string    `json:"name"`
	Description       string    `json:"description"`
	Website           string    `json:"website"`
	Industry          string    `json:"industry"`
	SubIndustry       string    `json:"subIndustry"`
	IndustryGroup     string    `json:"industryGroup"`
	TargetAudience    string    `json:"targetAudience"`
	ValueProposition  string    `json:"valueProposition"`
	IsPublic          bool      `json:"isPublic"`
	Employees         int64     `json:"employees"`
	Market            string    `json:"market"`
	LastFundingRound  string    `json:"lastFundingRound"`
	LastFundingAmount string    `json:"lastFundingAmount"`
}

func NewOrganizationUpdateEvent(aggregate eventstore.Aggregate, organizationFields *models.OrganizationFields, updatedAt time.Time, ignoreEmptyFields bool) (eventstore.Event, error) {
	eventData := OrganizationUpdateEvent{
		IgnoreEmptyFields: ignoreEmptyFields,
		Tenant:            organizationFields.Tenant,
		Name:              organizationFields.OrganizationDataFields.Name,
		Description:       organizationFields.OrganizationDataFields.Description,
		Website:           organizationFields.OrganizationDataFields.Website,
		Industry:          organizationFields.OrganizationDataFields.Industry,
		SubIndustry:       organizationFields.OrganizationDataFields.SubIndustry,
		IndustryGroup:     organizationFields.OrganizationDataFields.IndustryGroup,
		TargetAudience:    organizationFields.OrganizationDataFields.TargetAudience,
		ValueProposition:  organizationFields.OrganizationDataFields.ValueProposition,
		IsPublic:          organizationFields.OrganizationDataFields.IsPublic,
		Employees:         organizationFields.OrganizationDataFields.Employees,
		Market:            organizationFields.OrganizationDataFields.Market,
		LastFundingRound:  organizationFields.OrganizationDataFields.LastFundingRound,
		LastFundingAmount: organizationFields.OrganizationDataFields.LastFundingAmount,
		UpdatedAt:         updatedAt,
		SourceOfTruth:     organizationFields.Source.SourceOfTruth,
	}

	if err := validator.GetValidator().Struct(eventData); err != nil {
		return eventstore.Event{}, err
	}

	event := eventstore.NewBaseEvent(aggregate, OrganizationUpdateV1)
	if err := event.SetJsonData(&eventData); err != nil {
		return eventstore.Event{}, err
	}
	return event, nil
}

type OrganizationLinkPhoneNumberEvent struct {
	Tenant        string    `json:"tenant" validate:"required"`
	UpdatedAt     time.Time `json:"updatedAt"`
	PhoneNumberId string    `json:"phoneNumberId" validate:"required"`
	Label         string    `json:"label"`
	Primary       bool      `json:"primary"`
}

func NewOrganizationLinkPhoneNumberEvent(aggregate eventstore.Aggregate, tenant, phoneNumberId, label string, primary bool, updatedAt time.Time) (eventstore.Event, error) {
	eventData := OrganizationLinkPhoneNumberEvent{
		Tenant:        tenant,
		UpdatedAt:     updatedAt,
		PhoneNumberId: phoneNumberId,
		Label:         label,
		Primary:       primary,
	}

	if err := validator.GetValidator().Struct(eventData); err != nil {
		return eventstore.Event{}, err
	}

	event := eventstore.NewBaseEvent(aggregate, OrganizationPhoneNumberLinkV1)
	if err := event.SetJsonData(&eventData); err != nil {
		return eventstore.Event{}, err
	}
	return event, nil
}

type OrganizationLinkEmailEvent struct {
	Tenant    string    `json:"tenant" validate:"required"`
	UpdatedAt time.Time `json:"updatedAt"`
	EmailId   string    `json:"emailId" validate:"required"`
	Label     string    `json:"label"`
	Primary   bool      `json:"primary"`
}

func NewOrganizationLinkEmailEvent(aggregate eventstore.Aggregate, tenant, emailId, label string, primary bool, updatedAt time.Time) (eventstore.Event, error) {
	eventData := OrganizationLinkEmailEvent{
		Tenant:    tenant,
		UpdatedAt: updatedAt,
		EmailId:   emailId,
		Label:     label,
		Primary:   primary,
	}

	if err := validator.GetValidator().Struct(eventData); err != nil {
		return eventstore.Event{}, err
	}

	event := eventstore.NewBaseEvent(aggregate, OrganizationEmailLinkV1)
	if err := event.SetJsonData(&eventData); err != nil {
		return eventstore.Event{}, err
	}
	return event, nil
}

type OrganizationLinkDomainEvent struct {
	Tenant string `json:"tenant" validate:"required"`
	Domain string `json:"domain" validate:"required"`
}

func NewOrganizationLinkDomainEvent(aggregate eventstore.Aggregate, tenant, domain string) (eventstore.Event, error) {
	eventData := OrganizationLinkDomainEvent{
		Tenant: tenant,
		Domain: domain,
	}

	if err := validator.GetValidator().Struct(eventData); err != nil {
		return eventstore.Event{}, err
	}

	event := eventstore.NewBaseEvent(aggregate, OrganizationLinkDomainV1)
	if err := event.SetJsonData(&eventData); err != nil {
		return eventstore.Event{}, err
	}
	return event, nil
}

type OrganizationAddSocialEvent struct {
	Tenant        string    `json:"tenant" validate:"required"`
	SocialId      string    `json:"socialId" validate:"required"`
	PlatformName  string    `json:"platformName" validate:"required"`
	Url           string    `json:"url" validate:"required"`
	Source        string    `json:"source"`
	SourceOfTruth string    `json:"sourceOfTruth"`
	AppSource     string    `json:"appSource"`
	CreatedAt     time.Time `json:"createdAt"`
	UpdatedAt     time.Time `json:"updatedAt"`
}

func NewOrganizationAddSocialEvent(aggregate eventstore.Aggregate, tenant, socialId, platformName, url, source, sourceOfTruth, appSource string, createdAt time.Time, updatedAt time.Time) (eventstore.Event, error) {
	eventData := OrganizationAddSocialEvent{
		Tenant:        tenant,
		SocialId:      socialId,
		PlatformName:  platformName,
		Url:           url,
		Source:        source,
		SourceOfTruth: sourceOfTruth,
		AppSource:     appSource,
		CreatedAt:     createdAt,
		UpdatedAt:     updatedAt,
	}

	if err := validator.GetValidator().Struct(eventData); err != nil {
		return eventstore.Event{}, err
	}

	event := eventstore.NewBaseEvent(aggregate, OrganizationAddSocialV1)
	if err := event.SetJsonData(&eventData); err != nil {
		return eventstore.Event{}, err
	}
	return event, nil
}
