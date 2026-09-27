# slices-basics

Go-তে **slice প্রাইমারি** — array-এর window-view `primes[1:4]`, কোনো data copy নয়।

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

### Lines 5–8

```go
func main() {
	primes := [6]int{2, 3, 5, 7, 11, 13}
	var s []int = primes[1:4]
	fmt.Println(s)
}
```

- **Slice expression** — `primes[1:4]` = element 1,2,3 (4 বাদ) → `[3 5 7]`।
- Slice তে data অনুলিপি হয় না — শুধু array-র **view**।

---

## Expected Output

```
[3 5 7]
```

## মূল শিক্ষা / Key Takeaways

1. **`a[low:high]`** — half-open range: high বাদে।
2. **`[]int`** — array-এর মতো, কিন্তু size-bound নয়।
3. **View, not copy** — ভেতরের data একই।
4. **Derived type** — array-থেকে slice তৈরি।

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

### Lines 5–8

```go
func main() {
	primes := [6]int{2, 3, 5, 7, 11, 13}
	var s []int = primes[1:4]
	fmt.Println(s)
}
```

- **Slice expression** — `primes[1:4]` = elements 1,2,3 (4 excluded) → `[3 5 7]`.
- A slice does not copy data — it's just a **view** of the array.

---

## Expected Output

```
[3 5 7]
```

## Key Takeaways

1. **`a[low:high]`** — a half-open range: high is excluded.
2. **`[]int`** — like an array, but no fixed size bound.
3. **View, not copy** — the underlying data is shared.
4. **Derived type** — a slice built from an array.