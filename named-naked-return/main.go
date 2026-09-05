package main

import (
	"fmt"
)

func rectangleInfo(w, h float64) (area, perimeter float64) {
	area = w * h
	perimeter = 2 * (w + h)
	return
}

func main() {
	area, perimeter := rectangleInfo(10, 15)

	fmt.Println("area:", area)
	fmt.Println("perimeter:", perimeter)
}
