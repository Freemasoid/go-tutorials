package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {

	revenue, err1 := getUserInput("Revenue: ")
	expenses, err2 := getUserInput("Expenses: ")
	taxRate, err3 := getUserInput("Tax Rate: ")

	if err1 != nil || err2 != nil || err3 != nil {
		fmt.Println(err1)
		return
	}

	earningsBeforeTax, earningsAfterTax, ratio := calculator(revenue, expenses, taxRate)

	fmt.Printf("Earnings Before Tax: %.1f\n", earningsBeforeTax)
	fmt.Printf("Earnings After Tax: %.1f\n", earningsAfterTax)
	fmt.Printf("Profit ratio: %.3f\n", ratio)
	storeResults(earningsBeforeTax, earningsAfterTax, ratio)
}

func storeResults(ebt, profit, ratio float64) {
	results := fmt.Sprintf("EBT: %.1f\nProfit: %.1f\nRatio: %.3f\n", ebt, profit, ratio)
	os.WriteFile("results.txt", []byte(results), 0644)
}

func getUserInput(label string) (float64, error) {
	var userInput float64
	fmt.Print(label)
	fmt.Scan(&userInput)

	if userInput <= 0 {
		return 0, errors.New("value must be greater than zero")
	}

	return userInput, nil
}

func calculator(revenue, expenses, taxRate float64) (float64, float64, float64) {
	earningsBeforeTax := revenue - expenses
	earningsAfterTax := float64(earningsBeforeTax) * (1 - taxRate/100)
	ratio := float64(earningsBeforeTax) / earningsAfterTax

	return earningsBeforeTax, earningsAfterTax, ratio
}
