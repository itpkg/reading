package config_test

import (
	"testing"

	"github.com/itpkg/reading/api/config"
)

func TestElacsicSearch(t *testing.T) {
	cfg := config.Model{
		ElasticSearch: &config.ElasticSearch{
			Host:  "localhost",
			Port:  9200,
			Index: "test",
		},
	}
	_, err := cfg.OpenElasic()
	if err != nil {
		t.Errorf("bad in open elacsicsearch: %v", err)
	}
}
