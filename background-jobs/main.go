package main

import "fmt"

func processJob(jobID string) (bool, error) {
	if jobID == "" {
		return false, fmt.Errorf("empty job id")
	}
	return true, nil
}

func main() {
	jobIDs := []string{
		"job-101",
		"job-102",
		"job-103",
	}

	for _, id := range jobIDs {
		ok, err := processJob(id)

		if err != nil {
			fmt.Println("Failed:", err)
			continue
		}
		fmt.Println("Processed:", id, " success:", ok)
	}

	statusCount := map[string]int{
		"done":   3,
		"failed": 0,
	}

	if _, exists := statusCount["retrying"]; !exists {
		fmt.Println("No jobs are currently retrying")
	}
}
