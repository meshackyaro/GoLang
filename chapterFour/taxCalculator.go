package main

import "fmt"

func main() {
	citizens := []string{"John", "Mary", "Bob"}
	earnings := make([]float64, 3)
	taxes := make([]float64, 3)

	for count := 0; count < 3; count++ {
		fmt.Printf("Enter %s's earnings for the year: $", citizens[count])
		fmt.Scanln(&earnings[count])

		if earnings[count] <= 30000 {
			taxes[count] = 0.15 * earnings[count]
		} else {
			taxes[count] = 0.15*30000 + 0.20*(earnings[count]-30000)
		}

		fmt.Printf("%s's total tax is: $%.2f\n", citizens[count], taxes[count])

	}

}
