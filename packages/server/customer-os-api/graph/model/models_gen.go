// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"
	"time"
)

type Action interface {
	IsAction()
}

type ExtensibleEntity interface {
	IsNode()
	IsExtensibleEntity()
	GetID() string
	GetDefinition() *EntityDefinition
}

type Node interface {
	IsNode()
	GetID() string
}

// Describes the number of pages and total elements included in a query response.
// **A `response` object.**
type Pages interface {
	IsPages()
	// The total number of pages included in the query response.
	// **Required.**
	GetTotalPages() int
	// The total number of elements included in the query response.
	// **Required.**
	GetTotalElements() int64
}

type CallAction struct {
	ID        string    `json:"id"`
	StartedAt time.Time `json:"startedAt"`
	EndedAt   time.Time `json:"endedAt"`
}

func (CallAction) IsAction() {}

func (CallAction) IsNode()            {}
func (this CallAction) GetID() string { return this.ID }

type ChatAction struct {
	ID        string    `json:"id"`
	StartedAt time.Time `json:"startedAt"`
}

func (ChatAction) IsAction() {}

func (ChatAction) IsNode()            {}
func (this ChatAction) GetID() string { return this.ID }

type Company struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type CompanyInput struct {
	ID *string `json:"id"`
	// The name of the company associated with a Contact.
	// **Required.**
	Name *string `json:"name"`
}

type CompanyPage struct {
	Content       []*Company `json:"content"`
	TotalPages    int        `json:"totalPages"`
	TotalElements int64      `json:"totalElements"`
}

func (CompanyPage) IsPages() {}

// The total number of pages included in the query response.
// **Required.**
func (this CompanyPage) GetTotalPages() int { return this.TotalPages }

// The total number of elements included in the query response.
// **Required.**
func (this CompanyPage) GetTotalElements() int64 { return this.TotalElements }

// Describes the relationship a Contact has with a Company.
// **A `return` object**
type CompanyPosition struct {
	ID string `json:"id"`
	// Company associated with a Contact.
	// **Required.**
	Company *Company `json:"company"`
	// The Contact's job title.
	JobTitle *string `json:"jobTitle"`
}

// Describes the relationship a Contact has with a Company.
// **A `create` object**
type CompanyPositionInput struct {
	Company *CompanyInput `json:"company"`
	// The Contact's job title.
	JobTitle *string `json:"jobTitle"`
}

// A contact represents an individual in customerOS.
// **A `response` object.**
type Contact struct {
	// The unique ID associated with the contact in customerOS.
	// **Required**
	ID string `json:"id"`
	// The title associate with the contact in customerOS.
	Title *PersonTitle `json:"title"`
	// The first name of the contact in customerOS.
	// **Required**"
	FirstName string `json:"firstName"`
	// The last name of the contact in customerOS.
	// **Required**
	LastName string `json:"lastName"`
	// An ISO8601 timestamp recording when the contact was created in customerOS.
	// **Required**
	CreatedAt time.Time `json:"createdAt"`
	// A user-defined label applied against a contact in customerOS.
	Label *string `json:"label"`
	// User-defined notes associated with a contact in customerOS.
	Notes *string `json:"notes"`
	// User-defined field that defines the relationship type the contact has with your business.  `Customer`, `Partner`, `Lead` are examples.
	ContactType *ContactType `json:"contactType"`
	// `companyName` and `jobTitle` of the contact if it has been associated with a company.
	// **Required.  If no values it returns an empty array.**
	CompanyPositions []*CompanyPosition `json:"companyPositions"`
	// Identifies any contact groups the contact is associated with.
	//  **Required.  If no values it returns an empty array.**
	Groups []*ContactGroup `json:"groups"`
	// All phone numbers associated with a contact in customerOS.
	// **Required.  If no values it returns an empty array.**
	PhoneNumbers []*PhoneNumber `json:"phoneNumbers"`
	// All email addresses assocaited with a contact in customerOS.
	// **Required.  If no values it returns an empty array.**
	Emails []*Email `json:"emails"`
	// User defined metadata appended to the contact record in customerOS.
	// **Required.  If no values it returns an empty array.**
	CustomFields []*CustomField `json:"customFields"`
	FieldSets    []*FieldSet    `json:"fieldSets"`
	// Definition of the contact in customerOS.
	Definition *EntityDefinition `json:"definition"`
	// Contact owner (user)
	Owner         *User             `json:"owner"`
	Conversations *ConversationPage `json:"conversations"`
	Actions       []Action          `json:"actions"`
}

