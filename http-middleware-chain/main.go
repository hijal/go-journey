package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"time"
)

type healthHandler struct {
	version string
}

func (h healthHandler) ServeHTTP(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, `{"status":"ok","version":%q}`, h.version)
}

type Middleware func(http.Handler) http.Handler

type statusRecorder struct {
	http.ResponseWriter
	status int
}

func (r *statusRecorder) WriteHeader(code int) {
	r.status = code
	r.ResponseWriter.WriteHeader(code)
}

func Logging(logger *log.Logger) Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			rec := &statusRecorder{ResponseWriter: w, status: http.StatusOK}
			next.ServeHTTP(rec, r)

			logger.Printf("%s %s -> %d (%s)", r.Method, r.URL.Path, rec.status, time.Since(start).Round(time.Millisecond))
		})
	}
}

func RequireAPIKey(key string) Middleware {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Header.Get("X-API-Key") != key {
				http.Error(w, "invalid API key", http.StatusUnauthorized)
				return
			}
			h.ServeHTTP(w, r)
		})
	}
}

func Chain(h http.Handler, mws ...Middleware) http.Handler {
	for i := len(mws) - 1; i >= 0; i-- {
		h = mws[i](h)
	}
	return h
}

func main() {
	logger := log.New(os.Stdout, "[http] ", 0)

	mux := http.NewServeMux()

	mux.Handle("/health", healthHandler{version: "1.1.2"})
	mux.HandleFunc("/admin/reports", func(w http.ResponseWriter, _ *http.Request) {
		io.WriteString(w, "quarterly report")
	})

	app := Chain(mux, Logging(logger), RequireAPIKey("s3cr3t"))

	for _, tc := range []struct{ path, key string }{
		{"/health", "s3cr3t"},
		{"/admin/reports", "wrong"},
		{"/missing", "s3cr3t"},
	} {
		req := httptest.NewRequest(http.MethodGet, tc.path, nil)
		req.Header.Set("X-API-Key", tc.key)
		rec := httptest.NewRecorder()
		app.ServeHTTP(rec, req)
		fmt.Printf("  body=%q\n", rec.Body.String())
	}
}
