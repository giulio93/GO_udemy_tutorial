package main

import "fmt"

func main() {
	cards := newDeck()
	// hand, erDeck := deal(cards, 5)
	// hand.print()
	// erDeck.print()
	// cards.saveToFile("my_cards")
	// newDeckFromFile("my_cards").print()
	cards.print()
	fmt.Println("==SHUFFLE===")
	cards.shuffle()
	cards.print()

}