func (Contact) IsExtensibleEntity()                   {}
func (this Contact) GetID() string                    { return this.ID }
func (this Contact) GetDefinition() *EntityDefinition { return this.Definition }

func (Contact) IsNode() {}

// A collection of groups that a Contact belongs to.  Groups are user-defined entities.
// **A `return` object.**
type ContactGroup struct {
	// The unique ID associated with the `ContactGroup` in customerOS.
	// **Required**
	ID string `json:"id"`
	// The name of the `ContactGroup`.
	// **Required**
	Name     string        `json:"name"`
	Contacts *ContactsPage `json:"contacts"`
}

// Create a groups that can be associated with a `Contact` in customerOS.
// **A `create` object.**
type ContactGroupInput struct {
	// The name of the `ContactGroup`.
	// **Required**
	Name string `json:"name"`
}

// Specifies how many pages of `ContactGroup` information has been returned in the query response.
// **A `response` object.**
type ContactGroupPage struct {
	// A collection of groups that a Contact belongs to.  Groups are user-defined entities.
	// **Required.  If no values it returns an empty array.**
	Content []*ContactGroup `json:"content"`
	// Total number of pages in the query response.
	// **Required.**
	TotalPages int `json:"totalPages"`
	// Total number of elements in the query response.
	// **Required.**
	TotalElements int64 `json:"totalElements"`
}

func (ContactGroupPage) IsPages() {}

// The total number of pages included in the query response.
// **Required.**
func (this ContactGroupPage) GetTotalPages() int { return this.TotalPages }

// The total number of elements included in the query response.
// **Required.**
func (this ContactGroupPage) GetTotalElements() int64 { return this.TotalElements }

// Update a group that can be associated with a `Contact` in customerOS.
// **A `update` object.**
type ContactGroupUpdateInput struct {
	// The unique ID associated with the `ContactGroup` in customerOS.
	// **Required**
	ID string `json:"id"`
	// The name of the `ContactGroup`.
	// **Required**
	Name string `json:"name"`
}

// Create an individual in customerOS.
// **A `create` object.**
type ContactInput struct {
	// The unique ID associated with the definition of the contact in customerOS.
	DefinitionID *string `json:"definitionId"`
	// The title of the contact.
	Title *PersonTitle `json:"title"`
	// The first name of the contact.
	// **Required.**
	FirstName string `json:"firstName"`
	// The last name of the contact.
	// **Required.**
	LastName string `json:"lastName"`
	// A user-defined label attached to contact.
	Label *string `json:"label"`
	// User-defined notes associated with contact.
	Notes *string `json:"notes"`
	// User-defined field that defines the relationship type the contact has with your business.  `Customer`, `Partner`, `Lead` are examples.
	ContactTypeID *string `json:"contactTypeId"`
	// User defined metadata appended to contact.
	// **Required.**
	CustomFields []*CustomFieldInput `json:"customFields"`
	FieldSets    []*FieldSetInput    `json:"fieldSets"`
	// An email addresses associted with the contact.
	Email *EmailInput `json:"email"`
	// A phone number associated with the contact.
	PhoneNumber *PhoneNumberInput `json:"phoneNumber"`
	// Id of the contact owner (user)
	OwnerID *string `json:"ownerId"`
}

type ContactType struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type ContactTypeInput struct {
	Name string `json:"name"`
}

type ContactTypeUpdateInput struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// Updates data fields associated with an existing customer record in customerOS.
// **An `update` object.**
type ContactUpdateInput struct {
	// The unique ID associated with the contact in customerOS.
	// **Required.**
	ID string `json:"id"`
	// The title associate with the contact in customerOS.
	Title *PersonTitle `json:"title"`
	// The first name of the contact in customerOS.
	// **Required.**
	FirstName string `json:"firstName"`
	// The last name of the contact in customerOS.
	// **Required.**
	LastName string `json:"lastName"`
	// A user-defined label applied against a contact in customerOS.
	Label *string `json:"label"`
	// User-defined notes associated with contact.
	Notes *string `json:"notes"`
	// User-defined field that defines the relationship type the contact has with your business.  `Customer`, `Partner`, `Lead` are examples.
	ContactTypeID *string `json:"contactTypeId"`
	// Id of the contact owner (user)
	OwnerID *string `json:"ownerId"`
}

