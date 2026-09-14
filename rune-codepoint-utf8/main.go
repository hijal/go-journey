package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "বই"
	fmt.Println(len(s))
	fmt.Println(utf8.RuneCountInString(s))

	for i, r := range s {
		fmt.Printf("byte of offset %d: %c (U+%04X)\n", i, r, r)
	}
}
