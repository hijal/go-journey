package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"time"
)

func main() {
	stock := map[string]int{"keyboard": 12, "mouse": 40}
	mux := http.NewServeMux()

	mux.HandleFunc("GET /stock/{item}", func(w http.ResponseWriter, r *http.Request) {
		item := r.PathValue("item")
		qty, ok := stock[item]

		if !ok {
			http.Error(w, "unknown item: "+item, http.StatusNotFound)
			return
		}

		fmt.Fprintf(w, "%s: %d in stock", item, qty)
	})

	mux.HandleFunc("POST /stock/{item}/order", func(w http.ResponseWriter, r *http.Request) {
		item := r.PathValue("item")
		qty, ok := stock[item]

		if !ok || qty == 0 {
			http.Error(w, "out of stock: "+item, http.StatusConflict)
			return
		}

		stock[item] = qty - 1
		fmt.Fprintf(w, "ordered 1 %s, remaining %d", item, qty-1)
	})

	server := &http.Server{Addr: "127.0.0.1:8085", Handler: mux}

	go server.ListenAndServe()

	time.Sleep(100 * time.Millisecond)

	fetch := func(method, url string) string {
		req, err := http.NewRequest(method, url, nil)

		if err != nil {
			return "build failed: " + err.Error()
		}

		resp, err := http.DefaultClient.Do(req)

		if err != nil {
			return "request failed: " + err.Error()
		}

		defer resp.Body.Close()
		body, err := io.ReadAll(resp.Body)

		if err != nil {
			return "read failed: " + err.Error()
		}
		return string(body)
	}

	fmt.Println("GET ->", fetch(http.MethodGet, "http://127.0.0.1:8085/stock/keyboard"))
	fmt.Println("POST ->", fetch(http.MethodPost, "http://127.0.0.1:8085/stock/keyboard/order"))
	fmt.Println("GET ->", fetch(http.MethodGet, "http://127.0.0.1:8085/stock/keyboard"))

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)

	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		fmt.Println("shutdown:", err)
	}
}
