package main

import "fmt"

func main() {

	firstNumber := 0
	secondNumber := 0

	for count := 1; count < 2; count++ {

		fmt.Print("Enter first number: ")
		fmt.Scanf("%d", &firstNumber)

		fmt.Println("Enter second number: ")
		fmt.Scanf("%d", &secondNumber)

		if firstNumber == secondNumber {
			fmt.Println("0")
		} else if firstNumber > secondNumber {
			fmt.Println("1")
		} else if firstNumber < secondNumber {
			fmt.Println("-1")
		}

	}
}
