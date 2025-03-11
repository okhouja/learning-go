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

	// fmt.Print("Invesment Amount: ")
	outputText("Invesment Amount: ")
	fmt.Scan(&invesmentAmount)

	// fmt.Print("Expected Return Rate: ")
	outputText("Expected Return Rate: ")
	fmt.Scan(&expectingReturnRate)

	fmt.Print("Years: ")
	fmt.Scan(&years)

	futureValue := invesmentAmount * math.Pow(1+expectingReturnRate/100, years)
	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	// Create Formatted String
	formattedFV := fmt.Sprintf("Future Value: %.1f\n", futureValue)
	formattedRFV := fmt.Sprintf("Future Value (adjusted for inflation): %.1f\n", futureRealValue)

	// Output Information
	// fmt.Println("Future Value:",futureValue)	
	// fmt.Println("Future Value (adjusted for inflation):",futureRealValue)
	// fmt.Printf("Future Value: %.1f\nFuture Value (adjusted for inflation): %.1f", futureValue,futureRealValue)
	
	// multiline String (Working with Backtick :``)
	// fmt.Printf(`Future Value: %.1f\n
	// Future Value (adjusted for inflation): %.1f`, futureValue,futureRealValue)

	fmt.Print(formattedFV, formattedRFV)
	
}
	func outputText(text string){
		fmt.Println(text)
}
