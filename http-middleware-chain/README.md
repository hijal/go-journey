# http-middleware-chain

Go-তে **HTTP middleware chain (`Middleware func(http.Handler) http.Handler` + `Chain`)** শেখার ছোট example — logging + API-key + recorder।

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

### Lines 3–11

```go
import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"time"
)
```

- `httptest` — test request/response recorder।
- `io` — `WriteString`।
- `time` — timing।

### Lines 13–20

```go
type healthHandler struct {
	version string
}

func (h healthHandler) ServeHTTP(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, `{"status":"ok","version":%q}`, h.version)
}
```

**Handler struct** — `ServeHTTP`-এ valid JSON (struct-মধ্যস্থ version)।

### Line 22

```go
type Middleware func(http.Handler) http.Handler
```

**Middleware contract** — একটা handler নিয়ে, একটা handler দেয়।

### Lines 24–32

```go
type statusRecorder struct {
	http.ResponseWriter
	status int
}

func (r *statusRecorder) WriteHeader(code int) {
	r.status = code
	r.ResponseWriter.WriteHeader(code)
}
```

**ResponseWriter wrapper (embedding)** — `WriteHeader` intercept করে status capture; বাকিটা inner `ResponseWriter` থেকে। (নয়-pass status)।

### Lines 34–44

```go
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
```

**Factory → closure** — `Logging(logger)` দেয় middleware; inner chain + সময় measure + status log।

### Lines 46–56

```go
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
```

**Guard middleware** — key না মিললে 401 + early-return (chain আর এগোয় না)।

### Lines 58–63

```go
func Chain(h http.Handler, mws ...Middleware) http.Handler {
	for i := len(mws) - 1; i >= 0; i-- {
		h = mws[i](h)
	}
	return h
}
```

**Reverse wrap** — শেষ middleware আগে; প্রথম (Logging) সর্ব-বাইরে (execution order ঠিক রাখতে)।

### Lines 65–87

```go
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
```

- `app := Chain(mux, Logging, RequireAPIKey)` — chain: Logging → RequireAPIKey → mux।
- ৩টা case: success / 401 guard / 404 mux।

---

## Expected Output

```
[http] GET /health -> 200 (0s)
  body="{\"status\":\"ok\",\"version\":\"1.1.2\"}"
[http] GET /admin/reports -> 401 (0s)
  body="invalid API key\n"
[http] GET /missing -> 404 (0s)
  body="404 page not found\n"
```

## মূল শিক্ষা / Key Takeaways

1. **`Middleware func(http.Handler) http.Handler`** — স্ট্যান্ডার্ড composable contract।
2. **`Chain` reverse-wrap** — execution order উল্টো wrap-এ ঠিক।
3. **ResponseWriter embedding + override** — status capture।
4. **Factory-returning closure** — logger/key-এ parameterized middleware।
5. **Guard + early-return** — chain short-circuit।

---

---

<a name="english"></a>

## 🇬🇧 English Version

### Line 1

```go
package main
```

Declares an executable program (`main` package), runnable via `go run`.

### Lines 3–11

```go
import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"time"
)
```

- `httptest` — test requests/response records.
- `io` — `WriteString`.
- `time` — timing.

### Lines 13–20

```go
type healthHandler struct {
	version string
}

func (h healthHandler) ServeHTTP(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, `{"status":"ok","version":%q}`, h.version)
}
```

**A handler struct** — valid JSON in `ServeHTTP` (version embedded in the struct).

### Line 22

```go
type Middleware func(http.Handler) http.Handler
```

**The middleware contract** — takes a handler, returns a handler.

### Lines 24–32

```go
type statusRecorder struct {
	http.ResponseWriter
	status int
}

func (r *statusRecorder) WriteHeader(code int) {
	r.status = code
	r.ResponseWriter.WriteHeader(code)
}
```

**A ResponseWriter wrapper (embedding)** — intercepts `WriteHeader` to capture the status; everything else comes from the inner `ResponseWriter`. (Pass-by-pass status).

### Lines 34–44

```go
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
```

**Factory → closure** — `Logging(logger)` returns a middleware; the inner chain + timing + status log.

### Lines 46–56

```go
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
```

**A guard middleware** — a bad key → 401 + early return (the chain stops).

### Lines 58–63

```go
func Chain(h http.Handler, mws ...Middleware) http.Handler {
	for i := len(mws) - 1; i >= 0; i-- {
		h = mws[i](h)
	}
	return h
}
```

**Reverse wrapping** — the last middleware wraps first; the first (Logging) ends up outermost (keeping execution order correct).

### Lines 65–87

```go
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
```

- `app := Chain(mux, Logging, RequireAPIKey)` — the chain: Logging → RequireAPIKey → mux.
- 3 cases: success / the 401 guard / the mux 404.

---

## Expected Output

```
[http] GET /health -> 200 (0s)
  body="{\"status\":\"ok\",\"version\":\"1.1.2\"}"
[http] GET /admin/reports -> 401 (0s)
  body="invalid API key\n"
[http] GET /missing -> 404 (0s)
  body="404 page not found\n"
```

## Key Takeaways

1. **`Middleware func(http.Handler) http.Handler`** — the standard composable contract.
2. **`Chain` reverse-wrapping** — correct execution order via reversed wrapping.
3. **ResponseWriter embedding + override** — status capture.
4. **Factory-returning closures** — logger/key-parameterized middleware.
5. **Guard + early return** — chain short-circuiting.