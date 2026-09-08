package main

import (
	"errors"
	"fmt"
	"time"
)

var errPaymentDeclined = errors.New("payment declined")
var gatewayCalls int

func charge(amount int) error {
	gatewayCalls++
	if gatewayCalls < 3 {
		return fmt.Errorf("charge %d BDT: %w", amount, errPaymentDeclined)
	}
	return nil
}

func main() {
	const amount = 1500
	backoff := 100 * time.Millisecond

	remainingAttempts := 5
	for remainingAttempts > 0 {
		remainingAttempts--

		err := charge(amount)
		if err == nil {
			fmt.Printf("payment of %d BDT succeeded\n", amount)
			break
		}

		if !errors.Is(err, errPaymentDeclined) {
			fmt.Println("unexpected error - aborting:", err)
			return
		}

		fmt.Printf("charge failed (%v); retrying in %v\n", err, backoff)
		time.Sleep(backoff)
		backoff *= 2
	}

	if remainingAttempts == 0 {
		fmt.Println("gave up: payment could not be completed")
	}
}
