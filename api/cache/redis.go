package cache

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"net/http"

	"github.com/garyburd/redigo/redis"
)

type RedisProvider struct {
	Redis *redis.Pool `inject:""`
}

func (p *RedisProvider) Page(wrt http.ResponseWriter, req *http.Request, contentType string, minutes uint, callback func() ([]byte, error)) {
	var body []byte
	key := fmt.Sprintf("%s/%s", req.URL.Path, req.URL.Query().Get("locale"))
	if err := p.Get(key, &body); err != nil {
		if body, err = callback(); err != nil {
			wrt.WriteHeader(http.StatusInternalServerError)
			wrt.Write([]byte(err.Error()))
			return
		}
		p.Set(key, body, minutes)

		wrt.Header().Set("Content-Type", contentType)
		wrt.WriteHeader(http.StatusOK)
	} else {
		wrt.Header().Set("Content-Type", contentType)
		wrt.WriteHeader(http.StatusNotModified)
	}
	wrt.Write(body)
}

func (p *RedisProvider) key(k string) string {
	return fmt.Sprintf("cache://%s", k)
}

func (p *RedisProvider) GetOrSet(key string, val interface{}, cb func(interface{}) (uint, error)) error {
	if err := p.Get(key, val); err == nil {
		return nil
	}
	t, e := cb(val)
	if e == nil {
		return p.Set(key, val, t)
	} else {
		return e
	}
}

func (p *RedisProvider) Set(key string, val interface{}, minutes uint) error {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	if err := enc.Encode(val); err != nil {
		return err
	}
	rc := p.Redis.Get()
	defer rc.Close()
	_, err := rc.Do("SET", p.key(key), buf.Bytes(), "EX", minutes*60)
	return err
}

func (p *RedisProvider) Get(key string, val interface{}) error {
	rc := p.Redis.Get()
	defer rc.Close()
	bys, err := redis.Bytes(rc.Do("GET", p.key(key)))
	if err != nil {
		return err
	}
	var buf bytes.Buffer
	dec := gob.NewDecoder(&buf)
	buf.Write(bys)
	return dec.Decode(val)
}
