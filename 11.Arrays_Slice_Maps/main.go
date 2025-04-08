package main

import "fmt"

func main() {
	// userNames := []string{}
	userNames := make([]string, 2, 5) // 2 is the length, 5 is the capacity

	userNames[0] = "John"

	userNames = append(userNames, "Max")
	userNames = append(userNames, "Paul")

	fmt.Println(userNames)
}
