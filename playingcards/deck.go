package main

import (
	"fmt"       // for formatting
	"io/ioutil" //to write/read to/from a file
	"log"       //for writing logs
	"math/rand" //to grenerate random numbers
	"os"        //used to Exit
	"strings"   //used for strings, such as Split, etc.
	"time"      //used to get time for random number generation
)

//Create a new type of deck
type deck []string

// Generates a new deck of cards
func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Hearts", "Diamonds", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Joker", "Queen", "King"}

	// _ used when we don't need the slice index
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

// Receiver function - similar to extension in .net or default method in Java
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

func (d deck) shuffle() deck {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	l := len(d) - 1
	for i := range d {
		rnd := r.Intn(l)

		d[i], d[rnd] = d[rnd], d[i]
	}

	return d
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

func readDeckFromFile(fileName string) deck {
	bs, err := ioutil.ReadFile(fileName)

	if err != nil {
		log.Fatal(err)
		os.Exit(-1)
	}

	st := string(bs)

	cardsInString := strings.Split(st, "|")

	cards := deck(cardsInString)

	return cards
}
