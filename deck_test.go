package main

import (
	"os"
	"strings"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()
	if len(d) != 16 {
		t.Errorf("Deck lenght is %d. Expected %d", len(d), 16)
	}
	firstCard := d[0]
	if !strings.EqualFold(firstCard, "Ace of Spades") {
		t.Errorf("First card is %s. Expected Ace of Spades", firstCard)
	}
	lastCard := d[len(d)-1]
	if !strings.EqualFold(lastCard, "Four of Clubs") {
		t.Errorf("Last card is %s. Expected Four Of Clubs", lastCard)
	}
}

func TestSaveToFileAndLoadFromFile(t *testing.T) {
	os.Remove("_decktesting.csv")
	deck := newDeck()
	deck.saveToFile("_decktesting.csv")
	loadedDeck := loadFromFile("_decktesting.csv")
	if len(loadedDeck) != 16 {
		t.Errorf("Deck size is %d. Expected 16", len(loadedDeck))
	}
	os.Remove("_decktesting.csv")
}
