package main

import (
	"errors"
	"fmt"
)

var errInvalidAmount = errors.New("invalid amount")

type chargeResult struct {
	txnID string
	err   error
}

func charge(card string, amount int, out chan<- chargeResult) {
	if amount <= 0 {
		out <- chargeResult{
			err: fmt.Errorf("charge card %s: %w", card, errInvalidAmount),
		}
		return
	}
	out <- chargeResult{txnID: fmt.Sprintf("TXN-%s-%d", card, amount)}
}

func main() {
	results := make(chan chargeResult)

	for _, amount := range []int{1500, 0} {
		go charge("4242", amount, results)

		res := <-results

		if res.err != nil {
			fmt.Println("failed:", res.err, "| invalid amount?", errors.Is(res.err, errInvalidAmount))
			return
		}

		fmt.Println("charged:", res.txnID)
	}
}
