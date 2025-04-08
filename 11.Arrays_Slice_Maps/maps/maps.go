package main

import "fmt"

func main() {
	websites := map[string]string{
		"Google":              "https://google.com",
		"Amazon Web Services": "https://aws.com",
	}
	fmt.Println(websites)

	websites["Facebook"] = "https://facebook.com"
	fmt.Println(websites)

	// 
	delete(websites, "Google")
	fmt.Println(websites)
}
