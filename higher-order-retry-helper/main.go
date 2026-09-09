package main

import (
	"errors"
	"fmt"
	"time"
)

var errGatewayDown = errors.New("Payment gateway down")

func withRetry(attempts int, operation func() error) error {
	var lastErr error

	for i := range attempts {
		err := operation()

		if err == nil {
			fmt.Printf("attempt %d: success\n", i+1)
			return nil
		}

		lastErr = err

		fmt.Printf("attempt %d failed: %v\n", i+1, err)
		time.Sleep(100 * time.Millisecond)
	}
	return fmt.Errorf("withRetry: all %d attempts failed: %w", attempts, lastErr)
}

func main() {
	calls := 0
	flakyCharge := func() error {
		calls++

		if calls <= 2 {
			return fmt.Errorf("charge 4999 cents: %w", errGatewayDown)
		}
		return nil
	}

	if err := withRetry(5, flakyCharge); err != nil {
		fmt.Println("gave up:", err)
	} else {
		fmt.Println("payment captured")
	}

	fmt.Println("---")

	deadCharge := func() error {
		return fmt.Errorf("charge 1299 cents: %w", errGatewayDown)
	}

	if err := withRetry(3, deadCharge); err != nil {
		fmt.Println("final error:", err)
		if errors.Is(err, errGatewayDown) {
			fmt.Println("gateway down - enqueue job for later")
		}
	}
}
