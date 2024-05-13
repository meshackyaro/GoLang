package main

import "fmt"

func main() {
	var number int
	fmt.Println("Enter number: ")
	fmt.Scanln(&number)
	total := 1

	for input := number; input > 0; input-- {
		total *= input
	}
	fmt.Println("Factorial of", number, "is", total)
}
