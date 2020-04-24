package main

import "fmt"

func main() {
	cards := newDeckFromFile("my_deck")
	//cards.print()
	pickedCards, remainingCards := deal(cards, 5)
	fmt.Println("Printing length of remainig cards ", len(remainingCards))
	fmt.Println("***************Picked Cards****")
	pickedCards.print()
	pickedCards.shuffle()
	fmt.Println("***************After Shuffling****")
	pickedCards.print()
	//remainingCards.print()
	//card := "Ace of Spades"   // := asks the compiler to figure out the datatype of card based o n the value assigned
	//card = "Five of Diamonds" // for reassignng we need not assigning :=
	//var card string = "Ace of Spades"

	//card := newCard()
	//fmt.Println(card)

	//cards := []string{"Ace of Diamonds", newCard()}
	//cards := deck{"Ace of Diamonds", newCard()}
	//cards = append(cards, "Six of Spades") // does not alter the existing slice instead creates a new slice assigned to cards

	// for i, card := range cards { // everytime redecalred i, card

	// 	fmt.Println(i, card)
	// }
	//cards := newDeck()
	//cards.saveToFile("my_cards")
	//cards.print()

	//cards := newDeckFromFile("my")
	//cards.print()
	//fmt.Println(cards.toString())
	// hand, remainingCards := deal(cards, 5)

	////cards := newDeck()
	//cards.shuffle()
	//cards.print()

	// hand.print()
	// remainingCards.print()

}

// func newCard() string {
// 	return "Five of Diamonds"
// }
