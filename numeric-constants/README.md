# numeric-constants

Go-তে **untyped numeric constants + type compatibility** — `1<<100`-এর মতো giant constant-কে `int`/`float64`-এ use করার নিয়ম।

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

### Lines 5–8

```go
const (
	Big   = 1 << 100
	Small = Big >> 99
)
```

**Untyped constants** — type নির্দিষ্ট নয়, value-ই সব।

- `Big` — `1<<100` — int-এ রাখা **অসম্ভব** (huge value)।
- `Small` — `Big >> 99` = `1 << 1` = **2**।

### Lines 10–11

```go
func needInt(x int) int           { return x*10 + 1 }
func needFloat(x float64) float64 { return x * 0.1 }
```

দুইটা typed function: একটা-টা `int` চায়, আরেকটা `float64`।

### Lines 13–16

```go
func main() {
	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
}
```

- `needInt(Small)` — Small (2) `int`-এ মানায় → **21**।
- `needFloat(Small)` → **0.2**।
- `needFloat(Big)` — Big `float64`-এ মানায় → **~1.27e29**। (`needInt(Big)` করলে **compile error** — overflow!)

---

## Expected Output

```
21
0.2
1.2676506002282295e+29
```

## মূল শিক্ষা / Key Takeaways

1. **Untyped constant** — context-এর type-এ auto-adapt।
2. **Huge constants** — compile-time arbitrary precision।
3. **Overflow** — fit-না-করা স্থানে use-ই compile error।
4. **Shift-ভিত্তিক constants** — `1<<n` idiom।
5. **Type-driven selection** — কোন value কোন function-এ যায় — কনভার্শন-ই আসল বিষয়।

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

### Lines 5–8

```go
const (
	Big   = 1 << 100
	Small = Big >> 99
)
```

**Untyped constants** — no type specified, the value is everything.

- `Big` — `1<<100` — **impossible** to store in an int (a huge value).
- `Small` — `Big >> 99` = `1 << 1` = **2**.

### Lines 10–11

```go
func needInt(x int) int           { return x*10 + 1 }
func needFloat(x float64) float64 { return x * 0.1 }
```

Two typed functions — one wants an `int`, the other a `float64`.

### Lines 13–16

```go
func main() {
	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
}
```

- `needInt(Small)` — 2 fits an `int` → **21**.
- `needFloat(Small)` → **0.2**.
- `needFloat(Big)` — Big fits a `float64` → **~1.27e29**. (`needInt(Big)` would be a **compile error** — overflow!)

---

## Expected Output

```
21
0.2
1.2676506002282295e+29
```

## Key Takeaways

1. **Untyped constant** — auto-adapts to the context's type.
2. **Huge constants** — compile-time arbitrary precision.
3. **Overflow** — using it where it doesn't fit is a compile error.
4. **Shift-based constants** — the `1<<n` idiom.
5. **Type-driven selection** — which function each value goes to is what matters.