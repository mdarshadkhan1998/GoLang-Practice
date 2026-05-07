package main

import (
	"fmt"
)

func main() {
	cards := newDeck()
	cards.saveToFile("my_cards")
	hand, remainingCards := deal(cards, 13)
	fmt.Println(hand, remainingCards)

}
