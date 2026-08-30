package main

import "math"

func BDTToPaisa(bdt float64) int64 {
	return int64(math.Round(bdt * 100))
}
