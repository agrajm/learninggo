package main

import "testing"

import "os"

func TestNewDeck(t *testing.T) {

	d := newDeck()

	if len(d) != 52 {
		t.Errorf("Expected length of 52, got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first element Ace of Spades, got %v", d[0])
	}

	if d[len(d)-1] != "King of Clubs" {
		t.Errorf("Expected last element as King of Clubs, got %v", d[51])
	}

}

func TestSaveToDiskAndReadFromFile(t *testing.T) {

	os.Remove("_decktest")

	d := newDeck()
	d.saveToFile("_decktest")

	deck := readFromFile("_decktest")
	if len(deck) != 52 {
		t.Errorf("Failed to create deck properly, got %v cards", len(deck))
	}

	os.Remove("_decktest")

}
