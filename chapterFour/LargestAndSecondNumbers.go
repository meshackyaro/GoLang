package main

import "fmt"

func main() {
	counter := 1
	largest := 0
	secondLargest := 0

	for counter <= 10 {
		fmt.Print("Enter number: ")
		var number int
		_, err := fmt.Scanln(&number)
		if err != nil {
			return
		}
		if number > largest {
			secondLargest = largest
			largest = number
		} else if number > secondLargest {
			secondLargest = number
		}
		counter++
	}

	fmt.Println("The largest number is", largest)
	fmt.Println("The second largest number is", secondLargest)
}
