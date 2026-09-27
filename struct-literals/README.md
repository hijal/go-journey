# struct-literals

Go-তে **struct literal রূপভেদ** — positional, named, empty, pointer।

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

### Lines 9–14

```go
var (
	v1 = Vertex{1, 2}
	v2 = Vertex{X: 1}
	v3 = Vertex{}
	p  = &Vertex{1, 2}
)
```

**Literal-এর ৪ রূপ**:

- `Vertex{1, 2}` — positional: X=1, Y=2।
- `Vertex{X: 1}` — **named field**: বাদ-পড়া field-এর zero value → {1 0}।
- `Vertex{}` — সব zero → {0 0}।
- `&Vertex{1, 2}` — **pointer literal** → `&{1 2}`।

### Lines 16–17

```go
func main() {
	fmt.Println(v1, p, v2, v3)
}
```

---

## Expected Output

```
{1 2} &{1 2} {1 0} {0 0}
```

## মূল শিক্ষা / Key Takeaways

1. **Positional literal** — field-ক্রম-অনুযায়ী।
2. **Named `X:`** — নির্দিষ্ট field-ই সেট, বাকিগুলো zero।
3. **Empty `{}`** — সম্পূর্ণ zero struct।
4. **`&` literal** — heap-এ pointer-ও instant।

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

### Lines 9–14

```go
var (
	v1 = Vertex{1, 2}
	v2 = Vertex{X: 1}
	v3 = Vertex{}
	p  = &Vertex{1, 2}
)
```

**Four literal forms**:

- `Vertex{1, 2}` — positional: X=1, Y=2.
- `Vertex{X: 1}` — **named field**: the omitted field stays zero → {1 0}.
- `Vertex{}` — everything zero → {0 0}.
- `&Vertex{1, 2}` — a **pointer literal** → `&{1 2}`.

### Lines 16–17

```go
func main() {
	fmt.Println(v1, p, v2, v3)
}
```

---

## Expected Output

```
{1 2} &{1 2} {1 0} {0 0}
```

## Key Takeaways

1. **Positional literal** — values in field order.
2. **Named `X:`** — only that field is set; the rest are zero.
3. **Empty `{}`** — an all-zeros struct.
4. **`&` literal** — constructs a pointer in one step.