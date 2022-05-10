package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()
	if len(d) != 16 {
		t.Errorf("Expected deck lenght of 20, but got: %d", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Ace of Spades,but got %v", d[0])
	}

	if d[len(d)-1] != "Four of Clubs" {
		t.Errorf("Four of Clubs,but got %v", d[len(d)-1])
	}
}

func TestToSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")
	deck := newDeck()
	deck.saveToFile("_decktesting")
	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 16 {
		t.Errorf("Expected deck lenght of 20, but got: %v", len(loadedDeck))
	}

	os.Remove("_decktestig")

}
