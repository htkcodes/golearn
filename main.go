package main

func main() {

	cards := newDeck()

	cards.shuffle()

	//fmt.Print(cards.toString())

	//	cards.saveToFile("lol.txt")

	cards.print()

}
