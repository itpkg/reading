package site

import (
	"github.com/jinzhu/gorm"
)

type Dao struct {
	Db *gorm.DB `inject:""`
}
