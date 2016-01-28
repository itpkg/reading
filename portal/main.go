package main

import (
	"github.com/astaxie/beego"
	_ "github.com/astaxie/beego/cache/redis"
	"github.com/astaxie/beego/orm"
	_ "github.com/itpkg/reading/portal/models"
	_ "github.com/itpkg/reading/portal/models/cms"
	_ "github.com/itpkg/reading/portal/routers"
	_ "github.com/lib/pq"
)

func main() {
	beego.Run()
}

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
