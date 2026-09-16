package main

import (
	"errors"
	"fmt"
)

type DeclineError struct {
	Code string
	Msg  string
}

func (e *DeclineError) Error() string {
	return fmt.Sprintf("declined (code %s): %s", e.Code, e.Msg)
}

func authorize(amountCents int64) error {
	if amountCents > 500_000 {
		return &DeclineError{Code: "51", Msg: "insufficient funds"}
	}
	return nil
}

func main() {
	if err := authorize(100_000); err != nil {
		fmt.Println("unexpected:", err)
	} else {
		fmt.Println("approved, err =", err)
	}

	err := authorize(750_000)
	if err == nil {
		fmt.Println("unexpected approval")
		return
	}

	var decline *DeclineError

	if errors.As(err, &decline) {
		fmt.Println("retry with another card, code:", decline.Code)
	} else {
		fmt.Println("generic failure:", err)
	}
}
