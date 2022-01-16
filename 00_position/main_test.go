package main

import "testing"

func TestFindPos2_succes(t *testing.T) {
	got := findPos2("abcdefg", 'b')
	want := 1
	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func TestFindPos2_missing_char(t *testing.T) {
	got := findPos2("abcdefg", 'z')
	want := -1
	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func TestFindPos2_unicode(t *testing.T) {
	got := findPos2("kl☺fsdjpvo", 'd')
	want := 5
	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}
