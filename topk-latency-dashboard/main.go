package main

import (
	"fmt"
	"math/rand"
	"slices"
)

func main() {
	latencies := make([]int, 50)

	for i := range latencies {
		latencies[i] = 20 + rand.Intn(400)
	}

	slices.Sort(latencies)
	slowest := latencies[len(latencies)-3:]
	fmt.Print("slowest 3 requests (ms): ")
	for _, latency := range slowest {
		fmt.Print(latency, " ")
	}
	fmt.Println()

	p95 := latencies[int(float64(len(latencies))*0.95)-1]
	fmt.Println("p95 latency (ms):", p95)
}
