package auth

import (
	"fmt"
	"strings"
	"time"

	"github.com/itpkg/reading/api/core"
)

type User struct {
	core.Model

	Email        string `sql:"not null;index:idx_users_email"`
	Uid          string `sql:"not null;unique;type:char(36)"`
	Home         string `sql:"not null"`
	Logo         string `sql:"not null"`
	Name         string `sql:"not null"`
	ProviderType string `sql:"not null;default:'unknown';index:idx_users_provider_type"`
	ProviderId   string `sql:"not null;index:idx_users_provider_id"`

	LastSignIn  time.Time `sql:"not null"`
	SignInCount uint      `sql:"not null;default:0"`

	Permissions []Permission
}

func (p *User) SetGravatar() {
	p.Logo = fmt.Sprintf("https://gravatar.com/avatar/%s.png", core.Md5([]byte(strings.ToLower(p.Email))))
}

func (p User) String() string {
	return fmt.Sprintf("%s<%s>", p.Name, p.Email)
}

type Log struct {
	ID        uint      `gorm:"primary_key" json:"id"`
	UserID    uint      `sql:"not null" json:"-"`
	User      User      `json:"-"`
	Message   string    `sql:"not null" json:"message"`
	CreatedAt time.Time `sql:"not null;default:current_timestamp" json:"created_at"`
}

type Role struct {
	ID           uint   `gorm:"primary_key"`
	Name         string `sql:"not null;index:idx_roles_name"`
	ResourceType string `sql:"not null;default:'-';index:idx_roles_resource_type"`
	ResourceId   uint   `sql:"not null;default:0"`
}

func (p Role) String() string {
	return fmt.Sprintf("%s@%s://%d", p.Name, p.ResourceType, p.ResourceId)
}

type Permission struct {
	ID     uint `gorm:"primary_key"`
	User   User
	UserID uint `sql:"not null"`
	Role   Role
	RoleID uint      `sql:"not null"`
	Begin  time.Time `sql:"not null;default:current_date;type:date"`
	End    time.Time `sql:"not null;default:'1000-1-1';type:date"`
}

func (p *Permission) EndS() string {
	return p.End.Format("2006-01-02")
}
func (p *Permission) BeginS() string {
	return p.Begin.Format("2006-01-02")
}

func (p *Permission) Enable() bool {
	now := time.Now()
	return now.After(p.Begin) && now.Before(p.End)
}
