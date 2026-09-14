package main

import (
	"fmt"
	"maps"
	"slices"
)

func main() {
	statusCodes := []int{
		200, 404, 200, 500, 200, 404, 301, 200, 500, 200, 403, 404,
	}

	counts := make(map[int]int)

	for _, code := range statusCodes {
		counts[code]++
	}

	for _, code := range slices.Sorted(maps.Keys(counts)) {
		fmt.Printf("HTTP %d: %d hits\n", code, counts[code])
	}
	mostCode, mostHits := 0, 0

	for code, hits := range counts {
		if hits > mostHits {
			mostCode, mostHits = code, hits
		}
	}
	fmt.Printf("Most frequent: HTTP %d (%d hits)\n", mostCode, mostHits)
}
