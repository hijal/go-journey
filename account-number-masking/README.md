# account-number-masking

Go-তে **`strings.Repeat` + tail-slicing masking + edge-case guard** শেখার ছোট example — sensitive account number toggle-mask।

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
	"strings"
)
```

- `fmt` — `Println`।
- `strings` — `Repeat`।

### Lines 8–16

```go
func maskAccount(account string) string {
	if len(account) < 4 {
		return strings.Repeat("*", len(account))
	}

	visible := account[len(account)-4:]

	return strings.Repeat("*", len(account)-4) + visible
}
```

**Masking logic:**

- **Edge-guard**: `< 4` অক্ষর হলে পুরোটাই star (`*` × len) — শর্ট-নাম্বার-থেকে শেষ ৪টা দেখানোর কিছুই নেই।
- `account[len(account)-4:]` — **tail-slice**: শেষ ৪টা অক্ষর।
- `strings.Repeat("*", len(account)-4) + visible` — বাকি অংশ star + visible tail।

*(`strings.Repeat("*", n)` — n-সংখ্যক star-এর string তৈরি — সহজ শব্দ-বিল্ডিং, loop লাগে না।)*

### Lines 18–23

```go
func main() {
	accounts := []string{"1234567890", "98765", "42"}

	for _, acc := range accounts {
		fmt.Println(acc, "->", maskAccount(acc))
	}
}
```

৩টা নমুনা: long, medium, short:

- `1234567890` (10) → `******7890`
- `98765` (5) → `*8765`
- `42` (2, <4) → `**`

---

## Expected Output

```
1234567890 -> ******7890
98765 -> *8765
42 -> **
```

## মূল শিক্ষা / Key Takeaways

1. **Tail-slicing** — `account[len-4:]` last-4।
2. **`strings.Repeat`** — mask-run বানানো।
3. **Edge-guard** — `< 4` full-mask।
4. **Composition** — `Repeat + visible` concatenation।

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
	"strings"
)
```

- `fmt` — for `Println`.
- `strings` — for `Repeat`.

### Lines 8–16

```go
func maskAccount(account string) string {
	if len(account) < 4 {
		return strings.Repeat("*", len(account))
	}

	visible := account[len(account)-4:]

	return strings.Repeat("*", len(account)-4) + visible
}
```

**The masking logic:**

- **Edge-guard**: with `< 4` characters, mask the whole thing (`*` × len) — a short number has nothing worth showing.
- `account[len(account)-4:]` — a **tail-slice**: the last 4 characters.
- `strings.Repeat("*", len(account)-4) + visible` — the rest as stars + the visible tail.

*(`strings.Repeat("*", n)` — builds an n-asterisk string without a loop.)*

### Lines 18–23

```go
func main() {
	accounts := []string{"1234567890", "98765", "42"}

	for _, acc := range accounts {
		fmt.Println(acc, "->", maskAccount(acc))
	}
}
```

3 samples: long, medium, short:

- `1234567890` (10) → `******7890`
- `98765` (5) → `*8765`
- `42` (2, <4) → `**`

---

## Expected Output

```
1234567890 -> ******7890
98765 -> *8765
42 -> **
```

## Key Takeaways

1. **Tail-slicing** — `account[len-4:]` for the last 4.
2. **`strings.Repeat`** — building a mask run.
3. **Edge-guard** — full-mask under `< 4`.
4. **Composition** — the `Repeat + visible` concatenation.