package main

import (
	"fmt"
	"strings"
)

func maskAccount(account string) string {
	if len(account) < 4 {
		return strings.Repeat("*", len(account))
	}

	visible := account[len(account)-4:]

	return strings.Repeat("*", len(account)-4) + visible
}

func main() {
	accounts := []string{"1234567890", "98765", "42"}

	for _, acc := range accounts {
		fmt.Println(acc, "->", maskAccount(acc))
	}
}
