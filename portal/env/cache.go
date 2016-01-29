package env

import (
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/cache"
	_ "github.com/astaxie/beego/cache/redis"
)

var CACHE cache.Cache

func init() {
	port, err := beego.AppConfig.Int("redisPort")
	if err != nil {
		port = 6379
	}
	db, err := beego.AppConfig.Int("redisDb")
	if err != nil {
		db = 0
	}
	CACHE, err = cache.NewCache(
		"redis",
		fmt.Sprintf(
			`{"conn":"%s:%d","dbNum":"%d"}`,
			beego.AppConfig.String("redisHost"), port, db,
		))
	if err != nil {
		beego.Error(err)
	}
}
