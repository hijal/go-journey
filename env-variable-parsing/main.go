package main

import (
	"fmt"
	"strconv"
)

func main() {
	rawEnv := map[string]string{
		"APP_DEBUG":           "true",
		"MAX_WORKERS":         "12",
		"REQUEST_TIMEOUT_SEC": "",
	}

	var debugMode bool

	if v, ok := rawEnv["APP_DEBUG"]; ok {
		parsed, err := strconv.ParseBool(v)
		if err != nil {
			fmt.Println("invalid APP_DEBUG, keeping default false")
		} else {
			debugMode = parsed
		}
	}

	maxWorkers, err := strconv.Atoi(rawEnv["MAX_WORKERS"])
	if err != nil {
		maxWorkers = 4
	}

	var timeoutSec int
	if v, ok := rawEnv["REQUEST_TIMEOUT_SEC"]; ok && v != "" {
		timeoutSec, _ = strconv.Atoi(v)
	} else {
		timeoutSec = 30
	}

	fmt.Println("debugMode  :", debugMode)
	fmt.Println("maxWorkers :", maxWorkers)
	fmt.Println("timeoutSec :", timeoutSec)
}
