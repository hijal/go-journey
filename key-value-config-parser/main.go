package main

import (
	"fmt"
	"strings"
)

func main() {
	rawConfig := `
# app settings
server_port = 8080
  env = production 
`
	for _, line := range strings.Split(rawConfig, "\n") {
		line = strings.TrimSpace(line)

		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		key, value, found := strings.Cut(line, "=")

		if !found {
			fmt.Println("invalid line:", line)
			return
		}
		fmt.Printf("%s => %q\n", strings.TrimSpace(key), strings.TrimSpace(value))
	}
}
