package models

import (
	"time"

	"github.com/itpkg/reading/portal/models"
)

type Comment struct {
	Id      uint
	Body    string    `orm:"type(text)"`
	Created time.Time `orm:"auto_now_add;type(datetime)"`
	Updated time.Time `orm:"auto_now;type(datetime)"`

	User    *models.User `orm:"rel(fk)"`
	Article *Article     `orm:"rel(fk)"`
}