// Specifies how many pages of contact information has been returned in the query response.
// **A `response` object.**
type ContactsPage struct {
	// A contact entity in customerOS.
	// **Required.  If no values it returns an empty array.**
	Content []*Contact `json:"content"`
	// Total number of pages in the query response.
	// **Required.**
	TotalPages int `json:"totalPages"`
	// Total number of elements in the query response.
	// **Required.**
	TotalElements int64 `json:"totalElements"`
}

func (ContactsPage) IsPages() {}

// The total number of pages included in the query response.
// **Required.**
func (this ContactsPage) GetTotalPages() int { return this.TotalPages }

// The total number of elements included in the query response.
// **Required.**
func (this ContactsPage) GetTotalElements() int64 { return this.TotalElements }

type Conversation struct {
	ID        string    `json:"id"`
	StartedAt time.Time `json:"startedAt"`
	ContactID string    `json:"contactId"`
	Contact   *Contact  `json:"contact"`
	UserID    string    `json:"userId"`
	User      *User     `json:"user"`
}

func (Conversation) IsNode()            {}
func (this Conversation) GetID() string { return this.ID }

type ConversationInput struct {
	UserID    string  `json:"userId"`
	ContactID string  `json:"contactId"`
	ID        *string `json:"id"`
}

type ConversationPage struct {
	Content       []*Conversation `json:"content"`
	TotalPages    int             `json:"totalPages"`
	TotalElements int64           `json:"totalElements"`
}

func (ConversationPage) IsPages() {}

// The total number of pages included in the query response.
// **Required.**
func (this ConversationPage) GetTotalPages() int { return this.TotalPages }

// The total number of elements included in the query response.
// **Required.**
func (this ConversationPage) GetTotalElements() int64 { return this.TotalElements }

// Describes a custom, user-defined field associated with a `Contact`.
// **A `return` object.**
type CustomField struct {
	// The unique ID associated with the custom field.
	// **Required**
	ID string `json:"id"`
	// The name of the custom field.
	// **Required**
	Name string `json:"name"`
	// Datatype of the custom field.
	// **Required**
	Datatype CustomFieldDataType `json:"datatype"`
	// The value of the custom field.
	// **Required**
	Value      AnyTypeValue           `json:"value"`
	Definition *CustomFieldDefinition `json:"definition"`
}

func (CustomField) IsNode()            {}
func (this CustomField) GetID() string { return this.ID }

type CustomFieldDefinition struct {
	ID        string                    `json:"id"`
	Name      string                    `json:"name"`
	Type      CustomFieldDefinitionType `json:"type"`
	Order     int                       `json:"order"`
	Mandatory bool                      `json:"mandatory"`
	Length    *int                      `json:"length"`
	Min       *int                      `json:"min"`
	Max       *int                      `json:"max"`
}

func (CustomFieldDefinition) IsNode()            {}
func (this CustomFieldDefinition) GetID() string { return this.ID }

type CustomFieldDefinitionInput struct {
	Name      string                    `json:"name"`
	Type      CustomFieldDefinitionType `json:"type"`
	Order     int                       `json:"order"`
	Mandatory bool                      `json:"mandatory"`
	Length    *int                      `json:"length"`
	Min       *int                      `json:"min"`
	Max       *int                      `json:"max"`
}

// Describes a custom, user-defined field associated with a `Contact` of type String.
// **A `create` object.**
type CustomFieldInput struct {
	// The unique ID associated with the custom field.
	ID *string `json:"id"`
	// The name of the custom field.
	// **Required**
	Name string `json:"name"`
	// Datatype of the custom field.
	// **Required**
	Datatype CustomFieldDataType `json:"datatype"`
	// The value of the custom field.
	// **Required**
	Value        AnyTypeValue `json:"value"`
	DefinitionID *string      `json:"definitionId"`
}

