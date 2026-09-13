package main

import (
	"fmt"
	"slices"
)

type Job struct {
	ID      string
	Payload string
	Failed  bool
}

func main() {
	queue := []Job{
		{"J1", "send-email:1123", false},
		{"J2", "resize-image:8891", true},
		{"J3", "charge-card:5521", false},
		{"J4", "generate-pdf:3320", true},
		{"J5", "sync-inventory:7788", false},
	}

	queue = slices.DeleteFunc(queue, func(j Job) bool {
		return j.Failed
	})

	fmt.Println("remaining in queue:")
	for _, j := range queue {
		fmt.Printf("  %s -> %s\n", j.ID, j.Payload)
	}

	failedIDs := []string{"J2", "J4"}
	fmt.Println("to retry later:", failedIDs)
}
