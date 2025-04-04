package main

import "fmt"

func main() {
	age:= 32 // regular variable

	var agePointer *int
	agePointer = &age // pointer variable

	fmt.Println("Age:", *agePointer)

	adultYears := getAdultYears(agePointer) // function call
	fmt.Println("Adult years:", adultYears)	
}

func getAdultYears(age *int) int {
	return *age - 18

}