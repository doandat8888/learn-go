package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
)

type deck []string

func newDeck() deck {
	cards := deck{}

	cardsSuits := []string{"Spaces", "Diamonds", "Hearts", "Clubs"}
	cardsValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardsSuits { //_: underscores, use to define a variable but we don't want to use it
		for _, value := range cardsValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

//Save to file

func (d deck) toString() string {
	return strings.Join(d, ", ")
}

func (d deck) saveToFile(filename string) error {
	return os.WriteFile(filename, []byte(d.toString()), 0666)
}

//Create new deck from a file

func (d deck) readFromFile(filename string) ([]byte, error) {
	return os.ReadFile(filename)
}

func (d deck) newDeckFromFile(filename string) []string {
	cards := deck{}
	cardsByte, error := d.readFromFile(filename)
	if error != nil {
		fmt.Println("Error reading deck: ", error)
		os.Exit(1)
	}
	cards = strings.Split((string(cardsByte)), ", ")
	return cards
}

func (d deck) shuffleCard(cards deck) {
	rand.Shuffle(len(cards), func(i, j int) {
		cards[i], cards[j] = cards[j], cards[i]
	})
}
