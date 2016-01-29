package env

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/itpkg/reading/portal/models"
	"github.com/itpkg/reading/portal/utils"
)

type Dao struct {
	aes *utils.Aes
}

func (p *Dao) Get(o orm.Ormer, k string, v interface{}) error {
	s := models.Setting{Key: k}
	err := o.Read(&s)
	if err != nil {
		return err
	}
	buf := []byte(s.Val)
	if s.Flag {
		if buf, err = p.aes.Decrypt(buf); err != nil {
			return err
		}
	}
	return utils.FromJson(buf, v)
}

func (p *Dao) Set(o orm.Ormer, k string, v interface{}, f bool) error {
	buf, err := utils.ToJson(v)
	if err != nil {
		return err
	}
	if f {
		buf, err = p.aes.Encrypt(buf)
		if err != nil {
			return err
		}
	}
	s := models.Setting{Key: k}
	err = o.Read(&s)
	switch err {
	case orm.ErrNoRows:
		s.Val = string(buf)
		s.Flag = f
		_, err = o.Insert(&s)
	case nil:
		s.Val = string(buf)
		s.Flag = f
		_, err = o.Update(&s)
	}
	return err
}

//=============================================================================
var DAO *Dao

func init() {
	cip, err := utils.NewAesCipher([]byte(beego.AppConfig.String("key")[50:82]))
	if err != nil {
		beego.Error(err)
	}
	DAO = &Dao{aes: &utils.Aes{Cip: cip}}
}
