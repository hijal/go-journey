# structs-basics

Go-তে **struct definition + construction** — typed field-collection-এর প্রথম ধাপ।

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
type Vertex struct {
	X int
	Y int
}
```

**Struct** — এক-একটা-field-এর collection। `Vertex`-এর দুইটা `int` field: `X`, `Y`।

### Lines 10–11

```go
func main() {
	fmt.Println(Vertex{1, 2})
}
```

`Vertex{1, 2}` — positional struct literal: X=1, Y=2। Print → **{1 2}**।

---

## Expected Output

```
{1 2}
```

## মূল শিক্ষা / Key Takeaways

1. **`type Vertex struct`** — custom composite type।
2. **Fields** — `X`, `Y` প্রত্যেকে `int`।
3. **Positional literal** — ক্রম অনুযায়ী value।
4. **Default print** — `{X Y}` format-এ।

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
type Vertex struct {
	X int
	Y int
}
```

**A struct** — a collection of fields. `Vertex` has two `int` fields: `X`, `Y`.

### Lines 10–11

```go
func main() {
	fmt.Println(Vertex{1, 2})
}
```

`Vertex{1, 2}` — a positional struct literal: X=1, Y=2. Prints **{1 2}**.

---

## Expected Output

```
{1 2}
```

## Key Takeaways

1. **`type Vertex struct`** — a custom composite type.
2. **Fields** — `X` and `Y`, both `int`.
3. **Positional literal** — values in declaration order.
4. **Default print** — rendered as `{X Y}`.