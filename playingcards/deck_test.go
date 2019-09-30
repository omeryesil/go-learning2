package main

import "testing"

func TestNewDeck(t *testing.T) {
	d := newDeck()

	expected := 52
	actual := len(d)

	if actual != 52 {
		t.Errorf("Expected: %d. Actual: %d", expected, actual)
	}

}
