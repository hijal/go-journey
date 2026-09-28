# select-default-nonblocking

Go-তে **`select` + `default` = non-blocking** — queue ভরা হলে message drop।

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

`fmt` (print)।

### Lines 5–15

```go
func main() {
	alerts := make(chan string, 1)

	for _, msg := range []string{"cpu high", "disk full"} {
		select {
		case alerts <- msg:
			fmt.Println("queued:", msg)
		default:
			fmt.Println("dropped (queue full):", msg)
		}
	}
}
```

- **`make(chan string, 1)`** — 1-slot queue।
- **`select` with `default`** — কোনো case ready না হলে `default` চলে, **block হয় না**:
  - ১ম message: slot খালি → `queued: cpu high`।
  - ২য় message: slot ভরা (কেউ receive করেনি) → `dropped (queue full): disk full`।
- শুধু `select`-এ `default` থাকলেই এটা non-blocking select।

---

## Expected Output

```
queued: cpu high
dropped (queue full): disk full
```

## মূল শিক্ষা / Key Takeaways

1. **`default` clause** — ready case না থাকলেও চলে।
2. **Non-blocking** — কোনো অপেক্ষা নেই।
3. **Load-shedding** — queue ভরা হলে drop করা pattern।
4. **Worker ছাড়াই** এই select-এর জন্য goroutine লাগে না।

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

`fmt` (print).

### Lines 5–15

```go
func main() {
	alerts := make(chan string, 1)

	for _, msg := range []string{"cpu high", "disk full"} {
		select {
		case alerts <- msg:
			fmt.Println("queued:", msg)
		default:
			fmt.Println("dropped (queue full):", msg)
		}
	}
}
```

- **`make(chan string, 1)`** — a 1-slot queue.
- **`select` with `default`** — if no case is ready, `default` runs and it **never blocks**:
  - 1st message: slot free → `queued: cpu high`.
  - 2nd message: slot full (nobody received) → `dropped (queue full): disk full`.
- This is a non-blocking select purely because `default` is present.

---

## Expected Output

```
queued: cpu high
dropped (queue full): disk full
```

## Key Takeaways

1. **The `default` clause** — runs when no case is ready.
2. **Non-blocking** — no waiting at all.
3. **Load shedding** — the drop-when-full pattern.
4. **No worker needed** for this kind of select.