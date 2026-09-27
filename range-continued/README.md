# range-continued

Go-তে **range-এ skip** — index-শুধু `for i := range pow`, value-শুধু `for _, value := range`।

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

### Lines 5–14

```go
func main() {
	pow := make([]int, 10)

	for i := range pow {
		pow[i] = 1 << uint(i)
	}

	for _, value := range pow {
		fmt.Printf("%d\n", value)
	}
}
```

- `for i := range pow` — শুধু index (value skip)।
- `pow[i] = 1 << uint(i)` — `1<<0..9` → 2-এর ঘাত।
- `for _, value := range` — index discard, value print।

---

## Expected Output

```
1
2
4
8
16
32
64
128
256
512
```

## মূল শিক্ষা / Key Takeaways

1. **Index-only range** — value ignore।
2. **Value-only (`_`)** — index discard।
3. **`1 << uint(i)`** — bit-shift দিয়ে ঘাত।
4. **Write-then-read** — দুই-pass pattern।

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

### Lines 5–14

```go
func main() {
	pow := make([]int, 10)

	for i := range pow {
		pow[i] = 1 << uint(i)
	}

	for _, value := range pow {
		fmt.Printf("%d\n", value)
	}
}
```

- `for i := range pow` — index only (value skipped).
- `pow[i] = 1 << uint(i)` — `1<<0..9` → powers of two.
- `for _, value := range` — index discarded, value printed.

---

## Expected Output

```
1
2
4
8
16
32
64
128
256
512
```

## Key Takeaways

1. **Index-only range** — ignore the value.
2. **Value-only (`_`)** — discard the index.
3. **`1 << uint(i)`** — a power via bit shift.
4. **Write then read** — a two-pass pattern.