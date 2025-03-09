package main

import (
	"fmt"
	"math"
)

func main() {
	// invesmentAmount, years, expectingReturnRate := 1000.0, 10.0, 5.5 // Shortcut version
	var invesmentAmount float64 = 1000
	expectingReturnRate := 5.5
	years := 10.0

	futureValue := invesmentAmount * math.Pow(1+expectingReturnRate/100, years)
	fmt.Print(futureValue)
}
