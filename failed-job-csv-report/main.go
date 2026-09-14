package main

import (
	"fmt"
	"strings"
)

func extractField(line, key string) (string, bool) {
	for _, field := range strings.Fields(line) {
		if k, v, ok := strings.Cut(field, "="); ok && k == key {
			return v, true
		}
	}
	return "", false
}

func main() {
	jobLog := `2026-01-05T09:00:01 job=invoice-gen status=done duration=812
2026-01-05T09:00:04 job=report-mail status=failed duration=640
2026-01-05T09:00:09 job=invoice-gen status=failed duration=905
2026-01-05T09:00:12 job=db-cleanup status=done duration=430`

	var report strings.Builder
	report.WriteString("job,duration_ms\n")
	for _, line := range strings.Split(jobLog, "\n") {
		if !strings.Contains(line, "status=failed") {
			continue
		}

		job, jobOK := extractField(line, "job")
		duration, durationOK := extractField(line, "duration")

		if !jobOK || !durationOK {
			continue
		}
		report.WriteString(job)
		report.WriteString(",")
		report.WriteString(duration)
		report.WriteString("\n")
	}
	fmt.Println("failed jobs (CSV):")
	fmt.Print(strings.TrimSpace(report.String()))
}
