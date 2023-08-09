package rawentity

import (
	"time"
)

type RawContact struct {
	RawId     string    `gorm:"column:raw_id;default:gen_random_uuid();primaryKey"`
	Data      string    `gorm:"type:jsonb;gorm:not null"`
	EmittedAt time.Time `gorm:"type:timestamp with time zone;default:now()"`
}

func (RawContact) TableName() string {
	return "_openline_raw_contacts"
}
