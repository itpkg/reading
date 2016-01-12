package epub_test

import (
	"testing"

	"github.com/itpkg/reading/api/epub"
)

func TestEpub(t *testing.T) {
	if err := epub.Walk("tmp", func(name string) error {
		t.Logf("Find epub file %s", name)
		book, err := epub.Open(name)
		if err == nil {
			t.Logf("Book: %+v", book)
			t.Logf("Container: %+v", book.Container)
			for _, rf := range book.Container.RootFiles {
				t.Logf(" Root file: %+v", rf)
				t.Logf(" Opf %+v", rf.Opf)
			}
		} else {
			t.Errorf("bad in open epub: %v", err)
		}
		return err
	}); err != nil {
		t.Errorf("Bad in test epub %v", err)
	}

}
