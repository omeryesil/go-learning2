package main

import (
	"fmt"
)

func main() {
	cards := newDeck()

	cards.writeToFile("hello.txt")

	fmt.Println("success")
}
