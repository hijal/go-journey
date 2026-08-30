package main

import (
	"fmt"
	"unicode/utf8"
)

func validateUsername(name string) error {
	const maxChars = 15
	byteLen := len(name)
	runeLen := utf8.RuneCountInString(name)
	fmt.Printf("%q -> bytes=%d, runes=%d\n", name, byteLen, runeLen)

	if runeLen > maxChars {
		return fmt.Errorf("username %q has %d characters, max is %d", name, runeLen, maxChars)
	}

	return nil
}

func main() {
	names := []string{
		"GoLang",
		"বাংলা",
		"José_García",
	}

	for _, name := range names {
		if err := validateUsername(name); err != nil {
			fmt.Println(" ->", err)
		} else {
			fmt.Println(" -> accepted")
		}
	}

	fmt.Println("\nDecoding GoLang rune by rune:")
	for i, r := range "GoLang" {
		fmt.Printf("  byte offset %2d -> rune %q (code point U+%04X)\n", i, r, r)
	}

	fmt.Println("\nDecoding বাংলা rune by rune:")
	for i, r := range "বাংলা" {
		fmt.Printf("  byte offset %2d -> rune %q (code point U+%04X)\n", i, r, r)
	}

	fmt.Println("\nDecoding José_García rune by rune:")
	for i, r := range "José_García" {
		fmt.Printf("  byte offset %2d -> rune %q (code point U+%04X)\n", i, r, r)
	}
}
