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
	years := 10.0

	fmt.Scan(&invesmentAmount)

	futureValue := invesmentAmount * math.Pow(1+expectingReturnRate/100, years)
	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	fmt.Println(futureValue)
	fmt.Println(futureRealValue)
}
