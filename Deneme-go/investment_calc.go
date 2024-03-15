package main

import "fmt"

func main() {
	var balance float64
	var rate float64
	var days float64
	var i float64

	fmt.Println("Balance : ")
	fmt.Scanln(&balance)
	fmt.Println("Rate : ")
	fmt.Scanln(&rate)
	fmt.Println("Days : ")
	fmt.Scanln(&days)

	for i = 0; i < days; i++ {
		balance = (balance * rate / 100) + balance
	}

	fmt.Println("Final Balance:", balance)
}
