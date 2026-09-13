package main

import "fmt"

func main() {
	stock := [][]int{
		{120, 95, 140, 110},  // warehouse A
		{60, 70, 0, 55},      // warehouse B
		{200, 180, 190, 205}, // warehouse C
	}
	warehouses := []string{"A", "B", "C"}

	for i, row := range stock {
		total := 0
		for _, item := range row {
			total += item
		}
		fmt.Printf("warehouse %s total: %d\n", warehouses[i], total)
	}

	for w := range 4 {
		weekTotal := 0
		for _, row := range stock {
			weekTotal += row[w]
		}
		fmt.Printf("week %d total: %d\n", w+1, weekTotal)
	}
}
