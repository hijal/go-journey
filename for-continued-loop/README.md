# for-continued-loop

Go-তে **`for`-এর while-style ব্যবহার** — condition-কেন্দ্রিক লুপ, `sum += sum` দিয়ে doubling।

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

### Lines 5–11

```go
func main() {
	sum := 1

	for sum < 1000 {
		sum += sum
	}
	fmt.Println("sum:", sum)
}
```

**`for sum < 1000 { ... }`** — Go-র সিনট্যাক্সে `while`-এর জায়গায় শুধু `for` + condition। ধাপ:

1. `sum=1` → doubling: 1, 2, 4, 8, ... 512, **1024**।
2. `1024 >= 1000` হওয়ার সঙ্গে-ই লুপ থামে।
3. Print → **1024**।

---

## Expected Output

```
sum: 1024
```

## মূল শিক্ষা / Key Takeaways

1. **`for` = while** — কোনো জটিলতা নেই, শুধু condition।
2. **No `;` needed** — শুধু condition-এর লুপ।
3. **Doubling** — 1 থেকে 1000-এর উপরে যাওয়া = 10 iteration।
4. **exit condition** — প্লেইন উদাহরণে sum ≥ 1000।

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

### Lines 5–11

```go
func main() {
	sum := 1

	for sum < 1000 {
		sum += sum
	}
	fmt.Println("sum:", sum)
}
```

**`for sum < 1000 { ... }`** — the place of `while` in Go is just `for` + a condition.

1. `sum=1` → doubles: 1, 2, 4, 8, ... 512, **1024**.
2. As soon as `1024 >= 1000`, the loop stops.
3. Print → **1024**.

---

## Expected Output

```
sum: 1024
```

## Key Takeaways

1. **`for` = while** — no ceremony, just a condition.
2. **No `;` needed** — a lone-condition loop syntax.
3. **Doubling** — 1 to past 1000 takes 10 iterations.
4. **Exit condition** — the loop ends when `sum ≥ 1000`.