# arrays-basics

Go-তে **fixed-size array** — `[n]T` type, index-এ access, literal।

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

### Lines 5–13

```go
func main() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)
}
```

- `var a [2]string` — **fixed-size array**: ২টা string।
- Index-এ assign/read, পুরোটা print → `[Hello World]`।
- `[6]int{...}` — literal-সহ initialization → primes।

---

## Expected Output

```
Hello World
[Hello World]
[2 3 5 7 11 13]
```

## মূল শিক্ষা / Key Takeaways

1. **`[2]string`** — size type-এর অংশ।
2. **Fixed length** — runtime-এ বাড়ে না।
3. **Array literal** — `[6]int{...}`।
4. **Print whole array** — `[v1 v2 ...]` format-এ।

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

### Lines 5–13

```go
func main() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)
}
```

- `var a [2]string` — a **fixed-size array**: two strings.
- Assign/read by index, and print the whole thing → `[Hello World]`.
- `[6]int{...}` — initialized with a literal → primes.

---

## Expected Output

```
Hello World
[Hello World]
[2 3 5 7 11 13]
```

## Key Takeaways

1. **`[2]string`** — the size is part of the type.
2. **Fixed length** — cannot grow at runtime.
3. **Array literal** — `[6]int{...}`.
4. **Whole-array print** — rendered as `[v1 v2 ...]`.