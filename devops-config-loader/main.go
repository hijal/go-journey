package main

import "fmt"

func main() {
	fmt.Println("environment:", currentEnvironment)
	fmt.Println("max retries:", maxRetries)
	printBanner()
}

var currentEnvironment = detectEnvironment()

const maxRetries = 5

func detectEnvironment() string {
	return "development"
}

func printBanner() {
	fmt.Println("---deployment config loaded---")
}
