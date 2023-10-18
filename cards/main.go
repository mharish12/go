package main

import (
	"fmt"
)

func main() {
	cards := newDeck()
	// first, second := deal(cards, 5)
	// first, second := cards.deal(5)

	// first.print()

	// fmt.Println(" ")
	// second.print()

	// cards.print()

	fmt.Println(cards.toString())
	cards.saveToFile("cards.txt")
}
