package main

import (
	"fmt"
	"strconv"
	"strings"
)

func validateSKU(sku string) error {
	parts := strings.Split(sku, "-")
	if len(parts) != 2 {
		return fmt.Errorf("SKU must look like CAT-NUMBER, got %q", sku)
	}

	category, idText := parts[0], parts[1]
	if len(category) != 3 {
		return fmt.Errorf("category must be 3 letters, got %q", category)
	}

	id, err := strconv.Atoi(idText)
	if err != nil {
		return fmt.Errorf("product number %q is not a valid integer", idText)
	}
	if id < 1 {
		return fmt.Errorf("product number must be positive, got %d", id)
	}
	return nil
}

func main() {
	for _, sku := range []string{"BOOK-1024", "TOY-7", "BAD", "PEN-abc"} {
		if err := validateSKU(sku); err != nil {
			fmt.Println(sku, "REJECTED:", err)
		} else {
			fmt.Println(sku, "ACCEPTED")
		}
	}
}
