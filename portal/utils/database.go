package utils

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/lib/pq"
)

func init() {
	orm.RegisterDriver("postgres", orm.DRPostgres)

	maxIdle, err := beego.AppConfig.Int("databaseMaxIdle")
	if err != nil {
		beego.Error(err)
		return
	}
	maxOpen, err := beego.AppConfig.Int("databaseMaxOpen")
	if err != nil {
		beego.Error(err)
		return
	}

	orm.RegisterDataBase(
		"default", "postgres",
		beego.AppConfig.String("databaseUrl"),
		maxIdle, maxOpen,
	)

	orm.RunCommand()
}
