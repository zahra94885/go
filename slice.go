package main

import "fmt"

func main() {

	cardsAce := []string{}

	cards := []string{"Ace Pik!", card()}

	//Adding to slices
	cardsAce = append(cardsAce, "Ace khesht!!!")

	fmt.Println(cards)
	fmt.Println(cardsAce)

}

func card() string {
	return "Ace Dell!!"
}
