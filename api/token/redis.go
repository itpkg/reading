package token

import (
	"crypto/rand"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/garyburd/redigo/redis"
	"github.com/pborman/uuid"
)

type RedisProvider struct {
	Redis *redis.Pool `inject:""`
}

func (p *RedisProvider) ParseFromRequest(req *http.Request) (map[string]interface{}, error) {
	return p.parse(func(fn jwt.Keyfunc) (*jwt.Token, error) {
		return jwt.ParseFromRequest(req, fn)
	})
}

func (p *RedisProvider) Parse(str string) (map[string]interface{}, error) {
	return p.parse(func(fn jwt.Keyfunc) (*jwt.Token, error) {
		return jwt.Parse(str, fn)
	})
}

func (p *RedisProvider) parse(fn parseFunc) (map[string]interface{}, error) {
	token, err := fn(func(token *jwt.Token) (interface{}, error) {
		rc := p.Redis.Get()
		defer rc.Close()
		return redis.Bytes(rc.Do("GET", p.key(token.Header["kid"].(string))))
	})
	if err == nil {
		if token.Valid {
			//delete(token.Claims, "exp")
			return token.Claims, nil
		} else {
			return nil, errors.New("token is not valid.")
		}
	} else {
		return nil, err
	}
}

func (p *RedisProvider) New(data map[string]interface{}, minutes uint) (string, error) {
	token := jwt.New(jwt.SigningMethodHS512)
	for k, v := range data {
		token.Claims[k] = v
	}
	token.Claims["exp"] = time.Now().Add(time.Minute * time.Duration(minutes)).Unix()
	kid := uuid.New()
	key := make([]byte, 32)
	if _, err := rand.Read(key); err != nil {
		return "", err
	}
	rc := p.Redis.Get()
	defer rc.Close()
	if _, err := rc.Do("SET", p.key(kid), key, "EX", minutes*60); err != nil {
		return "", err
	}
	token.Header["kid"] = kid
	return token.SignedString(key)
}

func (p *RedisProvider) key(kid string) string {
	return fmt.Sprintf("token://%s", kid)
}
