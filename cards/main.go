package main

const handSize = 2

func main() {
	cards := newDeckFromFile("cards_file")
	cards.shuffle()
	cards.print()
}
