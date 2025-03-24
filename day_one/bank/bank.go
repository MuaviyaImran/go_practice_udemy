package main

import (
	"fmt"
	"example.com/bank/utils"
)

const accountBalanceFile = "balance.txt"

func main() {
	accountBalanace, err := utils.ReadFloatFromFile(accountBalanceFile)
	if err != nil {
		fmt.Println(err)
		// panic(err)
	}
	fmt.Println("\nWelcome to the bank!")
	for {
		menu()
		var choice int
		fmt.Print("Enter your choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Printf("Your account balance is: %v\n", accountBalanace)
		case 2:
			// Deposit
			var depositAmount float64
			fmt.Print("Enter deposit amount: ")
			fmt.Scan(&depositAmount)
			if depositAmount <= 0 {
				fmt.Println("Invalid deposit amount must be graeter than 0")
			}
			accountBalanace += depositAmount
			fmt.Printf("Balance updated! Your account balance is: %v", accountBalanace)
			utils.WriteFloatToFile(accountBalanace, accountBalanceFile)
		case 3:
			//Withdraw
			var withdrawAmount float64
			fmt.Print("Enter withdraw amount: ")
			fmt.Scan(&withdrawAmount)
			if withdrawAmount > accountBalanace {
				fmt.Println("Insufficient balance!")
			} else {
				accountBalanace -= withdrawAmount
				fmt.Printf("Balance updated! Your account balance is: %v", accountBalanace)
			}
			utils.WriteFloatToFile(accountBalanace, accountBalanceFile)
		default:
			fmt.Println("\nThank you for banking with us!")
			return
		}
	}
}
