package main

import "testing"

const (
	ToReverse         = "Alice sends a message to Bobb"
	ToReverseExpected = "bboB ot egassem a sdnes ecilA"
)

func TestReverse(t *testing.T) {
	reversed := Reverse(ToReverse)
	if reversed != ToReverseExpected {
		t.Errorf("Expected %s but was %s ", ToReverseExpected, reversed)
	}
}
