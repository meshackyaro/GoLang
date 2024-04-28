package main

import (
	"fmt"
	"math"
)

func main() {
	number := 0
	largestNumber := math.MinInt64

	for counter := 1; counter <= 10; counter++ {

		fmt.Print("Enter a number: ")
		fmt.Scan(&number)

		if number > largestNumber {
			largestNumber = number
		}
	}
	fmt.Printf("The largest number is: %d\n", largestNumber)
}
