package main

import "fmt"

func fetchPage(page int) []string {
	if page > 2 {
		return nil
	}
	return []string{
		fmt.Sprintf("order-%d-a", page),
		fmt.Sprintf("order-%d-b", page),
	}
}

func main() {
	var allOrders []string
	page := 1

	// fmt.Printf("%v\n", fetchPage(page))
	for {
		orders := fetchPage(page)
		if len(orders) == 0 {
			break
		}

		allOrders = append(allOrders, orders...)
		fmt.Printf("page %d: fetched %d orders\n", page, len(orders))
		page++
	}

	fmt.Println("total orders:", len(allOrders))
}
