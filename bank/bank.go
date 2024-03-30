package main

import (
	"fmt"
	"go_tutorials/bank/fileOps"
)

const accBalanceFile = "balance.txt"

func main() {
	var curBalance, err = fileOps.GetFloatFromFile(accBalanceFile)
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
		menuOptions()

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
			fileOps.WriteFloatToFile(curBalance, accBalanceFile)
		case 3:
			fmt.Println("How much do you want to withdraw?")
			fmt.Scan(&enterVal)
			if enterVal > curBalance {
				fmt.Println("Insufficient funds.")
			} else {
				curBalance -= enterVal
				fmt.Println("Your current balance:", curBalance, "EUR")
				fileOps.WriteFloatToFile(curBalance, accBalanceFile)
			}
		case 4:
			fmt.Println("Goodbye! Have a nice day!")
			return
		default:
			fmt.Println("Invalid option")
		}

	}

}
