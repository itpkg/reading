package auth

import (
	"errors"
	"net/http"

	"github.com/itpkg/reading/api/token"
	"github.com/jinzhu/gorm"
)

type Session struct {
	Db    *gorm.DB       `inject:""`
	Dao   *Dao           `inject:""`
	Token token.Provider `inject:""`
}

func (p *Session) Admin(req *http.Request) (*User, error) {
	u, e := p.User(req)
	if e != nil {
		return nil, e
	}
	if p.Dao.Is(u.ID, "admin") {
		return u, nil
	} else {
		return nil, errors.New("bad role")
	}
}

func (p *Session) User(req *http.Request) (*User, error) {
	d, e := p.Token.ParseFromRequest(req)
	if e != nil {
		return nil, e
	}
	var u User
	e = p.Db.Where("uid = ?", d["id"]).First(&u).Error
	return &u, e
}
