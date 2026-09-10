package main

import "fmt"

func main() {
	weeklyLoad := [7]float64{412.5, 398.2, 441.7, 460.1, 455.9, 380.4, 372.8}
	days := [7]string{"Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"}

	total := 0.0
	peakDay, lowDay := 0, 0

	for day, load := range weeklyLoad {
		total += load

		if load > weeklyLoad[peakDay] {
			peakDay = day
		}

		if load < weeklyLoad[lowDay] {
			lowDay = day
		}
	}

	fmt.Printf("weekly total: %.1f MW\n", total)
	fmt.Printf("average: %.2f MW\n", total/float64(len(weeklyLoad)))
	fmt.Printf("peak: %s %.1f MW\n", days[peakDay], weeklyLoad[peakDay])
	fmt.Printf("low: %s %.1f MW\n", days[lowDay], weeklyLoad[lowDay])
}
