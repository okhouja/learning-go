package main

import (
	"fmt"	
)

func main(){
	// Define a 4-element array of product names (only first one initialized)
	var productNames [4]string = [4]string{"A Book"}

	// Define an array of 4 prices
	prices := [4]float64{10.99,9.99, 45.99, 20.0}

	// Print the prices array
	fmt.Println(prices)

	// Set the 3rd product name
	productNames[2] = "A Carpet"

	// Print product names
	fmt.Println(productNames)

	// Print the 3rd price
	fmt.Println(prices[2])

	// Slices
		// featuredPrices := prices[1:3]
		// featuredPrices := prices[:3]
		// Create a slice from prices, starting from index 1
		featuredPrices := prices[1:]

		// Modify first element of the slice
		featuredPrices[0] = 199.99

		// Create a smaller slice from featuredPrices (just the first element)
		highlightedPrices := featuredPrices[:1]

	// Print slices	
	fmt.Println("Slice")
	fmt.Println(featuredPrices)
	fmt.Println(prices)	
	fmt.Println(highlightedPrices)
	fmt.Println(len(highlightedPrices), cap(highlightedPrices))

	// Extend highlightedPrices slice
	highlightedPrices = highlightedPrices[:3]

	// Print updated slice
	fmt.Println(highlightedPrices)
	fmt.Println(len(featuredPrices), cap(featuredPrices))
}