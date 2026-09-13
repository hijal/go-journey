package main

import "fmt"

func main() {
	latencies := []float64{120, 130, 125, 140, 200, 180, 90, 95, 100}
	const window = 3

	avgs := make([]float64, 0, len(latencies))

	for end := range len(latencies) {
		start := max(0, end+1-window)
		sum := 0.0
		for _, v := range latencies[start : end+1] {
			sum += v
		}
		avgs = append(avgs, sum/float64(end+1-start))
	}

	for i, a := range avgs {
		fmt.Printf("t=%d: moving avg = %.2f\n", i, a)
	}
}
