package main

import (
	"os"
	"testing"
)

// Text function names must be Capitalized (which makes them public)
func TestNewDeck(t *testing.T) {
	d := newDeck()

	expected := 52
	actual := len(d)

	if actual != 52 {
		t.Errorf("Expected: %v. Actual: %v", expected, actual)
	}
}

func TestWriteToFileAndReadDeckFromFile(t *testing.T) {
	const fileName = "deckfile_testing.txt"

	os.Remove(fileName)

	d := newDeck()
	d.writeToFile(fileName)

	//check if the file exists
	_, err := os.Stat(fileName)
	if os.IsNotExist(err) {
		t.Errorf("writeToFile fails")
	}

	//check if the right data is written to the file
	expectedCard := d[len(d)-1]
	d = readDeckFromFile(fileName)
	actualCard := d[len(d)-1]
	if actualCard != expectedCard {
		t.Errorf("Expected: %v. Actual: %s", actualCard, expectedCard)
	}

	//delete the test file
	os.Remove(fileName)
}
