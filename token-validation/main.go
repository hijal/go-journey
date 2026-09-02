package main

import (
	"errors"
	"fmt"
	"strings"
)

var errTokenExpired = errors.New("Token expired")

func validateToken(token string) (string, error) {
	if !strings.HasPrefix(token, "Bearer ") {
		return "", fmt.Errorf("malformed token %q: missing Bearer prefix", token)
	}
	payload := strings.TrimPrefix(token, "Bearer ")

	if payload == "expired.jwt" {
		return "", fmt.Errorf("token check failed: %w", errTokenExpired)
	}

	return payload, nil
}

func handleRequest(authHeader string) {
	if userID, err := validateToken(authHeader); err != nil {
		if errors.Is(err, errTokenExpired) {
			fmt.Println("401: please refresh your token")
		} else {
			fmt.Println("401:", err)
		}
	} else {
		fmt.Printf("200: serving data for user %q\n", userID)
	}
}

func main() {
	handleRequest("Bearer user-42.jwt")
	handleRequest("Bearer expired.jwt")
	handleRequest("Basic abc")
}
