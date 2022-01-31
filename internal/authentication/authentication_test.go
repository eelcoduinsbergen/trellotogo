package trellotogo

import (
	"testing"
)

func TestAuthorize(t *testing.T) {
	want := false
	var result = Authorize()
	if result != want {
		t.Fatalf("Expected true")
	}
}
