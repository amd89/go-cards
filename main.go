package main

func main() {
	// var card string
	// var card string = "Ace of Spades"  // long form veriable declaration  | use for global
	// card := newCard() //:= is only used for initiation regular = used otherwise  |  doesnt work for global

	//cards := []string{newCard(), "Ace of Diamonds"} //array declaration
	cards := newDeck()
	// cards = append(cards, "Six of Spades") //append returns a new list

	// hand, remainingDeck := deal(cards, 5)

	// hand.print()
	// remainingDeck.print()

	//cards.saveToFile("test.txt")
	cards.shuffle()
	cards.print()

}
