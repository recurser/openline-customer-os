package models

import (
	"fmt"
	commonModels "github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform/domain/common/models"
	"time"
)

type EmailValidation struct {
	ValidationError string `json:"validationError"`
}

type Email struct {
	ID              string              `json:"id"`
	RawEmail        string              `json:"rawEmail"`
	Email           string              `json:"email"`
	Source          commonModels.Source `json:"source"`
	CreatedAt       time.Time           `json:"createdAt"`
	UpdatedAt       time.Time           `json:"updatedAt"`
	EmailValidation EmailValidation     `json:"emailValidation"`
}

func (p *Email) String() string {
	return fmt.Sprintf("Email{ID: %s, RawEmail: %s, Email: %s, Source: %s, CreatedAt: %s, UpdatedAt: %s}", p.ID, p.RawEmail, p.Email, p.Source, p.CreatedAt, p.UpdatedAt)
}

func NewEmail() *Email {
	return &Email{}
}
