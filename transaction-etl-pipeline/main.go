package main

import (
	"fmt"
	"strconv"
	"strings"
)

type tx struct {
	id     string
	amount int
}

func source(lines []string) <-chan string {
	out := make(chan string)

	go func() {
		defer close(out)
		for _, line := range lines {
			out <- line
		}
	}()

	return out
}

func parseStage(in <-chan string) (<-chan tx, <-chan error) {
	out := make(chan tx)

	errc := make(chan error, 16)

	go func() {
		defer close(out)
		defer close(errc)

		for line := range in {
			id, raw, ok := strings.Cut(line, ",")

			if !ok {
				errc <- fmt.Errorf("malformed line %q", line)
				continue
			}

			amount, err := strconv.Atoi(raw)
			if err != nil {
				errc <- fmt.Errorf("parse amount of %s: %w", id, err)
				continue
			}

			out <- tx{id: id, amount: amount}
		}
	}()
	return out, errc
}

func filter(in <-chan tx, threshold int) <-chan tx {
	out := make(chan tx)

	go func() {
		defer close(out)

		for t := range in {
			if t.amount >= threshold {
				out <- t
			}
		}
	}()
	return out
}

func main() {
	lines := []string{
		"TX-1001,2500",
		"TX-1002,180",
		"TX-1003,notanumber",
		"TX-1004,9900",
		"brokenline",
		"TX-1005,430",
	}

	txs, errc := parseStage(source(lines))

	highValue := filter(txs, 500)

	total, count := 0, 0

	for t := range highValue {
		count++
		total += t.amount

		fmt.Println("high-value:", t.id, t.amount)
	}

	for err := range errc {
		fmt.Println("skipped:", err)
	}

	fmt.Printf("%d high-value transactions, total %d BDT\n", count, total)
}
