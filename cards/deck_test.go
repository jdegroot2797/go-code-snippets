package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 52 { 
		t.Errorf("Expected deck length of 52, but instead got: %d", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card to be Ace of Sapdes, but instead got %v", d[0])
	}
}

func TestSaveToDeckAndNewDeckTestFromFile(t *testing.T){
	os.Remove("_decktesting")

	deck := newDeck()
	deck.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 52 {
		t.Errorf("Expected deck length of 52, but instead got: %v", len(loadedDeck))
	}

	os.Remove("_decktesting")
}