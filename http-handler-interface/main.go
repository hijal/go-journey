package main

import (
	"encoding/json"
	"log/slog"
	"net/http"
)

type AccountService interface {
	Balance(id string) (int64, error)
}

type memService struct{}

func (memService) Balance(id string) (int64, error) {
	return 125_000, nil
}

type accountAPI struct {
	svc    AccountService
	logger *slog.Logger
}

func (a *accountAPI) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	cents, err := a.svc.Balance(id)

	if err != nil {
		a.logger.Error("balance lookup failed", "id", id, "err", err)
		http.Error(w, "service unavailable", http.StatusBadGateway)
	}

	a.logger.Info("balance served", "id", id)
	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(map[string]any{
		"id":            id,
		"balance_cents": cents,
	}); err != nil {
		a.logger.Error("encode response", "err", err)
	}
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("ok"))
}

func main() {
	logger := slog.Default()

	api := &accountAPI{svc: memService{}, logger: logger}
	mux := http.NewServeMux()
	mux.Handle("GET /balance", api)
	mux.HandleFunc("GET /health", healthHandler)
	logger.Info("listening on :8080")

	if err := http.ListenAndServe(":8080", mux); err != nil {
		logger.Error("server stopped", "err", err)
	}
}
