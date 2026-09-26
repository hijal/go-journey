package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"sync"
	"time"
)

type stockResponse struct {
	SKU   string `json:"sku"`
	Units int    `json:"units"`
}

func main() {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		sku := r.URL.Query().Get("sku")
		if sku == "SKU-SLOW" {
			time.Sleep(500 * time.Millisecond)
		}

		w.Header().Set("Content-Type", "application/json")

		if err := json.NewEncoder(w).Encode(stockResponse{SKU: sku, Units: len(sku) * 3}); err != nil {
			slog.Error("encode stock response", "sku", sku, "err", err)
		}
	}))

	defer srv.Close()

	skus := []string{"SKU-1001", "SKU-1002", "SKU-SLOW", "SKU-1004", "SKU-1005"}
	units := make([]int, len(skus))
	errs := make([]error, len(skus))

	client := &http.Client{}
	sem := make(chan struct{}, 2)

	var wg sync.WaitGroup

	for i, sku := range skus {
		wg.Go(func() {
			sem <- struct{}{}

			defer func() { <-sem }()
			ctx, cancel := context.WithTimeout(context.Background(), 200*time.Millisecond)
			defer cancel()

			url := fmt.Sprintf("%s/stock?sku=%s", srv.URL, sku)
			req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)

			if err != nil {
				errs[i] = fmt.Errorf("build request %s: %w", sku, err)
				return
			}

			resp, err := client.Do(req)
			if err != nil {
				errs[i] = fmt.Errorf("fetch %s: %w", sku, err)
				return
			}

			defer resp.Body.Close()

			if resp.StatusCode != http.StatusOK {
				errs[i] = fmt.Errorf("fetch %s: unexpected status %s", sku, resp.Status)
				return
			}

			var out stockResponse

			if err := json.NewDecoder(resp.Body).Decode(&out); err != nil {
				errs[i] = fmt.Errorf("decode %s: %w", sku, err)
				return
			}

			units[i] = out.Units
		})
	}

	wg.Wait()

	for i, sku := range skus {
		if errs[i] != nil {
			fmt.Printf("%-9s FAILED (timeout=%t): %v\n",
				sku, errors.Is(errs[i], context.DeadlineExceeded), errs[i])
			continue
		}
		fmt.Printf("%-9s %d units\n", sku, units[i])
	}
}
