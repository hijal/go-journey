package main

import (
	"fmt"
	"strings"
)

var allowedExtensions = []string{".tar.gz", ".zip", ".log"}

func splitExt(name string) (string, string) {
	i := strings.LastIndex(name, ".")
	if i < 0 {
		return name, ""
	}
	return name[:i], name[i:]
}

func main() {
	uploaded := []string{"Backup JAN.tar.gz", "db Dump.ZIP", "notes.txt"}

	for _, f := range uploaded {
		clean := strings.ToLower(strings.ReplaceAll(f, " ", "_"))

		allowed := false

		for _, ext := range allowedExtensions {
			if strings.HasSuffix(clean, ext) {
				allowed = true
				break
			}
		}

		base, ext := splitExt(clean)

		if allowed {
			fmt.Println(clean, "-> ACCEPTED  (base:", base, "| ext:", ext+")")
		} else {
			fmt.Println(clean, "-> REJECTED (allowed:", strings.Join(allowedExtensions, ", ")+")")
		}
	}
}
