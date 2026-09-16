# http-handler-interface

Go-তে **`http.Handler` interface (`ServeHTTP`) + `http.ServeMux` + pattern-based route + `slog` structured logging** শেখার ছোট example — balance service API।

**📖 ভাষা নির্বাচন করুন / Choose language:**

[🇧🇩 বাংলা](#bangla) • [🇬🇧 English](#english)

---

<a name="bangla"></a>

## 🇧🇩 বাংলা সংস্করণ

### Line 1

```go
package main
```

একটা executable program (`main` package) declare করে, যা `go run` দিয়ে চালানো যায়।

### Lines 3–7

```go
import (
	"encoding/json"
	"log/slog"
	"net/http"
)
```

- `json` — `NewEncoder`।
- `log/slog` — structured logging।
- `net/http` — handlers, ServeMux।

### Lines 9–17

```go
type AccountService interface {
	Balance(id string) (int64, error)
}

type memService struct{}

func (memService) Balance(id string) (int64, error) {
	return 125_000, nil
}
```

**Service abstraction** — interface + dummy in-memory implementation।

### Lines 19–42

```go
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
```

**`http.Handler` interface** — `ServeHTTP(w, r)` method-ই contract (সব handler type-তে)। Handler-এ:

- `r.URL.Query().Get("id")` — query param।
- svc call → error হলে 502 + log।
- `slog` structured log + JSON encode (`json.NewEncoder(w).Encode`)।
- `map[string]any` — dynamic JSON object।

### Lines 44–47

```go
func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("ok"))
}
```

Plain `func(w, r)` handler — Function-as-Handler।

### Lines 49–61

```go
	logger := slog.Default()

	api := &accountAPI{svc: memService{}, logger: logger}
	mux := http.NewServeMux()
	mux.Handle("GET /balance", api)
	mux.HandleFunc("GET /health", healthHandler)
	logger.Info("listening on :8080")

	if err := http.ListenAndServe(":8080", mux); err != nil {
		logger.Error("server stopped", "err", err)
	}
```

- `http.NewServeMux` — router।
- `mux.Handle("GET /balance", api)` — **method+path pattern** + `Handler` (interface)।
- `mux.HandleFunc("GET /health", ...)` — func-version।
- `http.ListenAndServe(":8080", mux)` — serve (error → log)।

---

## Expected Output

**`curl /balance?id=u1`:**

```
{"balance_cents":125000,"id":"u1"}
```

**`curl /health`:**

```
ok
```

**Server log (slog):**

```
INFO listening on :8080
INFO balance served id=u1
```

## মূল শিক্ষা / Key Takeaways

1. **`http.Handler`** — `ServeHTTP(w,r)` = one-method interface (spelling penting!)।
2. **`ServeMux`** — method+path pattern route।
3. **`Handle` vs `HandleFunc`** — interface vs function handler।
4. **`json.NewEncoder`** — direct-to-writer JSON।
5. **`slog`** — structured key-value logging।

---

---

<a name="english"></a>

## 🇬🇧 English Version

### Line 1

```go
package main
```

Declares an executable program (`main` package), runnable via `go run`.

### Lines 3–7

```go
import (
	"encoding/json"
	"log/slog"
	"net/http"
)
```

- `json` — for `NewEncoder`.
- `log/slog` — structured logging.
- `net/http` — handlers, ServeMux.

### Lines 9–17

```go
type AccountService interface {
	Balance(id string) (int64, error)
}

type memService struct{}

func (memService) Balance(id string) (int64, error) {
	return 125_000, nil
}
```

**The service abstraction** — an interface + a dummy in-memory implementation.

### Lines 19–42

```go
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
```

**The `http.Handler` interface** — the `ServeHTTP(w, r)` method is the contract (every handler type fits). Inside:

- `r.URL.Query().Get("id")` — a query param.
- svc call → on error: 502 + log.
- `slog` structured logging + JSON encode (`json.NewEncoder(w).Encode`).
- `map[string]any` — a dynamic JSON object.

### Lines 44–47

```go
func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("ok"))
}
```

A plain `func(w, r)` handler — Function-as-Handler.

### Lines 49–61

```go
	logger := slog.Default()

	api := &accountAPI{svc: memService{}, logger: logger}
	mux := http.NewServeMux()
	mux.Handle("GET /balance", api)
	mux.HandleFunc("GET /health", healthHandler)
	logger.Info("listening on :8080")

	if err := http.ListenAndServe(":8080", mux); err != nil {
		logger.Error("server stopped", "err", err)
	}
```

- `http.NewServeMux` — the router.
- `mux.Handle("GET /balance", api)` — a **method+path pattern** + a `Handler` (interface).
- `mux.HandleFunc("GET /health", ...)` — the func variant.
- `http.ListenAndServe(":8080", mux)` — serve (errors logged).

---

## Expected Output

**`curl /balance?id=u1`:**

```
{"balance_cents":125000,"id":"u1"}
```

**`curl /health`:**

```
ok
```

**Server log (slog):**

```
INFO listening on :8080
INFO balance served id=u1
```

## Key Takeaways

1. **`http.Handler`** — `ServeHTTP(w,r)` is a one-method interface (mind the spelling!).
2. **`ServeMux`** — method+path pattern routing.
3. **`Handle` vs `HandleFunc`** — interface vs function handlers.
4. **`json.NewEncoder`** — JSON straight to the writer.
5. **`slog`** — structured key-value logging.