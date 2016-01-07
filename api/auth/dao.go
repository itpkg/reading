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

func (p *Dao) SaveUser(pty, pid, email, name, home, logo string) (*User, error) {
	db := p.Db
	user := User{}
	var err error
	if db.Where("provider_id = ? AND provider_type = ?", pid, pty).First(&user).RecordNotFound() {
		user = User{
			Name:         name,
			Email:        email,
			Home:         home,
			Logo:         logo,
			Uid:          core.Uuid(),
			ProviderId:   pid,
			ProviderType: pty,
			LastSignIn:   time.Now(),
			SignInCount:  1,
		}
		err = p.Db.Create(&user).Error
	} else {
		err = db.Model(&user).UpdateColumns(map[string]interface{}{
			"name":          name,
			"email":         email,
			"home":          home,
			"logo":          logo,
			"last_sign_in":  time.Now(),
			"sign_in_count": user.SignInCount + 1,
		}).Error
	}
	return &user, err
}

func (p *Dao) GetRole(name string, resource_type string, resource_id uint) (*Role, error) {
	var e error
	r := Role{}
	db := p.Db
	if db.Where("name = ? AND resource_type = ? AND resource_id = ?", name, resource_type, resource_id).First(&r).RecordNotFound() {
		r = Role{
			Name:         name,
			ResourceType: resource_type,
			ResourceId:   resource_id,
		}
		e = db.Create(&r).Error

	}
	return &r, e
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
