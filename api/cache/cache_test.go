package cache_test

import (
	"testing"

	"github.com/itpkg/reading/api/cache"
	"github.com/itpkg/reading/api/config"
)

type S struct {
	Val int
}

const key = "kkk"

func TestRedis(t *testing.T) {
	rc := config.Redis{Host: "localhost", Port: 6379}
	var cp cache.Provider
	cp = &cache.RedisProvider{Redis: rc.Open()}

	s := S{Val: 111}
	if err := cp.Set(key, &s, 60); err != nil {
		t.Errorf("Bad in set: %v", err)
	}
	var s1 S
	if err := cp.Get(key, &s1); err != nil {
		t.Errorf("Bad in get: %v", err)
	}
	if s.Val != s1.Val {
		t.Errorf("Wang %d, get %d", s.Val, s1.Val)
	}
}
