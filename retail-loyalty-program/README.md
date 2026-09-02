# retail-loyalty-program

Go-তে **`switch`-এ `fallthrough`**, **early return (guard)**, আর `fmt.Printf` দিয়ে loyalty tier-benefit (সুবিধা) বিতরণ শেখার ছোট example।

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

- `package main` — একটা executable program।
- `fmt` — output print করার জন্য।

### Lines 5–23

```go
func grantBenefits(tier string) {
    if tier != "GOLD" && tier != "SILVER" && tier != "BRONZE" {
        fmt.Println("unknown tier: no benefits")
        return
    }

    fmt.Printf("%s member gets:\n", tier)

    switch tier {
    case "GOLD":
        fmt.Println("- Priority support queue")
        fallthrough
    case "SILVER":
        fmt.Println("- Free monthly report")
        fallthrough
    case "BRONZE":
        fmt.Println("- 5% cashback")
    }
}
```

`grantBenefits` function — একটা tier (string) নেয় এবং সদস্যের সুবিধা print করে।

- **Guard check:** `if tier != "GOLD" && tier != "SILVER" && tier != "BRONZE"` — যদি tier এই তিনটির কোনোটা না হয়, তাহলে প্রিন্ট করে `unknown tier: no benefits` আর `return` দিয়ে function শেষ (invalid tier-এর জন্য কিছুই দেয় না)।
- `fmt.Printf("%s member gets:\n", tier)` — tier মানটা print করে।
- **`switch tier`** — tier-এর মান অনুযায়ী case চলে:
  - `case "GOLD"` → `- Priority support queue` প্রিন্ট করে, তারপর **`fallthrough`** → পরের case-ও চলে।
  - `case "SILVER"` → `- Free monthly report`, `fallthrough` → পরের case।
  - `case "BRONZE"` → `- 5% cashback` (fallthrough নাই, এখানে থামে)।

**`fallthrough`-এর প্রভাব:** GOLD পেলে ওইটা + SILVER + BRONZE-র সুবিধা (তিনটাই) পায়; SILVER পেলে SILVER + BRONZE; BRONZE পেলে শুধু BRONZE। অর্থাৎ উঁচু tier-এর সদস্যরা নিচের tier-গুলোর সব সুবিধা পায় (cascade benefit)।

### Line 25

```go
func main() {
```

Program-এর entry point।

### Lines 26–32

```go
grantBenefits("GOLD")
fmt.Println()
grantBenefits("SILVER")
fmt.Println()
grantBenefits("BRONZE")
fmt.Println()
grantBenefits("FREE")
```

চারটা tier-এর জন্য call করা হয়, প্রতিটির মাঝে `fmt.Println()` দিয়ে একটা ফাঁকা line:

1. `grantBenefits("GOLD")` → তিনটাই সুবিধা।
2. `grantBenefits("SILVER")` → SILVER + BRONZE।
3. `grantBenefits("BRONZE")` → শুধু BRONZE।
4. `grantBenefits("FREE")` → invalid tier ⇒ `unknown tier: no benefits`।

---

## Expected Output

```
GOLD member gets:
- Priority support queue
- Free monthly report
- 5% cashback

SILVER member gets:
- Free monthly report
- 5% cashback

BRONZE member gets:
- 5% cashback

unknown tier: no benefits
```

## মূল শিক্ষা / Key Takeaways

1. **Guard / early return** — invalid input-এ সাথে-সাথে `return`।
2. **`switch tier`** — একটি value-কে একাধিক case-এর সাথে compare।
3. **`fallthrough`** — পরের case-ও execute করায় — cascade benefit logic-এর জন্য।
4. **`fmt.Printf` / `%s`** — string placeholder-সহ format output।

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

### Lines 5–23

```go
func grantBenefits(tier string) {
    if tier != "GOLD" && tier != "SILVER" && tier != "BRONZE" {
        fmt.Println("unknown tier: no benefits")
        return
    }

    fmt.Printf("%s member gets:\n", tier)

    switch tier {
    case "GOLD":
        fmt.Println("- Priority support queue")
        fallthrough
    case "SILVER":
        fmt.Println("- Free monthly report")
        fallthrough
    case "BRONZE":
        fmt.Println("- 5% cashback")
    }
}
```

`grantBenefits` function — takes a tier (string) and prints the member's benefits.

- **Guard check:** `if tier != "GOLD" && tier != "SILVER" && tier != "BRONZE"` — if the tier isn't any of these three, print `unknown tier: no benefits` and `return` to end the function (giving nothing for invalid tiers).
- `fmt.Printf("%s member gets:\n", tier)` — prints the tier value.
- **`switch tier`** — runs the matching case:
  - `case "GOLD"` → prints `- Priority support queue`, then **`fallthrough`** → the next case also runs.
  - `case "SILVER"` → prints `- Free monthly report`, `fallthrough` → next case.
  - `case "BRONZE"` → prints `- 5% cashback` (no fallthrough, stops here).

**Effect of `fallthrough`:** GOLD members get GOLD + SILVER + BRONZE benefits (all three); SILVER gets SILVER + BRONZE; BRONZE gets only BRONZE. So higher-tier members inherit all lower-tier benefits (cascade benefit).

### Line 25

```go
func main() {
```

Program entry point.

### Lines 26–32

```go
grantBenefits("GOLD")
fmt.Println()
grantBenefits("SILVER")
fmt.Println()
grantBenefits("BRONZE")
fmt.Println()
grantBenefits("FREE")
```

Called for four tiers, with a blank line via `fmt.Println()` between each:

1. `grantBenefits("GOLD")` → all three benefits.
2. `grantBenefits("SILVER")` → SILVER + BRONZE.
3. `grantBenefits("BRONZE")` → only BRONZE.
4. `grantBenefits("FREE")` → invalid tier ⇒ `unknown tier: no benefits`.

---

## Expected Output

```
GOLD member gets:
- Priority support queue
- Free monthly report
- 5% cashback

SILVER member gets:
- Free monthly report
- 5% cashback

BRONZE member gets:
- 5% cashback

unknown tier: no benefits
```

## Key Takeaways

1. **Guard / early return** — `return` immediately on invalid input.
2. **`switch tier`** — comparing one value against multiple cases.
3. **`fallthrough`** — also executing the next case — for cascade benefit logic.
4. **`fmt.Printf` / `%s`** — formatted output with a string placeholder.