// Describes a custom, user-defined field associated with a `Contact`.
// **An `update` object.**
type CustomFieldUpdateInput struct {
	// The unique ID associated with the custom field.
	// **Required**
	ID string `json:"id"`
	// The name of the custom field.
	// **Required**
	Name string `json:"name"`
	// Datatype of the custom field.
	// **Required**
	Datatype CustomFieldDataType `json:"datatype"`
	// The value of the custom field.
	// **Required**
	Value AnyTypeValue `json:"value"`
}

// Describes an email address associated with a `Contact` in customerOS.
// **A `return` object.**
type Email struct {
	// The unique ID associated with the contact in customerOS.
	// **Required**
	ID string `json:"id"`
	// An email address assocaited with the contact in customerOS.
	// **Required.**
	Email string `json:"email"`
	// Describes the type of email address (WORK, PERSONAL, etc).
	// **Required.**
	Label EmailLabel `json:"label"`
	// Identifies whether the email address is primary or not.
	// **Required.**
	Primary bool `json:"primary"`
}

type EmailAction struct {
	ID        string    `json:"id"`
	StartedAt time.Time `json:"startedAt"`
}

func (EmailAction) IsAction() {}

func (EmailAction) IsNode()            {}
func (this EmailAction) GetID() string { return this.ID }

// Describes an email address associated with a `Contact` in customerOS.
// **A `create` object.**
type EmailInput struct {
	// An email address assocaited with the contact in customerOS.
	// **Required.**
	Email string `json:"email"`
	// Describes the type of email address (WORK, PERSONAL, etc).
	// **Required.**
	Label EmailLabel `json:"label"`
	// Identifies whether the email address is primary or not.
	// **Required.**
	Primary *bool `json:"primary"`
}

// Describes an email address associated with a `Contact` in customerOS.
// **An `update` object.**
type EmailUpdateInput struct {
	// An email address assocaited with the contact in customerOS.
	// **Required.**
	ID string `json:"id"`
	// An email address assocaited with the contact in customerOS.
	// **Required.**
	Email string `json:"email"`
	// Describes the type of email address (WORK, PERSONAL, etc).
	// **Required.**
	Label EmailLabel `json:"label"`
	// Identifies whether the email address is primary or not.
	// **Required.**
	Primary *bool `json:"primary"`
}

type EntityDefinition struct {
	ID           string                     `json:"id"`
	Version      int                        `json:"version"`
	Name         string                     `json:"name"`
	Extends      *EntityDefinitionExtension `json:"extends"`
	FieldSets    []*FieldSetDefinition      `json:"fieldSets"`
	CustomFields []*CustomFieldDefinition   `json:"customFields"`
	Added        time.Time                  `json:"added"`
}

func (EntityDefinition) IsNode()            {}
func (this EntityDefinition) GetID() string { return this.ID }

type EntityDefinitionInput struct {
	Name         string                        `json:"name"`
	Extends      *EntityDefinitionExtension    `json:"extends"`
	FieldSets    []*FieldSetDefinitionInput    `json:"fieldSets"`
	CustomFields []*CustomFieldDefinitionInput `json:"customFields"`
}

type FieldSet struct {
	ID           string              `json:"id"`
	Name         string              `json:"name"`
	Added        time.Time           `json:"added"`
	CustomFields []*CustomField      `json:"customFields"`
	Definition   *FieldSetDefinition `json:"definition"`
}

type FieldSetDefinition struct {
	ID           string                   `json:"id"`
	Name         string                   `json:"name"`
	Order        int                      `json:"order"`
	CustomFields []*CustomFieldDefinition `json:"customFields"`
}

func (FieldSetDefinition) IsNode()            {}
func (this FieldSetDefinition) GetID() string { return this.ID }

type FieldSetDefinitionInput struct {
	Name         string                        `json:"name"`
	Order        int                           `json:"order"`
	CustomFields []*CustomFieldDefinitionInput `json:"customFields"`
}

type FieldSetInput struct {
	ID           *string             `json:"id"`
	Name         string              `json:"name"`
	CustomFields []*CustomFieldInput `json:"customFields"`
	DefinitionID *string             `json:"definitionId"`
}

type FieldSetUpdateInput struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Filter struct {
	Not    *Filter     `json:"NOT"`
	And    []*Filter   `json:"AND"`
	Or     []*Filter   `json:"OR"`
	Filter *FilterItem `json:"filter"`
}

