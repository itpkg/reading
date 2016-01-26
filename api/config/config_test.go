package config_test

/*

/etc/nginx/sites-enabled/development.conf:
server {
  listen 80;

  server_name localhost;
  root /var/www;
}

*/

import (
	"testing"

	"github.com/itpkg/reading/api/config"
	"github.com/itpkg/reading/api/core"
)

func TestHttp(t *testing.T) {
	models := map[string]interface{}{
		"development": &config.Model{
			SecretsS: random(),
			Redis: &config.Redis{
				Host: "localhost",
				Port: 6379,
				Db:   1,
			},
			Database: &config.Database{
				Adapter: "postgres",
				Host:    "localhost",
				Port:    5432,
				Name:    "itpkg_reading_dev",
				User:    "postgres",
				Extra:   map[string]string{"sslmode": "disable"},
			},
			ElasticSearch: &config.ElasticSearch{
				Host:  "localhost",
				Port:  9200,
				Index: "itpkg_reading_dev",
			},
			Http: &config.Http{
				Domain: "localhost",
				Port:   3000,
			},
			Storage: &config.Storage{
				Type: "local",
				Extra: map[string]string{
					"url":  "http://localhost/attachments",
					"root": "/var/www/attachments",
				},
			},
		},
		"production": &config.Model{
			SecretsS: random(),
			Redis: &config.Redis{
				Host: "localhost",
				Port: 6379,
				Db:   2,
			},
			Database: &config.Database{
				Adapter: "postgres",
				Host:    "localhost",
				Port:    5432,
				Name:    "itpkg_reading_prod",
				User:    "reading",
				Extra:   map[string]string{"sslmode": "disable"},
			},
			ElasticSearch: &config.ElasticSearch{
				Host:  "localhost",
				Port:  9200,
				Index: "itpkg_reading_prod",
			},
			Http: &config.Http{
				Domain: "change-me",
				Port:   3000,
			},
			Storage: &config.Storage{
				Type: "local",
				Extra: map[string]string{
					"url":  "http://localhost/attachments",
					"root": "/var/www/attachments",
				},
			},
		},
		"test": &config.Model{
			SecretsS: random(),
			Redis: &config.Redis{
				Host: "localhost",
				Port: 6379,
				Db:   3,
			},
			Database: &config.Database{
				Adapter: "postgres",
				Host:    "localhost",
				Port:    5432,
				Name:    "itpkg_reading_test",
				User:    "postgres",
				Extra:   map[string]string{"sslmode": "disable"},
			},
			ElasticSearch: &config.ElasticSearch{
				Host:  "localhost",
				Port:  9200,
				Index: "itpkg_reading_test",
			},
			Http: &config.Http{
				Domain: "localhost",
				Port:   3000,
			},
			Storage: &config.Storage{
				Type: "local",
				Extra: map[string]string{
					"url":  "http://localhost/attachments",
					"root": "/var/www/attachments",
				},
			},
		},
	}

	for env, mod := range models {
		if e := core.ToToml(env+".toml", mod); e != nil {
			t.Errorf("http error: %v", e)
		}
	}

}

func random() string {
	b, _ := core.RandomBytes(512)
	return core.ToBase64(b)
}
