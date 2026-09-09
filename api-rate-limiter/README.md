# api-rate-limiter

Go-তে **closure** দিয়ে per-client rate limiter শেখার ছোট example — `newRateLimiter` একটা function; এটা counter (`used`) capture করে প্রতিবার call-এ increment করে।

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

### Line 3

```go
import "fmt"
```

Console-এ output print করার জন্য `fmt` package import করা হয়।

### Lines 5–12

```go
func newRateLimiter(limit int) func(client string) bool {
	used := 0
	return func(client string) bool {
		used++
		fmt.Printf("%s -> request %d/%d\n", client, used, limit)
		return used <= limit
	}
}
```

`newRateLimiter` — একটা rate limiter factory:

- `limit int` — অনুমোদনযোগ্য request সংখ্যা।
- `used := 0` — **closure variable**: return-করা function-টা এটাকে capture করে।
- return typed function: `func(client string) bool` — client নাম নেয়, allowed কি না বলে।
- ভেতরের function-টা প্রতিটা call-এ:
  - `used++` — **shared counter increment** (state closure-এ থাকে, function ends হয় না)।
  - print করে request count।
  - `used <= limit` return — limit-টা crossing করলে `false`।

**লক্ষ্য করো:** `used`, `limit` দুটোই **call-এর মাঝে survive** করে— প্রতিটা `newRateLimiter` call একটা **fresh independent state** তৈরি করে। এটাই closure-এর মূল শক্তি: stateful function factory।

### Line 14

```go
func main() {
```

Program-এর entry point।

### Lines 15–19

```go
serve := newRateLimiter(3)
serve("client-A")
serve("client-A")
serve("client-B")
fmt.Println("client-B allowed?", serve("client-B"))
```

- `serve := newRateLimiter(3)` — limit 3-এ একটা limiter (shared counter)।
- 3টা request: 1/3 (A), 2/3 (A), 3/3 (B) — সব allow।
- 4th request (B): 4/3 → `4 <= 3` = **false** → `client-B allowed? false`।

**গুরুত্বপূর্ণ:** counter-টা সব client-এর জন্য **shared** — limit-টা per-instance (`serve`), per-client না। এটা global-per-API key rate limit-এর মতো।

### Lines 21–23

```go
burst := newRateLimiter(1)
burst("client-C")
fmt.Println("client-C allowed?", burst("client-C"))
```

- `burst := newRateLimiter(1)` — **আলাদা** limiter, limit 1, fresh counter।
- 1/1 (allowed), তারপর 2/1 → **false** → `client-C allowed? false`।

**শো:** দুটো `newRateLimiter` call দুটো **independent** state — `serve`-এর counter `burst`-কে affect করে না।

### Line 24

```go
}
```

Closing brace — `main` function শেষ হয়।

---

## Expected Output

```
client-A -> request 1/3
client-A -> request 2/3
client-B -> request 3/3
client-B -> request 4/3
client-B allowed? false
client-C -> request 1/1
client-C -> request 2/1
client-C allowed? false
```

## মূল শিক্ষা / Key Takeaways

1. **Closure factory** — `newRateLimiter` return-করা function-এ state (`used`) capture।
2. **State persistence** — counter-টা call-এর মাঝে বেঁচে থাকে।
3. **Independent state** — প্রতিটা factory call fresh counter।
4. **Bounded limit** — `used <= limit` limit crossing detect।
5. **Stateful function** — closure-র প্রাইভেট state (টা global-এ ছড়ায় না)।

---

---

<a name="english"></a>

## 🇬🇧 English Version

### Line 1

```go
package main
```

Declares an executable program (`main` package), runnable via `go run`.

### Line 3

```go
import "fmt"
```

Imports the `fmt` package for console output.

### Lines 5–12

```go
func newRateLimiter(limit int) func(client string) bool {
	used := 0
	return func(client string) bool {
		used++
		fmt.Printf("%s -> request %d/%d\n", client, used, limit)
		return used <= limit
	}
}
```

`newRateLimiter` — a rate limiter factory:

- `limit int` — the number of allowed requests.
- `used := 0` — a **closure variable**: the returned function captures it.
- Returns a typed function: `func(client string) bool` — takes a client name, says whether it's allowed.
- On each call the inner function:
  - `used++` — **increments a shared counter** (the state lives in the closure, it isn't lost when the function returns).
  - prints the request count.
  - returns `used <= limit` — `false` once the limit is crossed.

Note how `used` and `limit` **survive between calls** — every `newRateLimiter` call creates a **fresh independent state**. That's the essence of closures: stateful function factories.

### Line 14

```go
func main() {
```

Program entry point.

### Lines 15–19

```go
serve := newRateLimiter(3)
serve("client-A")
serve("client-A")
serve("client-B")
fmt.Println("client-B allowed?", serve("client-B"))
```

- `serve := newRateLimiter(3)` — a limiter with limit 3 (one shared counter).
- Three requests: 1/3 (A), 2/3 (A), 3/3 (B) — all allowed.
- 4th request (B): 4/3 → `4 <= 3` = **false** → `client-B allowed? false`.

**Important:** the counter is **shared** across all clients — the limit is per-instance (`serve`), not per-client. Like a global per-API-key rate limit.

### Lines 21–23

```go
burst := newRateLimiter(1)
burst("client-C")
fmt.Println("client-C allowed?", burst("client-C"))
```

- `burst := newRateLimiter(1)` — a **separate** limiter, limit 1, fresh counter.
- 1/1 (allowed), then 2/1 → **false** → `client-C allowed? false`.

**Shows:** two `newRateLimiter` calls produce two **independent** states — `serve`'s counter doesn't affect `burst`.

### Line 24

```go
}
```

Closing brace — ends the `main` function.

---

## Expected Output

```
client-A -> request 1/3
client-A -> request 2/3
client-B -> request 3/3
client-B -> request 4/3
client-B allowed? false
client-C -> request 1/1
client-C -> request 2/1
client-C allowed? false
```

## Key Takeaways

1. **Closure factory** — `newRateLimiter`'s returned function captures state (`used`).
2. **State persistence** — the counter lives across calls.
3. **Independent state** — each factory call gets a fresh counter.
4. **Bounded limit** — `used <= limit` detects crossing the limit.
5. **Stateful function** — closure-private state (not scattered globals).