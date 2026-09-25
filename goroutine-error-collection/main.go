package main

import (
	"errors"
	"fmt"
	"sync"
)

var errCardExpired = errors.New("card expired")

func chargeCard(cardID string) error {
	if cardID == "card-expired" {
		return errCardExpired
	}
	return nil
}

func main() {
	cards := []string{"card-ok-1", "card-expired", "card-ok-2"}

	errs := make([]error, len(cards))

	var wg sync.WaitGroup

	for i, card := range cards {
		wg.Go(func() {
			if err := chargeCard(card); err != nil {
				errs[i] = fmt.Errorf("charge %s: %w", card, err)
			}
		})
	}

	wg.Wait()

	if err := errors.Join(errs...); err != nil {
		fmt.Println("some charged failed:")
		fmt.Println(err)
		fmt.Println("any expired card?", errors.Is(err, errCardExpired))
	}
}
