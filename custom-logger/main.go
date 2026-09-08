package main

import (
	"fmt"
	"strings"
)

func Log(level string, message ...any) {
	strMsgs := make([]string, len(message))
	for i, msg := range message {
		strMsgs[i] = fmt.Sprint(msg)
	}
	finalMsg := strings.Join(strMsgs, " ")
	fmt.Printf("[%s] %s\n", strings.ToUpper(level), finalMsg)
}

func main() {
	Log("info", "User", 101, "logged in from", "192.168.1.1")
	Log("error", "Failed to connect to database")
	Log("debug", "Processing request", "with parameters:", map[string]string{"id": "123", "action": "update"})
	fmt.Println("Logging completed.")
}