type FilterItem struct {
	Property      string             `json:"property"`
	Operation     ComparisonOperator `json:"operation"`
	Value         AnyTypeValue       `json:"value"`
	CaseSensitive *bool              `json:"caseSensitive"`
}

type PageViewAction struct {
	ID             string    `json:"id"`
	StartedAt      time.Time `json:"startedAt"`
	EndedAt        time.Time `json:"endedAt"`
	PageTitle      *string   `json:"pageTitle"`
	PageURL        *string   `json:"pageUrl"`
	Application    string    `json:"application"`
	OrderInSession int       `json:"orderInSession"`
	EngagedTime    int       `json:"engagedTime"`
	SessionID      string    `json:"sessionId"`
}

func (PageViewAction) IsAction() {}

func (PageViewAction) IsNode()            {}
func (this PageViewAction) GetID() string { return this.ID }

// If provided as part of the request, results will be filtered down to the `page` and `limit` specified.
type Pagination struct {
	// The results page to return in the response.
	// **Required.**
	Page int `json:"page"`
	// The maximum number of results in the response.
	// **Required.**
	Limit int `json:"limit"`
}

// Describes a phone number associated with a `Contact` in customerOS.
// **A `return` object.**
type PhoneNumber struct {
	// The unique ID associated with the phone number.
	// **Required**
	ID string `json:"id"`
	// The phone number in e164 format.
	// **Required**
	E164 string `json:"e164"`
	// Defines the type of phone number.
	// **Required**
	Label PhoneNumberLabel `json:"label"`
	// Determines if the phone number is primary or not.
	// **Required**
	Primary bool `json:"primary"`
}

// Describes a phone number associated with a `Contact` in customerOS.
// **A `create` object.**
type PhoneNumberInput struct {
	// The phone number in e164 format.
	// **Required**
	E164 string `json:"e164"`
	// Defines the type of phone number.
	// **Required**
	Label PhoneNumberLabel `json:"label"`
	// Determines if the phone number is primary or not.
	// **Required**
	Primary *bool `json:"primary"`
}

// Describes a phone number associated with a `Contact` in customerOS.
// **An `update` object.**
type PhoneNumberUpdateInput struct {
	// The unique ID associated with the phone number.
	// **Required**
	ID string `json:"id"`
	// The phone number in e164 format.
	// **Required**
	E164 string `json:"e164"`
	// Defines the type of phone number.
	// **Required**
	Label PhoneNumberLabel `json:"label"`
	// Determines if the phone number is primary or not.
	// **Required**
	Primary *bool `json:"primary"`
}

// Describes the success or failure of the GraphQL call.
// **A `return` object**
type Result struct {
	// The result of the GraphQL call.
	// **Required.**
	Result bool `json:"result"`
}

type SortBy struct {
	By            string           `json:"by"`
	Direction     SortingDirection `json:"direction"`
	CaseSensitive *bool            `json:"caseSensitive"`
}

// Describes the User of customerOS.  A user is the person who logs into the Openline platform.
// **A `return` object**
type User struct {
	// The unique ID associated with the customerOS user.
	// **Required**
	ID string `json:"id"`
	// The first name of the customerOS user.
	// **Required**
	FirstName string `json:"firstName"`
	// The last name of the customerOS user.
	// **Required**
	LastName string `json:"lastName"`
	// The email address of the customerOS user.
	// **Required**
	Email string `json:"email"`
	// Timestamp of user creation.
	// **Required**
	CreatedAt     time.Time         `json:"createdAt"`
	Conversations *ConversationPage `json:"conversations"`
}

// Describes the User of customerOS.  A user is the person who logs into the Openline platform.
// **A `create` object.**
type UserInput struct {
	// The first name of the customerOS user.
	// **Required**
	FirstName string `json:"firstName"`
	// The last name of the customerOS user.
	// **Required**
	LastName string `json:"lastName"`
	// The email address of the customerOS user.
	// **Required**
	Email string `json:"email"`
}

// Specifies how many pages of `User` information has been returned in the query response.
// **A `return` object.**
type UserPage struct {
	// A `User` entity in customerOS.
	// **Required.  If no values it returns an empty array.**
	Content []*User `json:"content"`
	// Total number of pages in the query response.
	// **Required.**
	TotalPages int `json:"totalPages"`
	// Total number of elements in the query response.
	// **Required.**
	TotalElements int64 `json:"totalElements"`
}

