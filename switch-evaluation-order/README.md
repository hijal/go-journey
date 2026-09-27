# switch-evaluation-order

Go-তে **switch-এর top-down case evaluation** — `time.Saturday`-কে `today+N` দিয়ে track করা।

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

`time` (Weekday)।

### Lines 8–21

```go
func main() {
	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()

	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}
}
```

**Value switch**: `time.Saturday`-এর সাথে case-গুলো **উপর-নিচে** পরীক্ষা করে। Case-expression আজকের দিন থেকে offline (today+0/+1/+2)। আজ রবিবার হলে Saturday ৬ দিন দূরে → কেবল `default`।

---

## Expected Output (depending on today)

```
When's Saturday?
Today.      / Tomorrow. / In two days. / Too far away.
```

## মূল শিক্ষা / Key Takeaways

1. **Top-down evaluation** — case-গুলো ক্রমে পরীক্ষিত।
2. **First-match wins** — অগ্রিম মেল আটকে দেয়।
3. **Default safety** — কোনো মিল না-পেলে fallback।
4. **Time-dependent** — উদাহরণটি আপেক্ষিক-আজকে।

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

`time` (Weekday).

### Lines 8–21

```go
func main() {
	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()

	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}
}
```

**A value switch**: tests `time.Saturday` against each case expression **top-down**. The cases are offsets from today (today+0/+1/+2). If today is Sunday, Saturday is 6 days away → only `default` matches.

---

## Expected Output (depending on today)

```
When's Saturday?
Today.      / Tomorrow. / In two days. / Too far away.
```

## Key Takeaways

1. **Top-down evaluation** — cases are tried in order.
2. **First-match wins** — the earliest match stops the search.
3. **Default safety** — a fallback when nothing matches.
4. **Time-dependent output** — the example is relative to today.