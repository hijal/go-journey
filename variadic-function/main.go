package main

import "fmt"

func sum(nums ...int) int {
	total := 0
	for _, n := range nums {
		total += n
	}
	return total
}

func main() {
	result := sum(1, 2, 3, 4)
	fmt.Println("total sum:", result)
}