func (UserPage) IsPages() {}

// The total number of pages included in the query response.
// **Required.**
func (this UserPage) GetTotalPages() int { return this.TotalPages }

// The total number of elements included in the query response.
// **Required.**
func (this UserPage) GetTotalElements() int64 { return this.TotalElements }

type ActionType string

const (
	ActionTypePageView ActionType = "PAGE_VIEW"
	ActionTypeCall     ActionType = "CALL"
	ActionTypeEmail    ActionType = "EMAIL"
	ActionTypeChat     ActionType = "CHAT"
)

var AllActionType = []ActionType{
	ActionTypePageView,
	ActionTypeCall,
	ActionTypeEmail,
	ActionTypeChat,
}

func (e ActionType) IsValid() bool {
	switch e {
	case ActionTypePageView, ActionTypeCall, ActionTypeEmail, ActionTypeChat:
		return true
	}
	return false
}

func (e ActionType) String() string {
	return string(e)
}

func (e *ActionType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = ActionType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid ActionType", str)
	}
	return nil
}

func (e ActionType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type ComparisonOperator string

const (
	ComparisonOperatorEq       ComparisonOperator = "EQ"
	ComparisonOperatorContains ComparisonOperator = "CONTAINS"
)

var AllComparisonOperator = []ComparisonOperator{
	ComparisonOperatorEq,
	ComparisonOperatorContains,
}

func (e ComparisonOperator) IsValid() bool {
	switch e {
	case ComparisonOperatorEq, ComparisonOperatorContains:
		return true
	}
	return false
}

func (e ComparisonOperator) String() string {
	return string(e)
}

func (e *ComparisonOperator) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = ComparisonOperator(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid ComparisonOperator", str)
	}
	return nil
}

func (e ComparisonOperator) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type CustomFieldDataType string

const (
	CustomFieldDataTypeText     CustomFieldDataType = "TEXT"
	CustomFieldDataTypeBool     CustomFieldDataType = "BOOL"
	CustomFieldDataTypeDatetime CustomFieldDataType = "DATETIME"
	CustomFieldDataTypeInteger  CustomFieldDataType = "INTEGER"
	CustomFieldDataTypeDecimal  CustomFieldDataType = "DECIMAL"
)

var AllCustomFieldDataType = []CustomFieldDataType{
	CustomFieldDataTypeText,
	CustomFieldDataTypeBool,
	CustomFieldDataTypeDatetime,
	CustomFieldDataTypeInteger,
	CustomFieldDataTypeDecimal,
}

func (e CustomFieldDataType) IsValid() bool {
	switch e {
	case CustomFieldDataTypeText, CustomFieldDataTypeBool, CustomFieldDataTypeDatetime, CustomFieldDataTypeInteger, CustomFieldDataTypeDecimal:
		return true
	}
	return false
}

func (e CustomFieldDataType) String() string {
	return string(e)
}

func (e *CustomFieldDataType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = CustomFieldDataType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid CustomFieldDataType", str)
	}
	return nil
}

func (e CustomFieldDataType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type CustomFieldDefinitionType string

const (
	CustomFieldDefinitionTypeText CustomFieldDefinitionType = "TEXT"
)

var AllCustomFieldDefinitionType = []CustomFieldDefinitionType{
	CustomFieldDefinitionTypeText,
}

func (e CustomFieldDefinitionType) IsValid() bool {
	switch e {
	case CustomFieldDefinitionTypeText:
		return true
	}
	return false
}

func (e CustomFieldDefinitionType) String() string {
	return string(e)
}

func (e *CustomFieldDefinitionType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = CustomFieldDefinitionType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid CustomFieldDefinitionType", str)
	}
	return nil
}

func (e CustomFieldDefinitionType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

// Describes the type of email address (WORK, PERSONAL, etc).
// **A `return` object.
type EmailLabel string

const (
	EmailLabelMain  EmailLabel = "MAIN"
	EmailLabelWork  EmailLabel = "WORK"
	EmailLabelHome  EmailLabel = "HOME"
	EmailLabelOther EmailLabel = "OTHER"
)

var AllEmailLabel = []EmailLabel{
	EmailLabelMain,
	EmailLabelWork,
	EmailLabelHome,
	EmailLabelOther,
}

func (e EmailLabel) IsValid() bool {
	switch e {
	case EmailLabelMain, EmailLabelWork, EmailLabelHome, EmailLabelOther:
		return true
	}
	return false
}

func (e EmailLabel) String() string {
	return string(e)
}

func (e *EmailLabel) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = EmailLabel(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid EmailLabel", str)
	}
	return nil
}

