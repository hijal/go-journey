# group-variable

Go-তে **grouped `const` with `iota`** আর **grouped `var` declaration** শেখার ছোট example — একাধিক সম্পর্কিত constant/variable এক block-এ রাখা।

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

### Lines 5–11

```go
const (
	StatusPending = iota
	StatusPaid
	StatusShipped
	StatusDelivered
	StatusCancelled
)
```

Grouped **constant block** — order status-গুলোর একটা enum-এর মতো sequence:

- `StatusPending = iota` — প্রথমটায় `iota` assign করা হয় (0 থেকে শুরু)।
- `StatusPaid` — পরেরটা automatic-এ `iota` = 1 পায়।
- `StatusShipped` = 2, `StatusDelivered` = 3, `StatusCancelled` = 4।

`iota` হলো compiler-চালিত counter যা `const` block-এর প্রতিটা line-এ 0, 1, 2, ... বাড়ে। কোনো হাতের type না দিলেও এগুলো সব `int` constant।

### Line 13

```go
func main() {
```

Program-এর entry point।

### Lines 14–18

```go
var (
	orderID   = "ORD-1023"
	status    = StatusPending
	itemCount int
)
```

Grouped **`var` declaration** — তিনটা সম্পর্কিত variable এক block-এ:

- `orderID = "ORD-1023"` — type-inference: string।
- `status = StatusPending` — type-inference: এটা `StatusPending`-এর সাথে match, `int` (value 0)।
- `itemCount int` — শুধু type, value দেওয়া হয়নি — তাই **zero value** `0`।

### Line 20

```go
itemCount = 3
```

`itemCount`-এ `3` assign করা হয় (আগে declare হয়ে গেছে, এটা `=` assignment)।

### Lines 21–23

```go
fmt.Println("Order ID:", orderID)
fmt.Println("Status code:", status)
fmt.Println("Item count:", itemCount)
```

- `Order ID: ORD-1023`
- `Status code: 0` — status এখনও `StatusPending` → numeric 0।
- `Item count: 3`

### Lines 25–26

```go
status = StatusPaid
fmt.Println("Updated status code:", status)
```

`status`-কে **update**: `StatusPaid` assign করে (numeric 1)। Print: `Updated status code: 1`।

### Line 27

```go
}
```

Closing brace — `main` function শেষ হয়।

---

## Expected Output

```
Order ID: ORD-1023
Status code: 0
Item count: 3
Updated status code: 1
```

## মূল শিক্ষা / Key Takeaways

1. **`const` grouping** — `(...)` block-এ সম্পর্কিত constant-গুলো একসাথে; `iota` sequence-এর জন্য।
2. **`iota`** — block-ের line-সংখ্যা অনুযায়ী auto-increment counter।
3. **`var` grouping** — related variable-গুলোকে grouped declaration এর read-able.
4. **Zero value** — `itemCount int` (initial value ছাড়া) `0` দিয়ে শুরু।
5. **Assignment vs declaration** — `=` আগের variable-এ নতুন value দেয়; `:=` নতুন variable বানায়।

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

### Lines 5–11

```go
const (
	StatusPending = iota
	StatusPaid
	StatusShipped
	StatusDelivered
	StatusCancelled
)
```

Grouped **constant block** — an enum-like sequence of order statuses:

- `StatusPending = iota` — the first one gets `iota` (starting at 0).
- `StatusPaid` — the next automatically gets `iota` = 1.
- `StatusShipped` = 2, `StatusDelivered` = 3, `StatusCancelled` = 4.

`iota` is a compiler-driven counter that increments 0, 1, 2, ... for each line in a `const` block. Even without an explicit type, these are all `int` constants.

### Line 13

```go
func main() {
```

Program entry point.

### Lines 14–18

```go
var (
	orderID   = "ORD-1023"
	status    = StatusPending
	itemCount int
)
```

Grouped **`var` declaration** — three related variables in one block:

- `orderID = "ORD-1023"` — type-inferred: string.
- `status = StatusPending` — type-inferred: matches `StatusPending`, an `int` (value 0).
- `itemCount int` — type only, no value — so it starts at the **zero value** `0`.

### Line 20

```go
itemCount = 3
```

Assigns `3` to `itemCount` (already declared, so this is `=` assignment).

### Lines 21–23

```go
fmt.Println("Order ID:", orderID)
fmt.Println("Status code:", status)
fmt.Println("Item count:", itemCount)
```

- `Order ID: ORD-1023`
- `Status code: 0` — status is still `StatusPending` → numeric 0.
- `Item count: 3`

### Lines 25–26

```go
status = StatusPaid
fmt.Println("Updated status code:", status)
```

**Updates** `status`: assigns `StatusPaid` (numeric 1). Prints: `Updated status code: 1`.

### Line 27

```go
}
```

Closing brace — ends the `main` function.

---

## Expected Output

```
Order ID: ORD-1023
Status code: 0
Item count: 3
Updated status code: 1
```

## Key Takeaways

1. **`const` grouping** — a `(...)` block keeps related constants together; `iota` for sequences.
2. **`iota`** — an auto-increment counter based on the line within the block.
3. **`var` grouping** — grouped declarations keep related variables readable.
4. **Zero value** — `itemCount int` (no initial value) starts at `0`.
5. **Assignment vs declaration** — `=` gives an existing variable a new value; `:=` creates a new one.