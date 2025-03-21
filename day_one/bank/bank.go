package main

import "fmt"

func main() {
	accountBalanace := 1000.0
	fmt.Println("\nWelcome to the bank!")
	for {
		menu()

		var choice int
		fmt.Print("Enter your choice: ")
		fmt.Scan(&choice)

		if choice == 1 {
			showBalanace(accountBalanace)
		} else if choice == 2 {
			depositeAmount(accountBalanace)
		} else if choice == 3 {
			withdrawAmount(accountBalanace)

		} else if choice == 4 {
			fmt.Println("\nThank you for banking with us!")
			return
		} else {
			fmt.Println("\nThank you for banking with us!")
			continue
		}
	}
}

func depositeAmount(accountBalanace float64) {
	var depositAmount float64
	fmt.Print("Enter deposit amount: ")
	fmt.Scan(&depositAmount)
	if depositAmount <= 0 {
		fmt.Println("Invalid deposit amount must be graeter than 0")
		return
	}
	accountBalanace += depositAmount
	fmt.Printf("Balance updated! Your account balance is: %v", accountBalanace)
}
func showBalanace(accountBalanace float64) {
	fmt.Printf("Your account balance is: %v\n", accountBalanace)
}
func withdrawAmount(accountBalanace float64) {
	var withdrawAmount float64
	fmt.Print("Enter withdraw amount: ")
	fmt.Scan(&withdrawAmount)
	if withdrawAmount > accountBalanace {
		fmt.Println("Insufficient balance!")
		return
	} else {
		fmt.Printf("Balance updated! Your account balance is: %v", accountBalanace)
	}
}
func menu() {
	fmt.Println("\nWhat do you want to do?")
	fmt.Println(("1. Check balance"))
	fmt.Println(("2. Deposit Money"))
	fmt.Println(("3. Withdraw Money"))
	fmt.Println(("4. Exit"))
}
