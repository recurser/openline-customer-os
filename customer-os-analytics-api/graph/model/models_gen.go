// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"
	"time"
)

type PagedResult interface {
	IsPagedResult()
	GetTotalPages() int
	GetTotalElements() int64
}

type AppSession struct {
	ID             string    `json:"id"`
	Country        string    `json:"country"`
	Region         string    `json:"region"`
	City           string    `json:"city"`
	ReferrerSource string    `json:"referrerSource"`
	UtmCampaign    string    `json:"utmCampaign"`
	UtmContent     string    `json:"utmContent"`
	UtmMedium      string    `json:"utmMedium"`
	UtmSource      string    `json:"utmSource"`
	UtmNetwork     string    `json:"utmNetwork"`
	UtmTerm        string    `json:"utmTerm"`
	StartedAt      time.Time `json:"startedAt"`
	EndedAt        time.Time `json:"endedAt"`
	EngagedTime    int       `json:"engagedTime"`
}

type AppSessionsDataFilter struct {
	Field  AppSessionField `json:"Field"`
	Action Operation       `json:"Action"`
	Value  string          `json:"Value"`
}

type AppSessionsPage struct {
	Content       []*AppSession `json:"content"`
	TotalPages    int           `json:"totalPages"`
	TotalElements int64         `json:"totalElements"`
}

func (AppSessionsPage) IsPagedResult()               {}
func (this AppSessionsPage) GetTotalPages() int      { return this.TotalPages }
func (this AppSessionsPage) GetTotalElements() int64 { return this.TotalElements }

type PaginationFilter struct {
	Page  int `json:"page"`
	Limit int `json:"limit"`
}

type TimeFilter struct {
	TimePeriod TimePeriod `json:"timePeriod"`
	From       *time.Time `json:"from"`
	To         *time.Time `json:"to"`
}

type AppSessionField string

const (
	AppSessionFieldCountry        AppSessionField = "COUNTRY"
	AppSessionFieldCity           AppSessionField = "CITY"
	AppSessionFieldRegion         AppSessionField = "REGION"
	AppSessionFieldReferrerSource AppSessionField = "REFERRER_SOURCE"
	AppSessionFieldUtmCampaign    AppSessionField = "UTM_CAMPAIGN"
	AppSessionFieldUtmContent     AppSessionField = "UTM_CONTENT"
	AppSessionFieldUtmMedium      AppSessionField = "UTM_MEDIUM"
	AppSessionFieldUtmSource      AppSessionField = "UTM_SOURCE"
	AppSessionFieldUtmNetwork     AppSessionField = "UTM_NETWORK"
	AppSessionFieldUtmTerm        AppSessionField = "UTM_TERM"
)

var AllAppSessionField = []AppSessionField{
	AppSessionFieldCountry,
	AppSessionFieldCity,
	AppSessionFieldRegion,
	AppSessionFieldReferrerSource,
	AppSessionFieldUtmCampaign,
	AppSessionFieldUtmContent,
	AppSessionFieldUtmMedium,
	AppSessionFieldUtmSource,
	AppSessionFieldUtmNetwork,
	AppSessionFieldUtmTerm,
}

func (e AppSessionField) IsValid() bool {
	switch e {
	case AppSessionFieldCountry, AppSessionFieldCity, AppSessionFieldRegion, AppSessionFieldReferrerSource, AppSessionFieldUtmCampaign, AppSessionFieldUtmContent, AppSessionFieldUtmMedium, AppSessionFieldUtmSource, AppSessionFieldUtmNetwork, AppSessionFieldUtmTerm:
		return true
	}
	return false
}

func (e AppSessionField) String() string {
	return string(e)
}

func (e *AppSessionField) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = AppSessionField(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid AppSessionField", str)
	}
	return nil
}

func (e AppSessionField) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type Operation string

const (
	OperationEquals   Operation = "EQUALS"
	OperationContains Operation = "CONTAINS"
)

var AllOperation = []Operation{
	OperationEquals,
	OperationContains,
}

func (e Operation) IsValid() bool {
	switch e {
	case OperationEquals, OperationContains:
		return true
	}
	return false
}

func (e Operation) String() string {
	return string(e)
}

func (e *Operation) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = Operation(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid Operation", str)
	}
	return nil
}

func (e Operation) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type TimePeriod string

const (
	TimePeriodToday       TimePeriod = "TODAY"
	TimePeriodLastHour    TimePeriod = "LAST_HOUR"
	TimePeriodLast24Hours TimePeriod = "LAST_24_HOURS"
	TimePeriodLast7Days   TimePeriod = "LAST_7_DAYS"
	TimePeriodLast30Days  TimePeriod = "LAST_30_DAYS"
	TimePeriodMonthToDate TimePeriod = "MONTH_TO_DATE"
	TimePeriodYearToDate  TimePeriod = "YEAR_TO_DATE"
	TimePeriodDaily       TimePeriod = "DAILY"
	TimePeriodMonthly     TimePeriod = "MONTHLY"
	TimePeriodAllTime     TimePeriod = "ALL_TIME"
	TimePeriodCustom      TimePeriod = "CUSTOM"
)

var AllTimePeriod = []TimePeriod{
	TimePeriodToday,
	TimePeriodLastHour,
	TimePeriodLast24Hours,
	TimePeriodLast7Days,
	TimePeriodLast30Days,
	TimePeriodMonthToDate,
	TimePeriodYearToDate,
	TimePeriodDaily,
	TimePeriodMonthly,
	TimePeriodAllTime,
	TimePeriodCustom,
}

func (e TimePeriod) IsValid() bool {
	switch e {
	case TimePeriodToday, TimePeriodLastHour, TimePeriodLast24Hours, TimePeriodLast7Days, TimePeriodLast30Days, TimePeriodMonthToDate, TimePeriodYearToDate, TimePeriodDaily, TimePeriodMonthly, TimePeriodAllTime, TimePeriodCustom:
		return true
	}
	return false
}

func (e TimePeriod) String() string {
	return string(e)
}

func (e *TimePeriod) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = TimePeriod(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid TimePeriod", str)
	}
	return nil
}

func (e TimePeriod) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
