package lib01

import (
	"testing"
)

func TestEcho(t *testing.T) {
	want := "hello"
	got := Echo("hello")

	if got != want {
		t.Errorf("Echo was incorrect, got: %q, want: %q.", got, want)
	}
}
