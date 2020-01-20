package main

import "fmt"

func main() {

	cards := newDeck()

	hand, remainingCards := cards.deal(8)
	//hand.print()
	// remainingCards.print()

	fmt.Println(hand.toString())
	remainingCards.toString()

	myhand := readFromFile("handcards")
	myhand.shuffle().print()

}