func (e EmailLabel) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type EntityDefinitionExtension string

const (
	EntityDefinitionExtensionContact EntityDefinitionExtension = "CONTACT"
)

var AllEntityDefinitionExtension = []EntityDefinitionExtension{
	EntityDefinitionExtensionContact,
}

func (e EntityDefinitionExtension) IsValid() bool {
	switch e {
	case EntityDefinitionExtensionContact:
		return true
	}
	return false
}

func (e EntityDefinitionExtension) String() string {
	return string(e)
}

func (e *EntityDefinitionExtension) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = EntityDefinitionExtension(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid EntityDefinitionExtension", str)
	}
	return nil
}

func (e EntityDefinitionExtension) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

// The honorific title of an individual.
// **A `response` object.**
type PersonTitle string

const (
	// For men, regardless of marital status.
	PersonTitleMr PersonTitle = "MR"
	// For married women.
	PersonTitleMrs PersonTitle = "MRS"
	// For girls, unmarried women, and married women who continue to use their maiden name.
	PersonTitleMiss PersonTitle = "MISS"
	// For women, regardless of marital status, or when marital status is unknown.
	PersonTitleMs PersonTitle = "MS"
	// For the holder of a doctoral degree.
	PersonTitleDr PersonTitle = "DR"
)

var AllPersonTitle = []PersonTitle{
	PersonTitleMr,
	PersonTitleMrs,
	PersonTitleMiss,
	PersonTitleMs,
	PersonTitleDr,
}

func (e PersonTitle) IsValid() bool {
	switch e {
	case PersonTitleMr, PersonTitleMrs, PersonTitleMiss, PersonTitleMs, PersonTitleDr:
		return true
	}
	return false
}

func (e PersonTitle) String() string {
	return string(e)
}

func (e *PersonTitle) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = PersonTitle(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid PersonTitle", str)
	}
	return nil
}

func (e PersonTitle) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

// Defines the type of phone number.
// **A `response` object. **
type PhoneNumberLabel string

const (
	PhoneNumberLabelMain   PhoneNumberLabel = "MAIN"
	PhoneNumberLabelWork   PhoneNumberLabel = "WORK"
	PhoneNumberLabelHome   PhoneNumberLabel = "HOME"
	PhoneNumberLabelMobile PhoneNumberLabel = "MOBILE"
	PhoneNumberLabelOther  PhoneNumberLabel = "OTHER"
)

var AllPhoneNumberLabel = []PhoneNumberLabel{
	PhoneNumberLabelMain,
	PhoneNumberLabelWork,
	PhoneNumberLabelHome,
	PhoneNumberLabelMobile,
	PhoneNumberLabelOther,
}

func (e PhoneNumberLabel) IsValid() bool {
	switch e {
	case PhoneNumberLabelMain, PhoneNumberLabelWork, PhoneNumberLabelHome, PhoneNumberLabelMobile, PhoneNumberLabelOther:
		return true
	}
	return false
}

func (e PhoneNumberLabel) String() string {
	return string(e)
}

func (e *PhoneNumberLabel) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = PhoneNumberLabel(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid PhoneNumberLabel", str)
	}
	return nil
}

func (e PhoneNumberLabel) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type SortingDirection string

const (
	SortingDirectionAsc  SortingDirection = "ASC"
	SortingDirectionDesc SortingDirection = "DESC"
)

var AllSortingDirection = []SortingDirection{
	SortingDirectionAsc,
	SortingDirectionDesc,
}

func (e SortingDirection) IsValid() bool {
	switch e {
	case SortingDirectionAsc, SortingDirectionDesc:
		return true
	}
	return false
}

func (e SortingDirection) String() string {
	return string(e)
}

func (e *SortingDirection) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = SortingDirection(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid SortingDirection", str)
	}
	return nil
}

func (e SortingDirection) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
