package main

import (
	"cards-server/card"
	"fmt"
)

func main() {
	deck := card.StandardDeck()
	deck.Shuffle()
	fmt.Println(deck)
}
