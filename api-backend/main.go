package main

import "fmt"

type PaymentAPIClient struct {
	BaseURL    string
	APIKey     string
	httpClient string
}

func NewPaymentAPIClient(baseURL, apiKey string) *PaymentAPIClient {
	return &PaymentAPIClient{
		BaseURL:    baseURL,
		APIKey:     apiKey,
		httpClient: "default-http-client",
	}
}

func (c *PaymentAPIClient) ID() string {
	return c.APIKey
}

func main() {
	client := NewPaymentAPIClient(
		"https://api.pay.example.com",
		"sk_xxxxxxxxxxxxxxxxxx",
	)

	fmt.Println("Base URL:", client.BaseURL)
	fmt.Println("Client ID:", client.ID())
}
