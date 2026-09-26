# bounded-concurrent-http

Go-তে **semaphore + per-request timeout দিয়ে bounded HTTP fan-out** — 2-concurrent fetch, ধীর endpoint timeout-এ FAILED।

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

### Lines 3–13

```go
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
```

`httptest` (local test server) + `net/http` + JSON।

### Lines 15–18

```go
type stockResponse struct {
	SKU   string `json:"sku"`
	Units int    `json:"units"`
}
```

JSON payload struct।

### Lines 21–32

```go
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		sku := r.URL.Query().Get("sku")
		if sku == "SKU-SLOW" {
			time.Sleep(500 * time.Millisecond)
		}
		...
```

**Test server** — `SKU-SLOW` endpoint 500ms ঘুমায় (real-world slow upstream simulation); বাকিটা JSON ফেরায়।

### Lines 40–45

```go
	client := &http.Client{}
	sem := make(chan struct{}, 2)

	var wg sync.WaitGroup

	for i, sku := range skus {
		wg.Go(func() {
			sem <- struct{}{}
			defer func() { <-sem }()
```

**Semaphore cap 2** — যেকোনো-মুহূর্তে সর্বোচ্চ 2টা request।

### Lines 50–65

```go
			ctx, cancel := context.WithTimeout(context.Background(), 200*time.Millisecond)
			defer cancel()

			url := fmt.Sprintf("%s/stock?sku=%s", srv.URL, sku)
			req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
			...
			resp, err := client.Do(req)
			if err != nil {
				errs[i] = fmt.Errorf("fetch %s: %w", sku, err)
				return
			}
```

**Per-request 200ms ctx** — `SKU-SLOW` (500ms) → `context deadline exceeded`।

### Lines 87–94

```go
	for i, sku := range skus {
		if errs[i] != nil {
			fmt.Printf("%-9s FAILED (timeout=%t): %v\n",
				sku, errors.Is(errs[i], context.DeadlineExceeded), errs[i])
			continue
		}
		fmt.Printf("%-9s %d units\n", sku, units[i])
	}
```

Per-index error report — timeout-বনাম-অন্য error আলাদা।

---

## Expected Output

```
SKU-1001  24 units
SKU-1002  24 units
SKU-SLOW  FAILED (timeout=true): fetch SKU-SLOW: Get "http://127.0.0.1:PORT/stock?sku=SKU-SLOW": context deadline exceeded
SKU-1004  24 units
SKU-1005  24 units
```

## মূল শিক্ষা / Key Takeaways

1. **Bounded concurrency** — semaphore সর্বোচ্চ-চাপ নিয়ন্ত্রণে।
2. **Per-request timeout** — ধীর upstream-এর জন্য সুরক্ষা।
3. **`NewRequestWithContext`** — cancel-ও-সচেতন HTTP call।
4. **`errors.Is(DeadlineExceeded)`** — নির্দিষ্ট ফেল-রূপ।
5. **Index-aligned errs** — goroutine-সেফ ফলাফল।

---

---

<a name="english"></a>

## 🇬🇧 English Version

### Line 1

```go
package main
```

Declares an executable program (`main` package), runnable via `go run`.

### Lines 3–13

```go
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
```

`httptest` (local test server) + `net/http` + JSON.

### Lines 15–18

```go
type stockResponse struct {
	SKU   string `json:"sku"`
	Units int    `json:"units"`
}
```

The JSON payload struct.

### Lines 21–32

```go
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		sku := r.URL.Query().Get("sku")
		if sku == "SKU-SLOW" {
			time.Sleep(500 * time.Millisecond)
		}
		...
```

**Test server** — the `SKU-SLOW` endpoint sleeps 500ms (a simulated slow upstream); everything else returns JSON.

### Lines 40–45

```go
	client := &http.Client{}
	sem := make(chan struct{}, 2)

	var wg sync.WaitGroup

	for i, sku := range skus {
		wg.Go(func() {
			sem <- struct{}{}
			defer func() { <-sem }()
```

**Semaphore cap 2** — at most 2 requests in flight at any moment.

### Lines 50–65

```go
			ctx, cancel := context.WithTimeout(context.Background(), 200*time.Millisecond)
			defer cancel()

			url := fmt.Sprintf("%s/stock?sku=%s", srv.URL, sku)
			req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
			...
			resp, err := client.Do(req)
			if err != nil {
				errs[i] = fmt.Errorf("fetch %s: %w", sku, err)
				return
			}
```

**Per-request 200ms ctx** — `SKU-SLOW` (500ms) → `context deadline exceeded`.

### Lines 87–94

```go
	for i, sku := range skus {
		if errs[i] != nil {
			fmt.Printf("%-9s FAILED (timeout=%t): %v\n",
				sku, errors.Is(errs[i], context.DeadlineExceeded), errs[i])
			continue
		}
		fmt.Printf("%-9s %d units\n", sku, units[i])
	}
```

Per-index error reporting — timeouts distinguished from other failures.

---

## Expected Output

```
SKU-1001  24 units
SKU-1002  24 units
SKU-SLOW  FAILED (timeout=true): fetch SKU-SLOW: Get "http://127.0.0.1:PORT/stock?sku=SKU-SLOW": context deadline exceeded
SKU-1004  24 units
SKU-1005  24 units
```

## Key Takeaways

1. **Bounded concurrency** — the semaphore keeps peak pressure in check.
2. **Per-request timeout** — defends against a slow upstream.
3. **`NewRequestWithContext`** — cancel-aware HTTP calls.
4. **`errors.Is(DeadlineExceeded)`** — names the failure kind.
5. **Index-aligned errs** — goroutine-safe result collection.