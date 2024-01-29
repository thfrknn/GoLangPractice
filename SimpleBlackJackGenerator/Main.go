package main

import (
	"fmt"
	"time"
)

func main() {
	// Kart destesi oluşturuluyor ve karıştırılıyor
	cards := newDeck()
	cards = cards.shuffle()

	// Oyuncuların el kartları
	player1Hand := make(deck, 0)
	player2Hand := make(deck, 0)

	// Player 1'e 2 kart ver ve göster
	for i := 0; i < 2; i++ {
		card, remainingCards := dealOneCard(&cards)
		player1Hand = append(player1Hand, card)
		cards = remainingCards
	}
	fmt.Println("Player 1'in eli:")
	player1Hand.print()

	// Player 1'in sırası
	player1TurnCh := make(chan bool)
	go playerTurn(&player1Hand, &cards, player1TurnCh)
	player1Turn := <-player1TurnCh

	// Eğer 15 saniye içinde bir hamle yapılmazsa, otomatik olarak pas geç
	if !player1Turn {
		fmt.Println("Player 1 süre doldu. Pas geçildi.")
	} else {
		// Player 2'ye 2 kart ver ve göster
		for i := 0; i < 2; i++ {
			card, remainingCards := dealOneCard(&cards)
			player2Hand = append(player2Hand, card)
			cards = remainingCards
		}
		fmt.Println("\nPlayer 2'nin eli:")
		player2Hand.print()

		// Player 2'nin sırası
		player2TurnCh := make(chan bool)
		go playerTurn(&player2Hand, &cards, player2TurnCh)

		// Player 2'nin süresini kontrol et
		select {
		case player2Turn := <-player2TurnCh:
			// Eğer 15 saniye içinde bir hamle yapılmazsa, otomatik olarak pas geç
			if !player2Turn {
				fmt.Println("Player 2 süre doldu. Pas geçildi.")
			} else {
				// Oyun sonuçlarını göster
				showResults(player1Hand, player2Hand)
			}
		case <-time.After(15 * time.Second):
			fmt.Println("Player 2 süre doldu. Pas geçildi.")
		}
	}
}

// Oyuncunun sırasında kart çekme ve pas geçme işlemleri
func playerTurn(hand *deck, remainingCards *deck, turnCh chan bool) {
	var choice int
	timer := time.NewTimer(15 * time.Second)

	fmt.Println("1. Kart Çek")
	fmt.Println("2. Pas Geç")

	select {
	case <-timer.C:
		fmt.Println("Süre doldu. Pas geçiliyor.")
		turnCh <- false
		return
	default:
		fmt.Print("Seçiminiz: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			card, newRemainingCards := dealOneCard(remainingCards)
			*hand = append(*hand, card)
			*remainingCards = newRemainingCards
			fmt.Println("Yeni kart:", card)
			hand.print()

			// Skoru kontrol et, eğer 21'i geçtiyse As'ın değerini 1 olarak al
			if calculateScore(*hand) > 21 {
				for i, card := range *hand {
					if card == "Ace of Spades" || card == "Ace of Diamonds" || card == "Ace of Hearts" || card == "Ace of Clubs" {
						(*hand)[i] = "One of Ace"
						break
					}
				}
			}

			turnCh <- true
		case 2:
			fmt.Println("Pas geçildi.")
			turnCh <- false
		default:
			fmt.Println("Geçersiz seçim. Pas geçiliyor.")
			turnCh <- false
		}
	}
}

// Sonuçları gösterme
func showResults(player1Hand, player2Hand deck) {
	player1Score := calculateScore(player1Hand)
	player2Score := calculateScore(player2Hand)

	fmt.Println("\nPlayer 1'in eli:")
	player1Hand.print()
	fmt.Println("Player 1'in skoru:", player1Score)

	fmt.Println("\nPlayer 2'nin eli:")
	player2Hand.print()
	fmt.Println("Player 2'nin skoru:", player2Score)

	switch {
	case player1Score > 21 && player2Score > 21:
		fmt.Println("\nBerabere! Her iki oyuncu da kaybetti.")
	case player1Score > 21:
		fmt.Println("\nPlayer 2 Kazandı! Player 1 Bust oldu.")
	case player2Score > 21:
		fmt.Println("\nPlayer 1 Kazandı! Player 2 Bust oldu.")
	case player1Score > player2Score:
		fmt.Println("\nPlayer 1 Kazandı!")
	case player1Score < player2Score:
		fmt.Println("\nPlayer 2 Kazandı!")
	default:
		fmt.Println("\nBerabere!")
	}
}
