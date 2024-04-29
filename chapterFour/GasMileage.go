package main

import "fmt"

func main() {

	milesDrivenPerTrip := 0
	gallonsPerTankful := 0
	totalMilesDriven := 0
	totalGallons := 0
	trip := 0

	for milesDrivenPerTrip != -1 {

		fmt.Println("Enter miles per trip or -1 to exit: ")
		fmt.Scan(&milesDrivenPerTrip)

		fmt.Print("Enter gallons per tankful: ")
		fmt.Scan(&gallonsPerTankful)

		milesPerGallon := float64(milesDrivenPerTrip / gallonsPerTankful)
		fmt.Println("Miles driven per gallon is: ", milesPerGallon)

		totalMilesDriven += milesDrivenPerTrip
		totalGallons += gallonsPerTankful
		trip++

		if trip > 0 {
			combinedMilesPerGallon := float64(totalMilesDriven / totalGallons)
			fmt.Println("Combined miles per gallon is ", combinedMilesPerGallon)
		}

	}

}
