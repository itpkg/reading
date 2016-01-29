package utils_test

import (
	"testing"

	"github.com/itpkg/reading/api/core"
)

var hello = "hello, it-package!"
var key, _ = core.RandomBytes(32)

func TestHmac(t *testing.T) {

	hm := core.Hmac{
		Fn:  core.NewHmacHash(),
		Key: key,
	}

	dest1 := hm.Sum([]byte(hello))
	dest2 := hm.Sum([]byte(hello))

	t.Logf("HMAC1(%d): %x", len(dest1), dest1)
	t.Logf("HMAC2(%d): %x", len(dest2), dest2)
	if !hm.Equal(dest1, dest2) {
		t.Errorf("HMAC FAILED!")
	}

}
