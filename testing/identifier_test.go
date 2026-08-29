package main

import (
	"go/token"
	"testing"
)

func TestIsValidIdentifier(t *testing.T) {
	cases := []struct {
		name     string
		input    string
		expected bool
	}{
		{"simple lowercase", "amount", true},
		{"simple exported", "Amount", true},
		{"with underscore", "amount_cents", true},
		{"starts with underscore", "_temp", true},
		{"blank identifier", "_", true},
		{"starts with digit", "1amount", false},
		{"contains space", "amount cents", false},
		{"contains hyphen", "amount-cents", false},
		{"empty string", "", false},
		{"go keyword", "func", false},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := token.IsIdentifier(tc.input)
			if got != tc.expected {
				t.Errorf("IsIdentifier(%q) = %v, want %v", tc.input, got, tc.expected)
			}
		})
	}
}
