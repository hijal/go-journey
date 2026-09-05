package main

import (
	"errors"
	"fmt"
)

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return (a / b), nil
}

func main() {
	result, err := divide(10, 3)
	if err != nil {
		fmt.Println("divide error:", err)
	}
	fmt.Println(result)
}
