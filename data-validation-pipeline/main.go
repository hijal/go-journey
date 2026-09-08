package main

import (
	"fmt"
	"strings"
)

type Validator func(string) error

func validate(data string, validators ...Validator) []error {
	var errs []error

	for _, v := range validators {
		if err := v(data); err != nil {
			errs = append(errs, err)
		}
	}
	return errs
}

func main() {
	isLongEnough := func(s string) error {
		if len(s) < 5 {
			return fmt.Errorf("too short")
		}
		return nil
	}

	hasNumber := func(s string) error {
		if !strings.ContainsAny(s, "1234567890") {
			return fmt.Errorf("needs a number")
		}
		return nil
	}

	errs := validate("abc", isLongEnough, hasNumber)

	for _, err := range errs {
		fmt.Println(err)
	}
}
