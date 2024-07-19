package models

import "time"

type ID struct {
	ID uint `gorm:"primary_key" json:"id,omitempty"`
}

type TimeStamps struct {
	CreatedAt time.Time `gorm:"not null" json:"created_at" sql:"DEFAULT:CURRENT_TIMESTAMP"`
	CreatedBy string    `type:text" json:"created_by"`
	UpdatedAt time.Time `gorm:"not null" json:"updated_at" sql:"DEFAULT:CURRENT_TIMESTAMP"`
	UpdatedBy string    `type:text" json:"updated_by"`
}
type SoftDeletes struct {
	DeletedAt *time.Time `gorm:"index"  json:"deleted_at,omitempty"`
}
