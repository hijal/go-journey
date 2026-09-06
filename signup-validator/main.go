package main

import (
	"errors"
	"fmt"
	"strings"
)

type Validator func(string) error

func notEmpty(value string) error {
	if strings.TrimSpace(value) == "" {
		return errors.New("value is empty")
	}
	return nil
}

func minLength(n int) Validator {
	return func(value string) error {
		if len(value) < n {
			return fmt.Errorf("need at least %d characters", n)
		}
		return nil
	}
}

func runValidators(field string, value string, checks ...Validator) []error {
	var errs []error

	for _, check := range checks {
		if err := check(value); err != nil {
			errs = append(errs, fmt.Errorf("%s: %w", field, err))
		}
	}
	return errs
}

func main() {
	password := "hi"
	errs := runValidators("password", password, notEmpty, minLength(8))

	if len(errs) == 0 {
		fmt.Println("password accepted")
	} else {
		for _, err := range errs {
			fmt.Println("rejected:", err)
		}
	}

	var check Validator = minLength(4)
	fmt.Println("username check:", check("raf"))
}
