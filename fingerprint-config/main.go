package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	config := `
server:
  port: 8080
  timeout: 30s
`
	rawBytes := []byte(config)

	var checksum [32]byte = sha256.Sum256(rawBytes)
	fmt.Printf("Config fingerprint: %x\n", checksum)

	rawBytes[1] = 'X'
	fmt.Println("Mutated bytes did not change the original string")
	fmt.Println("  original string still starts with:", config[:5])

	backToString := string(rawBytes)
	fmt.Println("Bytes converted back to string, first line differs:", backToString[:5])
}
