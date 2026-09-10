# basic-array

Go-তে **fixed-size array** (`[4]int`) + ৪ ধরণের `range` loop শেখার ছোট example — index/value iteration, in-place mutation, value-only, আর Go 1.22-এর int-range।

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

### Lines 3–5

```go
import (
	"fmt"
)
```

`fmt` — `Println`।

### Lines 7–8

```go
func main() {
	a := [4]int{5, 10, 15, 20}
```

- `[4]int{...}` — **fixed-size array**: ঠিক 4টা element। Slice-এর মতো dynamic না; size type-এর অংশ।
- Element: `5, 10, 15, 20`।

### Lines 10–12

```go
	for i, v := range a {
		fmt.Println(i, v)
	}
```

১ম loop — **index + value**:

- `range a` — প্রতিটা iteration-এ `i` (index) আর `v` (value copy)।
- `Println(i, v)` — `0 5`, `1 10`, ...।

### Lines 13–17

```go
	fmt.Println("-----")
	for i := range a {
		a[i] *= 2
		fmt.Println(a[i])
	}
```

২য় loop — **index-only + in-place mutation**:

- `for i := range a` — শুধু index।
- `a[i] *= 2` — array-র element-টা সরাসরি double: `5→10, 10→20, 15→30, 20→40`।
- Type-এর **fixed-size** বলে `a[i]`-কে সরাসরি mutate করা যায় (len জানা)।

### Lines 18–21

```go
	fmt.Println("-----")
	for _, v := range a {
		fmt.Println(v)
	}
```

৩য় loop — **value-only**:

- `_` দিয়ে index ignore, `v` print — page-টা doubled value: `10, 20, 30, 40`।

### Lines 22–25

```go
	fmt.Println("-----")
	for i := range len(a) {
		fmt.Println(i)
	}
}
```

৪র্থ loop — **int-range (Go 1.22+)**:

- `for i := range len(a)` — range-টা একটা **integer**: `0` থেকে `len(a)-1` পর্যন্ত। এখানে `0 1 2 3`।
- শুধু index ইতর Loop-running-এর জন্য সবচেয়ে concise — `len`-এ implicit.

---

## Expected Output

```
0 5
1 10
2 15
3 20
-----
10
20
30
40
-----
10
20
30
40
-----
0
1
2
3
```

## মূল শিক্ষা / Key Takeaways

1. **Fixed-size array** — `[4]int`, size type-এর অংশ (slice নয়)।
2. **`for i, v := range`** — index + value।
3. **In-place mutation** — `a[i] *= 2`, index-এর মাধ্যমে element-টা বদলানো।
4. **Value-only** — `_` দিয়ে index discard।
5. **Int-range** — `for i := range len(a)` (Go 1.22+), শুধু index-এর জন্য।

---

---

<a name="english"></a>

## 🇬🇧 English Version

### Line 1

```go
package main
```

Declares an executable program (`main` package), runnable via `go run`.

### Lines 3–5

```go
import (
	"fmt"
)
```

`fmt` — for `Println`.

### Lines 7–8

```go
func main() {
	a := [4]int{5, 10, 15, 20}
```

- `[4]int{...}` — a **fixed-size array**: exactly 4 elements. Unlike a slice, the size is part of the type.
- Elements: `5, 10, 15, 20`.

### Lines 10–12

```go
	for i, v := range a {
		fmt.Println(i, v)
	}
```

1st loop — **index + value**:

- `range a` — each iteration yields `i` (index) and `v` (a copy of the value).
- `Println(i, v)` — `0 5`, `1 10`, ....

### Lines 13–17

```go
	fmt.Println("-----")
	for i := range a {
		a[i] *= 2
		fmt.Println(a[i])
	}
```

2nd loop — **index-only + in-place mutation**:

- `for i := range a` — index only.
- `a[i] *= 2` — doubles the array element directly: `5→10, 10→20, 15→30, 20→40`.
- Because the size is **fixed**, `a[i]` can be mutated in place (the length is known).

### Lines 18–21

```go
	fmt.Println("-----")
	for _, v := range a {
		fmt.Println(v)
	}
```

3rd loop — **value-only**:

- `_` discards the index; `v` prints the now-doubled values: `10, 20, 30, 40`.

### Lines 22–25

```go
	fmt.Println("-----")
	for i := range len(a) {
		fmt.Println(i)
	}
}
```

4th loop — **int-range (Go 1.22+)**:

- `for i := range len(a)` — ranges over an **integer**: from `0` to `len(a)-1`. Here: `0 1 2 3`.
- Index-only iteration is most concise with this — implicit `len`.

---

## Expected Output

```
0 5
1 10
2 15
3 20
-----
10
20
30
40
-----
10
20
30
40
-----
0
1
2
3
```

## Key Takeaways

1. **Fixed-size array** — `[4]int`, the size is part of the type (not a slice).
2. **`for i, v := range`** — index + value.
3. **In-place mutation** — `a[i] *= 2`, modifying elements via index.
4. **Value-only** — discard the index with `_`.
5. **Int-range** — `for i := range len(a)` (Go 1.22+), for index-only iteration.