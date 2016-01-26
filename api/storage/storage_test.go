package storage_test

import (
	"testing"

	"bytes"
	"github.com/itpkg/reading/api/storage"
)

var name = "hello.txt"
var body = []byte("Hello, Storage!!!")

func TestLocal(t *testing.T) {
	var p storage.Provider
	var buf bytes.Buffer
	buf.Write(body)

	p = &storage.LocalProvider{Root: "", Url: "http://localhost/attachments"}
	if url, _, err := p.Store(name, &buf); err == nil {
		t.Logf("file=%s, url=%s", name, url)
		if err := p.Delete(url); err != nil {
			t.Errorf("Bad in delete: %v", err)
		}
	} else {
		t.Errorf("Bad in store: %v", err)
	}
}
