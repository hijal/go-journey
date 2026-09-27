# range-basics

Go-তে **`for range` with index+value** — `2**i` আউটপুট।

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

### Line 5

```go
var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}
```

2-এর পাওয়ার-গুলো।

### Lines 7–10

```go
func main() {
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}
}
```

- **`for i, v := range pow`** — প্রতিটি element-এ `i` index, `v` value।
- `2**%d` → 2-এর ঘাত।

---

## Expected Output

```
2**0 = 1
2**1 = 2
2**2 = 4
...
2**7 = 128
```

## মূল শিক্ষা / Key Takeaways

1. **Two return values** — `i` index, `v` copy-of-value।
2. **`range` any slice/array** — sequential visit।
3. **`%d` verb** — numeric formatting।
4. **Power table** — 2-এর ঘাত full।

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

### Line 5

```go
var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}
```

Powers of two.

### Lines 7–10

```go
func main() {
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}
}
```

- **`for i, v := range pow`** — `i` is the index, `v` the value at each element.
- `2**%d` → an exponent of 2.

---

## Expected Output

```
2**0 = 1
2**1 = 2
2**2 = 4
...
2**7 = 128
```

## Key Takeaways

1. **Two return values** — `i` index and `v` a value copy.
2. **`range` over a slice** — a sequential visit.
3. **`%d` verb** — numeric formatting.
4. **Power table** — full list of 2's powers.