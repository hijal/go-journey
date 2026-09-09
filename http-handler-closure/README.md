# http-handler-closure

Go-তে **closure over a map** দিয়ে HTTP handlers-এ state শেখার ছোট example — `net/http` mux, `r.PathValue`, `http.Server` + `Shutdown`।

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

### Lines 3–9

```go
import (
	"context"
	"fmt"
	"io"
	"net/http"
	"time"
)
```

- `context` — `context.WithTimeout` (shutdown-এর জন্য)।
- `fmt` — `Fprintf`, `Println`।
- `io` — `io.ReadAll`।
- `net/http` — HTTP server/client।
- `time` — sleep, timeout।

### Line 11

```go
func main() {
```

Program-এর entry point।

### Line 12

```go
stock := map[string]int{"keyboard": 12, "mouse": 40}
```

**State map** — `stock`: item → available qty। এই map-টা handlers-এর **closure-এ capture** — handlers dispatch-এর সময় একই state-এ access।

### Line 13

```go
mux := http.NewServeMux()
```

**`http.NewServeMux()`** — একটা multiplexer (router)।

### Lines 15–25

```go
mux.HandleFunc("GET /stock/{item}", func(w http.ResponseWriter, r *http.Request) {
	item := r.PathValue("item")
	qty, ok := stock[item]

	if !ok {
		http.Error(w, "unknown item: "+item, http.StatusNotFound)
		return
	}

	fmt.Fprintf(w, "%s: %d in stock", item, qty)
})
```

GET handler — `Functional pattern + pattern path`:

- `"GET /stock/{item}"` — method + **path pattern**: `{item}` একটা **wildcard**।
- `r.PathValue("item")` — URL-থেকে wildcard value extract।
- `qty, ok := stock[item]` — map lookup; `!ok` হলে `http.Error` 404।
- success-এ `fmt.Fprintf(w, ...)` — response body।

**Closure:** `stock` handler-এ capture-করা, লজিক প্রতিটা request-বদলে চলে, কিন্তু map একটাই (shared state)।

### Lines 27–38

```go
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
```

POST handler — item অর্ডার:

- `"POST /stock/{item}/order"` — path-এ wildcard।
- `!ok || qty == 0` — out-of-stock check → `http.StatusConflict` (409)।
- সফল: `stock[item] = qty - 1` — **map mutate** (বৈধ write; shared state প্রভাব)।
- Response: `"ordered 1 %s, remaining %d"`।

**State mutation:** কেনো handler-টা লিখবে — concurrent goroutines যতক্ষণ শেয়ার lapse না, single-threaded server-এ ঠিক (এখানে hardcoded-এর সর্বনিম্ন, সত্যিকারের app-এ mutex লাগবে)।

### Line 40

```go
server := &http.Server{Addr: "127.0.0.1:8085", Handler: mux}
```

`http.Server` struct — address + handler।

### Line 42

```go
go server.ListenAndServe()
```

**`go`** — goroutine-এ server run (main block-এ ব্লক না)। Listen error ignored এখানে।

### Line 44

```go
time.Sleep(100 * time.Millisecond)
```

Server ready হওয়ার জন্য short wait (real app-এ proper probe করতে হবে)।

### Lines 46–66

```go
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
```

`fetch` — closure helper: HTTP request করে response body-টা string-এ return করে (error-carry wrappers)। `http.DefaultClient.Do(req)` — request execute; `io.ReadAll(resp.Body)` — body read।

### Lines 68–70

```go
fmt.Println("GET ->", fetch(http.MethodGet, "http://127.0.0.1:8085/stock/keyboard"))
fmt.Println("POST ->", fetch(http.MethodPost, "http://127.0.0.1:8085/stock/keyboard/order"))
fmt.Println("GET ->", fetch(http.MethodGet, "http://127.0.0.1:8085/stock/keyboard"))
```

তিনটা request:

- `GET /stock/keyboard` → `keyboard: 12 in stock`
- `POST /stock/keyboard/order` → `ordered 1 keyboard, remaining 11` (stock কমলো)
- `GET /stock/keyboard` (আবার) → `keyboard: 11 in stock` (**state persisted**: order-এর পর কমে গেছে)

**Proves:** post-টা map-টা mutate করল, পরের GET-টা সেই mutated state দেখাল — closure-shared state-র লাইভ demo।

### Lines 72–78

```go
ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
defer cancel()

if err := server.Shutdown(ctx); err != nil {
	fmt.Println("shutdown:", err)
}
```

Graceful server shutdown — `context.WithTimeout` 2s limit; `Shutdown(ctx)` awaiting in-flight।

### Line 79

```go
}
```

Closing brace — `main` function শেষ হয়।

---

## Expected Output

```
GET -> keyboard: 12 in stock
POST -> ordered 1 keyboard, remaining 11
GET -> keyboard: 11 in stock
```

## মূল শিক্ষা / Key Takeaways

