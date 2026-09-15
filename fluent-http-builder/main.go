package main

import (
	"fmt"
	"time"
)

type HTTPClientBuilder struct {
	baseURL string
	timeout time.Duration
	headers map[string]string
}

func NewHTTPClientBuilder() *HTTPClientBuilder {
	return &HTTPClientBuilder{
		headers: make(map[string]string),
	}
}

func (b *HTTPClientBuilder) BaseURL(url string) *HTTPClientBuilder {
	b.baseURL = url
	return b
}

func (b *HTTPClientBuilder) Timeout(d time.Duration) *HTTPClientBuilder {
	b.timeout = d
	return b
}

func (b *HTTPClientBuilder) Header(key, value string) *HTTPClientBuilder {
	b.headers[key] = value
	return b
}

func (b *HTTPClientBuilder) Build() string {
	return fmt.Sprintf("Client{baseURL=%s, timeout=%s, headers=%v}", b.baseURL, b.timeout, b.headers)
}

func main() {
	client := NewHTTPClientBuilder().
		BaseURL("https://api.example.com").
		Timeout(10*time.Second).
		Header("Authorization", "Bearer token123").
		Build()
	fmt.Println(client)
}
