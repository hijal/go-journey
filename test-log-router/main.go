package main

import "fmt"

func logDestination(level string) string {
	switch level {
	case "DEBUG", "TRACE":
		return "local file (rotated daily)"
	case "INFO":
		return "stdout"
	case "WARN", "ERROR":
		return "stdout + alerting pipeline"
	case "FATAL":
		return "stdout + alerting pipeline + pager"
	default:
		return "unknown level: drop and log a config warning"
	}
}

func main() {
	fmt.Println(logDestination("INFO"))
	fmt.Println(logDestination("DEBUG"))
	fmt.Println(logDestination("WARN"))
	fmt.Println(logDestination("ERROR"))
	fmt.Println(logDestination("FATAL"))
	fmt.Println(logDestination("VERBOSE"))
}
