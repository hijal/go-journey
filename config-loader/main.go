package main

import "fmt"

func main() {
	supportedRegions := []string{
		"BD",
		"IN",
		"SG",
	}

	regionCount := len(supportedRegions)

	fmt.Println("region count:", regionCount)

	{
		len := regionCount

		fmt.Println("Shadowed len value:", len)
	}

	fmt.Println("builtin len still works:", len(supportedRegions))
}
