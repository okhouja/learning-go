package main

import "fmt"

func main() {
	age:= 32 // regular variable

	var agePointer *int
	agePointer = &age // pointer variable

	fmt.Println("Age:", *agePointer)

	editAgetoAdultYears(agePointer) // function call
	fmt.Println("Adult years:", age)	
}

func editAgetoAdultYears(age *int)  {
	// return *age - 18
	*age = *age - 18 
}