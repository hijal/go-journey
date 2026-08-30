package main

import (
	"fmt"
	"strconv"
)

func main() {
	primary, backup := "server-a.internal", "server-b.internal"
	fmt.Println("Before failover — primary:", primary, "backup:", backup)

	primary, backup = backup, primary
	fmt.Println("After failover  — primary:", primary, "backup:", backup)

	portStr := "8080"
	port, err := strconv.Atoi(portStr)

	if err != nil {
		fmt.Println("invalid port:", err)
		return
	}

	fmt.Println("Parsed port:", port)

	invalidPort := "not-a-number"
	_, err = strconv.Atoi(invalidPort)
	fmt.Println("Expected parse error:", err)
}
