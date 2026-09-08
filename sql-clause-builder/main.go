package main

import (
	"fmt"
	"strings"
)

func BuilderClause(column string, values ...any) (string, []any) {
	if len(values) == 0 {
		return "", nil
	}

	placeHolders := make([]string, len(values))

	for i := range values {
		placeHolders[i] = "?"
	}
	query := fmt.Sprintf("%s IN (%s)", column, strings.Join(placeHolders, ", "))
	return query, values
}

func main() {
	query, args := BuilderClause("user_id", 101, 102, 103)
	fmt.Println("Query:", query)
	fmt.Println("Args:", args)
}
