# pointers-basics

Go-তে **pointer + dereference** — `&` address, `*`-এর মাধ্যমে value read/write।

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

### Lines 5–18

```go
func main() {
	i := 42

	p := &i
	fmt.Println("*p:", *p)

	*p = 21
	fmt.Println("i:", i)

	j := 2701
	p = &j

	*p = *p / 37
	fmt.Println("j:", j)
}
```

- `&i` — `i`-র **address** pointer (`p`)।
- `*p` — **dereference**, অর্থাৎ `i`-র value (42)।
- `*p = 21` — pointer-এর মাধ্যমে `i`-তে লিখে → i = 21।
- `p = &j` — pointer অন্য variable-কে নির্দেশ।
- `*p / 37` → `2701/37 = 73`।

---

## Expected Output

```
*p: 42
i: 21
j: 73
```

## মূল শিক্ষা / Key Takeaways

1. **`&` operator** — variable-এর address।
2. **`*` dereference** — সেই address-এর value।
3. **Assign via pointer** — মূল variable-ই বদলায়।
4. **Reassign pointer** — pointer-কে নতুন target-এ নির্দেশ।

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

### Lines 5–18

```go
func main() {
	i := 42

	p := &i
	fmt.Println("*p:", *p)

	*p = 21
	fmt.Println("i:", i)

	j := 2701
	p = &j

	*p = *p / 37
	fmt.Println("j:", j)
}
```

- `&i` — the **address** of `i`, stored in pointer `p`.
- `*p` — **dereference**, i.e. the value of `i` (42).
- `*p = 21` — writes through the pointer, so `i` becomes 21.
- `p = &j` — the pointer now targets another variable.
- `*p / 37` → `2701/37 = 73`.

---

## Expected Output

```
*p: 42
i: 21
j: 73
```

## Key Takeaways

1. **`&` operator** — the address of a variable.
2. **`*` dereference** — the value at that address.
3. **Assign via pointer** — the original variable changes.
4. **Reassign the pointer** — point it at a new target.