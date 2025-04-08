package main

import "fmt"

func main() {
	// 1) Create a new array (!) that contains three hobbies you have
	// 		Output (print) that array in the command line.
	fmt.Println("Task 1")
	hobbies := [3]string{"reading", "writing", "coding"}
	fmt.Println(hobbies)

	// 2) Also output more data about that array:
	//		- The first element (standalone)
	//		- The second and third element combined as a new list
	fmt.Println("Task 2")

	fmt.Println(hobbies[0])
	updateHobbies := hobbies[1:3]
	fmt.Println(updateHobbies)

	// 3) Create a slice based on the first element that contains
	//		the first and second elements.
	//		Create that slice in two different ways (i.e. create two slices in the end)
	fmt.Println("Task 3")

	mainHobbies := hobbies[:2]
	fmt.Println(mainHobbies)

	// 4) Re-slice the slice from (3) and change it to contain the second
	//		and last element of the original array.
	fmt.Println("Task 4")

	fmt.Println(cap(mainHobbies))
	mainHobbies = mainHobbies[1:3]
	fmt.Println(mainHobbies)

	// 5) Create a "dynamic array" that contains your course goals (at least 2 goals)
	fmt.Println("Task 5")

	courseGoals := []string{"learn Go!", "learn PostgreSQL"}
	fmt.Println(courseGoals)

	// 6) Set the second goal to a different one AND then add a third goal to that existing dynamic array
	fmt.Println("Task 6")

	courseGoals[1] = "learn Docker"
	courseGoals = append(courseGoals, "learn Kubernetes")
	fmt.Println(courseGoals)

	// 7) Bonus: Create a "Product" struct with title, id, price and create a
	//		dynamic list of products (at least 2 products).
	//		Then add a third product to the existing list of products.
	fmt.Println("Task 7")

	type Product struct {
		id    int
		title string
		price float64
	}
	products := []Product{
		{
			1,
			"product 1",
			12.99,
		},
		{
			2,
			"product 2",
			15.99,
		},
	}
	fmt.Println(products)

	newProduct := Product{
		3,
		"product 3",
		19.99,
	}

	products = append(products, newProduct)
	fmt.Println(products)

}

// Time to practice what you learned!
