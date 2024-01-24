// main.go

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	cards := newDeck()
	cards = cards.shuffle()

	player1Hand := make(deck, 0)
	player2Hand := make(deck, 0)

	for i := 0; i < 2; i++ {
		card, remainingCards := dealOneCard(&cards)
		player1Hand = append(player1Hand, card)
		cards = remainingCards

		card, remainingCards = dealOneCard(&cards)
		player2Hand = append(player2Hand, card)
		cards = remainingCards
	}

	player1Score := calculateScore(player1Hand)
	player2Score := calculateScore(player2Hand)

	fmt.Println("Player 1's hand:")
	player1Hand.print()
	fmt.Println("Player 1's score:", player1Score)

	fmt.Println("\nPlayer 2's hand:")
	player2Hand.print()
	fmt.Println("Player 2's score:", player2Score)
}
