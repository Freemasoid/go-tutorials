package main

import "fmt"

func main() {
	var curBalance float64 = 1000.0
	var choice int
	var enterVal float64

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
		case 3:
			fmt.Println("How much do you want to withdraw?")
			fmt.Scan(&enterVal)
			if enterVal > curBalance {
				fmt.Println("Insufficient funds.")
			} else {
				curBalance -= enterVal
				fmt.Println("Your current balance:", curBalance, "EUR")
			}
		case 4:
			fmt.Println("Goodbye! Have a nice day!")
			return
		default:
			fmt.Println("Invalid option")
		}

	}

}
