package auth

import (
	"time"

	"github.com/itpkg/reading/api/core"
	"github.com/jinzhu/gorm"
)

type Dao struct {
	Db *gorm.DB `inject:""`
}

func (p *Dao) Log(user uint, message string) error {
	return p.Db.Create(&Log{UserID: user, Message: message}).Error
}

func (p *Dao) ConfirmUser(id uint) error {
	return p.Db.Model(&User{}).Where("id = ?", id).UpdateColumn("confirmed_at", time.Now()).Error
}

func (p *Dao) SetUserSignIn(u *User) error {
	return p.Db.Model(&User{}).Where("id = ?", u.ID).UpdateColumns(map[string]interface{}{
		"last_sign_in":  time.Now(),
		"sign_in_count": u.SignInCount + 1,
	}).Error
}

func (p *Dao) SetUserPassword(id uint, password string) error {
	passwd, err := core.Ssha512([]byte(password), 8)
	if err != nil {
		return err
	}
	return p.Db.Model(&User{}).Where("id = ?", id).UpdateColumn("password", passwd).Error
}

func (p *Dao) LockUser(id uint, flag bool) error {
	var t *time.Time
	if flag {
		n := time.Now()
		t = &n
	} else {
		t = nil
	}
	return p.Db.Model(&User{}).Where("id = ?", id).UpdateColumn("locked_at", t).Error
}

func (p *Dao) NewEmailUser(name, email, password string) (*User, error) {
	passwd, err := core.Ssha512([]byte(password), 8)
	if err != nil {
		return nil, err
	}

	u := User{
		Name:       name,
		Email:      email,
		Password:   passwd,
		Uid:        core.Uuid(),
		ProviderId: email,
	}
	u.SetGravatar()
	if err = p.Db.Create(&u).Error; err != nil {
		return nil, err
	}
	return &u, nil
}

func (p *Dao) NewRole(name string, resource_type string, resource_id uint) (*Role, error) {
	r := Role{
		Name:         name,
		ResourceType: resource_type,
		ResourceId:   resource_id,
	}
	if err := p.Db.Create(&r).Error; err != nil {
		return nil, err
	}
	return &r, nil
}

func (p *Dao) Apply(role uint, user uint, dur time.Duration) error {
	begin := time.Now()
	end := begin.Add(dur)
	var count int
	p.Db.Model(Permission{}).Where("role_id = ? AND user_id = ?", role, user).Count(&count)
	if count == 0 {
		return p.Db.Create(&Permission{
			UserID: user,
			RoleID: role,
			Begin:  begin,
			End:    end,
		}).Error
	} else {
		return p.Db.Model(&Permission{}).Where("role_id = ? AND user_id = ?", role, user).Updates(map[string]interface{}{"begin": begin, "end": end}).Error
	}
}
