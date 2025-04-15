package main

import "fmt"

type transformFN func(int) int //

// type anotherFN func(int,[]string,map{string}[]int)([]int,string) int // Function type sample show how complex function type can be
// Custom type (Type Aliases) for function that takes an int and returns an int

func main() {
	numbers := []int{1, 2, 3, 4} // Original slice of integers

	moreNumbers := []int{5, 1, 2}

	doubled := transformNumbers(&numbers, double) // Apply the 'double' function to each element in numbers

	tripled := transformNumbers(&numbers, triple) // Apply the 'triple' function to each element in numbers

	// Print the results
	fmt.Println(doubled) // [2 4 6 8]
	fmt.Println(tripled) // [3 6 9 12]

	transformerFn1 := getTransformerFunction(&numbers)
	transformerFn2 := getTransformerFunction(&moreNumbers)

	transformedNumbers := transformNumbers(&numbers, transformerFn1)
	moreTransformedNumbers := transformNumbers(&moreNumbers, transformerFn2)

	fmt.Println(transformedNumbers)
	fmt.Println(moreTransformedNumbers)
}

// transformNumbers takes a pointer to a slice of integers
// and a function (transformFN) to apply to each element.
// It returns a new slice with the transformed values.
func transformNumbers(numbers *[]int, transform transformFN) []int {
	dNumbers := []int{}

	// Loop through the slice and apply the transform function
	for _, val := range *numbers {
		dNumbers = append(dNumbers, transform(val))
	}

	return dNumbers
}

func getTransformerFunction(numbers *[]int) transformFN {
	if (*numbers)[0] == 1 {
		return double
	} else {
		return triple
	}
}

// double returns the given number multiplied by 2.
func double(n int) int {
	return n * 2
}

// triple returns the given number multiplied by 3.
func triple(n int) int {
	return n * 3
}
