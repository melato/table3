package table

import (
	"testing"
)

func TestWidth(t *testing.T) {
	if width("a") != 1 {
		t.Fail()
	}
	w := width("λ")
	if w != 1 {
		t.Errorf("expected width = 1")
	}
	if len("λ") == 1 {
		t.Errorf("expected len != 1")
	}
}
