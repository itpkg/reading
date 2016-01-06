package core

import (
	"time"
)

type Model struct {
	ID        uint      `gorm:"primary_key"`
	CreatedAt time.Time `sql:"not null;default:current_timestamp"`
	UpdatedAt time.Time `sql:"not null"`
}
