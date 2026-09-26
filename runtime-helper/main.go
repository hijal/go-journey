package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("CPU cores:", runtime.NumCPU())
	fmt.Println("GOMAXPROCS:", runtime.GOMAXPROCS(0))
	fmt.Println("goroutines now:", runtime.NumGoroutine())
}
