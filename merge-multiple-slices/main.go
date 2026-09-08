package main

import "fmt"

func MergeSlices(slices ...[]int) []int {
	var result []int
	for _, slice := range slices {
		result = append(result, slice...)
	}
	return result
}

func main() {
	s1 := []int{1, 2}
	s2 := []int{3, 4, 5}
	s3 := []int{6}

	fmt.Println(MergeSlices(s1, s2, s3))
}
