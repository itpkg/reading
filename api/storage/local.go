package storage

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"
)

type LocalProvider struct {
	Root string
	Url  string
}

func (p *LocalProvider) Delete(url string) error {
	return os.Remove(p.realPath(url[len(p.Url)+1:]))
}

func (p *LocalProvider) Store(name string, body []byte) (string, error) {
	name = fmt.Sprintf("%s%s", time.Now().Format("2006-01-02-15-04-05-999999"), strings.ToLower(filepath.Ext(name)))
	return fmt.Sprintf("%s/%s", p.Url, name), ioutil.WriteFile(p.realPath(name), body, 0600)
}

func (p *LocalProvider) realPath(name string) string {
	return path.Join(p.Root, name)
}
