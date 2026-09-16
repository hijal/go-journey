package main

import (
	"cmp"
	"fmt"
	"slices"
	"sort"
)

type Txn struct {
	ID     string
	Amount int64
}

type statement []Txn

func (s statement) Len() int           { return len(s) }
func (s statement) Less(i, j int) bool { return s[i].Amount < s[j].Amount }
func (s statement) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

var _ sort.Interface = (*statement)(nil)

func main() {
	st := statement{
		{"T-1", 250_00},
		{"T-2", -90_00},
		{"T-3", 1_000_00},
		{"T-4", -15_50},
	}

	sort.Sort(st)
	fmt.Println("ascending by amount (sort.Interface):")
	for _, t := range st {
		fmt.Printf("  %s ৳%.2f\n", t.ID, float64(t.Amount)/100)
	}

	slices.SortFunc(st, func(a, b Txn) int {
		return cmp.Compare(b.Amount, a.Amount)
	})

	fmt.Println("descending by amount (slices.SortFunc):")
	for _, t := range st {
		fmt.Printf("  %s ৳%.2f\n", t.ID, float64(t.Amount)/100)
	}
}
