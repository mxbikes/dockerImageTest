package main

import "testing"

func TestGreet(t *testing.T) {
	got := Greet("Wonka")
	want := "Hello, Wonka"

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}
