# fluent-http-builder

Go-তে **builder pattern (fluent chain setter) + multi-config HTTP client scaffold** শেখার ছোট example।

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

### Lines 3–6

```go
import (
	"fmt"
	"time"
)
```

- `fmt` — `Println`, `Sprintf`।
- `time` — `Duration`।

### Lines 8–12

```go
type HTTPClientBuilder struct {
	baseURL string
	timeout time.Duration
	headers map[string]string
}
```

Builder-এর ৩টা unexported config field।

### Lines 14–18

```go
func NewHTTPClientBuilder() *HTTPClientBuilder {
	return &HTTPClientBuilder{
		headers: make(map[string]string),
	}
}
```

Constructor — map initializes (nil-map-এ write panic করত)।

### Lines 20–33

```go
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
```

**Fluent setters** — প্রতিটা `return b` (self-pointer) → chaining possible। field set-এর পর আবার builder ফেরত দেয়।

### Lines 35–37

```go
func (b *HTTPClientBuilder) Build() string {
	return fmt.Sprintf("Client{baseURL=%s, timeout=%s, headers=%v}", b.baseURL, b.timeout, b.headers)
}
```

`Build()` — শেষে সমন্বিত string (`%v` map-print-ও দেখায়)।

### Lines 39–45

```go
	client := NewHTTPClientBuilder().
		BaseURL("https://api.example.com").
		Timeout(10*time.Second).
		Header("Authorization", "Bearer token123").
		Build()
	fmt.Println(client)
```

**Method chain** — dot-এ dot; `10*time.Second` → `timeout=10s`-হয়ে print।

---

## Expected Output

```
Client{baseURL=https://api.example.com, timeout=10s, headers=map[Authorization:Bearer token123]}
```

## মূল শিক্ষা / Key Takeaways

1. **Builder pattern** — constructor + step-set।
2. **Setters return self** — fluent chain।
3. **Constructor initializes** — headers map non-nil রাখা।
4. **`time.Duration`** — `10*time.Second` → `10s` print।

---

---

<a name="english"></a>

## 🇬🇧 English Version

### Line 1

```go
package main
```

Declares an executable program (`main` package), runnable via `go run`.

### Lines 3–6

```go
import (
	"fmt"
	"time"
)
```

- `fmt` — for `Println`, `Sprintf`.
- `time` — for `Duration`.

### Lines 8–12

```go
type HTTPClientBuilder struct {
	baseURL string
	timeout time.Duration
	headers map[string]string
}
```

The builder's 3 unexported config fields.

### Lines 14–18

```go
func NewHTTPClientBuilder() *HTTPClientBuilder {
	return &HTTPClientBuilder{
		headers: make(map[string]string),
	}
}
```

The constructor — the map is initialized (writing to a nil map would panic).

### Lines 20–33

```go
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
```

**Fluent setters** — each `return b` (the self pointer) enables chaining: set a field, hand the builder back.

### Lines 35–37

```go
func (b *HTTPClientBuilder) Build() string {
	return fmt.Sprintf("Client{baseURL=%s, timeout=%s, headers=%v}", b.baseURL, b.timeout, b.headers)
}
```

`Build()` — produces the final aggregated string (`%v` prints the map too).

### Lines 39–45

```go
	client := NewHTTPClientBuilder().
		BaseURL("https://api.example.com").
		Timeout(10*time.Second).
		Header("Authorization", "Bearer token123").
		Build()
	fmt.Println(client)
```

A **method chain** — dot after dot; `10*time.Second` prints as `timeout=10s`.

---

## Expected Output

```
Client{baseURL=https://api.example.com, timeout=10s, headers=map[Authorization:Bearer token123]}
```

## Key Takeaways

1. **Builder pattern** — constructor + step-set.
2. **Setters return self** — the fluent chain.
3. **Constructor-side initialization** — keeps the headers map non-nil.
4. **`time.Duration`** — `10*time.Second` prints as `10s`.