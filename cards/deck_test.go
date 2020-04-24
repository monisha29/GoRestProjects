package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	//t needs to be told if something goes wrong
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("Expected deck length of 16 , but got %v ", len(d))
	}

	if d[0] != "Spades of Ace" {
		t.Errorf("Expected Ace of Spades , but got %v ", d[0])
	}
}

// cleanup needs to be done manually
func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_deckTesting")

	d := newDeck()
	d.saveToFile("_deckTesting")

	loadedDeck := newDeckFromFile("_deckTesting")

	if len(loadedDeck) != 16 {
		t.Errorf("Expected deck length of 16 , but got %v ", len(loadedDeck))
	}

	os.Remove("_deckTesting")

}
