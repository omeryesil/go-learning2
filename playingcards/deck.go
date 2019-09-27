package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

//Create a new type of deck
type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Hearts", "Diamonds", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Joker", "Queen", "King"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

func (d deck) print() { //by convention, best practices, mostly first or the first 2 letters are used (d for deck)
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(d deck, handsize int) (deck, deck) {
	return d[:handsize], d[handsize:]
}

func (d deck) toString() string {
	return strings.Join(([]string(d)), "|")
}

func (d deck) writeToFile(fileName string) bool {
	dataInBytes := []byte(d.toString())

	err := ioutil.WriteFile(fileName, dataInBytes, 0644)

	if err != nil {
		log.Fatal(err)
		return false
	}

	return true
}
