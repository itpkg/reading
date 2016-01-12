package core

import (
	"time"
)

type Model struct {
	ID        uint      `gorm:"primary_key" json:"id"`
	CreatedAt time.Time `sql:"not null;default:current_timestamp" json:"created_at"`
	UpdatedAt time.Time `sql:"not null" json:"updated_at"`
}
