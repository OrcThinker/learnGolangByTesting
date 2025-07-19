package main

import (
	"testing"
)

func TestHello(t *testing.T) {
	got := hello("Chriss")
	want := "Hello, Chriss"

	if got != want {
		t.Errorf("Got %q want %q",got,want)
		t.Fail()
	}
}
