package main

import (
	"fmt"
	"slices"
)

func main() {
	logs := []string{
		"INFO service started",
		"INFO request received",
		"ERROR db connection refused",
		"INFO retrying",
		"ERROR timeout after 3s",
	}

	tail := logs[len(logs)-3:] // last 3 logs
	fmt.Println("tail:", tail)

	i := slices.IndexFunc(logs, func(line string) bool {
		return len(line) > 5 && line[:5] == "ERROR"
	})

	if i >= 0 {
		fmt.Printf("first error at index %d: %q\n", i, logs[i])
	}

	tail[0] = "INFO tail overwritten"

	fmt.Println("original[2:]", logs[2])
}
