package main

import "fmt"

func main() {
	i := 42

	p := &i
	fmt.Println("*p:", *p)

	*p = 21
	fmt.Println("i:", i)

	j := 2701
	p = &j
	
	*p = *p / 37
	fmt.Println("j:", j)
}
