package main

import (
	"fmt"
)

func main() {
	cards := newDeck(6)
	cards.shuffle()
	// cards.print()
	// player := cards.deal()
	// player1 := cards.deal()
	// fmt.Println(player)
	var player []string
	var dealer []string
	player = append(player, cards.deal()...)
	player = append(player, cards.deal()...)
	dealer = append(dealer, cards.deal()...)
	dealer = append(dealer, cards.deal()...)

	fmt.Println(player)
	// fmt.Println(dealer)
	// player = append(player, "Ace of Spade")
	// player = append(player, "Ace of Diamonds")
	var playerHand int
	var dealerHand int
	playerHand = cards.checkValueOfHand(player, playerHand)
	dealerHand = cards.checkValueOfHand(dealer, dealerHand)
	fmt.Println(playerHand)
	//fmt.Println(dealerHand)

	player = append(player, cards.deal()...)

	fmt.Println(player)
	playerHand = cards.checkValueOfHand(player, playerHand)
	fmt.Println(playerHand)
}
