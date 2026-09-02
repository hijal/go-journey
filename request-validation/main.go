package main

import (
	"fmt"
	"strings"
)

func registerUser(email, password string) string {
	email = strings.TrimSpace(email)

	if email == "" {
		return "400: email is required"
	}

	if password == "" {
		return "400: password is required"
	}

	if len(password) < 8 {
		return "400: password must be at least 8 characters long"
	}

	if !strings.Contains(email, "@") {
		return "400: email must be valid"
	}
	return "Registration successful"
}

func main() {
	fmt.Println(registerUser("  ", "secret123"))
	fmt.Println(registerUser("rata@example.com", "short"))
	fmt.Println(registerUser("rataexample.com", "secret123"))
	fmt.Println(registerUser("rata@example.com", "secret123"))
}
