package main

import (
	"errors"
	"fmt"
)

type Timeout struct {
	Read  int
	Write int
}

type ServerConfig struct {
	Timeout
	Host string
	Port int
}

var ErrInvalidConfig = errors.New("invalid configs")

func NewServerConfig(host string, port, readTo, writeTo int) (*ServerConfig, error) {
	if host == "" {
		return nil, fmt.Errorf("%w: host is required", ErrInvalidConfig)
	}

	if port < 1 || port > 65535 {
		return nil, fmt.Errorf("%w: port %d out of range", ErrInvalidConfig, port)
	}

	return &ServerConfig{
		Timeout: Timeout{
			Read:  readTo,
			Write: writeTo,
		},
		Host: host,
		Port: port,
	}, nil
}

func main() {
	cfg, err := NewServerConfig("", 8080, 5, 10)
	if err != nil {
		fmt.Println("rejected:", err)
		fmt.Println("is config error:", errors.Is(err, ErrInvalidConfig))
		return
	}

	cfg, err = NewServerConfig("localhost", 8080, 5, 10)
	if err != nil {
		fmt.Println("unexpected:", err)
		return
	}

	fmt.Printf("listening on %s:%d (read %ds / write %ds)\n",
		cfg.Host, cfg.Port, cfg.Read, cfg.Write)
}
