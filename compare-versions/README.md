# compare-versions

Go-তে **dotted-version numeric comparison + missing-part = 0 + `max` loop + `Atoi` guard** শেখার ছোট example — semver-style compare (-1/0/1)।

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

### Lines 3–7

```go
import (
	"fmt"
	"strconv"
	"strings"
)
```

- `fmt` — `Printf`, `Errorf`।
- `strconv` — `Atoi`।
- `strings` — `Split`।

### Lines 9–42

```go
func compareVersions(a, b string) (int, error) {

	as := strings.Split(a, ".")
	bs := strings.Split(b, ".")

	for i := range max(len(as), len(bs)) {
		av, bv := 0, 0

		if i < len(as) {
			n, err := strconv.Atoi(as[i])
			if err != nil {
				return 0, fmt.Errorf("bad version %q: part %q is not a number", a, as[i])
			}
			av = n
		}

		if i < len(bs) {
			n, err := strconv.Atoi(bs[i])
			if err != nil {
				return 0, fmt.Errorf("bad version %q: part %q is not a number", b, bs[i])
			}
			bv = n
		}

		switch {
		case av < bv:
			return -1, nil
		case av > bv:
			return 1, nil
		}
	}

	return 0, nil
}
```

**The comparison logic:**

- `strings.Split(a, ".")` — dot-partition।
- `for i := range max(len(as), len(bs))` — long অংশ বরাবর loop (`max` builtin)।
- **Missing part → 0**: যদি `i` slice-এর বাইরে, `av`/`bv` ডিফল্ট `0`। তাই `2.0` vs `2.0.0` → `0`।
- `strconv.Atoi` — প্রতিটি অংশ numeric নয়-হলে error (`"1.x"` → reject)।
- `switch` — প্রতিটি position-এ **numeric** compare: `<` → `-1`, `>` → `1`।
- সব position সমান → `0`।

*(String-compare হলে `"1.10.0" < "1.9"` ভুল দেখাতো; numeric compare-ই `10 > 9`।)*

### Lines 44–59

```go
func main() {
	pairs := [][2]string{
		{"1.10.0", "1.9.5"},
		{"2.0", "2.0.0"},
		{"0.9", "1.0"},
		{"1.x", "1.0"},
	}

	for _, p := range pairs {
		result, err := compareVersions(p[0], p[1])
		if err != nil {
			fmt.Printf("%s vs %s -> error: %v\n", p[0], p[1], err)
			continue
		}
		fmt.Printf("%s vs %s -> %d\n", p[0], p[1], result)
	}
}
```

৪টা pair:

- `1.10.0` vs `1.9.5` — numeric: `10 > 9` → `1`
- `2.0` vs `2.0.0` — missing-part 0 → `0`
- `0.9` vs `1.0` — `0 < 1` → `-1`
- `1.x` vs `1.0` — "x" numeric নয় → error

---

## Expected Output

```
1.10.0 vs 1.9.5 -> 1
2.0 vs 2.0.0 -> 0
0.9 vs 1.0 -> -1
1.x vs 1.0 -> error: bad version "1.x": part "x" is not a number
```

## মূল শিক্ষা / Key Takeaways

1. **Numeric (not lexicographic) compare** — `10 > 9` ঠিকভাবে।
2. **Missing-part → 0** — অসম-দৈর্ঘ্য version handle।
3. **`max` builtin** — loop দৈর্ঘ্য।
4. **`strconv.Atoi` + `%q` error** — invalid part diagnose।
5. **Early `-1/1` return** — প্রথম পার্থক্যই result।

---

---

<a name="english"></a>

## 🇬🇧 English Version

### Line 1

```go
package main
```

Declares an executable program (`main` package), runnable via `go run`.

### Lines 3–7

```go
import (
	"fmt"
	"strconv"
	"strings"
)
```

- `fmt` — for `Printf`, `Errorf`.
- `strconv` — for `Atoi`.
- `strings` — for `Split`.

### Lines 9–42

```go
func compareVersions(a, b string) (int, error) {

	as := strings.Split(a, ".")
	bs := strings.Split(b, ".")

	for i := range max(len(as), len(bs)) {
		av, bv := 0, 0

		if i < len(as) {
			n, err := strconv.Atoi(as[i])
			if err != nil {
				return 0, fmt.Errorf("bad version %q: part %q is not a number", a, as[i])
			}
			av = n
		}

		if i < len(bs) {
			n, err := strconv.Atoi(bs[i])
			if err != nil {
				return 0, fmt.Errorf("bad version %q: part %q is not a number", b, bs[i])
			}
			bv = n
		}

		switch {
		case av < bv:
			return -1, nil
		case av > bv:
			return 1, nil
		}
	}

	return 0, nil
}
```

**The comparison logic:**

- `strings.Split(a, ".")` — split on dots.
- `for i := range max(len(as), len(bs))` — iterate over the longer side (the `max` builtin).
- **Missing part → 0**: if `i` is beyond a slice, `av`/`bv` stay `0`. So `2.0` vs `2.0.0` → `0`.
- `strconv.Atoi` — a non-numeric part (`"1.x"`) returns an error.
- `switch` — a **numeric** compare at each position: `<` → `-1`, `>` → `1`.
- All positions equal → `0`.

*(A string compare would wrongly say `"1.10.0" < "1.9"`; numeric compare gets `10 > 9`.)*

### Lines 44–59

```go
func main() {
	pairs := [][2]string{
		{"1.10.0", "1.9.5"},
		{"2.0", "2.0.0"},
		{"0.9", "1.0"},
		{"1.x", "1.0"},
	}

	for _, p := range pairs {
		result, err := compareVersions(p[0], p[1])
		if err != nil {
			fmt.Printf("%s vs %s -> error: %v\n", p[0], p[1], err)
			continue
		}
		fmt.Printf("%s vs %s -> %d\n", p[0], p[1], result)
	}
}
```

4 pairs:

- `1.10.0` vs `1.9.5` — numerically `10 > 9` → `1`
- `2.0` vs `2.0.0` — missing part is 0 → `0`
- `0.9` vs `1.0` — `0 < 1` → `-1`
- `1.x` vs `1.0` — "x" is not numeric → error

---

## Expected Output

```
1.10.0 vs 1.9.5 -> 1
2.0 vs 2.0.0 -> 0
0.9 vs 1.0 -> -1
1.x vs 1.0 -> error: bad version "1.x": part "x" is not a number
```

## Key Takeaways

1. **Numeric (not lexicographic) compare** — `10 > 9` handled.
2. **Missing part → 0** — handles unequal-length versions.
3. **`max` builtin** — loop length.
4. **`strconv.Atoi` + `%q` error** — diagnosing bad parts.
5. **Early `-1/1` return** — first difference decides.