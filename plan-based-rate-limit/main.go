package main

import (
	"fmt"
	"maps"
	"slices"
)

func main() {
	rateLimits := map[string]int{
		"free":       100,
		"pro":        1000,
		"enterprise": 10000,
	}

	requestedPlan := "pro"

	limit, exist := rateLimits[requestedPlan]

	if !exist {
		fmt.Println("unknown plan:", requestedPlan)
		return
	}
	fmt.Println(requestedPlan, "plan allows", limit, "requests/hour")

	rateLimits["pro"] = 1500
	delete(rateLimits, "free")

	for _, plan := range slices.Sorted(maps.Keys(rateLimits)) {
		fmt.Println(plan, "->", rateLimits[plan])
	}
}
