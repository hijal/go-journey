# free-shipping

Go-তে **boolean logic** (`||`, `&&`, `!`) দিয়ে business rules শেখার ছোট example — free shipping, bulk order, upsell-এর শর্ত evaluation।

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

### Lines 6–8

```go
cartTotal := 1499.0
itemCount := 2
isPremium := false
```

তিনটা input-মান:

- `cartTotal` = 1499.0 (float) — cart-এর মোট মূল্য
- `itemCount` = 2 (int) — কয়টা item
- `isPremium` = false (bool) — premium সদস্য কি না

### Line 10

```go
freeShipping := cartTotal >= 1500 || (isPremium && itemCount > 0)
```

**`||` (OR)** — দুটো শর্তের যেকোনোটাই true হলে true:

- Shorthand-1: `cartTotal >= 1500` — 1500+ হলে। এখানে `1499 >= 1500` = **false**
- Shorthand-2: `isPremium && itemCount > 0` — premium সদস্য **এবং** কমপক্ষে ১টা item। এখানে `false && true` = **false**
- overall: `false || false` = **false**

Result: `freeShipping = false`। **Short-circuit:** Go left-to-right evaluate করে। `cartTotal >= 1500` false হলেও বা true হলেও — `||`-এর বাম false হওয়ায় ডান side-টাও check হয় (কারণ result এখনো জানা যায়নি)।

### Line 11

```go
bulkOrder := itemCount >= 10
```

`itemCount >= 10` = `2 >= 10` = **false**। সাধারণ condition।

### Line 12

```go
upSellCandidate := cartTotal < 1500 && !isPremium
```

**`&&` (AND)** + **`!` (NOT)**:

- `cartTotal < 1500` = `1499 < 1500` = **true**
- `!isPremium` = `!false` = **true** (NOT-এর সাথে flip)
- `true && true` = **true**

Result: `upSellCandidate = true` — customer-কে free-shipping-এর জন্য upsell দেখানো উচিত।

### Lines 14–16

```go
fmt.Println("free shipping:", freeShipping)
fmt.Println("bulk order:", bulkOrder)
fmt.Println("show free-shipping upsell banner:", upSellCandidate)
```

Output:

- `free shipping: false`
- `bulk order: false`
- `show free-shipping upsell banner: true`

### Line 17

```go
}
```

Closing brace — `main` function শেষ হয়।

---

## Expected Output

```
free shipping: false
bulk order: false
show free-shipping upsell banner: true
```

## মূল শিক্ষা / Key Takeaways

1. **`||` OR** — যেকোনো একটা শর্ত true হলেই true; otherwise false (true || anything = true)।
2. **`&&` AND** — দুটোই true হলে true; false হলে false।
3. **`!` NOT** — boolean value-কে flip (true↔false)।
4. **Boolean expressions** — একটা boolean value-র মধ্যে complex business rule encapsulate।
5. **Evaluation** — Go-র always evaluates সঠিক result পেতে — আপনি rule order-এ freedom পাবেন।

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

### Lines 6–8

```go
cartTotal := 1499.0
itemCount := 2
isPremium := false
```

Three input values:

- `cartTotal` = 1499.0 (float) — the cart total
- `itemCount` = 2 (int) — how many items
- `isPremium` = false (bool) — premium membership?

### Line 10

```go
freeShipping := cartTotal >= 1500 || (isPremium && itemCount > 0)
```

**`||` (OR)** — true if either condition is true:

- Side 1: `cartTotal >= 1500` — 1500+. Here `1499 >= 1500` = **false**
- Side 2: `isPremium && itemCount > 0` — premium **and** at least 1 item. Here `false && true` = **false**
- overall: `false || false` = **false**

Result: `freeShipping = false`. **Short-circuit:** Go evaluates left-to-right. When the left side is false for `||`, the right side still needs checking (the result isn't known yet).

### Line 11

```go
bulkOrder := itemCount >= 10
```

`itemCount >= 10` = `2 >= 10` = **false**. A simple comparison.

### Line 12

```go
upSellCandidate := cartTotal < 1500 && !isPremium
```

**`&&` (AND)** + **`!` (NOT)**:

- `cartTotal < 1500` = `1499 < 1500` = **true**
- `!isPremium` = `!false` = **true** (NOT flips the value)
- `true && true` = **true**

Result: `upSellCandidate = true` — the customer should see a free-shipping upsell.

### Lines 14–16

```go
fmt.Println("free shipping:", freeShipping)
fmt.Println("bulk order:", bulkOrder)
fmt.Println("show free-shipping upsell banner:", upSellCandidate)
```

Output:

- `free shipping: false`
- `bulk order: false`
- `show free-shipping upsell banner: true`

### Line 17

```go
}
```

Closing brace — ends the `main` function.

---

## Expected Output

```
free shipping: false
bulk order: false
show free-shipping upsell banner: true
```

## Key Takeaways

1. **`||` OR** — true if any condition is true; false only if all are false.
2. **`&&` AND** — true only if both are true.
3. **`!` NOT** — flips a boolean value (true↔false).
4. **Boolean expressions** — encapsulate complex business rules in a single boolean value.
5. **Evaluation** — Go evaluates until the result is known; order of rules is flexible.