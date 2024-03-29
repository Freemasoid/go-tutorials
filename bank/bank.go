package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

const accBalanceFile = "balance.txt"

func writeBalanceToFile(value float64) {
	balance := fmt.Sprint(value)
	os.WriteFile(accBalanceFile, []byte(balance), 0644)
}

func getBalanceFromFile() (float64, error) {
	data, err := os.ReadFile(accBalanceFile)

	if err != nil {
		return 1000, errors.New("failed to find balance.txt")
	}

	balanceText := string(data)
	balance, err := strconv.ParseFloat(balanceText, 64)

	if err != nil {
		return 1000, errors.New("failed to parse stored balance value")
	}

	return balance, nil
}

func main() {
	var curBalance, err = getBalanceFromFile()
	var choice int
	var enterVal float64

	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Println("-----------------")
	}

	fmt.Println("Welcome to Go Bank!")
	fmt.Println("What do you want to do?")

	for {
		fmt.Println("")
		fmt.Println("1. Check balance")
		fmt.Println("2. Deposit money")
		fmt.Println("3. Withdraw money")
		fmt.Println("4. Exit")

		fmt.Print("Enter a number of operation: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Println("Your current balance:", curBalance, "EUR")
		case 2:
			fmt.Println("How much do you want to deposit?")
			fmt.Scan(&enterVal)
			curBalance += enterVal
			fmt.Println("Your current balance:", curBalance, "EUR")
			writeBalanceToFile(curBalance)
		case 3:
			fmt.Println("How much do you want to withdraw?")
			fmt.Scan(&enterVal)
			if enterVal > curBalance {
				fmt.Println("Insufficient funds.")
			} else {
				curBalance -= enterVal
				fmt.Println("Your current balance:", curBalance, "EUR")
				writeBalanceToFile(curBalance)
			}
		case 4:
			fmt.Println("Goodbye! Have a nice day!")
			return
		default:
			fmt.Println("Invalid option")
		}

	}

}
