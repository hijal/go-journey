package main

import (
	"fmt"
	"strings"
)

func JoinPath(parts ...string) string {
	var cleanParts []string

	for _, p := range parts {
		if p != "" {
			cleanParts = append(cleanParts, strings.Trim(p, "/"))
		}
	}
	return "/" + strings.Join(cleanParts, "/")
}

func main() {
	fmt.Println(JoinPath("/api/", "v1/", "/users", "101/"))
}
