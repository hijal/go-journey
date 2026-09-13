package main

import "fmt"

func pageItems(all []int, page, perPage int) []int {
	start := (page - 1) * perPage
	if start >= len(all) || start < 0 {
		return []int{}
	}

	end := start + perPage
	if end > len(all) {
		end = len(all)
	}

	return all[start:end]
}

func main() {
	catalog := []int{}

	for id := range 25 {
		catalog = append(catalog, 1000+id)
	}

	page1 := pageItems(catalog, 1, 10)
	page3 := pageItems(catalog, 3, 10)
	page9 := pageItems(catalog, 9, 10)

	fmt.Println("page 1:", page1)
	fmt.Println("page 3:", page3)
	fmt.Println("page 9:", page9, "(empty, out of range)")
}
