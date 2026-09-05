package main

import "fmt"

type Operator func(int, int) int

func multiply(a, b int) int {
	return a * b
}

func main() {
	var op Operator = multiply
	fmt.Println("function value:", op(3, 4))
}
