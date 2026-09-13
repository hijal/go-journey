package main

import (
	"fmt"
	"slices"
)

func buildReplicaSet(base []string, replicas int) []string {
	merged := slices.Clone(base)
	for i := range replicas {
		merged = append(merged, fmt.Sprintf("pod-%02d", i))
	}
	return merged
}

func main() {
	defaultConfig := []string{"--port=8080", "--log-level=info"}

	prod := buildReplicaSet(defaultConfig, 3)
	prod[0] = "--port=9090"

	fmt.Println("default config:", defaultConfig)
	fmt.Println("production:", prod)
	fmt.Println("equal?", slices.Equal(defaultConfig, []string{"--port=8080", "--log-level=info"}))
}
