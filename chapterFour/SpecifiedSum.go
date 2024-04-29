package main

import "fmt"

func main() {

	otherInputs := 0
	sum := 0

	fmt.Print("Enter firstInput: ")
	var firstInput int
	fmt.Scanln(&firstInput)

	for sum < firstInput {
		fmt.Print("Enter otherInput: ")
		fmt.Scanln(&otherInputs)
		sum += otherInputs
	}
	fmt.Println(sum)
}
