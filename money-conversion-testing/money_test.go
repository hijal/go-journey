package main

import "testing"

func TestBDTToPaisa(t *testing.T) {
	tests := []struct {
		name string
		bdt  float64
		want int64
	}{
		{"whole number", 100.0, 10000},
		{"two decimal places", 1499.50, 149950},
		{"rounds up", 19.995, 2000},
		{"rounds down", 19.994, 1999},
		{"zero", 0.0, 0},
		{"negative (refund)", -50.25, -5025},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := BDTToPaisa(tt.bdt)

			if got != tt.want {
				t.Errorf("BDTToPaisa(%v) = %d; want %d", tt.bdt, got, tt.want)
			}
		})
	}
}
