package main

var deckSize int

func main() {
	// card := "Ace of Spades"
	cards := deck{"Ace of Spades", newCard()}
	cards = append(cards, "Six of Spades")
	cards.print()
}

func newCard() string {
	return "Five of Diamonds"
}
