package main

import (
	"fmt"
	"slices"
)

func main() {
	events := []int{1700000100, 1700000250, 1700000400, 1700000550, 1700000700, 1700000850, 1700001000}

	recent := events[len(events)-min(5, len(events)):]
	newestFirst := slices.Clone(recent)
	slices.Reverse(newestFirst)

	fmt.Println("newest first:", newestFirst)
	fmt.Println("original untouched:", events[:2], "...")
}
