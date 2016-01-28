package models

import (
	"time"
)

type Tag struct {
	Id      int
	Name    string    `orm:"unique"`
	Created time.Time `orm:"auto_now_add;type(datetime)"`
	Updated time.Time `orm:"auto_now;type(datetime)"`

	Articles []*Article `orm:"reverse(many)"`
}
