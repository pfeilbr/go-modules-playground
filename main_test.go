package main

import (
	"testing"
)

func TestEchoQuote(t *testing.T) {
	want := "\"hello\""
	got := EchoQuote("hello")

	if want != got {
		t.Errorf("EchoQuote was incorrect, got: %v, want: %v.", got, want)
	}
}
