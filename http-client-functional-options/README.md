# http-client-functional-options

Go-তে **Functional Options pattern** শেখার ছোট example — `type Option func(*APIClient)` + variadic options দিয়ে configurable-কিন্তু-default-মতো constructor।

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

- `fmt` — `Printf`।
- `time` — `time.Duration`।

### Lines 8–13

```go
type APIClient struct {
	BaseURL    string
	Timeout    time.Duration
	RetryCount int
	AuthToken  string
}
```

`APIClient` — configurable client:

- `BaseURL` — API-র root URL।
- `Timeout` — request timeout।
- `RetryCount` — retry-র সংখ্যা।
- `AuthToken` — authentication token।

### Line 15

```go
type Option func(*APIClient)
```

`Option` — **functional option**: একটা function যা `*APIClient`-কে modify করে। এটাই pattern-এর মূল।

### Lines 17–33

```go
func WithTimeout(d time.Duration) Option {
	return func(a *APIClient) {
		a.Timeout = d
	}
}

func WithRetry(count int) Option {
	return func(a *APIClient) {
		a.RetryCount = count
	}
}

func WithAuth(token string) Option {
	return func(a *APIClient) {
		a.AuthToken = token
	}
}
```

৩টা **option constructor** — প্রতিটা একটা closure return করে:

- `WithTimeout(d)` — closure-টা `a.Timeout = d` set করে।
- `WithRetry(count)` — `a.RetryCount = count`।
- `WithAuth(token)` — `a.AuthToken = token`।

**Closure capture:** প্রতিটা closure-টা `d`/`count`/`token` capture করে — call-এর সময় client-এ প্রয়োগ করে। Param-টা captured — আলাদা স্টেট storage-এর দরকার নেই।

### Lines 35–46

```go
func NewAPIClient(baseURL string, opts ...Option) *APIClient {
	client := &APIClient{
		BaseURL:    baseURL,
		Timeout:    30 * time.Second,
		RetryCount: 3,
	}

	for _, opt := range opts {
		opt(client)
	}
	return client
}
```

`NewAPIClient` — constructor:

- Defaults দিয়ে initialize: `Timeout` 30s, `RetryCount` 3 (AuthToken default — zero value, খালি)।
- `opt(client)` — প্রতিটা option-টা client-এ apply: **default-গুলোর ওপরে override**।
- Order-টা apply-এর ক্রম — একই option একাধিকবার দিলে শেষটা জিতবে।

### Lines 48–56

```go
func main() {
	client := NewAPIClient(
		"https://api.example.com",
		WithTimeout(10*time.Second),
		WithAuth("super-secret-token"),
	)

	fmt.Printf("client configured for %s with %vs timeout\n", client.BaseURL, client.Timeout.Seconds())
}
```

- Defaults-তে `WithTimeout(10s)` → 30s-র বদলে **10s**। AuthToken-এ secret set।
- Output-এ timeout 10s দেখা যায় — prove default override।

---

## Expected Output

```
client configured for https://api.example.com with 10s timeout
```

## মূল শিক্ষা / Key Takeaways

1. **`Option func(*T)`** — functional option এর type, receiver modify করে।
2. **Option closure** — `With*`-দের captured param-দিয়ে mutate।
3. **Defaults + overrides** — constructor default দিয়ে apply।
4. **Variadic `...Option`** — ঐচ্ছিক config, zero/অনেক option।
5. **Clean API** — long parameter-list এড়ানো, backward-compatible।

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

- `fmt` — for `Printf`.
- `time` — for `time.Duration`.

### Lines 8–13

```go
type APIClient struct {
	BaseURL    string
	Timeout    time.Duration
	RetryCount int
	AuthToken  string
}
```

`APIClient` — a configurable client:

- `BaseURL` — the API root URL.
- `Timeout` — the request timeout.
- `RetryCount` — the number of retries.
- `AuthToken` — the authentication token.

### Line 15

```go
type Option func(*APIClient)
```

`Option` — a **functional option**: a function that modifies the `*APIClient`. This is the heart of the pattern.

### Lines 17–33

```go
func WithTimeout(d time.Duration) Option {
	return func(a *APIClient) {
		a.Timeout = d
	}
}

func WithRetry(count int) Option {
	return func(a *APIClient) {
		a.RetryCount = count
	}
}

func WithAuth(token string) Option {
	return func(a *APIClient) {
		a.AuthToken = token
	}
}
```

Three **option constructors** — each returns a closure:

- `WithTimeout(d)` — its closure sets `a.Timeout = d`.
- `WithRetry(count)` — sets `a.RetryCount = count`.
- `WithAuth(token)` — sets `a.AuthToken = token`.

**Closure capture:** each closure captures `d`/`count`/`token` and applies them to the client at call time. No separate state storage needed for the parameters.

### Lines 35–46

```go
func NewAPIClient(baseURL string, opts ...Option) *APIClient {
	client := &APIClient{
		BaseURL:    baseURL,
		Timeout:    30 * time.Second,
		RetryCount: 3,
	}

	for _, opt := range opts {
		opt(client)
	}
	return client
}
```

`NewAPIClient` — the constructor:

- Initializes with defaults: `Timeout` 30s, `RetryCount` 3 (AuthToken defaults to the zero value, empty).
- `opt(client)` — applies each option on top of the defaults.
- Application order matters — the last option wins.

### Lines 48–56

```go
func main() {
	client := NewAPIClient(
		"https://api.example.com",
		WithTimeout(10*time.Second),
		WithAuth("super-secret-token"),
	)

	fmt.Printf("client configured for %s with %vs timeout\n", client.BaseURL, client.Timeout.Seconds())
}
```

- `WithTimeout(10s)` overrides the default → **10s** instead of 30s. The auth token is set to a secret.
- The output shows 10s timeout — proving the default override.

---

## Expected Output

```
client configured for https://api.example.com with 10s timeout
```

## Key Takeaways

1. **`Option func(*T)`** — the functional-option type, modifying the receiver.
2. **Option closures** — `With*` constructors mutate via captured params.
3. **Defaults + overrides** — defaults from the constructor, then options applied.
4. **Variadic `...Option`** — optional config, zero or many options.
5. **Clean API** — avoids long parameter lists and stays backward-compatible.