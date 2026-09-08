package main

import (
	"errors"
	"fmt"
	"strings"
)

type MultiError struct {
	Errors []error
}

func (m *MultiError) Error() string {
	msgs := make([]string, len(m.Errors))

	for i, err := range m.Errors {
		msgs[i] = err.Error()
	}

	return strings.Join(msgs, ";")
}

func CombineErrors(errs ...error) error {
	var mErr MultiError

	for _, err := range errs {
		if err != nil {
			mErr.Errors = append(mErr.Errors, err)
		}
	}
	if len(mErr.Errors) == 0 {
		return nil
	}
	return &mErr
}

func main() {
	err := CombineErrors(nil, errors.New("DB timeout"), errors.New("invalid input"))

	if err != nil {
		fmt.Println(err)
	}
}
