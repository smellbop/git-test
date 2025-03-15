package main

import (
	"fmt"
	"os"
	"os/exec"
)

const ledger = "ledger.txt"

func main() {
	for {
		//TODO display options
		clearTerminal()
		println("Welcome to the bank. Select an option.")
		println("1. View balance")
		println("2. Deposit")
		println("3. Withdraw")
		println("4. Exit")
		//TODO collect user input
		userInput := getUserInput("Select option.")

		//TODO prompt based on input
		switch userInput {
		case 1:
			viewBalance()
		case 2:
			deposit()
		case 3:
			withdraw()
		case 4:
			clearTerminal()
			println("Bye!")
			return
		default:
			clearTerminal()
			println("Invalid input.")
			getUserInput("any key to menu")
		}
	}
}

func getUserInput(prompt string) float64 {
	println(prompt)
	var userInput float64
	fmt.Scan(&userInput)
	return userInput
}
func clearTerminal() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()

}

func withdraw() {
	clearTerminal()
	println("Withdrawal screen")
}

func deposit() {
	clearTerminal()
	println("deposit screen")
	//TODO get deposit amount
	userDeposit := getUserInput("deposit amount")
	//TODO add deposit to ledger
	addDeposit(userDeposit)
}

func addDeposit(userDeposit float64) {

}

func viewBalance() {
	clearTerminal()
	balance := getCurrentBalance()
	fmt.Printf("Current Balance: %.2f\n", balance)
	getUserInput("any key to return to menu")
}

func getCurrentBalance() float64 {
	//TODO read balance from ledger
	return 9000.0
}
