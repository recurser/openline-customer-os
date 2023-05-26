package model

type WorkspaceInput struct {
	Name      string  `json:"name"`
	Provider  string  `json:"provider"`
	AppSource *string `json:"appSource"`
}

type EmailInput struct {
	Email     string  `json:"email"`
	Primary   bool    `json:"primary"`
	AppSource *string `json:"appSource"`
}

type PersonInput struct {
	IdentityId string  `json:"identityId"`
	Email      string  `json:"email"`
	Provider   string  `json:"provider"`
	AppSource  *string `json:"appSource"`
}

type UserInput struct {
	FirstName string      `json:"firstName"`
	LastName  string      `json:"lastName"`
	Email     EmailInput  `json:"email"`
	Person    PersonInput `json:"person"`
}

type TenantInput struct {
	Name      string  `json:"name"`
	AppSource *string `json:"appSource"`
}
