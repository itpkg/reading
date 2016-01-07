package assets_test

import (
	"fmt"
	"testing"

	"github.com/itpkg/reading/api/assets"
	"github.com/itpkg/reading/api/config"
	"github.com/itpkg/reading/api/core"
)

func TestDao(t *testing.T) {
	cfg := config.Model{
		ElasticSearch: &config.ElasticSearch{
			Host:  "localhost",
			Port:  9200,
			Index: "test",
		},
	}
	c, e := cfg.OpenElastic()
	if e != nil {
		t.Logf("error on open: %v", e)
	}
	dao := assets.Dao{Client: c, Cfg: &cfg}
	ty := "text/plain"

	title := "ttt"
	if e := dao.Set(ty, title, "测试啊"); e != nil {
		t.Errorf("error on set: %v", e)
		return
	}

	id := core.Md5([]byte(title))

	b, e := dao.Get(id)
	if e != nil {
		t.Errorf("error on get: %v", e)
		return
	}
	t.Logf("GET: %+v", b)
	if b.Title != title {
		t.Errorf("wang %s, get %s", title, b.Title)
		return
	}

	if err := dao.Del(id); err != nil {
		t.Errorf("bad in del %v", err)
		return
	}

	for i := 1; i < 10; i++ {
		dao.Set(ty, fmt.Sprintf("title / %d", i), fmt.Sprintf("body %d", i))
	}
	total, ids, err := dao.List(ty+"111", 0, 5)
	if err != nil {
		t.Errorf("bad in list: %v", err)

	}
	if len(ids) == 0 {
		t.Errorf("ids is empty")
	} else {
		t.Logf("GET IDS(%d of %d): %v", len(ids), total, ids)
	}

}
