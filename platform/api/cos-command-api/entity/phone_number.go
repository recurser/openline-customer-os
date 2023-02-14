package entity

import (
	"fmt"
	"time"
)

type PhoneNumberEntity struct {
	Id            string
	E164          string
	Label         string
	Primary       bool
	Source        DataSource
	SourceOfTruth DataSource
	CreatedAt     time.Time
}

func (phone PhoneNumberEntity) ToString() string {
	return fmt.Sprintf("id: %s\ne164: %s\nlabel: %s", phone.Id, phone.E164, phone.Label)
}

type PhoneNumberEntities []PhoneNumberEntity

func (phone PhoneNumberEntity) Labels() []string {
	return []string{"PhoneNumber"}
}