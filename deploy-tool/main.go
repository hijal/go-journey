package main

import (
	"fmt"
	"time"
)

func deployRelease(version string) {
	start := time.Now()

	defer func() {
		fmt.Printf("[audit] deploy %s took %v\n", version, time.Since(start).Round(time.Millisecond))
	}()

	fmt.Printf("deploying version %s...\n", version)
	time.Sleep(120 * time.Millisecond)
	fmt.Printf("deployment finished\n")
}

func main() {
	deployRelease("v2.4.1")
	fmt.Println("main continues after deploy")
}
