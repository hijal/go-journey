package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
)

type Config struct {
	Port   int    `json:"port"`
	Env    string `json:"env"`
	DbHost string `json:"db_host"`
}

func configPath() string {
	_, file, _, ok := runtime.Caller(0)
	if !ok {
		return "config.json"
	}
	return filepath.Join(filepath.Dir(file), "config.json")
}

func ParseConfig(filePath string) (*Config, error) {
	data, err := os.ReadFile(filePath)

	if err != nil {
		return nil, fmt.Errorf("failed to read file %s: %w", filePath, err)
	}

	var cfg Config

	if err := json.Unmarshal(data, &cfg); err != nil {
		return nil, fmt.Errorf("failed to parse JSON: %w", err)
	}

	return &cfg, nil
}

func main() {
	cfg, err := ParseConfig(configPath())

	if err != nil {
		fmt.Printf("fatal error loading config: %v\n", err)
		return
	}

	fmt.Printf("loaded config for env: %v\n", cfg)
}
