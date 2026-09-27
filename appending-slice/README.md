# appending-slice

Go-তে **`append` dynamic growth** — capacity-এর বাইরে গেলে slice auto-বড় হয়।

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

### Lines 5–17

```go
func main() {
	var s []int
	printSlice(s)

	s = append(s, 0)
	printSlice(s)

	s = append(s, 1)
	printSlice(s)

	s = append(s, 2, 3, 4)
	printSlice(s)
}
```

- `var s []int` — nil থেকে শুরু।
- **`append(s, ...)`** — প্রায়ই result-পুনঃassign `s =` (capacity-বদলে নতুন slice)।
- Growth: cap 0→1→2→**6** (0,1 append-এ, তারপর 3-element-এ বড় হয়ে)।

### Lines 19–21

```go
func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
```

---

## Expected Output

```
len=0 cap=0 []
len=1 cap=1 [0]
len=2 cap=2 [0 1]
len=5 cap=6 [0 1 2 3 4]
```

## মূল শিক্ষা / Key Takeaways

1. **`append(s, v...)`** — element যোগ।
2. **Assign back `s =`** — growth-এর পর নতুন slice।
3. **Capacity doubling-ish** — অতিরিক্ত ক্ষমতা realloc।
4. **Multi-append** — `append(s, 2, 3, 4)`।

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

### Lines 5–17

```go
func main() {
	var s []int
	printSlice(s)

	s = append(s, 0)
	printSlice(s)

	s = append(s, 1)
	printSlice(s)

	s = append(s, 2, 3, 4)
	printSlice(s)
}
```

- Starts from `var s []int` (nil).
- **`append(s, ...)`** — often reassigns back with `s =` (the capacity may change).
- Growth: cap 0→1→2→**6** (grew for the single appends, then again for three more).

### Lines 19–21

```go
func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
```

---

## Expected Output

```
len=0 cap=0 []
len=1 cap=1 [0]
len=2 cap=2 [0 1]
len=5 cap=6 [0 1 2 3 4]
```

## Key Takeaways

1. **`append(s, v...)`** — adds elements.
2. **Assign back `s =`** — the slice may change after growth.
3. **Capacity growth** — extra room behind a reallocation.
4. **Multi-append** — `append(s, 2, 3, 4)`.