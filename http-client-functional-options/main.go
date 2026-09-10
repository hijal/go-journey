package main

import (
	"fmt"
	"time"
)

type APIClient struct {
	BaseURL    string
	Timeout    time.Duration
	RetryCount int
	AuthToken  string
}

type Option func(*APIClient)

func WithTimeout(d time.Duration) Option {
	return func(a *APIClient) {
		a.Timeout = d
	}
}

func WithRetry(count int) Option {
	return func(a *APIClient) {
		a.RetryCount = count
	}
}

func WithAuth(token string) Option {
	return func(a *APIClient) {
		a.AuthToken = token
	}
}

func NewAPIClient(baseURL string, opts ...Option) *APIClient {
	client := &APIClient{
		BaseURL:    baseURL,
		Timeout:    30 * time.Second,
		RetryCount: 3,
	}

	for _, opt := range opts {
		opt(client)
	}
	return client
}

func main() {
	client := NewAPIClient(
		"https://api.example.com",
		WithTimeout(10*time.Second),
		WithAuth("super-secret-token"),
	)

	fmt.Printf("client configured for %s with %vs timeout\n", client.BaseURL, client.Timeout.Seconds())
}
