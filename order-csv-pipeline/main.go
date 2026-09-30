package main

import (
	"fmt"
	"strconv"
	"strings"
)

type order struct {
	id    string
	net   int
	total int
	err   error
}

func parse(lines []string) <-chan order {
	out := make(chan order)

	go func() {
		defer close(out)
		for _, line := range lines {
			id, amt, found := strings.Cut(line, ",")

			if !found {
				out <- order{id: line, err: fmt.Errorf("parse %q: missing comma", line)}
				continue
			}
			n, err := strconv.Atoi(amt)
			if err != nil {
				out <- order{id: id, err: fmt.Errorf("parse %q: %w", line, err)}
				continue
			}
			out <- order{id: id, net: n}
		}
	}()
	return out
}

func addVAT(in <-chan order) <-chan order {
	out := make(chan order)

	go func() {
		defer close(out)

		for o := range in {
			if o.err == nil {
				o.total = o.net + o.net*15/100
			}
			out <- o
		}
	}()

	return out
}

func main() {
	lines := []string{
		"ORD-1,1000",
		"ORD-2,abc",
		"ORD-3,2400",
		"broken-line",
	}

	var revenue int
	for o := range addVAT(parse(lines)) {
		if o.err != nil {
			fmt.Println("skip:", o.err)
			continue
		}
		fmt.Printf("%s net=%d total=%d\n", o.id, o.net, o.total)
		revenue += o.total
	}
	fmt.Println("revenue:", revenue)
}
