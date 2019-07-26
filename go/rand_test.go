package main

import "testing"

func TestRand(t *testing.T) {
	if Rand() != "hi" {
		t.Error("Expected hi")
	}
}
