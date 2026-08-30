package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("int32 max value:", math.MaxInt32)
	fmt.Println("int64 max value:", math.MaxInt64)

	var total32 int32 = math.MaxInt32 - 2
	deltas := []int32{1, 1, 1, 1}

	for _, d := range deltas {
		total32 += d
	}
	fmt.Println("total32 after overflow:", total32)

	var total64 int64 = math.MaxInt32 - 2
	for _, d := range deltas {
		total64 += int64(d)
	}
	fmt.Println("total64 (safe)      :", total64)
}
