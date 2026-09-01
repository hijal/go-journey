package main

import (
	"cmp"
	"fmt"
)

func main() {
	yourScore, yourLevel := 1500, 9
	theirScore, theirLevel := 1500, 7

	rank := cmp.Compare(yourScore, theirScore)

	switch {
	case rank < 0:
		fmt.Println("they are ahead on score")
	case rank > 0:
		fmt.Println("you are ahead on score")
	default:
		levelRank := cmp.Compare(yourLevel, theirLevel)
		switch {
		case levelRank < 0:
			fmt.Println("they win on level tie-break")
		case levelRank > 0:
			fmt.Println("you win on level tie-break")
		default:
			fmt.Println("perfect tie - share first place")
		}
	}
}
