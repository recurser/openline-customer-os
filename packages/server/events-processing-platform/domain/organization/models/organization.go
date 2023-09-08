package models

import (
	"fmt"
	common_models "github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform/domain/common/models"
	"time"
)

type RenewalLikelihoodProbability string

const (
	RenewalLikelihoodHIGH   RenewalLikelihoodProbability = "HIGH"
	RenewalLikelihoodMEDIUM RenewalLikelihoodProbability = "MEDIUM"
	RenewalLikelihoodLOW    RenewalLikelihoodProbability = "LOW"
	RenewalLikelihoodZERO   RenewalLikelihoodProbability = "ZERO"
)

func (r RenewalLikelihoodProbability) CamelCaseString() string {
	switch r {
	case RenewalLikelihoodHIGH:
		return "High"
	case RenewalLikelihoodMEDIUM:
		return "Medium"
	case RenewalLikelihoodLOW:
		return "Low"
	case RenewalLikelihoodZERO:
		return "Zero"
	}
	return ""
}

type Social struct {
	PlatformName string `json:"platformName"`
	Url          string `json:"url"`
}

type Organization struct {
	ID                string                             `json:"id"`
	Name              string                             `json:"name"`
	Hide              bool                               `json:"hide"`
	Description       string                             `json:"description"`
	Website           string                             `json:"website"`
	Industry          string                             `json:"industry"`
	SubIndustry       string                             `json:"subIndustry"`
	IndustryGroup     string                             `json:"industryGroup"`
	TargetAudience    string                             `json:"targetAudience"`
	ValueProposition  string                             `json:"valueProposition"`
	IsPublic          bool                               `json:"isPublic"`
	Employees         int64                              `json:"employees"`
	Market            string                             `json:"market"`
	LastFundingRound  string                             `json:"lastFundingRound"`
	LastFundingAmount string                             `json:"lastFundingAmount"`
	Source            common_models.Source               `json:"source"`
	CreatedAt         time.Time                          `json:"createdAt,omitempty"`
	UpdatedAt         time.Time                          `json:"updatedAt,omitempty"`
	PhoneNumbers      map[string]OrganizationPhoneNumber `json:"phoneNumbers"`
	Emails            map[string]OrganizationEmail       `json:"emails"`
	Domains           []string                           `json:"domains,omitempty"`
	Socials           map[string]Social                  `json:"socials,omitempty"`
	RenewalLikelihood RenewalLikelihood                  `json:"renewalLikelihood,omitempty"`
	RenewalForecast   RenewalForecast                    `json:"renewalForecast,omitempty"`
	BillingDetails    BillingDetails                     `json:"billingDetails,omitempty"`
}

type RenewalLikelihood struct {
	RenewalLikelihood RenewalLikelihoodProbability `json:"renewalLikelihood,omitempty"`
	Comment           *string                      `json:"comment,omitempty"`
	UpdatedAt         time.Time                    `json:"updatedAt,omitempty"`
	UpdatedBy         string                       `json:"updatedBy,omitempty"`
}

type RenewalForecast struct {
	Amount    *float64  `json:"amount,omitempty"`
	Comment   *string   `json:"comment,omitempty"`
	UpdatedAt time.Time `json:"updatedAt,omitempty"`
	UpdatedBy string    `json:"updatedBy,omitempty"`
}

type BillingDetails struct {
	Amount            *float64   `json:"amount,omitempty"`
	UpdatedBy         string     `json:"updatedBy,omitempty"`
	Frequency         string     `json:"frequency,omitempty"`
	RenewalCycle      string     `json:"renewalCycle,omitempty"`
	RenewalCycleStart *time.Time `json:"renewalCycleStart,omitempty"`
	RenewalCycleNext  *time.Time `json:"renewalCycleNext,omitempty"`
}

type OrganizationPhoneNumber struct {
	Primary bool   `json:"primary"`
	Label   string `json:"label"`
}

type OrganizationEmail struct {
	Primary bool   `json:"primary"`
	Label   string `json:"label"`
}

func (o *Organization) String() string {
	return fmt.Sprintf("Organization{ID: %s, Name: %s, Description: %s, Website: %s, Industry: %s, IsPublic: %t, Source: %s, CreatedAt: %s, UpdatedAt: %s}", o.ID, o.Name, o.Description, o.Website, o.Industry, o.IsPublic, o.Source, o.CreatedAt, o.UpdatedAt)
}
