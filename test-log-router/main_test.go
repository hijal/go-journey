package main

import "testing"

func TestLogDestination(t *testing.T) {
	tests := []struct {
		name  string
		level string
		want  string
	}{
		{name: "debug goes to local file", level: "DEBUG", want: "local file (rotated daily)"},
		{name: "trace groups with debug", level: "TRACE", want: "local file (rotated daily)"},
		{name: "info goes to stdout", level: "INFO", want: "stdout"},
		{name: "error triggers alerting", level: "ERROR", want: "stdout + alerting pipeline"},
		{name: "fatal pages on-call", level: "FATAL", want: "stdout + alerting pipeline + pager"},
		{name: "unknown level falls to default", level: "VERBOSE", want: "unknown level: drop and log a config warning"},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := logDestination(tc.level)
			if got != tc.want {
				t.Errorf("logDestination(%q) = %q, want %q", tc.level, got, tc.want)
			}
		})
	}
}
