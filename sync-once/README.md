# sync-once

Go-তে **`sync.OnceValue`** — costly init-কে lazily, thread-safe — একবারই চালানো।

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
	"sync"
)
```

`sync` (`OnceValue`, `WaitGroup`)।

### Lines 8–11

```go
var loadConfig = sync.OnceValue(func() map[string]string {
	fmt.Println("reading config from disk (happens once)")
	return map[string]string{"currency": "BDT"}
})
```

**`sync.OnceValue(fn)`** — `fn` lazily-called; return করা value cache হয়; যে-কয়বার-ই ডাকা হোক, `fn` মাত্র **একবার** চলে। (Variant: `Once` no-value, `OnceFunc` fire-only, `OnceValues` two-return।)

### Lines 13–20

```go
func main() {
	var wg sync.WaitGroup
	for range 3 {
		wg.Go(func() {
			fmt.Println("currency:", loadConfig()["currency"])
		})
	}
	wg.Wait()
}
```

3 goroutine **concurrently** একই `loadConfig()` কল করে — কোনো race নেই, disk-read একবার-ই।

---

## Expected Output

```
reading config from disk (happens once)
currency: BDT
currency: BDT
currency: BDT
```

## মূল শিক্ষা / Key Takeaways

1. **`sync.OnceValue`** — lazy + cached + thread-safe।
2. **Exactly-once semantics** — concurrent-কলেও init একবার।
3. **`sync.Once*` family** — `Once`, `OnceFunc`, `OnceValue`, `OnceValues`।
4. **No external locking** — `Once` নিজেই synchronization handle করে।
5. **Idiomatic for** — config parsing, singleton connection, resource init।

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
	"sync"
)
```

`sync` (`OnceValue`, `WaitGroup`).

### Lines 8–11

```go
var loadConfig = sync.OnceValue(func() map[string]string {
	fmt.Println("reading config from disk (happens once)")
	return map[string]string{"currency": "BDT"}
})
```

**`sync.OnceValue(fn)`** — `fn` is called lazily; the returned value is cached; no matter how often it's called, `fn` runs **exactly once**. (Variants: `Once` no-value, `OnceFunc` fire-only, `OnceValues` two-value.)

### Lines 13–20

```go
func main() {
	var wg sync.WaitGroup
	for range 3 {
		wg.Go(func() {
			fmt.Println("currency:", loadConfig()["currency"])
		})
	}
	wg.Wait()
}
```

3 goroutines call the same `loadConfig()` **concurrently** — no race and a single disk-read.

---

## Expected Output

```
reading config from disk (happens once)
currency: BDT
currency: BDT
currency: BDT
```

## Key Takeaways

1. **`sync.OnceValue`** — lazy + cached + thread-safe.
2. **Exactly-once semantics** — a single init even under concurrent calls.
3. **The `sync.Once*` family** — `Once`, `OnceFunc`, `OnceValue`, `OnceValues`.
4. **No external locking** — `Once` handles synchronization itself.
5. **Idiomatic for** — config parsing, singleton connections, resource init.