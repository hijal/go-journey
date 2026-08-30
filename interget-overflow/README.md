# interget-overflow

Go-তে **integer overflow** বোঝার ছোট example — fixed-size `int32`-এ value-টা max-এর বাইরে গেলে কী হয়, আর `int64` দিয়ে কীভাবে বাঁচা যায়।

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

### Lines 3–6

```go
import (
	"fmt"
	"math"
)
```

- `fmt` — console-এ output print করার জন্য।
- `math` — platform-নির্ভর int max constants-এর জন্য (`math.MaxInt32`, `math.MaxInt64`), যা `math` package-এ signed int type-এর limit হিসেবে define।

### Line 8

```go
func main() {
```

Program-এর entry point।

### Lines 9–10

```go
fmt.Println("int32 max value:", math.MaxInt32)
fmt.Println("int64 max value:", math.MaxInt64)
```

- `math.MaxInt32` — 32-bit signed int-এর সর্বোচ্চ মান: `2147483647`।
- `math.MaxInt64` — 64-bit signed int-এর সর্বোচ্চ মান: `9223372036854775807`।

### Lines 12–13

```go
var total32 int32 = math.MaxInt32 - 2
deltas := []int32{1, 1, 1, 1}
```

`total32` — `int32` type, `MaxInt32 - 2` দিয়ে শুরু (2147483645)। `deltas` — ৪টা `int32` value-র slice। প্রতিটা add করলে `total32` দুইবার max-এর উপর চলে যাবে।

### Lines 15–17

```go
for _, d := range deltas {
	total32 += d
}
```

প্রতিটা `d` যোগ করা হয়:

- Start: 2147483645
- `+1` → 2147483646
- `+1` → 2147483647 (max!)
- `+1` → **overflow**: 2147483648 fit হয় না `int32`-তে → মান ঘুরে গিয়ে **-2147483648** হয়
- `+1` → -2147483647

Go-তে runtime-এ integer overflow **নীরবে** ঘটে (panic/error নেই) — signed-এ result mathematically ভুল হয়।

### Line 18

```go
fmt.Println("total32 after overflow:", total32)
```

Output: `total32 after overflow: -2147483647` — ভুল, কারণ আসলে `2147483647 + 2`-এর কথা। এটাই overflow-এর নীরব বিপদ।

### Lines 20–23

```go
var total64 int64 = math.MaxInt32 - 2
for _, d := range deltas {
	total64 += int64(d)
}
```

একই হিসাব, কিন্তু এবার `total64` **`int64`**:

- Start: 2147483645
- ৪টা `+1` → 2147483649

`int64`-এর max অনেক বড় (9223372036854775807), তাই এই range-এ কোনো overflow-ই নেই। লক্ষ্য করো `int64(d)` — `d`-কে explicit `int64`-এ convert করা হয়েছে, নাহলে `total64 += d` compile করত না (mismatched types)।

### Line 24

```go
fmt.Println("total64 (safe)      :", total64)
```

Output: `total64 (safe)      : 2147483649` — সঠিক উত্তর।

---

## Expected Output

```
int32 max value: 2147483647
int64 max value: 9223372036854775807
total32 after overflow: -2147483647
total64 (safe)      : 2147483649
```

## মূল শিক্ষা / Key Takeaways

1. **Fixed-size integers** — `int32`/`int64`-এর নির্দিষ্ট range আছে; `int` machine-dependent।
2. **Silent overflow** — Go runtime-এ overflow detect/panic হয় না — quiet ভুল result।
3. **Wrap-around** — max-এর বেশি গেলে value negative-এ ঘুরে যায় (two's complement)।
4. **চওড়া type নিরাপত্তা** — accumulator-এ কম-বেশি margin-এর জন্য `int64` ব্যবহার করা।
5. **Type conversion** — `int64(d)` দিয়ে slice element-কে target type-এ match করানো।

---

---

<a name="english"></a>

## 🇬🇧 English Version

### Line 1

```go
package main
```

Declares an executable program (`main` package), runnable via `go run`.

### Lines 3–6

```go
import (
	"fmt"
	"math"
)
```

- `fmt` — for printing output to the console.
- `math` — for the platform-agnostic max constants (`math.MaxInt32`, `math.MaxInt64`), which the `math` package defines as limits for signed int types.

### Line 8

```go
func main() {
```

Program entry point.

### Lines 9–10

```go
fmt.Println("int32 max value:", math.MaxInt32)
fmt.Println("int64 max value:", math.MaxInt64)
```

- `math.MaxInt32` — the highest 32-bit signed int: `2147483647`.
- `math.MaxInt64` — the highest 64-bit signed int: `9223372036854775807`.

### Lines 12–13

```go
var total32 int32 = math.MaxInt32 - 2
deltas := []int32{1, 1, 1, 1}
```

`total32` — type `int32`, starting at `MaxInt32 - 2` (2147483645). `deltas` — a slice of four `int32` values. Adding each one pushes `total32` past max twice.

### Lines 15–17

```go
for _, d := range deltas {
	total32 += d
}
```

Each `d` gets added:

- Start: 2147483645
- `+1` → 2147483646
- `+1` → 2147483647 (max!)
- `+1` → **overflow**: 2147483648 doesn't fit `int32` → the value wraps around to **-2147483648**
- `+1` → -2147483647

In Go, integer overflow at runtime happens **silently** (no panic, no error) — for signed types the result is simply mathematically wrong.

### Line 18

```go
fmt.Println("total32 after overflow:", total32)
```

Output: `total32 after overflow: -2147483647` — wrong, since it should be `2147483647 + 2`. That's the silent danger of overflow.

### Lines 20–23

```go
var total64 int64 = math.MaxInt32 - 2
for _, d := range deltas {
	total64 += int64(d)
}
```

Same math, but now `total64` is **`int64`**:

- Start: 2147483645
- Four `+1`s → 2147483649

`int64`'s max is far larger (9223372036854775807), so there's no overflow in this range. Note `int64(d)` — the element is explicitly converted to `int64`; without it, `total64 += d` wouldn't compile (mismatched types).

### Line 24

```go
fmt.Println("total64 (safe)      :", total64)
```

Output: `total64 (safe)      : 2147483649` — the correct answer.

---

## Expected Output

```
int32 max value: 2147483647
int64 max value: 9223372036854775807
total32 after overflow: -2147483647
total64 (safe)      : 2147483649
```

## Key Takeaways

1. **Fixed-size integers** — `int32`/`int64` have fixed ranges; `int` is machine-dependent.
2. **Silent overflow** — Go's runtime doesn't detect or panic on overflow — a quiet wrong result.
3. **Wrap-around** — crossing max wraps the value to negative (two's complement).
4. **Wider types are safer** — use `int64` for accumulators needing margin.
5. **Type conversion** — `int64(d)` matches slice elements to the target type.
