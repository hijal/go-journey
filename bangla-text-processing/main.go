package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	sms := "গো শিখি"
	byteLen := len(sms)
	runes := utf8.RuneCountInString(sms)

	fmt.Printf("bytes: %d, characters: %d\n", byteLen, runes)

	for i, r := range sms {
		fmt.Printf("byte position %d: %c\n", i, r)
	}
}
