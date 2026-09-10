package main

import (
	"fmt"
)

func main() {
	a := [4]int{5, 10, 15, 20}

	for i, v := range a {
		fmt.Println(i, v)
	}
	fmt.Println("-----")
	for i := range a {
		a[i] *= 2
		fmt.Println(a[i])
	}
	fmt.Println("-----")
	for _, v := range a {
		fmt.Println(v)
	}
	fmt.Println("-----")
	for i := range len(a) {
		fmt.Println(i)
	}

}
