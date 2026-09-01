package main

import "fmt"

func pageCount(total, limit int) int {
	if limit <= 0 {
		return 0
	}
	return (total + limit - 1) / limit
}

func main() {
	totalProducts, limit, page := 95, 10, 3

	pages := pageCount(totalProducts, limit)
	offset := (page - 1) * limit

	fmt.Println("total pages:", pages)
	fmt.Println("page", page, "show items", offset+1, "to", min(offset+limit, totalProducts))
	fmt.Println("has next page:", page < pages)
	fmt.Println("has prev page:", page > 1)
}
