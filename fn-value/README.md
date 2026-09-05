# fn-value

Go-তে **function as a value** (function type) শেখার ছোট example — custom function type তৈরি করে function reference variable-এ store করা।

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

- `package main` — একট�া executable program。
- `fmt` — output print করতে。

### Line 5

```go
type Operator func(int, int) int
```

একট�া custom function type তৈরি:

- `type Operator` — `Operator` নামে একট�া নতুন type।
- `func(int, int) int` — এই type-টাকে সংজ্ঞা দেয়: দুট�া `int` parameter, একট�া `int` return value।

এখন `Operator` হলো একট�া function类型 যার signature `(int, int) int`。

### Lines 7–9

```go
func multiply(a, b int) int {
    return a * b
}
```

একট�া normal named function `multiply` — দুট�া int নিয়ে product দেয়।

### Line 12

```go
var op Operator = multiply
```

`multiply` function-ট�ো `op` variable-এ assign করা হচ্ছ — function value as a variable。

### Line 13

```go
fmt.Println("function value:", op(3, 4))
```

`op(3, 4)` মানে `multiply(3, 4)` = `12`।

---

## Expected Output

```
function value: 12
```

## মূল শিক্ষা / Key Takeaways

1. **Function type** — `type Operator func(int, int) int` দিয়ে custom function signature type তৈরি করা যায়。
2. **Function as value** — function-কে variable-এ store করে পরে call করা যায়।
3. **Assignability** — `multiply`-র signature `func(int, int) int`, তাই `Operator`-এ assign করা যায়।
4. **Function variables** — function pointers এর বদলে Go-তে সরাসরি function value variable-এ রাখা যায়。

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

### Line 5

```go
type Operator func(int, int) int
```

Defines a custom function type:

- `type Operator` — a new type named `Operator`.
- `func(int, int) int` — this type represents any function that takes two `int`s and returns an `int`.

Now `Operator` is a function type with the signature `(int, int) int`。

### Lines 7–9

```go
func multiply(a, b int) int {
    return a * b
}
```

A normal named function `multiply` that takes two ints and returns their product.

### Line 12

```go
var op Operator = multiply
```

The `multiply` function is assigned to the `op` variable — a function value stored in a variable。

### Line 13

```go
fmt.Println("function value:", op(3, 4))
```

`op(3, 4)` means `multiply(3, 4)` = `12`。

---

## Expected Output

```
function value: 12
```

## Key Takeaways

1. **Function type** — `type Operator func(int, int) int` lets you define a custom function signature type.
2. **Function as value** — functions can be stored in variables and called through them.
3. **Assignability** — `multiply`'s signature matches `func(int, int) int`, so it can be assigned to `Operator`.
4. **Function variables** — Go stores function values directly in variables, not pointers.
