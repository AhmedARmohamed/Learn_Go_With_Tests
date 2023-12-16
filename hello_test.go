package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("ahmed")
	want := "Hello, ahmed"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
