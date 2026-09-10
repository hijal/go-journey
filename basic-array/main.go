package main

import (
	"fmt"
)

type Coordinate [2]float64

type Point struct {
	x, y int
}

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

	b := a

	b[0] = 100
	b[1] = 200
	b[2] = 300
	b[3] = 400
	fmt.Println("-----")
	for i, v := range b {
		fmt.Println(i, v)
	}
	fmt.Println("-----")
	for i, v := range a {
		fmt.Println(i, v)
	}
	fmt.Println("-----")
	var board [3][3]string
	board[1][1] = "X"
	fmt.Println(len(board), len(board[0]))

	var pickup Coordinate
	pickup[0] = 37.7749
	pickup[1] = -122.4194
	fmt.Println(pickup)

	arr := [2]int{1, 2}
	p := &arr
	p[0] = 10
	p[1] = 20
	fmt.Println(arr)

	arr1 := [4]int{1, 2, 3, 4}

	s := arr1[:]
	s[0] = 99
	fmt.Println(arr1)
	fmt.Println(s)

	pt := Point{x: 1, y: 2}
	fmt.Printf("value(%v): %v\n", pt, pt)
	fmt.Printf("Plus (%+v):   %+v\n", pt, pt)
	fmt.Printf("Syntax (%#v): %#v\n", pt, pt)
	fmt.Printf("Type (%T):    %T\n", pt, pt)
}
