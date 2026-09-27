# mutating-maps

Go-তে **map mutation** — insert, update, delete, comma-ok lookup।

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

### Lines 5–17

```go
func main() {
	m := make(map[string]int)
	m["Answer"] = 42
	fmt.Println("The value:", m["Answer"])

	m["Answer"] = 48
	fmt.Println("The value:", m["Answer"])

	delete(m, "Answer")
	fmt.Println("The value:", m["Answer"])

	v, ok := m["Answer"]
	fmt.Println("The value:", v, "present?", ok)
}
```

- `m[k] = v` — insert/update।
- **`delete(m, k)`** — key মুছে ফেলে।
- Absent-key read → **zero value** (0)।
- **`v, ok := m[k]`** — `ok`-এ presence জানায় → false।

---

## Expected Output

```
The value: 42
The value: 48
The value: 0
The value: 0 present? false
```

## মূল শিক্ষা / Key Takeaways

1. **Insert & update** — একই operation।
2. **`delete`** — key সরানো।
3. **Zero on missing** — read কখনো panic করে না।
4. **Comma-ok** — presence-check-এর idiomatic রূপ।

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

### Lines 5–17

```go
func main() {
	m := make(map[string]int)
	m["Answer"] = 42
	fmt.Println("The value:", m["Answer"])

	m["Answer"] = 48
	fmt.Println("The value:", m["Answer"])

	delete(m, "Answer")
	fmt.Println("The value:", m["Answer"])

	v, ok := m["Answer"]
	fmt.Println("The value:", v, "present?", ok)
}
```

- `m[k] = v` — insert/update.
- **`delete(m, k)`** — removes the key.
- Reading a missing key yields the **zero value** (0).
- **`v, ok := m[k]`** — `ok` tells presence → false.

---

## Expected Output

```
The value: 42
The value: 48
The value: 0
The value: 0 present? false
```

## Key Takeaways

1. **Insert and update** — the same operation.
2. **`delete`** — remove a key.
3. **Zero on missing** — reads never panic.
4. **Comma-ok** — the idiomatic presence check.