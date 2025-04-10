package main

import "fmt"

type floatMap map[string]float64 // Custom type (Type Aliases)for map[string]float64

func (m floatMap) output() {
	fmt.Println(m)
}
func main() {
	// userNames := []string{}
	userNames := make([]string, 2, 5) // 2 is the length, 5 is the capacity

	userNames[0] = "John"

	userNames = append(userNames, "Max")
	userNames = append(userNames, "Paul")

	fmt.Println(userNames)

	courseRatings := make(floatMap, 3) // 3 is the capacity

	courseRatings["go"] = 4.7
	courseRatings["react"] = 4.8
	courseRatings["angular"] = 4.9
	courseRatings["vue"] = 4.9

	courseRatings.output()
	
	// fmt.Println(cap(userNames)) // Prints the capacity of the slice
	// fmt.Println(len(courseRatings)) // Prints the number of key-value pairs in the map
}
