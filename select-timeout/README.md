# select-timeout

Go-তে **`select` race with timeout** — দুই source + `time.After` guard, যেটা আগে আসে সেটিই।

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

`time` (Sleep/After)।

### Lines 8–29

```go
func main() {
	primary := make(chan string)
	secondary := make(chan string)

	go func() {
		time.Sleep(time.Millisecond * 50)
		primary <- "primary-db result"
	}()

	go func() {
		time.Sleep(time.Millisecond * 10)
		secondary <- "secondary-db result"
	}()

	select {
	case r := <-primary:
		fmt.Println(r)
	case r := <-secondary:
		fmt.Println(r)
	case <-time.After(time.Millisecond * 100):
		fmt.Println("timeout")
	}
}
```

- দুইটা **unbuffered** channel, দুইটা goroutine।
- `primary` → 50ms, `secondary` → 10ms → **secondary আগে**, তাই সেটিই জেতে।
- **`time.After(100ms)`** — তৃতীয় case: দুটোই না এলে timeout।
- `select` একই সঙ্গে সব case ready হলে **random** একটা বেছে নেয় (priority নেই)।

---

## Expected Output

```
secondary-db result
```

## মূল শিক্ষা / Key Takeaways

1. **`select` = race** — যে channel আগে ready, সেটিই।
2. **`time.After`** — deadline হিসেবে একটা channel।
3. **Deterministic এখানে** — 10ms সবসময় 50ms-এর আগে।
4. **কেউ না এলে** — `timeout` branch।

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

`time` (Sleep/After).

### Lines 8–29

```go
func main() {
	primary := make(chan string)
	secondary := make(chan string)

	go func() {
		time.Sleep(time.Millisecond * 50)
		primary <- "primary-db result"
	}()

	go func() {
		time.Sleep(time.Millisecond * 10)
		secondary <- "secondary-db result"
	}()

	select {
	case r := <-primary:
		fmt.Println(r)
	case r := <-secondary:
		fmt.Println(r)
	case <-time.After(time.Millisecond * 100):
		fmt.Println("timeout")
	}
}
```

- Two **unbuffered** channels and two goroutines.
- `primary` takes 50ms, `secondary` takes 10ms → **secondary wins**.
- **`time.After(100ms)`** is the third case: a timeout if neither arrives.
- If several cases are ready at once, `select` picks one **at random** (no priority).

---

## Expected Output

```
secondary-db result
```

## Key Takeaways

1. **`select` is a race** — whichever channel is ready first.
2. **`time.After`** — a deadline expressed as a channel.
3. **Deterministic here** — 10ms always beats 50ms.
4. **If nothing arrives** — the `timeout` branch.