package auth

import (
	"net/http"

	"github.com/itpkg/reading/api/core"
	"github.com/itpkg/reading/api/token"
	"github.com/jinzhu/gorm"
)

type Controller struct {
	core.Controller
}

func (p *Controller) CurrentUser(token token.Provider, db *gorm.DB, req *http.Request) (*User, error) {
	d, e := token.ParseFromRequest(req)
	if e != nil {
		return nil, e
	}
	var u User
	e = db.Where("uid = ?", d["id"]).First(&u).Error
	return &u, e
}
