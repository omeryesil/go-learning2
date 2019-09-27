package main

import "fmt"
import "ioutil"

func main() {
	cards := newDeck()

	hand, remainingDeck := deal(cards, 5)

	hand.print()
	fmt.Println("------------------------")
	remainingDeck.print()
}
