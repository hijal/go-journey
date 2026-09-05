# variadic-function

Go-তে **variadic function** (`...int`) শেখার ছোট example — function যেটা যেকোনো সংখ্যক argument নিতে পারে।

**📖 ভাষা নির্বাচন করুন / Choose language:**

[🇧🇩 বানলা](#bangla) • [🇬🇧 English](#english)

---

<a name="bangla"></a>

## 🇧🇩 বানলা সংস্করণ

### Lines 1–3

```go
package main

import "fmt"
```

- `package main` — একট�া executable program。
- `fmt` — output print করতে。

### Lines 5–11

```go
func sum(nums ...int) int {
    total := 0
    for _, n := range nums {
        total += n
    }
    return total
}
```

Variadic function:

- `nums ...int` — এই parameter-টি যেকোনো সংখ্যক `int` accept করতে পারে; function-এর ভিতরে এটা একট�া `[]int` slice-এ পরিণত হয়।
- `total := 0` — accumulator, শুরুতে zero value。
- `for _, n := range nums` — প্রতিট�া element-এ loop করে。
- `total += n` — সব numbers যোগ করে。

### Lines 13–16

```go
func main() {
    result := sum(1, 2, 3, 4)
    fmt.Println("total sum:", result)
}
```

- `sum(1, 2, 3, 4)` — চারট�া argument দিন; variadic function automatically `[]int{1, 2, 3, 4}` তৈরি করে。
- `result` = `10`
- `fmt.Println("total sum:", result)` → `total sum: 10`

---

## Expected Output

```
total sum: 10
```

## মূল শিক্ষা / Key Takeaways

1. **`...T` syntax** — variadic parameter; যেকোনো সংখ্যক arguments accept করে。
2. **Slice transformation** — variadic arguments function-এ `[]T` slice-এ পরিণত হয়。
3. **Flexible calls** — `sum()` (empty), `sum(1)`, `sum(1, 2, 3)` — সব কাজ করে。
4. **Range over nums** — `for _, n := range nums` দিয়ে slice-এর সব element iterate করা যায়。

---

---

<a name="english"></a>

## 🇬🇧 English Version

### Lines 1–3

```go
package main

import "fmt"
```

- `package main` — an executable program.
- `fmt` — for console output.

### Lines 5–11

```go
func sum(nums ...int) int {
    total := 0
    for _, n := range nums {
        total += n
    }
    return total
}
```

Variadic function:

- `nums ...int` — this parameter accepts any number of `int` arguments; inside the function it becomes a `[]int` slice.
- `total := 0` — accumulator starting at zero value。
- `for _, n := range nums` — loops over each element。
- `total += n` — adds all numbers together。

### Lines 13–16

```go
func main() {
    result := sum(1, 2, 3, 4)
    fmt.Println("total sum:", result)
}
```

- `sum(1, 2, 3, 4)` — four arguments; variadic function automatically creates `[]int{1, 2, 3, 4}`。
- `result` = `10`
- `fmt.Println("total sum:", result)` → `total sum: 10`

---

## Expected Output

```
total sum: 10
```

## Key Takeaways

1. **`...T` syntax** — variadic parameter; accepts any number of arguments。
2. **Slice transformation** — variadic arguments become a `[]T` slice inside the function。
3. **Flexible calls** — `sum()`, `sum(1)`, `sum(1, 2, 3)` — all work。
4. **Range over nums** — use `for _, n := range nums` to iterate over the slice。
