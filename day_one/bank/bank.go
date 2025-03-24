package main

import (
	"errors"
	"fmt"
	"os"
)

const accountBalanceFile = "balance.txt"

func writeBalanceToFile(balance float64) {
	balanceText := fmt.Sprintf("%v", balance)
	os.WriteFile(accountBalanceFile, []byte(balanceText), 0644)
}

func readBalanceFromFile() (float64, error) {
	data, err := os.ReadFile(accountBalanceFile)
	if err != nil {
		return 1000, errors.New("error reading balance file")
	}
	balanceText := string(data)
	var balance float64
	fmt.Sscan(balanceText, &balance)
	return balance, nil
}

func main() {
	accountBalanace, err := readBalanceFromFile()
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
			writeBalanceToFile(accountBalanace)
		case 3:
			var withdrawAmount float64
			fmt.Print("Enter withdraw amount: ")
			fmt.Scan(&withdrawAmount)
			if withdrawAmount > accountBalanace {
				fmt.Println("Insufficient balance!")
			} else {
				accountBalanace -= withdrawAmount
				fmt.Printf("Balance updated! Your account balance is: %v", accountBalanace)
			}
			writeBalanceToFile(accountBalanace)
		default:
			fmt.Println("\nThank you for banking with us!")
			return
		}
	}
}

func menu() {
	fmt.Println("\n\nWhat do you want to do?")
	fmt.Println(("1. Check balance"))
	fmt.Println(("2. Deposit Money"))
	fmt.Println(("3. Withdraw Money"))
	fmt.Println(("4. Exit"))
}
