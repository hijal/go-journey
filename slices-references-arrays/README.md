# slices-references-arrays

Go-তে **slice = shared view** — দুইটা slice-এর element-এ mutation মূল array-তে দেখা যায়।

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
	names := [4]string{
		"John", "Paul", "George", "Ringo",
	}

	fmt.Println(names)

	a := names[0:2]
	b := names[1:3]

	fmt.Println(a, b)

	b[0] = "XXX"
	fmt.Println(a, b)
	fmt.Println(names)
}
```

- `a = names[0:2]` → [John Paul], `b = names[1:3]` → [Paul George] — দুটোই একই `names` array-কে view করে।
- `b[0] = "XXX"` → b-এর প্রথমটা = a-এর **দ্বিতীয়টা** → `a`-তেও "XXX"।
- মূল `names`-ও বদলে যায়: **slice কোনো data copy নয়**।

---

## Expected Output

```
[John Paul George Ringo]
[John Paul] [Paul George]
[John XXX] [XXX George]
[John XXX George Ringo]
```

## মূল শিক্ষা / Key Takeaways

1. **Overlapping slices** — shared underlying array।
2. **Mutation visible everywhere** — আসল data সবার কাছে।
3. **No copy on slice** — slice হচ্ছে view।
4. **`b[0]` → `a[1]`** — overlap-এর সূক্ষ্মতা।

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
	names := [4]string{
		"John", "Paul", "George", "Ringo",
	}

	fmt.Println(names)

	a := names[0:2]
	b := names[1:3]

	fmt.Println(a, b)

	b[0] = "XXX"
	fmt.Println(a, b)
	fmt.Println(names)
}
```

- `a = names[0:2]` → [John Paul], `b = names[1:3]` → [Paul George] — both view the same `names` array.
- `b[0] = "XXX"` → the first of `b` is the **second of `a`** → "XXX" shows up in `a` too.
- The original `names` changes as well: **a slice copies no data**.

---

## Expected Output

```
[John Paul George Ringo]
[John Paul] [Paul George]
[John XXX] [XXX George]
[John XXX George Ringo]
```

## Key Takeaways

1. **Overlapping slices** — shared underlying array.
2. **Mutation visible everywhere** — the real data is shared.
3. **No copy on slice** — a slice is a view.
4. **`b[0]` → `a[1]`** — the subtlety of overlap.