# time-channels

Go-তে **ticker + deadline channel** — `time.NewTicker` + `time.After`, `select`-এ দুটো case।

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

`time` (NewTicker/After/Sleep)।

### Lines 8–22

```go
func main() {
	ticker := time.NewTicker(30 * time.Millisecond)
	defer ticker.Stop()

	deadline := time.After(100 * time.Millisecond)

	for {
		select {
		case <-ticker.C:
			fmt.Println("health check")
		case <-deadline:
			fmt.Println("stop checking")
			return
		}
	}
}
```

- **`time.NewTicker(30ms)`** — নিয়মিত tick পাঠায়; `defer ticker.Stop()` resource মুক্ত করে।
- **`time.After(100ms)`** — একবার-ই fire হয় এমন channel।
- **`select`** দুটোই দেখে:
  - 30ms, 60ms, 90ms → **৩বার** "health check"।
  - 100ms → deadline → "stop checking" + `return`।
- `ticker.C` ছাড়া loop কখনো শেষ হতো না — `deadline`-ই exit দেয়।

---

## Expected Output

```
health check
health check
health check
stop checking
```

## মূল শিক্ষা / Key Takeaways

1. **Ticker channel** — `ticker.C`-এ পুনরাবৃত্ত tick।
2. **`defer ticker.Stop()`** — resource cleanup।
3. **`time.After`** — one-shot channel।
4. **Tick সংখ্যা** — 100ms/30ms = 3 (deadline আসার আগে)।

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

`time` (NewTicker/After/Sleep).

### Lines 8–22

```go
func main() {
	ticker := time.NewTicker(30 * time.Millisecond)
	defer ticker.Stop()

	deadline := time.After(100 * time.Millisecond)

	for {
		select {
		case <-ticker.C:
			fmt.Println("health check")
		case <-deadline:
			fmt.Println("stop checking")
			return
		}
	}
}
```

- **`time.NewTicker(30ms)`** — sends a periodic tick; `defer ticker.Stop()` frees the resource.
- **`time.After(100ms)`** — a channel that fires exactly once.
- The **`select`** watches both:
  - At 30ms, 60ms, 90ms → "health check" **three times**.
  - At 100ms → the deadline wins → "stop checking" and `return`.
- Without `ticker.C` the loop would never end — the `deadline` is what exits it.

---

## Expected Output

```
health check
health check
health check
stop checking
```

## Key Takeaways

1. **Ticker channel** — recurring ticks on `ticker.C`.
2. **`defer ticker.Stop()`** — resource cleanup.
3. **`time.After`** — a one-shot channel.
4. **Tick count** — 100ms/30ms = 3 before the deadline.