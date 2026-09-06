package main

import (
	"errors"
	"fmt"
	"strings"
)

func validateConfig(host string, port int) (status string, err error) {
	if strings.TrimSpace(host) == "" {
		return "", errors.New("host is empty")
	}

	if port < 1 || port > 65535 {
		return "", fmt.Errorf("port %d out of range", port)
	}

	status = fmt.Sprintf("config OK: %s:%d", host, port)
	return status, nil
}

func main() {
	status, err := validateConfig("api.example.com", 8080)
	if err != nil {
		fmt.Println("invalid:", err)
	} else {
		fmt.Println(status)
	}

	status, err = validateConfig("", 8080)
	if err != nil {
		fmt.Println("invalid:", err)
	} else {
		fmt.Println(status)
	}

	status, err = validateConfig("api.example.com", 99999)
	if err != nil {
		fmt.Println("invalid:", err)
	} else {
		fmt.Println(status)
	}
}
