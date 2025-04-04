package main

import (
	"fmt"
)

func main() {
	
	revenue := getUserInput("Enter Revenue: ")
	expenses := getUserInput("Enter Expenses: ")	
	taxRate := getUserInput("Enter Tax Rate: ")

	ebt, profit, ratio := calculateFinacials(revenue, expenses, taxRate)
		
	// Output results
	fmt.Printf("%.1f\n", ebt)
	fmt.Printf("%.2f\n", profit)
	fmt.Printf("%.3f\n", ratio)
}

func calculateFinacials(revenue, expenses, taxRate float64) (float64, float64, float64){
	ebt := revenue - expenses // ebt = earning before tax
	profit := ebt * (1 - taxRate / 100)
	ratio := ebt / profit
	return ebt, profit, ratio
}

func getUserInput(infoText string) float64{
	var userInput float64
	fmt.Print(infoText)
	fmt.Scan(&userInput)
	return userInput
}