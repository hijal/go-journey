# basic-types

Go-র **basic data types** (`int`, `float64`, `string`, `bool`, `rune`, `byte`) শেখার ছোট example — প্রতিটা declared, তারপর print করা হয়। বিশেষ করে `rune`-টা কীভাবে numeric code point হিসেবে print হয়।

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

Console-এ output print করার জন্য `fmt` package import করা হয়।

### Line 5

```go
func main() {
```

Program-এর entry point।

### Line 6

```go
var accountCount int = 100
```

`var` + explicit type — `accountCount` হলো `int`, value `100`। **Explicit type declaration**, `:=` নেই। `var name type = value` সবচেয়ে স্পষ্ট form।

### Line 7

```go
var interestRate float64 = 11.11
```

`interestRate` হলো `float64` (double-precision floating point), value `11.11`। Fractional (দশমিক) সংখ্যার জন্য।

### Line 8

```go
var currency string = "BDT"
```

`currency` হলো `string`, value `"BDT"`। শুধু characters-এর sequence, double-quote-এর ভিতরে।

### Line 9

```go
var isActive bool = true
```

`isActive` হলো `bool` (boolean), value `true`। `true`/`false` — শুধু দুইটা।

### Line 10

```go
var currencySymbol rune = '৳'
```

`currencySymbol` হলো **`rune`** — type name `int32`-এর alias, single-quote-এর ভিতরে **একটা** Unicode character রাখে। এখানে `'৳'` (taka sign) — Bangla script-এর character। Rune-টা একটা numeric code point (Unicode value)।

### Line 11

```go
var firstByte = 'A'
```

`firstByte` — single-quote character literal `'A'`। Type ধরা হয়নি (inferred), তাই default `rune`। Ascii character `'A'`-র code point `65`।

### Line 12

```go
fmt.Println(accountCount, interestRate, currency, isActive, currencySymbol, firstByte)
```

`fmt.Println` — space-দিয়ে সব value print করে:

- `accountCount` → `100`
- `interestRate` → `11.11`
- `currency` → `BDT`
- `isActive` → `true`
- `currencySymbol` → `2547` (৳-র Unicode code point! `rune` print করলে **character না**, numeric value বের হয়)
- `firstByte` → `65` (`'A'`-র ASCII code point)

> **গুরুত্বপূর্ণ:** `rune` (alias `int32`) একটা integer টাইপ। `fmt.Println` দিয়ে print করলে character এর বদলে তার numeric code point দেখায়। Real character `৳` দেখতে হলে `%c` verb-এর সাথে `fmt.Printf` লাগবে।

### Line 13

```go
}
```

Closing brace — `main` function শেষ হয়।

---

## Expected Output

```
100 11.11 BDT true 2547 65
```

## মূল শিক্ষা / Key Takeaways

1. **Basic types** — `int`, `float64`, `string`, `bool`, `rune`, `byte` — initialize করার বিভিন্ন ভাবে।
2. **`var name type = value`** — explicit type-সহ declaration, `:=`-এর বদলে।
3. **`rune` একটা integer** — `int32`-এর alias; print করলে numeric code point দেখায় (৳ → 2547)।
4. **`byte`** — `uint8`-এর alias; `'A'` → `65`।
5. **Character literal** — single quotes `'x'`, string double quotes `"x"`।

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

Imports the `fmt` package for console output.

### Line 5

```go
func main() {
```

Program entry point.

### Line 6

```go
var accountCount int = 100
```

`var` + explicit type — `accountCount` is `int`, value `100`. An **explicit type declaration**, no `:=`. `var name type = value` is the most explicit form.

### Line 7

```go
var interestRate float64 = 11.11
```

`interestRate` is `float64` (double-precision floating point), value `11.11`. For fractional (decimal) numbers.

### Line 8

```go
var currency string = "BDT"
```

`currency` is `string`, value `"BDT"`. A sequence of characters, inside double quotes.

### Line 9

```go
var isActive bool = true
```

`isActive` is `bool` (boolean), value `true`. Only `true`/`false`.

### Line 10

```go
var currencySymbol rune = '৳'
```

`currencySymbol` is a **`rune`** — an alias for `int32`, holding **one** Unicode character in single quotes. Here `'৳'` (the taka sign) — a Bangla script character. A rune is a numeric code point (a Unicode value).

### Line 11

```go
var firstByte = 'A'
```

`firstByte` — a single-quote character literal `'A'`. No type was given (inferred), so it defaults to `rune`. The ASCII character `'A'` has code point `65`.

### Line 12

```go
fmt.Println(accountCount, interestRate, currency, isActive, currencySymbol, firstByte)
```

`fmt.Println` prints all values separated by spaces:

- `accountCount` → `100`
- `interestRate` → `11.11`
- `currency` → `BDT`
- `isActive` → `true`
- `currencySymbol` → `2547` (the Unicode code point of ৳! Printing a `rune` shows the numeric value, not the character)
- `firstByte` → `65` (the ASCII code point of `'A'`)

> **Important:** a `rune` (alias `int32`) is an integer type. Printing it with `fmt.Println` shows its numeric code point, not the character. To render the actual `৳`, you'd need `fmt.Printf` with the `%c` verb.

### Line 13

```go
}
```

Closing brace — ends the `main` function.

---

## Expected Output

```
100 11.11 BDT true 2547 65
```

## Key Takeaways

1. **Basic types** — `int`, `float64`, `string`, `bool`, `rune`, `byte` — several ways to initialize.
2. **`var name type = value`** — explicit declaration with a type, instead of `:=`.
3. **A `rune` is an integer** — alias for `int32`; printing shows the numeric code point (৳ → 2547).
4. **`byte`** — alias for `uint8`; `'A'` → `65`.
5. **Character literal** — single quotes `'x'`, string double quotes `"x"`.