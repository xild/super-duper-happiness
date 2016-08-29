package main

import "testing"

const (
	ToReverse         = "Âlica sends$ a message to Bob"
	ToReverseExpected = "boB ot egassem a $sdnes acilÂ"
)

func TestReverse(t *testing.T) {
	reversed := Reverse(ToReverse)
	if reversed != ToReverseExpected {
		t.Errorf("Expected [%s] but was [%s] ", ToReverseExpected, reversed)
	}
}
