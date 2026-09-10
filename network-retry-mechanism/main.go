package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func Retry(maxAttempts int, baseDelay time.Duration, operation func() error) error {
	var err error

	for attempt := 1; attempt <= maxAttempts; attempt++ {
		err = operation()

		if err == nil {
			return nil
		}

		if attempt < maxAttempts {
			delay := baseDelay * time.Duration(attempt)
			jitter := time.Duration(rand.Intn(100) * int(time.Millisecond))
			time.Sleep(delay + jitter)
			fmt.Printf("Attempt %d failed: %v. Retrying...\n", attempt, err)
		}
	}

	return fmt.Errorf("operation failed after %d attempts: %w", maxAttempts, err)
}

func main() {
	flakyAPICall := func() error {
		if rand.Float32() < 0.7 {
			return errors.New("connection timeout")
		}
		return nil
	}

	err := Retry(3, 100*time.Millisecond, flakyAPICall)

	if err != nil {
		fmt.Println("final failure:", err)
	} else {
		fmt.Println("API call succeed")
	}
}
