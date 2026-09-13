package main

import (
	"fmt"
	"math/rand"
)

func main() {
	servers := []string{"api-1", "api-2", "db-1", "cache-1", "worker-1"}

	failed := make([]string, 0, len(servers))

	for _, server := range servers {
		healthy := rand.Intn(100) < 80
		if !healthy {
			failed = append(failed, server)
		}
	}
	fmt.Println("healthy servers:", len(servers)-len(failed))
	fmt.Println("total servers:", len(servers))
	fmt.Println("failed servers:", failed)
}
