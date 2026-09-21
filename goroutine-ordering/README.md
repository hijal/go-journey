# goroutine-ordering

Go-তে **goroutine scheduling non-determinism + main-এর exit-এ pending goroutine kill** শেখার ছোট example।

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

`fmt` + `time` (`Sleep` — goroutine-দের চালানোর সময় দেওয়ার naive উপায়)।

### Lines 8–10

```go
func sendEmail(to string) {
	fmt.Println("email sent to", to)
}
```

সাধারণ func — `go` কীওয়ার্ড দিয়ে আলাদা goroutine-এ চালানো হবে।

### Lines 12–24

```go
func main() {
	go sendEmail("example@example.com")

	go func() {
		fmt.Println("i am anonymous function")
	}()

	fmt.Println("this is goroutine ordering")

	time.Sleep(100 * time.Millisecond)

	go fmt.Println("you will probably never see this")
	fmt.Println("main is done")
}
```

**The point:**

- `go sendEmail(...)` + `go func(){...}()` — ২টা goroutine spawn; **no ordering guarantee**।
- `fmt.Println("this is goroutine ordering")` — main-run-time-এ print, সাধারণত প্রথম (main শুরুর সময়েই চলে)।
- `time.Sleep(100ms)` — goroutine-দের অন-dayযোগ্য slack (না-হলে main বের হয়ে যেত)।
- `go fmt.Println("you will probably never see this")` — **sleep-এর পরে** spawn; এরপর main সাথে-সাথে শেষ → গোরুটীন-এ scheduling-এর সময়ই নেই → print **হয় না**।

### Expected Behavior

গোরুটীন order নির্ধারিত নয় — ২টা-য় 2টা যেকোনো ক্রমে (সাধারণত); শেষটা almost-never।

---

## Expected Output

```
this is goroutine ordering
i am anonymous function
email sent to example@example.com
main is done
```

(last 2 lines swap করতে পারে; `you will probably never see this` প্রায় নিশ্চিতভাবে absent)

## মূল শিক্ষা / Key Takeaways

1. **`go` কীওয়ার্ড** — goroutine launch।
2. **No ordering guarantee** — scheduler decides।
3. **main exit kill** — pending goroutine-দের work করার সুযোগ হয় না।
4. **`time.Sleep` as sync** — naive; proper tool: `sync.WaitGroup`/channel।
5. **Anonymous func + `go`** — closure goroutine।

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

`fmt` + `time` (`Sleep` — a naive way to give goroutines time to run).

### Lines 8–10

```go
func sendEmail(to string) {
	fmt.Println("email sent to", to)
}
```

A plain func — it will run in its own goroutine via the `go` keyword.

### Lines 12–24

```go
func main() {
	go sendEmail("example@example.com")

	go func() {
		fmt.Println("i am anonymous function")
	}()

	fmt.Println("this is goroutine ordering")

	time.Sleep(100 * time.Millisecond)

	go fmt.Println("you will probably never see this")
	fmt.Println("main is done")
}
```

**The point:**

- `go sendEmail(...)` + `go func(){...}()` — 2 goroutines spawned; **no ordering guarantee**.
- `fmt.Println("this is goroutine ordering")` — printed from the main goroutine, so usually first.
- `time.Sleep(100ms)` — slack so the goroutines get a chance (otherwise main exits first).
- `go fmt.Println("you will probably never see this")` — spawned **after the sleep**; main finishes right after, so there's no time to schedule it — it usually **never prints**.

### Expected Behavior

The goroutine order is not deterministic — the first two appear in either order; the last one almost always missing.

---

## Expected Output

```
this is goroutine ordering
i am anonymous function
email sent to example@example.com
main is done
```

(the last 2 lines may swap; `you will probably never see this` is almost certainly absent)

## Key Takeaways

1. **The `go` keyword** — launches a goroutine.
2. **No ordering guarantee** — the scheduler decides.
3. **main-exit kill** — pending goroutines never get to run.
4. **`time.Sleep` as sync** — naive; use `sync.WaitGroup`/channels instead.
5. **Anonymous func + `go`** — closure goroutines.