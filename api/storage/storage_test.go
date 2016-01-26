package storage_test

import (
	"testing"

	"github.com/itpkg/reading/api/storage"
)

var name = "hello.txt"
var body = []byte("Hello, Storage!!!")

func TestLocal(t *testing.T) {
	var p storage.Provider
	p = &storage.LocalProvider{Root: "", Url: "http://localhost/attachments"}
	if url, err := p.Store(name, body); err == nil {
		t.Logf("file=%s, url=%s", name, url)
		if err := p.Delete(url); err != nil {
			t.Errorf("Bad in delete: %v", err)
		}
	} else {
		t.Errorf("Bad in store: %v", err)
	}
}
