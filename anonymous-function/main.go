package main

import "fmt"

func main() {
	// IIFE
	result := func(x int) int {
		return x * x
	}(6)

	double := func(x int) int {
		return x * x
	}
	fmt.Println(result)
	fmt.Println(double(10))
}
