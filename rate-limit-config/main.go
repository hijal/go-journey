package main

import "fmt"

const (
	MaxRequestsPerMin = 120
	IdleTimeoutSec    = 30
	ReadTimeoutSec    = 5
	defaultZone       = "us-east-1"
)

func main() {
	fmt.Printf("Rate limit %d req/min\n", MaxRequestsPerMin)
	fmt.Printf("Idle timeout %ds | read timeout: %ds\n", IdleTimeoutSec, ReadTimeoutSec)
	fmt.Printf("Default deployment zone: %s\n", defaultZone)
}
