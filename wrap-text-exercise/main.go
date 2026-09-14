package main

import (
	"fmt"
	"strings"
)

func wrapText(text string, width int) string {
	var b strings.Builder

	lineLen := 0

	for _, word := range strings.Fields(text) {
		switch {
		case lineLen == 0:
			b.WriteString(word)
			lineLen = len(word)
		case lineLen+1+len(word) <= width:
			b.WriteString(" ")
			b.WriteString(word)
			lineLen += 1 + len(word)
		default:
			b.WriteString("\n")
			b.WriteString(word)
			lineLen = len(word)
		}
	}

	return b.String()
}

func main() {
	fmt.Println(wrapText("go is simple but strings are deep", 12))
}
