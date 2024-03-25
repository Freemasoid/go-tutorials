package main

import (
	"fmt"
)

func main() {

	var revenue int
	var expenses int
	var taxRate float64

	fmt.Print("Revenue: ")
	fmt.Scan(&revenue)
	fmt.Print("Expenses: ")
	fmt.Scan(&expenses)
	fmt.Print("Tax Rate: ")
	fmt.Scan(&taxRate)

	earningsBeforeTax := revenue - expenses
	earningsAfterTax := float64(earningsBeforeTax) * (1 - taxRate/100)
	ratio := float64(earningsBeforeTax) / earningsAfterTax

	fmt.Println("Earnings Before Tax: ", earningsBeforeTax)
	fmt.Println("Earnings After Tax: ", earningsAfterTax)
	fmt.Println("Profit ratio: ", ratio)
}
