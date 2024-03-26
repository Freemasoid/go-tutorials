package main

import (
	"fmt"
)

func main() {

	revenue := getUserInput("Revenue: ")
	expenses := getUserInput("Expenses: ")
	taxRate := getUserInput("Tax Rate: ")

	earningsBeforeTax, earningsAfterTax, ratio := calculator(revenue, expenses, taxRate)

	fmt.Printf("Earnings Before Tax: %.1f\n", earningsBeforeTax)
	fmt.Printf("Earnings After Tax: %.1f\n", earningsAfterTax)
	fmt.Printf("Profit ratio: %.3f\n", ratio)
}

func getUserInput(label string) float64 {
	var userInput float64
	fmt.Print(label)
	fmt.Scan(&userInput)
	return userInput
}

func calculator(revenue, expenses, taxRate float64) (float64, float64, float64) {
	earningsBeforeTax := revenue - expenses
	earningsAfterTax := float64(earningsBeforeTax) * (1 - taxRate/100)
	ratio := float64(earningsBeforeTax) / earningsAfterTax

	return earningsBeforeTax, earningsAfterTax, ratio
}
