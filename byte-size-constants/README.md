# byte-size-constants

Go-তে **iota** আর **bit shift (`<<`)** দিয়ে byte-size constants (KB, MB, GB, TB) তৈরি শেখার ছোট example — iota-র value-কে calculator হিসেবে ব্যবহার।

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

### Lines 5–12

```go
const (
	_  = iota             // 0 skipped
	KB = 1 << (10 * iota) // 1 << 10 = 1024
	MB = 1 << (10 * iota) // 1 << 20
	GB = 1 << (10 * iota) // 1 << 30
	TB = 1 << (10 * iota) // 1 << 40
	PB = 1 << (10 * iota) // 1 << 50
)
```

**iota + bit-shift দিয়ে byte units:** প্রতিটা line-এ `iota` auto-increase হয় (0, 1, 2, 3, 4, 5)। প্রতিটা expression `1 << (10 × iota)`:

- `_` → iota=0, `1 << 0` = 1 — **skipped** (blank identifier); প্রতি-1 মান (byte unit) লাগে না।
- `KB` → iota=1, `1 << 10` = 1024
- `MB` → iota=2, `1 << 20` = 1048576 (1024²)
- `GB` → iota=3, `1 << 30` = 1073741824 (1024³)
- `TB` → iota=4, `1 << 40` = 1099511627776 (1024⁴)
- `PB` → iota=5, `1 << 50` = 1125899906842624 (1024⁵)

> **কেন `<< (10 × iota)`:** binary-তে `1 << n` মানে `1 × 2ⁿ`। কিন্তু byte units 1024-base (2¹⁰), তাই `10 × iota` — 10-bit shifts, প্রতিটা unit-এ। (1KB = 2¹⁰ = 1024 byte.)

### Line 14

```go
func main() {
```

Program-এর entry point।

### Lines 15–16

```go
var maxLogSize = 12 * MB
var uploadLimit = 250 * MB
```

দুটো practical value:

- `maxLogSize = 12 * MB` = 12 × 1048576 = 12582912 byte
- `uploadLimit = 250 * MB` = 250 × 1048576 = 262144000 byte

### Lines 18–20

```go
fmt.Printf("max log size: %d bytes (%d MB)\n", maxLogSize, maxLogSize/MB)
fmt.Printf("upload limit: %d bytes (%d MB)\n", uploadLimit, uploadLimit/MB)
fmt.Printf("1 TB = %d bytes\n", TB)
```

Output:

- `max log size: 12582912 bytes (12 MB)`
- `upload limit: 262144000 bytes (250 MB)`
- `1 TB = 1099511627776 bytes`

### Line 21

```go
}
```

Closing brace — `main` function শেষ হয়।

---

## Expected Output

```
max log size: 12582912 bytes (12 MB)
upload limit: 262144000 bytes (250 MB)
1 TB = 1099511627776 bytes
```

## মূল শিক্ষা / Key Takeaways

1. **`_` = blank identifier** — iota-কে skip করার জন্য (block-এর first line-এ value drop)।
2. **`1 << (10 * iota)`** — binary bit-shift দিয়ে power-of-two byte constants।
3. **`iota` as calculator** — iota-র value arithmetic expression-এ ব্যবহার করা।
4. **Constant arithmetic** — `maxLogSize/MB` compile-time-এ constant, efficient।
5. **Readable config** — `250 * MB` সহজেই বোঝা যায় (262144000 না)।

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

### Lines 5–12

```go
const (
	_  = iota             // 0 skipped
	KB = 1 << (10 * iota) // 1 << 10 = 1024
	MB = 1 << (10 * iota) // 1 << 20
	GB = 1 << (10 * iota) // 1 << 30
	TB = 1 << (10 * iota) // 1 << 40
	PB = 1 << (10 * iota) // 1 << 50
)
```

**Byte units built with iota + bit shift:** each line, `iota` auto-increments (0, 1, 2, 3, 4, 5). Every expression is `1 << (10 × iota)`:

- `_` → iota=0, `1 << 0` = 1 — **skipped** (blank identifier); don't need the 1-byte unit.
- `KB` → iota=1, `1 << 10` = 1024
- `MB` → iota=2, `1 << 20` = 1048576 (1024²)
- `GB` → iota=3, `1 << 30` = 1073741824 (1024³)
- `TB` → iota=4, `1 << 40` = 1099511627776 (1024⁴)
- `PB` → iota=5, `1 << 50` = 1125899906842624 (1024⁵)

> **Why `<< (10 × iota)`:** `1 << n` in binary means `1 × 2ⁿ`. Byte units are 1024-based (2¹⁰), so `10 × iota` shifts by 10 bits per unit. (1KB = 2¹⁰ = 1024 bytes.)

### Line 14

```go
func main() {
```

Program entry point.

### Lines 15–16

```go
var maxLogSize = 12 * MB
var uploadLimit = 250 * MB
```

Two practical values:

- `maxLogSize = 12 * MB` = 12 × 1048576 = 12582912 bytes
- `uploadLimit = 250 * MB` = 250 × 1048576 = 262144000 bytes

### Lines 18–20

```go
fmt.Printf("max log size: %d bytes (%d MB)\n", maxLogSize, maxLogSize/MB)
fmt.Printf("upload limit: %d bytes (%d MB)\n", uploadLimit, uploadLimit/MB)
fmt.Printf("1 TB = %d bytes\n", TB)
```

Output:

- `max log size: 12582912 bytes (12 MB)`
- `upload limit: 262144000 bytes (250 MB)`
- `1 TB = 1099511627776 bytes`

### Line 21

```go
}
```

Closing brace — ends the `main` function.

---

## Expected Output

```
max log size: 12582912 bytes (12 MB)
upload limit: 262144000 bytes (250 MB)
1 TB = 1099511627776 bytes
```

## Key Takeaways

1. **`_` = blank identifier** — skips iota on the first line of a block.
2. **`1 << (10 * iota)`** — power-of-two byte constants via binary bit shift.
3. **`iota` as calculator** — iota's value fed into arithmetic expressions.
4. **Constant arithmetic** — `maxLogSize/MB` is computed at compile time, efficient.
5. **Readable config** — `250 * MB` is far easier to understand than `262144000`.