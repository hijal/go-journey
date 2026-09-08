# merge-multiple-slices

Go-তে **variadic `...[]int`** আর **`append` with `...` unpack** দিয়ে একাধিক slice merge শেখার ছোট example।

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

Console-এ output print করার জন্য `fmt` package import করা হয়।

### Lines 5–11

```go
func MergeSlices(slices ...[]int) []int {
	var result []int
	for _, slice := range slices {
		result = append(result, slice...)
	}
	return result
}
```

`MergeSlices` — একাধিক `[]int` slice-কে একটাতে merge:

- `slices ...[]int` — **variadic of slices**: একাধিক `[]int` argument, প্রতিটা নিজেই একটা slice (`[] []int`)।
- `var result []int` — খালি result slice (nil-সহ, `append` দিয়ে grow)।
- Loop-এ প্রতিটা `slice`:
  - `result = append(result, slice...)` — **`slice...`** unpack: slice-টার elements-গুলো আলাদা element হিসেবে `append`-তে। ছাড়া `...`-এ slice-টা force whole-slice-append করত না (append-এ দ্বিতীয় argument-এ single element নেয় — slice-টা whole-append করতে `...` require)।
- সব slice-র elements ক্রম-সহ একটাতে।

**`append` shape:** `append(dst, src...)` — সবচেয়ে common slice-merge idiom। Underlying array-র capacity প্রযোজ্য হলে grow, না হলে নতুন array।

### Line 13

```go
func main() {
```

Program-এর entry point।

### Lines 14–16

```go
s1 := []int{1, 2}
s2 := []int{3, 4, 5}
s3 := []int{6}
```

তিনটা slice: `[1 2]`, `[3 4 5]`, `[6]`।

### Line 18

```go
fmt.Println(MergeSlices(s1, s2, s3))
```

`MergeSlices(s1, s2, s3)` — merge করে `[1 2 3 4 5 6]`।

### Line 19

```go
}
```

Closing brace — `main` function শেষ হয়।

---

## Expected Output

```
[1 2 3 4 5 6]
```

## মূল শিক্ষা / Key Takeaways

1. **Variadic of slices** — `...[]int` — each argument is its own slice.
2. **`slice...` unpack** — slice-কে `append`-তে individual elements-এ বিচ্ছিন্ন।
3. **`append(dst, src...)`** — স্ট্যান্ডার্ড slice-merge idiom।
4. **Nil-ready** — `var result []int`-এ `append` নিরাপদে।

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

Imports the `fmt` package for console output.

### Lines 5–11

```go
func MergeSlices(slices ...[]int) []int {
	var result []int
	for _, slice := range slices {
		result = append(result, slice...)
	}
	return result
}
```

`MergeSlices` — merges multiple `[]int` slices into one:

- `slices ...[]int` — **variadic of slices**: multiple `[]int` arguments, each itself a slice (a `[] []int`).
- `var result []int` — an empty result slice (works from nil; grows via `append`).
- The loop over each `slice`:
  - `result = append(result, slice...)` — the **`slice...`** unpack: the slice's elements are passed as separate elements to `append`. Without `...`, a slice couldn't be whole-appended (append takes a single element as its second arg — you need `...` to unpack).

All elements of every slice are concatenated in order.

**`append` shape:** `append(dst, src...)` — the most common slice-merge idiom. It grows the underlying array when capacity runs out, otherwise reuses it.

### Line 13

```go
func main() {
```

Program entry point.

### Lines 14–16

```go
s1 := []int{1, 2}
s2 := []int{3, 4, 5}
s3 := []int{6}
```

Three slices: `[1 2]`, `[3 4 5]`, `[6]`.

### Line 18

```go
fmt.Println(MergeSlices(s1, s2, s3))
```

`MergeSlices(s1, s2, s3)` — merges into `[1 2 3 4 5 6]`.

### Line 19

```go
}
```

Closing brace — ends the `main` function.

---

## Expected Output

```
[1 2 3 4 5 6]
```

## Key Takeaways

1. **Variadic of slices** — `...[]int` — each argument is its own slice.
2. **`slice...` unpack** — spread a slice into individual elements for `append`.
3. **`append(dst, src...)`** — the standard slice-merge idiom.
4. **Nil-ready** — `append` is safe on `var result []int`.