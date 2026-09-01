package main

import "fmt"

func ipToUnit32(a, b, c, d byte) uint32 {
	return uint32(a)<<24 | uint32(b)<<16 | uint32(c)<<8 | uint32(d)
}

func main() {
	ip := ipToUnit32(192, 168, 1, 55)
	network := ipToUnit32(192, 168, 1, 0)
	mask := uint32(0xFFFFFF00)

	inSubnet := ip&mask == network&mask
	fmt.Printf("ip: %08x, masked: %08x\n", ip, ip&mask)
	fmt.Println("in subnet 192.168.1.0/24:", inSubnet)

	otherIP := ipToUnit32(192, 168, 7, 55)
	fmt.Println("192.168.7.55 in subnet:", otherIP&mask == network&mask)
}
