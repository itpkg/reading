package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

type User struct {
	Id           uint
	Uid          string `orm:"unique"`
	Name         string `orm:"index"`
	Logo         string
	Email        string `orm:"unique"`
	ProviderId   string
	ProviderType string `orm:"size(16);index"`
	SignInCount  uint
	LastSignIn   time.Time `orm:"type(datetime);null"`
	Updated      time.Time `orm:"auto_now;type(datetime)"`
	Created      time.Time `orm:"auto_now_add;type(datetime)"`

	Logs        []*Log        `orm:"reverse(many)"`
	Permissions []*Permission `orm:"reverse(many)"`
}

func (p *User) TableUnique() [][]string {
	return [][]string{
		[]string{"ProviderId", "ProviderType"},
	}
}

type Log struct {
	Id      uint
	Message string
	User    *User     `orm:"rel(fk)"`
	Created time.Time `orm:"auto_now_add;type(datetime)"`
}

type Role struct {
	Id           uint
	Name         string    `orm:"index"`
	ResourceId   uint      `orm:"null"`
	ResourceType string    `orm:"null;index"`
	Updated      time.Time `orm:"auto_now;type(datetime)"`
	Created      time.Time `orm:"auto_now_add;type(datetime)"`

	Permissions []*Permission `orm:"reverse(many)"`
}

type Permission struct {
	Id      uint
	User    *User     `orm:"rel(fk)"`
	Role    *Role     `orm:"rel(fk)"`
	Begin   time.Time `orm:"type(date)"`
	End     time.Time `orm:"type(date)"`
	Updated time.Time `orm:"auto_now;type(datetime)"`
	Created time.Time `orm:"auto_now_add;type(datetime)"`
}

func init() {
	orm.RegisterModel(new(User), new(Log), new(Role), new(Permission))
}
