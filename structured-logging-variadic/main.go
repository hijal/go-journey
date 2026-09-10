package main

import (
	"fmt"
	"strings"
	"time"
)

func Log(level string, message string, keysAndValues ...any) {
	if len(keysAndValues)%2 != 0 {
		fmt.Printf("warning: odd number of key-value arguments passed to logger")
	}

	timestamp := time.Now().Format(time.RFC3339)
	var contextData []string

	for i := 0; i < len(keysAndValues)-1; i += 2 {
		contextData = append(contextData, fmt.Sprintf("%v=%v", keysAndValues[i], keysAndValues[i+1]))
	}

	contextStr := ""

	if len(contextData) > 0 {
		contextStr = " | " + strings.Join(contextData, ", ")
	}

	fmt.Printf("[%s] [%s] %s%s\n", timestamp, strings.ToUpper(level), message, contextStr)
}

func main() {
	Log("info", "server started")
	Log("error", "database connection failed", "host", "localhost", "port", "5432", "retry", "true")
}
