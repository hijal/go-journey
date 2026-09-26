# runtime-helper

Go **`runtime` package-এর introspection helpers** — CPU core, GOMAXPROCS, live goroutine count।

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
	"runtime"
)
```

`runtime` — Go runtime-এর introspective API।

### Lines 8–11

```go
func main() {
	fmt.Println("CPU cores:", runtime.NumCPU())
	fmt.Println("GOMAXPROCS:", runtime.GOMAXPROCS(0))
	fmt.Println("goroutines now:", runtime.NumGoroutine())
}
```

- **`runtime.NumCPU()`** — machine-এ logical CPU core সংখ্যা।
- **`runtime.GOMAXPROCS(0)`** — `0` দিলে value শুধু query হয়, set হয় না — অ্যাক্টিভ parallel-capable thread limit।
- **`runtime.NumGoroutine()`** — এই মুহূর্তে live goroutine count (এখানে শুধু `main` → 1)।

---

## Expected Output

```
CPU cores: 16
GOMAXPROCS: 16
goroutines now: 1
```

## মূল শিক্ষা / Key Takeaways

1. **`NumCPU` vs `GOMAXPROCS`** — hardware limit vs runtime policy।
2. **`GOMAXPROCS(0)`** — query-only call (0 দিলে set হয় না)।
3. **`NumGoroutine`** — goroutine leak monitoring-এর প্রাথমিক টুল।
4. **Scheduler** — Go-র M×N scheduler-এ P-number control করে GOMAXPROCS।

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
	"runtime"
)
```

`runtime` — Go runtime's introspective API.

### Lines 8–11

```go
func main() {
	fmt.Println("CPU cores:", runtime.NumCPU())
	fmt.Println("GOMAXPROCS:", runtime.GOMAXPROCS(0))
	fmt.Println("goroutines now:", runtime.NumGoroutine())
}
```

- **`runtime.NumCPU()`** — the machine's logical CPU core count.
- **`runtime.GOMAXPROCS(0)`** — passing `0` only queries (doesn't set) — the active parallel-capable thread limit.
- **`runtime.NumGoroutine()`** — live goroutine count right now (here only `main` → 1).

---

## Expected Output

```
CPU cores: 16
GOMAXPROCS: 16
goroutines now: 1
```

## Key Takeaways

1. **`NumCPU` vs `GOMAXPROCS`** — hardware limit vs runtime policy.
2. **`GOMAXPROCS(0)`** — a query-only call (0 doesn't set).
3. **`NumGoroutine`** — a first-line tool for goroutine-leak monitoring.
4. **Scheduler** — in Go's M×N scheduler, GOMAXPROCS controls the P-number (parallelism).