1. **Closure over a map** — handlers-এ shared state (`stock`) capture।
2. **`http.NewServeMux` + method/path patterns** — `GET /x/{item}`।
3. **`r.PathValue`** — URL wildcard extract।
4. **State mutation** — POST-এ stock decrement; পরের GET-তে visible।
5. **`http.Server` + `Shutdown`** — listen + graceful stop।

---

---

<a name="english"></a>

## 🇬🇧 English Version

### Line 1

```go
package main
```

Declares an executable program (`main` package), runnable via `go run`.

### Lines 3–9

```go
import (
	"context"
	"fmt"
	"io"
	"net/http"
	"time"
)
```

- `context` — for `context.WithTimeout` (shutdown).
- `fmt` — for `Fprintf`, `Println`.
- `io` — for `io.ReadAll`.
- `net/http` — HTTP server/client.
- `time` — sleep, timeout.

### Line 11

```go
func main() {
```

Program entry point.

### Line 12

```go
stock := map[string]int{"keyboard": 12, "mouse": 40}
```

**State map** — `stock`: item → available quantity. This map is **captured** in the handlers' closure — the handlers access the same state at dispatch time.

### Line 13

```go
mux := http.NewServeMux()
```

**`http.NewServeMux()`** — a multiplexer (router).

### Lines 15–25

```go
mux.HandleFunc("GET /stock/{item}", func(w http.ResponseWriter, r *http.Request) {
	item := r.PathValue("item")
	qty, ok := stock[item]

	if !ok {
		http.Error(w, "unknown item: "+item, http.StatusNotFound)
		return
	}

	fmt.Fprintf(w, "%s: %d in stock", item, qty)
})
```

GET handler — pattern-based routing:

- `"GET /stock/{item}"` — method + **path pattern**: `{item}` is a **wildcard**.
- `r.PathValue("item")` — extracts the wildcard value from the URL.
- `qty, ok := stock[item]` — map lookup; if `!ok`, `http.Error` returns 404.
- On success, `fmt.Fprintf(w, ...)` writes the response body.

**Closure:** `stock` is captured by the handler; the logic runs per request but the map is shared (one state).

### Lines 27–38

```go
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
```

POST handler — orders an item:

- `"POST /stock/{item}/order"` — wildcard in the path.
- `!ok || qty == 0` — out-of-stock check → `http.StatusConflict` (409).
- On success: `stock[item] = qty - 1` — **mutates the map** (shared state effect).
- Response: `"ordered 1 %s, remaining %d"`.

**State mutation:** the handler writes — sharing state across requests. (Fine in this single-threaded demo; a real app would need a mutex.)

### Line 40

```go
server := &http.Server{Addr: "127.0.0.1:8085", Handler: mux}
```

The `http.Server` struct — address + handler.

### Line 42

```go
go server.ListenAndServe()
```

**`go`** — runs the server in a goroutine (so main doesn't block). The listen error is ignored here for simplicity.

### Line 44

```go
time.Sleep(100 * time.Millisecond)
```

A short wait for the server to be ready (a real app would probe/wait properly).

### Lines 46–66

```go
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
```

`fetch` — a closure helper: performs an HTTP request and returns the response body as a string (with error-carrying wrappers). `http.DefaultClient.Do(req)` executes the request; `io.ReadAll(resp.Body)` reads the body.

### Lines 68–70

```go
fmt.Println("GET ->", fetch(http.MethodGet, "http://127.0.0.1:8085/stock/keyboard"))
fmt.Println("POST ->", fetch(http.MethodPost, "http://127.0.0.1:8085/stock/keyboard/order"))
fmt.Println("GET ->", fetch(http.MethodGet, "http://127.0.0.1:8085/stock/keyboard"))
```

Three requests:

- `GET /stock/keyboard` → `keyboard: 12 in stock`
- `POST /stock/keyboard/order` → `ordered 1 keyboard, remaining 11` (stock decreased)
- `GET /stock/keyboard` (again) → `keyboard: 11 in stock` (**state persisted**: reduced after the order)

**Proves:** the POST mutated the map, and the following GET shows the mutated state — a live demo of closure-shared state.

### Lines 72–78

```go
ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
defer cancel()

if err := server.Shutdown(ctx); err != nil {
	fmt.Println("shutdown:", err)
}
```

Graceful server shutdown — `context.WithTimeout` with a 2s limit; `Shutdown(ctx)` waits for in-flight work.

### Line 79

```go
}
```

Closing brace — ends the `main` function.

---

## Expected Output

```
GET -> keyboard: 12 in stock
POST -> ordered 1 keyboard, remaining 11
GET -> keyboard: 11 in stock
```

## Key Takeaways

1. **Closure over a map** — shared state (`stock`) captured by the handlers.
2. **`http.NewServeMux` + method/path patterns** — `GET /x/{item}`.
3. **`r.PathValue`** — extracting URL wildcards.
4. **State mutation** — POST decrements stock; visible on the next GET.
5. **`http.Server` + `Shutdown`** — listen and graceful stop.