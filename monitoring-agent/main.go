package main

import (
	"fmt"
	"math"
)

func main() {
	sample32 := []float32{12.1, 45.667, 8.333, 99.9, 0.001}
	sample64 := []float64{12.1, 45.667, 8.333, 99.9, 0.001}

	var sum32 float32
	for _, v := range sample32 {
		sum32 += v
	}

	var sum64 float64
	for _, v := range sample64 {
		sum64 += v
	}

	fmt.Printf("float32 sum: %.10f\n", sum32)
	fmt.Printf("float64 sum: %.10f\n", sum64)

	target := 166.001
	const epsilon = 1e-9
	fmt.Println("epsilon:", epsilon)

	if math.Abs(sum64-target) < epsilon {
		fmt.Println("sum64 matches target (within tolerance)")
	} else {
		fmt.Printf("sum64 differs from target by %.15f\n", sum64-target)
	}

	a, b, c := 0.1, 0.2, 0.3

	fmt.Println("a + b == c ?", a+b == c)
	fmt.Printf("a + b = %.20f\n", a+b)
	fmt.Printf("c     = %.20f\n", c)
}
