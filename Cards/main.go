package main

import "fmt"

func main() {
	// selectHand()
	// saveToFileExample()
	// readFromFileExample()
	shuffleExample()
}

func selectHand() {
	cards := newDeck()
	hand, remainingCards := deal(cards, 5)
	fmt.Println("Hand")
	hand.print()
	fmt.Println("Remaining")
	remainingCards.print()
}

func saveToFileExample() {
	cards := newDeck()
	cards.saveToFile("my_file")
}

func readFromFileExample() {
	cards := newDeckFromFile("my_file")
	cards.print()
}

func shuffleExample() {
	cards := newDeck()
	cards.shuffle()
	cards.print()
}
