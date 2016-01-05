package sitemap_test

import (
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/itpkg/reading/api/sitemap"
)

func urls() []*sitemap.Url {
	urls := make([]*sitemap.Url, 0)
	for i := 0; i < 10; i++ {
		urls = append(urls,
			&sitemap.Url{
				Loc:        fmt.Sprintf("http://localhost/item/%d", i),
				LastMod:    sitemap.Time(time.Now()),
				ChangeFreq: sitemap.Daily,
				Priority:   sitemap.Priority(0.3),
			},
		)
	}
	return urls
}

func TestSitemap(t *testing.T) {
	if err := sitemap.Xml(os.Stdout, urls); err != nil {
		t.Errorf("Error on write xml: %v", err)
	}
	os.Stdout.Write([]byte("\n"))
}
