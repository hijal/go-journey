package main

import (
	"fmt"
	"maps"
	"slices"
)

func main() {
	jobs := []struct {
		Name   string
		Status string
	}{
		{"build-api", "passed"},
		{"lint", "failed"},
		{"unit-tests", "passed"},
		{"integration-tests", "failed"},
		{"deploy-staging", "pending"},
		{"build-web", "passed"},
	}

	byStatus := make(map[string][]string)

	for _, j := range jobs {
		byStatus[j.Status] = append(byStatus[j.Status], j.Name)
	}

	for _, status := range slices.Sorted(maps.Keys(byStatus)) {
		fmt.Printf("%s (%d):\n", status, len(byStatus[status]))

		for _, name := range byStatus[status] {
			fmt.Println("  -", name)
		}
	}
}
