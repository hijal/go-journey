package main

import "fmt"

const (
	_  = iota             // 0 skipped
	KB = 1 << (10 * iota) // 1 << 10 = 1024
	MB = 1 << (10 * iota) // 1 << 20
	GB = 1 << (10 * iota) // 1 << 30
	TB = 1 << (10 * iota) // 1 << 40
	PB = 1 << (10 * iota) // 1 << 50
)

func main() {
	var maxLogSize = 12 * MB
	var uploadLimit = 250 * MB

	fmt.Printf("max log size: %d bytes (%d MB)\n", maxLogSize, maxLogSize/MB)
	fmt.Printf("upload limit: %d bytes (%d MB)\n", uploadLimit, uploadLimit/MB)
	fmt.Printf("1 TB = %d bytes\n", TB)
}
