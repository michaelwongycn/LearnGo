package main

import (
	"fmt"
)

func printPrompt(currentCommand *uint8, currentNumber *float32) {
	fmt.Printf("Current number is %.2f\n", *currentNumber)
	fmt.Println("Command list")
	fmt.Println("1. Addition (+)")
	fmt.Println("2. Subtraction (-)")
	fmt.Println("3. Multiplication (*)")
	fmt.Println("4. Division (/)")
	fmt.Println("0. Stop")
}

func addition(currentNumber *float32, targetNumber *float32) {
	*currentNumber = *currentNumber + *targetNumber
}

func substraction(currentNumber *float32, targetNumber *float32) {
	*currentNumber = *currentNumber - *targetNumber
}

func multiplication(currentNumber *float32, targetNumber *float32) {
	*currentNumber = *currentNumber * *targetNumber
}

func division(currentNumber *float32, targetNumber *float32) {
	*currentNumber = *currentNumber / *targetNumber
}

func main() {

	var currentNumber float32

	for {
		var currentCommand uint8
		var targetNumber float32

		printPrompt(&currentCommand, &currentNumber)
		fmt.Print("Please Input your command : ")
		fmt.Scan(&currentCommand)

		if currentCommand == 0 {
			break
		}

		fmt.Print("Please Input number : ")
		fmt.Scan(&targetNumber)

		switch currentCommand {
		case 1:
			addition(&currentNumber, &targetNumber)
		case 2:
			substraction(&currentNumber, &targetNumber)
		case 3:
			multiplication(&currentNumber, &targetNumber)
		case 4:
			if targetNumber == 0 {
				fmt.Println("Can't Divide by 0")
			} else {
				division(&currentNumber, &targetNumber)
			}
		default:
			fmt.Println("Invalid Command")
		}
	}

	fmt.Printf("Final Number is %.2f", currentNumber)

}
