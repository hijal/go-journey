package main

import (
	"fmt"
	"log/slog"
	"os"
)

type Secret string

const redacted = "[REDACTED]"

func (Secret) String() string {
	return redacted
}

func (Secret) GoString() string {
	return `Secret("` + redacted + `")`
}

func (Secret) LogValue() slog.Value {
	return slog.StringValue(redacted)
}

func (s Secret) Reveal() string {
	return string(s)
}

type DatabaseConfig struct {
	Host     string
	Port     int
	User     string
	Password Secret
}

func (c DatabaseConfig) LogValue() slog.Value {
	return slog.GroupValue(
		slog.String("host", c.Host),
		slog.Int("port", c.Port),
		slog.String("user", c.User),
		slog.Any("password", c.Password),
	)
}

func (c DatabaseConfig) DSN() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%d/app", c.User, c.Password.Reveal(), c.Host, c.Port)
}

func main() {
	cfg := DatabaseConfig{Host: "db.internal", Port: 5432, User: "app", Password: "abc123"}

	fmt.Printf("%v\n", cfg)
	fmt.Printf("%+v\n", cfg)
	fmt.Printf("%#v\n", cfg)

	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		ReplaceAttr: func(_ []string, a slog.Attr) slog.Attr {
			if a.Key == slog.TimeKey {
				return slog.Attr{}
			}
			return a
		},
	}))

	logger.Info("connecting to database", "db", cfg)

	fmt.Println("dsn length:", len(cfg.DSN()))
}
