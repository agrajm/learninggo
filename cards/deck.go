package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

// Create a new type of deck
// which is a slice of strings
type deck []string

func newDeck() deck {
	cardSuits := []string{"Spades", "Heart", "Diamonds", "Clubs"}
	cardNumbers := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	cards := deck{}

	for _, suite := range cardSuits {
		for _, number := range cardNumbers {
			cards = append(cards, number+" of "+suite)
		}
	}
	return cards
}

// deal and split the cards
func (d deck) deal(handsize int) (deck, deck) {
	return d[:handsize], d[handsize:]
}

// Receiver Function -- applies to deck type
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

// create deck from file
func readFromFile(fileName string) deck {
	cardbytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(-1)
	}
	cardStr := string(cardbytes)
	return deck(strings.Split(cardStr, ","))
}

// save a deck to a file
func (d deck) saveToFile(fileName string) {
	ioutil.WriteFile(fileName, []byte(d.toString()), 0644)
}

// convert to a string representation
func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) shuffle() deck {

	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)

	for i := range d {
		newPos := r.Intn(len(d) - 1)
		// swap
		d[i], d[newPos] = d[newPos], d[i]
	}
	return d
}
