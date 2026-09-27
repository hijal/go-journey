# if-else-basics

Go-তে **`if` with init statement + `else` branch** — `pow` helper-এ scoped variable + condition।

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

### Lines 8–15

```go
func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	return lim
}
```

**`if` with short statement** — `v := math.Pow(...)` যেটা কেবল `if`/`else` block-এর ভেতরে দৃশ্যমান (scoped)।

- `v < lim` → ভেতরে থেকে return।
- নাহলে → `else`-এ print, তারপর `lim` return।

### Lines 17–21

```go
func main() {
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
}
```

- `pow(3,2,10)`: 9 < 10 → 9 (কোনো print নেই)।
- `pow(3,3,20)`: 27 ≥ 20 → **"27 >= 20"** print, ফলে 20।

---

## Expected Output

```
27 >= 20
9 20
```

## মূল শিক্ষা / Key Takeaways

1. **Init statement in `if`** — scope সংকুচিত রাখে।
2. **`else` branch** — সব pathway covered।
3. **Scoped variable** — `if`-এর বাইরে ব্যবহার করা যায় না।
4. **Mixed print/return** — শিক্ষণীয় side-effect।

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

### Lines 8–15

```go
func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	return lim
}
```

**An `if` with a short statement** — `v := math.Pow(...)` is visible only inside the `if`/`else` blocks (scoped).

- If `v < lim` → return from inside.
- Otherwise → the `else` prints, then `lim` is returned.

### Lines 17–21

```go
func main() {
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
}
```

- `pow(3,2,10)`: 9 < 10 → 9 (no print).
- `pow(3,3,20)`: 27 ≥ 20 → prints **"27 >= 20"**, then 20.

---

## Expected Output

```
27 >= 20
9 20
```

## Key Takeaways

1. **Init statement in `if`** — keeps the scope tight.
2. **The `else` branch** — every path is covered.
3. **Scoped variable** — unusable outside the `if`.
4. **Mixed print/return** — an instructive side-effect.