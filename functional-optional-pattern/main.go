package main

import "fmt"

type Server struct {
	Port    int
	Timeout int
	IsHTTPS bool
}

type Option func(*Server)

func WithPort(port int) Option {
	return func(s *Server) {
		s.Port = port
	}
}

func WithTimeout(t int) Option {
	return func(s *Server) {
		s.Timeout = t
	}
}

func WithHTTPS(isHTTPS bool) Option {
	return func(s *Server) {
		s.IsHTTPS = isHTTPS
	}
}

func NewServer(opts ...Option) *Server {
	s := &Server{
		Port:    8080,
		Timeout: 30,
	}
	for _, opt := range opts {
		opt(s)
	}
	return s
}

func main() {
	srv := NewServer(WithPort(9090), WithHTTPS(true), WithTimeout(50))
	fmt.Printf("%+v\n", srv)

	srv = NewServer()
	fmt.Printf("%+v\n", srv)

	srv = NewServer(WithTimeout(100))
	fmt.Printf("%+v\n", srv)
}
