package memory

import (
	"testing"

	"github.com/ironzhang/x-pearls/govern/testutil"
)

func TestDriver(t *testing.T) {
	d, err := Open("test", nil)
	if err != nil {
		t.Fatalf("open: %v", err)
	}
	defer d.Close()

	testutil.TestDriver(t, d)
}
