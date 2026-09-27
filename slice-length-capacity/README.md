# slice-length-capacity

Go-তে **len vs cap** — পুনঃ-স্লাইসিং `s[:0]`, `s[:4]`, `s[2:]`-এ capacity ধরে রাখে।

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

### Lines 5–21

```go
func main() {
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)

	s = s[:0]
	printSlice(s)

	s = s[:4]
	printSlice(s)

	s = s[2:]
	printSlice(s)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
```

- `len` = দৃশ্যমান element সংখ্যা, `cap` = ভেতরের array-র ধারণক্ষমতা।
- `s[:0]` → len 0, কিন্তু `cap` এখনও **6**।
- `s[:4]` → len 4 (উপরে-থেকে resize)।
- `s[2:]` → len 2, cap **4** (উপরে drop করলে capacity কমে)।

---

## Expected Output

```
len=6 cap=6 [2 3 5 7 11 13]
len=0 cap=6 []
len=4 cap=6 [2 3 5 7]
len=2 cap=4 [5 7]
```

## মূল শিক্ষা / Key Takeaways

1. **`len`** — slice-এ এখন element কয়টা।
2. **`cap`** — স্লাইস-start থেকে array-এর শেষ পর্যন্ত।
3. **Re-slice up** — `s[:4]` capacity বাঁচায়।
4. **Re-slice down** — `s[2:]` capacity সঙ্কুচিত।

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

### Lines 5–21

```go
func main() {
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)

	s = s[:0]
	printSlice(s)

	s = s[:4]
	printSlice(s)

	s = s[2:]
	printSlice(s)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
```

- `len` = elements currently visible, `cap` = the underlying array capacity.
- `s[:0]` → len 0, but `cap` is still **6**.
- `s[:4]` → len 4 (grown back up).
- `s[2:]` → len 2, cap **4** (dropping from the front shrinks capacity).

---

## Expected Output

```
len=6 cap=6 [2 3 5 7 11 13]
len=0 cap=6 []
len=4 cap=6 [2 3 5 7]
len=2 cap=4 [5 7]
```

## Key Takeaways

1. **`len`** — how many elements the slice has now.
2. **`cap`** — from the slice start to the end of the array.
3. **Re-slice upward** — `s[:4]` keeps its capacity.
4. **Re-slice downward** — `s[2:]` shrinks capacity.