package main

import (
	"fmt"
	"math/rand"
	"time"
)

type deck []string

// newDeck, bir kart destesi oluşturan ve geri döndüren yardımcı fonksiyon.
func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "King", "Joker", "Queen"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

// Kart destesini ekrana yazdıran yöntemi kurduk
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

// shuffle, kart destesini karıştıran ve geri döndüren yöntem.
func (d deck) shuffle() deck {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d {
		newPos := r.Intn(len(d) - 1)
		d[i], d[newPos] = d[newPos], d[i]
	}

	return d
}

// dealOneCard, bir kartı desteden çıkararak döndüren ve desteyi güncelleyen yöntem.
func dealOneCard(d *deck) (string, deck) {
	card := (*d)[0]
	*d = (*d)[1:]
	return card, *d
}

// calculateScore, bir elin skorunu hesaplayan fonksiyon.
func calculateScore(hand deck) int {
	score := 0
	for _, card := range hand {
		switch {
		case card == "Ace of Spades" || card == "Ace of Diamonds" || card == "Ace of Hearts" || card == "Ace of Clubs":
			score += 11
		case card == "King of Spades" || card == "King of Diamonds" || card == "King of Hearts" || card == "King of Clubs",
			card == "Queen of Spades" || card == "Queen of Diamonds" || card == "Queen of Hearts" || card == "Queen of Clubs",
			card == "Joker of Spades" || card == "Joker of Diamonds" || card == "Joker of Hearts" || card == "Joker of Clubs":
			score += 10
		default:
			score += 1
		}
	}

	return score
}
