package main

import "fmt"

const (
	read    = 1 << 2
	write   = 1 << 1
	execute = 1 << 0
)

func main() {
	perm := read | write
	fmt.Println("permission:", perm)
	fmt.Println("can read:", perm&read != 0)
	fmt.Println("can execute:", perm&execute != 0)

	perm |= execute
	fmt.Println("after chmod +x:", perm, "can execute:", perm&execute != 0)

	perm &^= write
	fmt.Println("after chmod -w:", perm, "can write:", perm&write != 0)
}
