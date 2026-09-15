package main

import (
	"fmt"
	"io"
)

type LogCollector struct {
	lines []string
}

func (l *LogCollector) Write(p []byte) (int, error) {
	l.lines = append(l.lines, string(p))
	return len(p), nil
}

func main() {
	collector := &LogCollector{}

	var w io.Writer = collector
	fmt.Fprintln(w, "server started")
	fmt.Fprintln(w, "connection accepted")

	for i, line := range collector.lines {
		fmt.Printf("log[%d]: %s", i, line)
	}
}
