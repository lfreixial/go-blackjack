package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

var playerScore int = 0
var dealerScore int = 0

func main() {
	var answer string
	//for {
	for {
		fmt.Printf("Your Score: %v\nDealer Score: %v\n", playerScore, dealerScore)
		fmt.Print("1: Play\n2: Quit\n->")
		fmt.Scanln(&answer)
		choice, err := strconv.ParseInt(answer, 5, 12)
		if err != nil || choice <= 0 || choice > 2 {
			for {
				fmt.Println("Error - Choice not valid. Please select one from the following")
				fmt.Print("1: Play\n2: Quit\n-> ")
				fmt.Scanln(&answer)
				choice, err = strconv.ParseInt(answer, 5, 12)
				if err == nil && choice > 0 && choice <= 2 {
					break
				}
				_ = choice
			}

		}

		if choice == 2 {
			os.Exit(0)
		} else {

			game()

		}
	}
}

func game() {
	cards := newDeck(6)
	cards.shuffle()

	var player []string
	var dealer []string
	player = append(player, cards.deal()...)
	player = append(player, cards.deal()...)
	//dealer = append(dealer, cards.deal()...)
	//dealer = append(dealer, cards.deal()...)

	var playerHand int
	var dealerHand int
	playerHand = cards.checkValueOfHand(player, playerHand)
	// dealerHand = cards.checkValueOfHand(dealer, dealerHand)

	fmt.Printf("Your hand:%v %v\n", player, playerHand)

	// loop to ask if the player wants an extra card

	var answer string
	//for {

	for {
		fmt.Print("1: Hit\n2: Stand\n->")
		fmt.Scanln(&answer)
		choice, err := strconv.ParseInt(answer, 5, 12)
		if err != nil || choice <= 0 || choice > 2 {
			for {
				fmt.Println("Error - Choice not valid. Please select one from the following")
				fmt.Print("1: Hit\n2: Stand\n-> ")
				fmt.Scanln(&answer)
				choice, err = strconv.ParseInt(answer, 5, 12)
				if err == nil && choice > 0 && choice <= 2 {
					break
				}
				_ = choice
			}

		}

		if choice == 2 {
			break
		} else {
			player = append(player, cards.deal()...)
			playerHand = cards.checkValueOfHand(player, playerHand)
			fmt.Printf("Your hand:%v %v\n", player, playerHand)
			if playerHand > 21 {
				fmt.Println("You Bust!")
				//os.Exit(0)
				dealerScore += 1
				main()

			}
		}
	}

	for {
		// dealer draw card and display - calculate hand
		dealer = append(dealer, cards.deal()...)
		dealerHand = cards.checkValueOfHand(dealer, dealerHand)
		fmt.Printf("Dealer hand:%v %v\n", dealer, dealerHand)
		//dealer draw card and display - calculate hand
		if dealerHand > 21 {
			fmt.Println("Dealer Bust! You Won!")
			playerScore += 1
			break
		}
		if dealerHand > playerHand {
			fmt.Println("Dealer Won!")
			dealerScore += 1
			break
		}
		time.Sleep(1 * time.Second)
		dealer = append(dealer, cards.deal()...)
		dealerHand = cards.checkValueOfHand(dealer, dealerHand)
		fmt.Printf("Dealer hand:%v %v\n", dealer, dealerHand)
		time.Sleep(1 * time.Second)
		if dealerHand > 21 {
			fmt.Println("Dealer Bust! You Won!")
			playerScore += 2
			break
		}
		if dealerHand > playerHand {
			fmt.Println("Dealer Won!")
			dealerScore += 1
			break
		}
		if dealerHand == 21 && playerHand == 21 {
			fmt.Println("It's a Draw!")
		}
	}
	// win or lose
}
