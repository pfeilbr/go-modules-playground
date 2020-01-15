package lib02

import (
	"testing"
)

func TestQuote(t *testing.T) {
	want := "\"hello\""
	got := Quote("hello")

	if got != want {
		t.Errorf("Quote was incorrect, got: %q, want: %q.", got, want)
	}
}
