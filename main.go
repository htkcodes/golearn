package main

import "fmt"

func main() {

	card := kek()
	fmt.Println(card)

	cards := []string{kek(), kek(), "kekito"}
	cards = append(cards, "lmao")
	fmt.Println(cards)

}

func kek() string {
	return "keklol"
}
