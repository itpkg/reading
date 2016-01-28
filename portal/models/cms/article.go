package models

import (
	"time"

	"github.com/itpkg/reading/portal/models"
)

type Article struct {
	Id      uint
	Uid     string `orm:"unique"`
	Lang    string `orm:"index"`
	Title   string
	Summary string    `orm:"size(500)"`
	Body    string    `orm:"type(text)"`
	Created time.Time `orm:"auto_now_add;type(datetime)"`
	Updated time.Time `orm:"auto_now;type(datetime)"`

	Tags []*Tag       `orm:"rel(m2m)"`
	User *models.User `orm:"rel(fk)"`
}
