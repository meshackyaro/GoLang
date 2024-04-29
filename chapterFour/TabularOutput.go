package main

import "fmt"

func main() {
	fmt.Println("N\tN^2\tN^3\tN^4")
	for counter := 1; counter <= 5; counter++ {
		fmt.Printf("%d\t%d\t%d\t%d\n", counter, counter*counter, counter*counter*counter, counter*counter*counter*counter)
	}
}
