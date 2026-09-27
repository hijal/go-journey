# slice-with-make

Go-তে **`make` দিয়ে zeroed slice** — `make([]T, len)` ও `make([]T, len, cap)`।

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

### Lines 5–22

```go
func main() {
	a := make([]int, 5)
	printSlice("a", a)

	b := make([]int, 0, 5)
	printSlice("b", b)

	c := b[:2]
	printSlice("c", c)

	d := c[2:5]
	printSlice("d", d)
}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}
```

- `make([]int, 5)` → len 5, cap 5, zero-filled।
- `make([]int, 0, 5)` → len 0, **cap 5**।
- `b[:2]` → len 2, cap ধরে রাখে।
- `c[2:5]` → cap-এর ভেতরে high-bound পর্যন্ত।

---

## Expected Output

```
a len=5 cap=5 [0 0 0 0 0]
b len=0 cap=5 []
c len=2 cap=5 [0 0]
d len=3 cap=3 [0 0 0]
```

## মূল শিক্ষা / Key Takeaways

1. **`make([]int, 5)`** — zeroed, len=cap=5।
2. **`make([]int, 0, 5)`** — pre-allocated capacity।
3. **Slice within cap** — `b[:2]`, `c[2:5]` কোনো growth ছাড়াই।
4. **`d`-র cap 3** — স্লাইস-start থেকে array-শেষ।

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

### Lines 5–22

```go
func main() {
	a := make([]int, 5)
	printSlice("a", a)

	b := make([]int, 0, 5)
	printSlice("b", b)

	c := b[:2]
	printSlice("c", c)

	d := c[2:5]
	printSlice("d", d)
}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}
```

- `make([]int, 5)` → len 5, cap 5, zero-filled.
- `make([]int, 0, 5)` → len 0, **cap 5**.
- `b[:2]` → len 2, capacity preserved.
- `c[2:5]` → index up to the capacity bound.

---

## Expected Output

```
a len=5 cap=5 [0 0 0 0 0]
b len=0 cap=5 []
c len=2 cap=5 [0 0]
d len=3 cap=3 [0 0 0]
```

## Key Takeaways

1. **`make([]int, 5)`** — zeroed, len=cap=5.
2. **`make([]int, 0, 5)`** — pre-allocated capacity.
3. **Slicing within cap** — `b[:2]`, `c[2:5]` without any growth.
4. **`d`'s cap is 3** — from slice start to array end.