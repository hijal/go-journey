package main

import "fmt"

func Swap(a, b *int) {
	*a, *b = *b, *a
}

func main() {
	x, y := 10, 20

	fmt.Println("Before:", x, y)
	Swap(&x, &y)
	fmt.Println("After:", x, y)

	prices := []int{40, 10, 30}
	Swap(&prices[0], &prices[1])
	fmt.Println("Swapped slice:", prices)
}
