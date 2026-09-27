# pointers-to-structs

Go-তে **struct pointer + implicit dereference** — `(*p).X`-এর জায়গায় `p.X`।

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

### Lines 9–15

```go
func main() {
	v := Vertex{1, 2}

	p := &v
	p.X = 1e9

	fmt.Println(v)
}
```

- `p := &v` — struct-এর pointer।
- `p.X = 1e9` — **implicit dereference**: `(*p).X`-এর সংক্ষিপ্ত রূপ, মূল `v`-তে প্রভাব।
- Print → `{1000000000 2}`।

---

## Expected Output

```
{1000000000 2}
```

## মূল শিক্ষা / Key Takeaways

1. **Struct pointer** — `&v` দিয়ে struct-কে নির্দেশ।
2. **Implicit dereference** — `p.X` = `(*p).X`।
3. **Mutation in place** — মূল struct বদলায়।
4. **`1e9`** — float literal, X int-এ 1000000000।

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

### Lines 9–15

```go
func main() {
	v := Vertex{1, 2}

	p := &v
	p.X = 1e9

	fmt.Println(v)
}
```

- `p := &v` — a pointer to the struct.
- `p.X = 1e9` — **implicit dereference**: shorthand for `(*p).X`, affecting `v` itself.
- Prints **{1000000000 2}**.

---

## Expected Output

```
{1000000000 2}
```

## Key Takeaways

1. **Struct pointer** — `&v` points to the struct.
2. **Implicit dereference** — `p.X` means `(*p).X`.
3. **In-place mutation** — the original struct changes.
4. **`1e9`** — a float literal; stored as 1000000000 in `X`.