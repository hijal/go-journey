# named-naked-return

Go-তে **named return values** এবং **naked return** (bare return) প্যাটার্ন শেখার ছোট example — function-এর return parameters আগে থেকে name করে তাদের সরাসরি assign করা।

**📖 ভাষা নির্বাচন করুন / Choose language:**

[🇧🇩 বাংলা](#bangla) • [🇬🇧 English](#english)

---

<a name="bangla"></a>

## 🇧🇩 বানলা সংস্করণ

### Lines 1–5

```go
package main

import (
    "fmt"
)
```

- `package main` — একট�া executable program。
- `fmt` — output print করতে।

### Lines 7–11

```go
func rectangleInfo(w, h float64) (area, perimeter float64) {
    area = w * h
    perimeter = 2 * (w + h)
    return
}
```

Named return values + naked return:

- `(area, perimeter float64)` — দুটো return parameter-এর name ও type আগে থেকে declare করা হয়েছে।
- `area = w * h` — `area` variable-এ সরাসরি assign করা হচ্ছে (return statement-এ name লাগে না)。
- `perimeter = 2 * (w + h)` — একইভাবে `perimeter`-এ assign。
- `return` — **naked return** — নাম ছাড়া return; named parameters-গুলো auto return হয়।

### Lines 13–18

```go
func main() {
    area, perimeter := rectangleInfo(10, 15)

    fmt.Println("area:", area)
    fmt.Println("perimeter:", perimeter)
}
```

- `rectangleInfo(10, 15)` → `area = 150`, `perimeter = 50`
- `fmt.Println("area:", area)` → `area: 150`
- `fmt.Println("perimeter:", perimeter)` → `perimeter: 50`

---

## Expected Output

```
area: 150
perimeter: 50
```

## মূল শিক্ষা / Key Takeaways

1. **Named return values** — `(area, perimeter float64)` parameter-এর name upfront দেওয়া; function-এর ভিতরে সরাসরি use করা যায়।
2. **Naked return** — `return` (no values) — named parameters auto-return হয়।
3. **Readability** — name করা return values function-এর intent clearer করে।
4. **Assignment style** — `area = w * h` means `area` is set before returning。

---

---

<a name="english"></a>

## 🇬🇧 English Version

### Lines 1–5

```go
package main

import (
    "fmt"
)
```

- `package main` — an executable program.
- `fmt` — for console output.

### Lines 7–11

```go
func rectangleInfo(w, h float64) (area, perimeter float64) {
    area = w * h
    perimeter = 2 * (w + h)
    return
}
```

Named return values + naked return:

- `(area, perimeter float64)` — two return parameters are named and typed upfront.
- `area = w * h` — directly assigns to the named `area` parameter.
- `perimeter = 2 * (w + h)` — same for `perimeter`.
- `return` — **naked return** — returns with no values; named parameters are auto-returned.

### Lines 13–18

```go
func main() {
    area, perimeter := rectangleInfo(10, 15)

    fmt.Println("area:", area)
    fmt.Println("perimeter:", perimeter)
}
```

- `rectangleInfo(10, 15)` → `area = 150`, `perimeter = 50`
- `fmt.Println("area:", area)` → `area: 150`
- `fmt.Println("perimeter:", perimeter)` → `perimeter: 50`

---

## Expected Output

```
area: 150
perimeter: 50
```

## Key Takeaways

1. **Named return values** — `(area, perimeter float64)` names parameters upfront; they can be used directly inside the function.
2. **Naked return** — `return` (no values) — named parameters are automatically returned.
3. **Readability** — named return values make the function's intent clearer.
4. **Assignment style** — `area = w * h` sets `area` before the return。
