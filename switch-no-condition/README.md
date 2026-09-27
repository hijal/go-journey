# switch-no-condition

Go-তে **condition-less `switch`** — `if/else-if/else`-এর সহজ বিকল্প।

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

`time` (Hour)।

### Lines 8–18

```go
func main() {
	t := time.Now()

	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon")
	default:
		fmt.Println("Good evening")
	}
}
```

**`switch {}`** — expression ছাড়া, প্রথম **true** case-টাই চালায় (if/else-if/else-equivalent)। প্রথম `true`-তেই থামে — `t.Hour()`-এর মানসাপেক্ষ।

---

## Expected Output (time-of-day dependent)

```
Good morning!   / Good afternoon / Good evening
```

## মূল শিক্ষা / Key Takeaways

1. **Condition-less switch** — cleaner chain-এর জন্য।
2. **First true wins** — শীর্ষ-অনুক্রম।
3. **Boolean cases** — প্রতিটি case নিজস্ব condition।

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

`time` (Hour).

### Lines 8–18

```go
func main() {
	t := time.Now()

	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon")
	default:
		fmt.Println("Good evening")
	}
}
```

**`switch {}`** — no expression; the first **true** case runs (the if/else-if/else equivalent). It stops at the first true — depends on the value of `t.Hour()`.

---

## Expected Output (time-of-day dependent)

```
Good morning!   / Good afternoon / Good evening
```

## Key Takeaways

1. **Condition-less switch** — a cleaner conditional chain.
2. **First true wins** — top-down order matters.
3. **Boolean cases** — no expression, just per-case conditions.