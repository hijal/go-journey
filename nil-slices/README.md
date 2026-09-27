# nil-slices

Go-তে **nil slice** — zero value, `len`/`cap` 0, `nil`-এর সঙ্গে তুলনা।

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

### Lines 5–10

```go
func main() {
	var s []int
	fmt.Println(s, len(s), cap(s))
	if s == nil {
		fmt.Println("nil!")
	}
}
```

- `var s []int` — কোনো literal ছাড়া → **nil slice**।
- `len`/`cap` = 0, `s == nil` → true।

---

## Expected Output

```
[] 0 0
nil!
```

## মূল শিক্ষা / Key Takeaways

1. **Zero value** — নন-আরম্ভিত slice `nil`।
2. **len/cap 0** — দৃশ্যমান element নেই।
3. **`s == nil`** — explicit-check।
4. **Append-safe** — `append` nil slice-এও কাজ করে।

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

### Lines 5–10

```go
func main() {
	var s []int
	fmt.Println(s, len(s), cap(s))
	if s == nil {
		fmt.Println("nil!")
	}
}
```

- `var s []int` — declared without a literal → a **nil slice**.
- `len`/`cap` are 0, and `s == nil` is true.

---

## Expected Output

```
[] 0 0
nil!
```

## Key Takeaways

1. **Zero value** — an uninitialized slice is `nil`.
2. **len/cap 0** — no elements visible.
3. **`s == nil`** — an explicit check.
4. **Append-safe** — `append` works fine on a nil slice.