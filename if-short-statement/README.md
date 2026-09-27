# if-short-statement

Go-তে **`if` short statement without `else`** — computed value ব্যবহার করে bypass-able early return।

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
	"math"
)
```

`math` (Pow)।

### Lines 8–13

```go
func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}
```

**Short statement + early return** — কোনো `else` ছাড়াই:

- `v < lim` → সরাসরি `v` return।
- নাহলে `lim` return (fallback)।

### Lines 15–19

```go
func main() {
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
}
```

- `pow(3,2,10)`: 9 < 10 → **9**।
- `pow(3,3,20)`: 27 ≥ 20 → **20**।

---

## Expected Output

```
9 20
```

## মূল শিক্ষা / Key Takeaways

1. **Guard clause style** — nesting কমায়।
2. **No `else`** — early return sole escaping।
3. **Scoped `v`** — if-block-এর ভেতরেই।
4. **Fallback return** — সব path-এ মান।

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
	"math"
)
```

`math` (Pow).

### Lines 8–13

```go
func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}
```

**A short statement + early return** — no `else` needed:

- If `v < lim` → return `v` directly.
- Otherwise return `lim` as the fallback.

### Lines 15–19

```go
func main() {
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
}
```

- `pow(3,2,10)`: 9 < 10 → **9**.
- `pow(3,3,20)`: 27 ≥ 20 → **20**.

---

## Expected Output

```
9 20
```

## Key Takeaways

1. **Guard-clause style** — less nesting.
2. **No `else`** — a single early return escapes.
3. **Scoped `v`** — lives inside the block.
4. **Fallback return** — a value on every path.