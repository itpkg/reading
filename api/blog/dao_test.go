package blog_test

import (
	"fmt"
	"testing"

	"github.com/itpkg/reading/api/blog"
	"github.com/itpkg/reading/api/config"
)

func TestDao(t *testing.T) {
	cfg := config.Model{
		ElasticSearch: &config.ElasticSearch{
			Host:  "localhost",
			Port:  9200,
			Index: "test",
		},
	}
	c, e := cfg.OpenElasic()
	if e != nil {
		t.Logf("error on open: %v", e)
	}
	dao := blog.Dao{Client: c, Cfg: &cfg}
	ty := "text/plain"

	title := "ttt"
	if e := dao.Set(ty, title, "测试啊"); e != nil {
		t.Errorf("error on set: %v", e)
		return
	}
	b, e := dao.Get(title)
	if e != nil {
		t.Errorf("error on get: %v", e)
		return
	}
	t.Logf("GET: %+v", b)
	if b.Title != title {
		t.Errorf("wang %s, get %s", title, b.Title)
		return
	}

	if err := dao.Del(title); err != nil {
		t.Errorf("bad in del %v", err)
		return
	}

	for i := 1; i < 10; i++ {
		dao.Set(ty, fmt.Sprintf("title / %d", i), fmt.Sprintf("body %d", i))
	}
	ids, err := dao.List(ty, 1, 5)
	if err != nil {
		t.Errorf("bad in list: %v", err)

	}
	if len(ids) == 0 {
		t.Errorf("ids is empty")
	} else {
		t.Logf("GET IDS(%d): %v", len(ids), ids)
	}

}
