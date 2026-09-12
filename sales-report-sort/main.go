package main

import (
	"cmp"
	"fmt"
	"slices"
)

type Sale struct {
	Region string
	Amount float64
	Rep    string
}

func main() {
	sales := []Sale{
		{Region: "Dhaka", Amount: 1200, Rep: "Ayesha"},
		{Region: "Chattogram", Amount: 800, Rep: "Rakib"},
		{Region: "Dhaka", Amount: 450, Rep: "Tanvir"},
		{Region: "Sylhet", Amount: 990, Rep: "Mitu"},
		{Region: "Chattogram", Amount: 310, Rep: "Farhan"},
	}

	slices.SortFunc(sales, func(a, b Sale) int {
		return cmp.Compare(b.Amount, a.Amount)
	})

	fmt.Println("Top sales (desc):")
	for _, s := range sales {
		fmt.Printf("  %-11s %-8s %8.2f\n", s.Region, s.Rep, s.Amount)
	}

	slices.SortFunc(sales, func(a, b Sale) int {
		if c := cmp.Compare(a.Region, b.Region); c != 0 {
			return c
		}
		return cmp.Compare(b.Amount, a.Amount)
	})

	fmt.Println("By region, then amount:")
	for _, s := range sales {
		fmt.Printf("  %-11s %-8s %8.2f\n", s.Region, s.Rep, s.Amount)
	}

	var summary struct {
		Orders  int
		Revenue float64
		Best    Sale
	}

	summary.Best = slices.MaxFunc(sales, func(a, b Sale) int {
		return cmp.Compare(a.Amount, b.Amount)
	})

	for _, s := range sales {
		summary.Orders++
		summary.Revenue += s.Amount
	}
	fmt.Printf("Summary: %d orders, revenue %.2f, best rep = %s\n",
		summary.Orders, summary.Revenue, summary.Best.Rep)
}
