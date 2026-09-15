# swap-function

Go-তে **pointer-based swap (`*a, *b = *b, *a`) + slice-element address** শেখার ছোট example।

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

`fmt` — `Println`।

### Lines 5–7

```go
func Swap(a, b *int) {
	*a, *b = *b, *a
}
```

**Pointer swap** — `*a, *b = *b, *a` parallel assignment:

- ডানে evaluate: `*b`, `*a`;
- বামে assign: `*a` → পুরনো `*b`, `*b` → পুরনো `*a`।

*(না-হলে temp variable লাগত; parallel assignment-এ একটা statement-ই।)*

### Lines 9–14

```go
	x, y := 10, 20

	fmt.Println("Before:", x, y)
	Swap(&x, &y)
	fmt.Println("After:", x, y)
```

`Swap(&x, &y)` — address pass করে; dereference-এ মূল variable-ই বদলায়।

### Lines 16–18

```go
	prices := []int{40, 10, 30}
	Swap(&prices[0], &prices[1])
	fmt.Println("Swapped slice:", prices)
```

**Slice-element address এও কাজ করে** — `&prices[0]`, `&prices[1]` (element-এরই address)।

---

## Expected Output

```
Before: 10 20
After: 20 10
Swapped slice: [10 40 30]
```

## মূল শিক্ষা / Key Takeaways

1. **Parallel assignment** — `*a, *b = *b, *a` (temp-ছাড়া)।
2. **Pointer dereference swap** — মূল variable mutate।
3. **Slice element address** — `&prices[i]`-ও pointer হিসেবে যায়।

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

`fmt` — for `Println`.

### Lines 5–7

```go
func Swap(a, b *int) {
	*a, *b = *b, *a
}
```

**The pointer swap** — `*a, *b = *b, *a` is a parallel assignment:

- Right side evaluated first: `*b`, `*a`;
- then assigned: `*a` → old `*b`, `*b` → old `*a`.

*(Otherwise you'd need a temp variable; parallel assignment does it in one statement.)*

### Lines 9–14

```go
	x, y := 10, 20

	fmt.Println("Before:", x, y)
	Swap(&x, &y)
	fmt.Println("After:", x, y)
```

`Swap(&x, &y)` — passing addresses; the dereference mutates the real variables.

### Lines 16–18

```go
	prices := []int{40, 10, 30}
	Swap(&prices[0], &prices[1])
	fmt.Println("Swapped slice:", prices)
```

**It works on slice-element addresses too** — `&prices[0]`, `&prices[1]` (addresses of the elements themselves).

---

## Expected Output

```
Before: 10 20
After: 20 10
Swapped slice: [10 40 30]
```

## Key Takeaways

1. **Parallel assignment** — `*a, *b = *b, *a` (no temp).
2. **Pointer dereference swap** — mutates the originals.
3. **Slice-element addresses** — `&prices[i]` works as a pointer too.