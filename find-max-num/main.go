package main

import "fmt"

func FindMaxNum(nums ...int) int {
	if len(nums) == 0 {
		return 0
	}
	max := nums[0]
	for _, num := range nums {
		if num > max {
			max = num
		}
	}
	return max
}

func main() {
	fmt.Println(FindMaxNum(10, 50, 30, 90, 20))
}
