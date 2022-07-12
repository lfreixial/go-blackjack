package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

type deck []string

func newDeck(numOfDecks int) deck {
	cards := deck{}

	cardSuit := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	for _, suit := range cardSuit {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	for i := 1; i < numOfDecks; i++ {
		for _, suit := range cardSuit {
			for _, value := range cardValues {
				cards = append(cards, value+" of "+suit)
			}
		}
	}
	return cards

}

func (d deck) shuffle() {

	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d {
		newPosition := r.Intn(len(d) - 1)

		d[i], d[newPosition] = d[newPosition], d[i]
	}
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i+1, card)
	}
}

func (d deck) deal() deck {
	card := d[:1]

	d.shuffle()
	return card
	// return d[:1], d[1:2]
}

func (d deck) checkValueOfHand(hand []string, score int) int {
	handSize := len(hand)
	score = 0

	for i := 0; i < handSize; i++ {
		if strings.Contains(hand[i], "Two") {
			score += 2
		} else if strings.Contains(hand[i], "Three") {
			score += 3
		} else if strings.Contains(hand[i], "Four") {
			score += 4
		} else if strings.Contains(hand[i], "Five") {
			score += 5
		} else if strings.Contains(hand[i], "Six") {
			score += 6
		} else if strings.Contains(hand[i], "Seven") {
			score += 7
		} else if strings.Contains(hand[i], "Eight") {
			score += 8
		} else if strings.Contains(hand[i], "Nine") {
			score += 9
		} else if strings.Contains(hand[i], "Ten") {
			score += 10
		} else if strings.Contains(hand[i], "King") {
			score += 10
		} else if strings.Contains(hand[i], "Jack") {
			score += 10
		} else if strings.Contains(hand[i], "Queen") {
			score += 10
		} else if strings.Contains(hand[i], "Ace") {
			if score+11 > 21 {
				score += 1
			} else {
				score += 11
			}
		}
	}

	return score
}

// func deal(d deck, handSize int) (deck, deck) {
// 	return d[:handSize], d[handSize:]

// }
