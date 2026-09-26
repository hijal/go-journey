package main

import (
	"fmt"
	"time"
)

func riskScore(cardBIN string, out chan<- int) {
	time.Sleep(80 * time.Millisecond)
	out <- len(cardBIN) * 7
}

func main() {
	scores := make(chan int)

	go riskScore("457173", scores)
	fmt.Println("validating billing address while the risk model runs")
	time.Sleep(30 * time.Millisecond)

	score := <-scores
	fmt.Println("risk score:", score)

	if score > 30 {
		fmt.Println("transaction flagged for manual review")
	}
}
