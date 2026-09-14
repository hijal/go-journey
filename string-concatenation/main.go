package main

import (
	"fmt"
	"slices"
	"strings"
)

func main() {
	var b strings.Builder

	for range 3 {
		b.WriteString("go ")
	}
	fmt.Println("b:", b.String())

	words := []string{"banana", "Apple", "apple"}
	slices.Sort(words)
	fmt.Println("words:", words)
}
