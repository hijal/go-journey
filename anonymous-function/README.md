# anonymous-function

Go-তে **anonymous function** (function literal) আর **IIFE** (Immediately Invoked Function Expression) প্যাটার্ন শেখার ছোট example — function declare করে সাথে সাথে invoke করা।

**📖 ভাষা নির্বাচন করুন / Choose language:**

[🇧🇩 বাংলা](#bangla) • [🇬🇧 English](#english)

---

<a name="bangla"></a>

## 🇧🇩 বাংলா সংস্করণ

### Lines 1–3

```go
package main

import "fmt"
```

- `package main` — একটା executable program।
- `fmt` — output print করার জন্য।

### Line 5

```go
func main() {
```

Program-এর entry point。

### Lines 6–9

```go
// IIFE
result := func(x int) int {
    return x * x
}(6)
```

একট�া anonymous function তৈরি করে সাথে সাথে invoke করা হচ্ছে (IIFE pattern):

- `func(x int) int` — একটি function literal যেটা `int` নেয় আর `int` রিটার্ন করে।
- `(6)` — সাথে সাথে `6` argument দিয়ে invoke করা হলো。
- `result :=` — result value `36` (`6*6`)।

### Lines 10–13

```go
double := func(x int) int {
    return x * x
}
```

একই anonymous function-টাকে `double` variable-এ store করা হচ্ছ — function value as a variable。

### Lines 14–15

```go
fmt.Println(result)
fmt.Println(double(10))
```

- `result` → `36`
- `double(10)` → `100` (`10*10`)

### Line 16

```go
}
```

Closing brace — `main` function শেষ হয়।

---

## Expected Output

```
36
100
```

## মূল শিক্ষা / Key Takeaways

1. **Anonymous function** — name ছাড়া function literal; inline তৈরি করা যায়।
2. **IIFE** — function literal-কে সাথে সাথে `()` দিয়ে invoke করা যায়।
3. **Function value as variable** — anonymous function-কে variable-এ store করে পরে call করা যায়。
4. **Type annotation** — `func(x int) int` — parameter আর return type উভয়ই explicitly define করা হয়।

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
func main() {
```

Program entry point.

### Lines 6–9

```go
// IIFE
result := func(x int) int {
    return x * x
}(6)
```

An anonymous function is created and immediately invoked (IIFE pattern):

- `func(x int) int` — a function literal that takes an `int` and returns an `int`.
- `(6)` — immediately invoked with argument `6`.
- `result :=` — captures the result value `36` (`6*6`).

### Lines 10–13

```go
double := func(x int) int {
    return x * x
}
```

The same anonymous function is stored in the `double` variable — function value as a variable.

### Lines 14–15

```go
fmt.Println(result)
fmt.Println(double(10))
```

- `result` → `36`
- `double(10)` → `100` (`10*10`)

### Line 16

```go
}
```

Closing brace — ends the `main` function.

---

## Expected Output

```
36
100
```

## Key Takeaways

1. **Anonymous function** — a function literal without a name; can be defined inline.
2. **IIFE** — a function literal can be immediately invoked with `()`.
3. **Function value as variable** — an anonymous function can be stored in a variable and called later.
4. **Type annotation** — `func(x int) int` explicitly defines both parameter and return types.
