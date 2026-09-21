package main

import (
	"errors"
	"fmt"
	"net/http"
)

type NotFoundError struct {
	Resource string
	ID       string
}

func (e *NotFoundError) Error() string {
	return fmt.Sprintf("%s %q not found", e.Resource, e.ID)
}

type ValidationError struct {
	Field string
	Err   error
}

func (e *ValidationError) Error() string {
	return e.Field + ": " + e.Err.Error()
}

func (e *ValidationError) Unwrap() error {
	return e.Err
}

var (
	ErrUnauthorized = errors.New("unauthorized")
	ErrNegative     = errors.New("must not be negative")
)

func GetAccount(id, token string) error {
	if token == "" {
		return ErrUnauthorized
	}

	if id != "acc-1" {
		return fmt.Errorf("account service: %w", &NotFoundError{Resource: "account", ID: id})
	}
	return nil
}

func Withdraw(amount int64) error {
	if amount < 0 {
		return fmt.Errorf("withdraw: %w", &ValidationError{Field: "amount", Err: ErrNegative})
	}
	return nil
}

func StatusFor(err error) int {
	var nf *NotFoundError
	var ve *ValidationError

	switch {
	case err == nil:
		return http.StatusOK
	case errors.Is(err, ErrUnauthorized):
		return http.StatusUnauthorized
	case errors.As(err, &nf):
		return http.StatusNotFound
	case errors.As(err, &ve):
		return http.StatusUnprocessableEntity
	default:
		return http.StatusInternalServerError
	}
}

func main() {
	cases := []error{
		GetAccount("acc-1", "tok"),
		GetAccount("acc-1", ""),
		GetAccount("acc-9", "tok"),
		Withdraw(-50),
		errors.New("disk full"),
	}

	for _, err := range cases {
		fmt.Printf("%d %v\n", StatusFor(err), err)
	}

	var nf *NotFoundError
	if err := GetAccount("acc-9", "tok"); errors.As(err, &nf) {
		fmt.Println("missing resource:", nf.Resource, nf.ID)
	}
	fmt.Println("is negative:", errors.Is(Withdraw(-1), ErrNegative))
}
