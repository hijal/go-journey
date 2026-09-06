package main

import (
	"fmt"
	"strings"
)

func buildPipeline() []func(string) string {
	trim := strings.TrimSpace

	removePlaceholder := func(s string) string {
		return strings.ReplaceAll(s, "[TICKET]", "")
	}

	collapseSpaces := func(s string) string {
		return strings.Join(strings.Fields(s), " ")
	}

	return []func(string) string{
		trim, removePlaceholder, collapseSpaces,
	}
}

func processTicket(raw string, steps []func(string) string) string {
	text := raw
	for _, step := range steps {
		text = step(text)
	}
	return text
}

func main() {
	raw := "  [TICKET]   My   payment   failed   twice  "

	steps := buildPipeline()
	fmt.Printf("raw: %q\n", raw)
	fmt.Printf("clean: %q\n", processTicket(raw, steps))
}
