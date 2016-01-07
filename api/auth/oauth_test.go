package auth_test

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/itpkg/reading/api/auth"
)

func TestGoogle(t *testing.T) {
	cf, err := os.Open("../config/google.oauth.json")
	if err != nil {
		t.Errorf("bad in open google file: %v", err)
		return
	}

	de := json.NewDecoder(cf)
	cfg := &auth.GoogleConf{}
	if err := de.Decode(&cfg); err != nil {
		t.Errorf("bad in parse config file: %v", err)
	}
	g := auth.NewGoogle(cfg)
	t.Logf("Google oauth url: %s", g.Url())
}
