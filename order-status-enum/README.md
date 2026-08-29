# order-status-enum

Go-তে **typed constant (enum-like)**, `iota`, আর **custom `String()` method** বুঝতে সহায়ক ছোট example।

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
type OrderStatus int
```

একটা নতুন **named type** `OrderStatus` define করা হয়, যার underlying type হলো `int`। এটা enum-এর মতো — order-এর status-কে একটা specific type হিসেবে represent করতে, যেন ভুল int value-তে গুলি না লাগে।

### Lines 7–13

```go
const (
    OrderPending OrderStatus = iota
    OrderPaid
    OrderShipped
    OrderDelivered
    OrderCanceled
)
```

একটা **constant block** যেখানে `iota` ব্যবহার করা হয়েছে। `iota` হলো একটা auto-incrementing counter যা প্রতি নতুন line-এ `0, 1, 2, ...` বাড়ে:

- `OrderPending` = 0
- `OrderPaid` = 1
- `OrderShipped` = 2
- `OrderDelivered` = 3
- `OrderCanceled` = 4

প্রথমটায় `OrderStatus = iota` লিখে type নির্ধারণ করা হয়; বাকিগুলো automatically নিচের line-এ `OrderStatus` type-ই পায় এবং `iota`-র successive value ধরে।

### Lines 15–28

```go
func (s OrderStatus) String() string {
    names := [...]string{
        "Pending",
        "Paid",
        "Shipped",
        "Delivered",
        "Canceled",
    }

    if int(s) < 0 || int(s) >= len(names) {
        return "Unknown"
    }
    return names[s]
}
```

এটা `OrderStatus`-এর একটা **method** `String()`, যা status number-কে readable নামে convert করে।

- `names := [...]string{...}` — একটা **array** (size `...` মানে compiler নিজে count করে নেয়) যেখানে প্রতিটা status-এর display নাম আছে, index-order-এ। Index 0 = "Pending", 1 = "Paid" ইত্যাদি।
- `if int(s) < 0 || int(s) >= len(names)` — range-check: যদি numeric value negative বা array-র বাইরে হয়, তাহলে `"Unknown"` return করে (অস্বাভাবিক value-র নিরাপদ handling)।
- নাহলে `return names[s]` — `s`-কে index হিসেবে use করে সঠিক নামটা return করে।

যখনই কোনো `OrderStatus` value-কে `%v`/`%s` format বা `fmt.Println`-এ print করা হয়, Go automatic-এ এই `String()` method call করে readable name দেখায়।

### Line 30

```go
func main() {
```

Program-এর entry point।

### Line 31

```go
current := OrderPaid
```

`current` নামে একটা variable, যার value `OrderPaid` (numeric value 1)। Type inference-এর ফলে type হয় `OrderStatus`।

### Line 33

```go
fmt.Println("current status:", current)
```

`current` print করে। যেহেতু `OrderStatus`-এর `String()` method আছে, এখানে numeric 1 না দেখিয়ে `Paid` দেখায়। Output: `current status: Paid`। (`fmt.Println` value-কে `%v` দিয়ে print করে এবং `String()` call করে।)

### Line 34

```go
fmt.Println("numeric value", int(current))
```

`int(current)` দিয়ে `current`-কে **type conversion** করে plain int-এ নিয়ে `1` print করে। Output: `numeric value 1`।

### Lines 36–38

```go
if current == OrderPaid {
    fmt.Println("your order is ready to go!")
}
```

যদি `current`-টি `OrderPaid`-এর সমান হয়, তাহলে মেসেজটা print করে। যেহেতু `current` = `OrderPaid`, তাই `your order is ready to go!` print হয়।

### Line 39

```go
}
```

Closing brace — `main` function শেষ হয়।

---

## Expected Output

```
current status: Paid
numeric value 1
your order is ready to go!
```

## মূল শিক্ষা / Key Takeaways

1. **Typed constants** — `type OrderStatus int` দিয়ে enum-এর মতো fixed set of values।
2. **`iota`** — auto-incrementing counter, consecutive constants বানাতে ব্যবহৃত।
3. **Custom `String()` method** — numeric value-কে readable নামে print করতে।
4. **Type conversion** — `int(current)` দিয়ে typed value-কে plain int-এ আনা।
5. **Type safety** — `OrderStatus`-এর সাথে plain int সরাসরি compare করা যায় না — enum-এর মতো behavior।

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
type OrderStatus int
```

Defines a new **named type** `OrderStatus` whose underlying type is `int`. This is similar to an enum — representing order statuses as a specific type so they can't be confused with arbitrary ints.

### Lines 7–13

```go
const (
    OrderPending OrderStatus = iota
    OrderPaid
    OrderShipped
    OrderDelivered
    OrderCanceled
)
```

A **constant block** using `iota`, Go's auto-incrementing counter that increases by `0, 1, 2, ...` on each new line:

- `OrderPending` = 0
- `OrderPaid` = 1
- `OrderShipped` = 2
- `OrderDelivered` = 3
- `OrderCanceled` = 4

The first one sets `OrderStatus = iota` to fix the type; the rest automatically inherit the `OrderStatus` type and take successive `iota` values.

### Lines 15–28

```go
func (s OrderStatus) String() string {
    names := [...]string{
        "Pending",
        "Paid",
        "Shipped",
        "Delivered",
        "Canceled",
    }

    if int(s) < 0 || int(s) >= len(names) {
        return "Unknown"
    }
    return names[s]
}
```

This is a **method** `String()` on `OrderStatus` that converts a status number to a readable name.

- `names := [...]string{...}` — an **array** (size `...` lets the compiler count) holding the display name for each status by index. Index 0 = "Pending", 1 = "Paid", etc.
- `if int(s) < 0 || int(s) >= len(names)` — a range check: if the numeric value is negative or outside the array, return `"Unknown"` (safe handling of unexpected values).
- Otherwise `return names[s]` — uses `s` as the index to return the correct name.

Whenever an `OrderStatus` value is printed with `%v`, `%s`, or via `fmt.Println`, Go automatically calls this `String()` method to show the readable name.

### Line 30

```go
func main() {
```

Program entry point.

### Line 31

```go
current := OrderPaid
```

A variable `current` with the value `OrderPaid` (numeric value 1). Type inference makes its type `OrderStatus`.

### Line 33

```go
fmt.Println("current status:", current)
```

Prints `current`. Because `OrderStatus` has a `String()` method, it shows `Paid` instead of the number 1. Output: `current status: Paid`. (`fmt.Println` prints the value with `%v`, which calls `String()`.)

### Line 34

```go
fmt.Println("numeric value", int(current))
```

Converts `current` to a plain int with `int(current)` and prints `1`. Output: `numeric value 1`.

### Lines 36–38

```go
if current == OrderPaid {
    fmt.Println("your order is ready to go!")
}
```

If `current` equals `OrderPaid`, it prints the message. Since `current` is `OrderPaid`, it prints `your order is ready to go!`.

### Line 39

```go
}
```

Closing brace — ends the `main` function.

---

## Expected Output

```
current status: Paid
numeric value 1
your order is ready to go!
```

## Key Takeaways

1. **Typed constants** — `type OrderStatus int` creates an enum-like fixed set of values.
2. **`iota`** — auto-incrementing counter for consecutive constants.
3. **Custom `String()` method** — prints readable names instead of numbers.
4. **Type conversion** — `int(current)` brings a typed value back to plain int.
5. **Type safety** — `OrderStatus` can't be directly compared with a plain int — enum-like behavior.
