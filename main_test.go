package main

import "testing"

func TestHello(t *testing.T) {
	if hello() != myStringHello {
		t.Fatal("Test fail!")
	}
}
