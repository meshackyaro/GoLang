package main

import "fmt"

func main() {

	weeklyPay := 200.00
	value := 0.00
	bonus := 0.09

	for value != -1 {
		fmt.Print("Enter item value: ")
		fmt.Scan(&value)

		newValue := (value * bonus) + weeklyPay
		fmt.Println("New Value: ", newValue)
	}

}
