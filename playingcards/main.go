package main

import "fmt"

const fileName = "deck.txt"

func main() {
	fmt.Println("Creating a new deck -------------------------")
	cards := newDeck()

	fmt.Println("Shuffling the deck -------------------------")
	cards = cards.shuffle()

	fmt.Println("Writing the shuffled deck to a file ---------------------")
	cards.writeToFile(fileName)

	fmt.Println("Reading the deck from a file -------------------------")
	cards = deckFromFile(fileName)
	cards.print()
}
