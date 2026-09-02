package main

import "fmt"

func logDestination(level string) string {
	switch level {
	case "DEBUG", "Trace":
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
	fmt.Println("DEBUG   ->", logDestination("DEBUG"))
	fmt.Println("INFO    ->", logDestination("INFO"))
	fmt.Println("WARN    ->", logDestination("WARN"))
	fmt.Println("FATAL   ->", logDestination("FATAL"))
	fmt.Println("VERBOSE ->", logDestination("VERBOSE"))
}
