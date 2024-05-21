package main

import "testing"

func TestHelloWorld(t *testing.T) {
	got := getMessage("Paul")
	want := "Hello Paul"

	if got != want {
		t.Fail()
	}
}

func TestHelloWorldFail(t *testing.T) {
	got := getMessage("Paul")
	want := "Hello James"
	if got != want {
		t.Fail()
	}
}
