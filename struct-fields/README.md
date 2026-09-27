# struct-fields

Go-তে **struct field access + mutation** — dot-notation দিয়ে read/write।

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

### Lines 5–7

```go
type Vertex struct {
	X, Y int
}
```

`X, Y int` — একই type-র একাধিক field এক-লাইনে।

### Lines 9–14

```go
func main() {
	v := Vertex{1, 2}
	fmt.Println(v)
	v.X = 4
	fmt.Println(v.X)
	fmt.Println(v.Y)
}
```

- `Vertex{1, 2}` → `v`।
- `v.X = 4` — **field mutation**।
- `v.X` (৪-এ) এবং `v.Y` (২) আলাদা print।

---

## Expected Output

```
{1 2}
4
2
```

## মূল শিক্ষা / Key Takeaways

1. **Dot access** — `v.X` field read।
2. **Field mutation** — struct variable-এর field-এ মান বদল।
3. **`X, Y int`** — multi-field declaration।
4. **Struct print** — `{1 2}` shorthand।

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

### Lines 5–7

```go
type Vertex struct {
	X, Y int
}
```

`X, Y int` — multiple fields of the same type on one line.

### Lines 9–14

```go
func main() {
	v := Vertex{1, 2}
	fmt.Println(v)
	v.X = 4
	fmt.Println(v.X)
	fmt.Println(v.Y)
}
```

- `Vertex{1, 2}` → `v`.
- `v.X = 4` — **field mutation**.
- `v.X` (now 4) and `v.Y` (still 2) printed separately.

---

## Expected Output

```
{1 2}
4
2
```

## Key Takeaways

1. **Dot access** — `v.X` reads a field.
2. **Field mutation** — struct values are mutable.
3. **`X, Y int`** — a compact multi-field declaration.
4. **Struct print** — the `{1 2}` shorthand.