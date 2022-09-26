package models

import "time"

type Base struct {
	ID        string     `json:"id" gorm:"type:uuid;default:uuid_generate_v4()"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	DeletedAt *time.Time `sql:"index" json:"deleted_at,omitempty"`
}
