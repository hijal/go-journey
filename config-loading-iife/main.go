package main

import (
	"fmt"
	"os"
	"strconv"
)

type Config struct {
	port  int
	env   string
	dbURL string
	debug bool
}

var cfg = func() Config {
	port, err := strconv.Atoi(os.Getenv("APP_PORT"))

	if err != nil {
		port = 8080
	}

	env := os.Getenv("APP_ENV")
	if env == "" {
		env = "development"
	}

	dbURL := os.Getenv("APP_DB_URL")

	if dbURL == "" {
		dbURL = "postgres://localhost:5432/app"
	}

	return Config{
		port:  port,
		env:   env,
		dbURL: dbURL,
		debug: env == "development",
	}
}()

func main() {
	fmt.Printf("starting on port %d (%s)\n", cfg.port, cfg.env)
	fmt.Println("db:", cfg.dbURL)
	fmt.Println("debug mode:", cfg.debug)
}
