package storage

import (
	"fmt"
	"io"
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

func (p *LocalProvider) Store(name string, reader io.Reader) (string, uint, error) {
	buf, err := ioutil.ReadAll(reader)
	if err != nil {
		return "", 0, err
	}

	name = fmt.Sprintf("%d%s", time.Now().UnixNano(), strings.ToLower(filepath.Ext(name)))
	return fmt.Sprintf("%s/%s", p.Url, name), uint(len(buf)), ioutil.WriteFile(p.realPath(name), buf, 0644)
}

func (p *LocalProvider) realPath(name string) string {
	return path.Join(p.Root, name)
}
