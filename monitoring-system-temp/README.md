# monitoring-system-temp

Go-তে **builtin `max`**, **`if/else if`**, **`switch` with initializer** আর `%.1f` formatting দিয়ে machine-temperature monitoring শেখার ছোট example।

**📖 ভাষা নির্বাচন করুন / Choose language:**

[🇧🇩 বাংলা](#bangla) • [🇬🇧 English](#english)

---

<a name="bangla"></a>

## 🇧🇩 বাংলা সংস্করণ

### Lines 1–3

```go
package main

import "fmt"
```

- `package main` — একটা executable program।
- `fmt` — output print করার জন্য।

### Lines 5–12

```go
func classify(celsius float64) string {
    if celsius < 38 {
        return "normal"
    } else if celsius < 42 {
        return "warning"
    }
    return "critical"
}
```

`classify` function — একটা temperature (`celsius`, float64) নেয় এবং status string return করে:

- `if celsius < 38` → `"normal"` (স্বাভাবিক)।
- `else if celsius < 42` → `"warning"` (সতর্কতা)।
- বাকি (≥ 42) → `"critical"` (জটিল/বিপজ্জনক)।

এটা **`if / else if`** chain — প্রথম true শর্তটাই চলে।

### Line 14

```go
func main() {
```

Program-এর entry point।

### Line 15

```go
sensorA, sensorB := 41.3, 39.6
```

এক লাইনে দুটো variable — **multiple assignment**। `sensorA = 41.3`, `sensorB = 39.6` (দুটো float64)।

### Line 16

```go
reading := max(sensorA, sensorB)
```

**builtin `max`** — Go-র builtin function, দুটো value-র বড়টা return করে। `max(41.3, 39.6) = 41.3`। এটা `reading`-এ রাখা হয়।

### Lines 18–25

```go
switch level := classify(reading); level {
case "normal":
    fmt.Println("machine-7: running normally")
case "warning":
    fmt.Printf("machine-7: notify maintenance (%.1f°C)\n", reading)
default:
    fmt.Printf("machine-7: emergency shutdown (%.1f°C)\n", reading)
}
```

**`switch` with initializer** — `level := classify(reading)` আগে execute হয় (classify-র result-টা `level`-এ), তারপর `; level` দিয়ে ওই value-টা switch করা হয়।

- `case "normal"` → running normally।
- `case "warning"` → `fmt.Printf("... (%.1f°C)\n", reading)` — `%.1f` দিয়ে দশমিকের ১ ঘরে temperature দেখায় (41.3°C)।
- `default` → emergency shutdown (critical এর জন্য)।

এখানে reading=41.3, classify দেয় "warning" ⇒ `machine-7: notify maintenance (41.3°C)`।

---

## Expected Output

```
machine-7: notify maintenance (41.3°C)
```

## মূল শিক্ষা / Key Takeaways

1. **Builtin `max`** — two values থেকে বড়টা বের করা।
2. **`if / else if`** — tiered classification।
3. **`switch` with initializer** — `switch x := expr; x { ... }` — expression evaluate + switch একসাথে।
4. **Multiple assignment** — `a, b := 1, 2` এক লাইনে।
5. **`%.1f`** — float-কে ১ দশমিকে format।

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

### Lines 5–12

```go
func classify(celsius float64) string {
    if celsius < 38 {
        return "normal"
    } else if celsius < 42 {
        return "warning"
    }
    return "critical"
}
```

`classify` function — takes a temperature (`celsius`, float64) and returns a status string:

- `if celsius < 38` → `"normal"`.
- `else if celsius < 42` → `"warning"`.
- Otherwise (≥ 42) → `"critical"`.

This is an **`if / else if`** chain — the first true condition runs.

### Line 14

```go
func main() {
```

Program entry point.

### Line 15

```go
sensorA, sensorB := 41.3, 39.6
```

Assigns two variables on one line — **multiple assignment**. `sensorA = 41.3`, `sensorB = 39.6` (both float64).

### Line 16

```go
reading := max(sensorA, sensorB)
```

The **builtin `max`** — a Go builtin returning the larger of two values. `max(41.3, 39.6) = 41.3`. Stored in `reading`.

### Lines 18–25

```go
switch level := classify(reading); level {
case "normal":
    fmt.Println("machine-7: running normally")
case "warning":
    fmt.Printf("machine-7: notify maintenance (%.1f°C)\n", reading)
default:
    fmt.Printf("machine-7: emergency shutdown (%.1f°C)\n", reading)
}
```

**`switch` with initializer** — `level := classify(reading)` runs first (storing the classify result in `level`), then `; level` switches on that value.

- `case "normal"` → running normally.
- `case "warning"` → `fmt.Printf("... (%.1f°C)\n", reading)` — shows the temperature with 1 decimal place (`%.1f`, 41.3°C).
- `default` → emergency shutdown (for critical).

Here reading=41.3, classify returns "warning" ⇒ `machine-7: notify maintenance (41.3°C)`.

---

## Expected Output

```
machine-7: notify maintenance (41.3°C)
```

## Key Takeaways

1. **Builtin `max`** — gets the larger of two values.
2. **`if / else if`** — tiered classification.
3. **`switch` with initializer** — `switch x := expr; x { ... }` — evaluate an expression and switch in one step.
4. **Multiple assignment** — `a, b := 1, 2` on one line.
5. **`%.1f`** — format a float to 1 decimal place.
