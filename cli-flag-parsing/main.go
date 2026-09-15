package main

import (
	"flag"
	"fmt"
)

func main() {
	verbose := flag.Bool("verbose", false, "enable verbose logging")
	workers := flag.Int("workers", 4, "number of worker goroutines")
	output := flag.String("output", "stdout", "output destination")
	flag.Parse()

	if *verbose {
		fmt.Println("Verbose mode enabled")
	}

	fmt.Printf("starting %d workers, writing to %s\n", *workers, *output)
}
