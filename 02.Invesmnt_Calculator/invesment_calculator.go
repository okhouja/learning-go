package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate = 2.5
	// invesmentAmount, years, expectingReturnRate := 1000.0, 10.0, 5.5 // Shortcut version
	var invesmentAmount float64 = 1000
	expectingReturnRate := 5.5
	// years := 10.0
	var years float64

	fmt.Print("Invesment Amount: ")
	fmt.Scan(&invesmentAmount)

	fmt.Print("Expected Return Rate: ")
	fmt.Scan(&expectingReturnRate)

	fmt.Print("Years: ")
	fmt.Scan(&years)

	futureValue := invesmentAmount * math.Pow(1+expectingReturnRate/100, years)
	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	// Output Information
	// fmt.Println("Future Value:",futureValue)
	fmt.Printf("Future Value: %v\nFuture Value (adjusted for inflation): %v", futureValue,futureRealValue)
	// fmt.Println("Future Value (adjusted for inflation):",futureRealValue)
}
