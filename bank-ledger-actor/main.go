package main

import (
	"errors"
	"fmt"
	"sync"
)

var errInsufficientFunds = errors.New("insufficient funds")

type opKind int

const (
	deposit opKind = iota
	withdraw
	balance
)

type response struct {
	balance int
	err     error
}

type request struct {
	kind    opKind
	account string
	amount  int
	reply   chan response
}

func ledger(requests <-chan request) {
	balances := map[string]int{}

	for req := range requests {
		switch req.kind {
		case deposit:
			balances[req.account] += req.amount
		case withdraw:
			if balances[req.account] < req.amount {
				req.reply <- response{
					balance: balances[req.account],
					err:     fmt.Errorf("withdraw %d from %s: %w", req.amount, req.account, errInsufficientFunds),
				}
				continue
			}
			balances[req.account] -= req.amount
		}
		req.reply <- response{
			balance: balances[req.account],
		}
	}
}

func call(requests chan<- request, kind opKind, account string, amount int) (int, error) {
	reply := make(chan response, 1)
	requests <- request{kind: kind, account: account, amount: amount, reply: reply}

	res := <-reply
	return res.balance, res.err
}

func main() {
	requests := make(chan request)
	go ledger(requests)

	var wg sync.WaitGroup
	for range 100 {
		wg.Add(1)

		go func() {
			defer wg.Done()

			if _, err := call(requests, deposit, "ACC-77", 10); err != nil {
				fmt.Println("deposit failed:", err)
			}
		}()
	}
	wg.Wait()

	bal, _ := call(requests, balance, "ACC-77", 0)
	fmt.Println("balance after deposits:", bal)

	if _, err := call(requests, withdraw, "ACC-77", 5000); err != nil {
		fmt.Println("error:", err, "| insufficient?", errors.Is(err, errInsufficientFunds))
	}

	bal, err := call(requests, withdraw, "ACC-77", 400)

	if err != nil {
		fmt.Println("unexpected error:", err)
		return
	}

	fmt.Println("balance after withdrawing 400:", bal)
	close(requests)
}
