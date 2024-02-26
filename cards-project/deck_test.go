package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()
	if len(d) != 16 {
		t.Errorf("Expected deck length of 16, but got %d", len(d))
	}
	if d[0] != "Ace of Spaces" {
		t.Errorf("Expected first element of deck is Ace of Spaces, but got %s", d[0])
	}
	if d[len(d)-1] != "Four of Clubs" {
		t.Errorf("Expected last element of deck is Four of Clubs, but got %s", d[len(d)-1])
	}
}

func TestSaveToDeckAndNewDeckTestFromFile(t *testing.T) {

	os.Remove("_decktesting")

	deck := newDeck()
	deck.saveToFile("_decktesting")

	listDesk := deck.newDeckFromFile("_decktesting")
	if len(listDesk) != 16 {
		t.Errorf("Expected deck length of 16, but got %d", len(listDesk))
	}

	os.Remove("_decktesting")
}
