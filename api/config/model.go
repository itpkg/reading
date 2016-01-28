package config

import (
	"crypto/cipher"
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/garyburd/redigo/redis"
	"github.com/itpkg/reading/api/core"
	"github.com/itpkg/reading/api/storage"
	"github.com/jinzhu/gorm"
	"gopkg.in/olivere/elastic.v3"
)

type Model struct {
	Env           string         `toml:"-"`
	SecretsS      string         `toml:"secret"`
	Secrets       []byte         `toml:"-"`
	Http          *Http          `toml:"http"`
	Storage       *Storage       `toml:"storage"`
	Database      *Database      `toml:"database"`
	Redis         *Redis         `toml:"redis"`
	ElasticSearch *ElasticSearch `toml:"elastic_search"`
}

func (p *Model) IsProduction() bool {
	return p.Env == "production"
}

func (p *Model) Home() string {
	if p.IsProduction() {
		if p.Http.Ssl {
			return fmt.Sprintf("https://www.%s", p.Http.Domain)
		} else {
			return fmt.Sprintf("http://www.%s", p.Http.Domain)
		}
	} else {
		return fmt.Sprintf("http://localhost:%d", p.Http.Port)
	}
}
func (p *Model) OpenStorage() (storage.Provider, error) {
	s := p.Storage
	switch s.Type {
	default:
		return &storage.LocalProvider{Url: s.Extra["url"], Root: s.Extra["root"]}, nil
	}
}
func (p *Model) OpenElastic() (*elastic.Client, error) {
	var err error
	var client *elastic.Client

	if p.IsProduction() {
		client, err = elastic.NewClient(elastic.SetURL(p.ElasticSearch.Url()))
	} else {
		//		file, err := os.OpenFile("tmp/elastic.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0600)
		//		if err!=nil{
		//			return nil, err
		//		}
		client, err = elastic.NewClient(
			elastic.SetURL(p.ElasticSearch.Url()),
			elastic.SetErrorLog(log.New(os.Stderr, "", log.LstdFlags)),
			elastic.SetInfoLog(log.New(os.Stdout, "", log.LstdFlags)),
			elastic.SetTraceLog(log.New(os.Stdout, "", log.LstdFlags)),
		)
	}

	/*
		if err != nil {
			return nil, err
		}

		if _, _, err := client.Ping().Do();err!=nil{
			return nil, err
		}
	*/
	/*
			exists, err := client.IndexExists(p.ElasticSearch.Index).Do()
			if err != nil {
				return nil, err
			}
			if !exists {
				if _, err := client.CreateIndex(p.ElasticSearch.Index).BodyString(
					`
		{
			"settings":{
				"number_of_shards":1,
				"number_of_replicas":0
			},
			"mappings":{}
		}
					`,
				).Do(); err != nil {
					return nil, err
				}
			}
	*/
	return client, err
}

func (p *Model) OpenRedis() *redis.Pool {
	return &redis.Pool{
		MaxIdle:     3,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, e := redis.Dial("tcp", fmt.Sprintf("%s:%d", p.Redis.Host, p.Redis.Port))
			if e != nil {
				return nil, e
			}
			if _, e = c.Do("SELECT", p.Redis.Db); e != nil {
				c.Close()
				return nil, e
			}
			return c, nil
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
}

func (p *Model) AesCipher() (cipher.Block, error) {
	return core.NewAesCipher(p.Secrets[60:92])
}

func (p *Model) OpenDatabase() (*gorm.DB, error) {
	var db gorm.DB
	var err error
	switch p.Database.Adapter {
	case "postgres":
		if db, err = gorm.Open("postgres", fmt.Sprintf("user=%s password=%s host=%s port=%d dbname=%s sslmode=%s", p.Database.User, p.Database.Password, p.Database.Host, p.Database.Port, p.Database.Name, p.Database.Extra["sslmode"])); err != nil {
			return nil, err
		}

	default:
		return nil, errors.New(fmt.Sprintf("Unsupported adapter %s", p.Database.Adapter))
	}
	if err := db.DB().Ping(); err == nil {
		db.DB().SetMaxIdleConns(10)
		db.DB().SetMaxOpenConns(100)
		db.LogMode(!p.IsProduction())
		return &db, nil
	} else {
		return nil, err
	}
}
