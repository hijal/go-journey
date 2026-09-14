package main

import "fmt"

func main() {
	var flags map[string]bool

	if enabled, ok := flags["new-checkout"]; ok {
		fmt.Println("new-checkout configured:", enabled)
	} else {
		fmt.Println("new-checkout: no flag configured yet")
	}

	flags = make(map[string]bool)

	flags["new-checkout"] = true
	flags["dark-mode"] = true
	flags["legacy-api"] = false

	fmt.Println("dark-mode enabled?  ", flags["dark-mode"])
	fmt.Println("legacy-api enabled? ", flags["legacy-api"])
	fmt.Println("beta-search enabled?", flags["beta-search"])

	if v, ok := flags["legacy-api"]; ok {
		fmt.Println("legacy-api explicitly configured as", v)
	}

	delete(flags, "legacy-api")
	fmt.Println("Flags remaining:", len(flags))
}
