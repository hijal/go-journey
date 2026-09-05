# devide-zero

Go-তে **division by zero** হ্যান্ডলিং করে error return করা শেখার ছোট example — custom error তৈরি করে safe divide function লেখা।

**📖 ভাষা নির্বাচন করুন / Choose language:**

[🇧🇩 বাংলা](#bangla) • [🇬🇧 English](#english)

---

<a name="bangla"></a>

## 🇧🇩 বাংলা সংস্করণ

### Lines 1–6

```go
package main

import (
    "errors"
    "fmt"
)
```

- `package main` — একট�া executable program。
- `errors` — `errors.New` দিয়ে custom error তৈরি করতে。
- `fmt` — output print করতে।

### Lines 8–13

```go
func divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, errors.New("division by zero")
    }
    return (a / b), nil
}
```

Custom `divide` function যা দুটো value return করে:

- `func divide(a, b float64) (float64, error)` — দুট�া input (`float64`), দুট�া output: `float64` (result) + `error` (nil হলে success)।
- `if b == 0` — divisor শূন্য হলে error return করে (`errors.New("division by zero")`)।
- `return (a / b), nil` — সঠিক case-এ result + nil error।

### Lines 15–21

```go
func main() {
    result, err := divide(10, 3)
    if err != nil {
        fmt.Println("divide error:", err)
    }
    fmt.Println(result)
}
```

- `result, err := divide(10, 3)` — normal division, error নেই, result `3.333...`。
- `if err != nil` — error থাকলে print করে, কিন্তু এখানে error নেই তাই skip।
- `fmt.Println(result)` → `3.3333333333333335`

---

## Expected Output

```
3.3333333333333335
```

## মূল শিক্ষা / Key Takeaways

1. **Multiple return values** — Go-তে একাধিক value return করা সাধারণ; এখানে `float64, error`।
2. **`errors.New()`** — simple string error তৈরি করতে stdlib থেকে。
3. **Zero division guard** — `b == 0` চেক করে safe divide করা।
4. **Error handling idiom** — `if err != nil` check করাই Go-তে standard pattern。

---

---

<a name="english"></a>

## 🇬🇧 English Version

### Lines 1–6

```go
package main

import (
    "errors"
    "fmt"
)
```

- `package main` — an executable program.
- `errors` — to create custom errors with `errors.New`.
- `fmt` — for console output.

### Lines 8–13

```go
func divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, errors.New("division by zero")
    }
    return (a / b), nil
}
```

A custom `divide` function that returns two values:

- `func divide(a, b float64) (float64, error)` — two inputs (`float64`), two outputs: `float64` (result) + `error` (`nil` means success).
- `if b == 0` — returns an error when divisor is zero (`errors.New("division by zero")`).
- `return (a / b), nil` — on success, returns the result and a `nil` error.

### Lines 15–21

```go
func main() {
    result, err := divide(10, 3)
    if err != nil {
        fmt.Println("divide error:", err)
    }
    fmt.Println(result)
}
```

- `result, err := divide(10, 3)` — normal division, no error, result is `3.333...`.
- `if err != nil` — would print the error, but here there is no error so it's skipped.
- `fmt.Println(result)` → `3.3333333333333335`

---

## Expected Output

```
3.3333333333333335
```

## Key Takeaways

1. **Multiple return values** — common in Go; here returning `float64, error`.
2. **`errors.New()`** — creates a simple string error from the standard library.
3. **Zero division guard** — checks `b == 0` for a safe divide.
4. **Error handling idiom** — `if err != nil` is the standard Go pattern.
