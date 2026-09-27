# function-closures

Go-তে **closure captures state** — `adder`-এর প্রতিটি instance-এর নিজস্ব `sum`।

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

### Lines 5–12

```go
func adder() func(int) int {
	sum := 0

	return func(i int) int {
		sum += i
		return sum
	}
}
```

**Closure** — ভেতরের function-টা `sum`-কে **capture** করে, বাইরের scope-এর variable বেঁচে থাকে।

### Lines 14–22

```go
func main() {
	pos, neg := adder(), adder()

	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}
}
```

- `adder()` দুইবার → **দুইটা স্বাধীন closure** (`pos`, `neg`)।
- `pos(i)` cumulative-sum: 0,1,3,6,10,...
- `neg(-2*i)` cumulative-negative: 0,-2,-6,-12,...

---

## Expected Output

```
0 0
1 -2
3 -6
6 -12
10 -20
15 -30
21 -42
28 -56
36 -72
45 -90
```

## মূল শিক্ষা / Key Takeaways

1. **Closure** — নিজের scope-র variable ধরে রাখে।
2. **Stateful function** — প্রতিটি call-এ মান বদলায়।
3. **Independent instances** — `pos` ও `neg` আলাদা sum।
4. **Function-type return** — `func(int) int`।

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

### Lines 5–12

```go
func adder() func(int) int {
	sum := 0

	return func(i int) int {
		sum += i
		return sum
	}
}
```

A **closure** — the inner function captures `sum`, keeping the outer variable alive.

### Lines 14–22

```go
func main() {
	pos, neg := adder(), adder()

	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}
}
```

- Calling `adder()` twice → **two independent closures** (`pos`, `neg`).
- `pos(i)` accumulates: 0,1,3,6,10,...
- `neg(-2*i)` accumulates negatives: 0,-2,-6,-12,...

---

## Expected Output

```
0 0
1 -2
3 -6
6 -12
10 -20
15 -30
21 -42
28 -56
36 -72
45 -90
```

## Key Takeaways

1. **Closure** — keeps variables from its own scope.
2. **Stateful function** — each call changes the captured value.
3. **Independent instances** — `pos` and `neg` have separate `sum`s.
4. **Function-type return** — `func(int) int`.