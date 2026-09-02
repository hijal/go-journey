package main

import "fmt"

func classify(celsius float64) string {
	if celsius < 38 {
		return "normal"
	} else if celsius < 42 {
		return "warning"
	}
	return "critical"
}

func main() {
	sensorA, sensorB := 41.3, 39.6
	reading := max(sensorA, sensorB)

	switch level := classify(reading); level {
	case "normal":
		fmt.Println("machine-7: running normally")
	case "warning":
		fmt.Printf("machine-7: notify maintenance (%.1f°C)\n", reading)
	default:
		fmt.Printf("machine-7: emergency shutdown (%.1f°C)\n", reading)
	}
}
