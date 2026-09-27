# function-values

Go-তে **functions as values** — closure-style variable, argument-passing, `math.Pow`।

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

`math` (Sqrt, Pow)।

### Lines 8–10

```go
func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}
```

**Function-type parameter** — যে-কোনো দুই-float-থেকে-float function গ্রহণ করে।

### Lines 12–20

```go
func main() {
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}

	fmt.Println(hypot(5, 12))

	fmt.Println(compute(hypot))
	fmt.Println(compute(math.Pow))
}
```

- `hypot := func(...)` — anonymous function variable-এ।
- `hypot(5,12)` → **13**।
- `compute(hypot)` → `hypot(3,4)` = **5**।
- `compute(math.Pow)` → `math.Pow(3,4)` = **81**।

---

## Expected Output

```
13
5
81
```

## মূল শিক্ষা / Key Takeaways

1. **First-class functions** — value হিসেবে ব্যবহার।
2. **Anonymous literal** — `func(...) {...}`।
3. **Function-type param** — `func(float64, float64) float64`।
4. **Library function passing** — `math.Pow` সরাসরি।

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

`math` (Sqrt, Pow).

### Lines 8–10

```go
func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}
```

A **function-type parameter** — accepts any two-float → float function.

### Lines 12–20

```go
func main() {
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}

	fmt.Println(hypot(5, 12))

	fmt.Println(compute(hypot))
	fmt.Println(compute(math.Pow))
}
```

- `hypot := func(...)` — an anonymous function in a variable.
- `hypot(5,12)` → **13**.
- `compute(hypot)` → `hypot(3,4)` = **5**.
- `compute(math.Pow)` → `math.Pow(3,4)` = **81**.

---

## Expected Output

```
13
5
81
```

## Key Takeaways

1. **First-class functions** — usable as values.
2. **Anonymous literal** — `func(...) {...}`.
3. **Function-type param** — `func(float64, float64) float64`.
4. **Passing a library function** — `math.Pow` directly.