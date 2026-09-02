# insurance-eligibility-rule

Go-তে **struct (nested)**, **boolean logic** (`&&`, `!`) আর **compound eligibility condition** দিয়ে insurance (বীমা) eligibility নিয়ম শেখার ছোট example।

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

### Lines 5–8

```go
type Claim struct {
    Month  string
    Amount int
}
```

একটা struct `Claim` define করে — একটা বীমা দাবি (claim): `Month` (কোন মাসে দাবি) আর `Amount` (দাবির পরিমাণ, int)।

### Lines 10–15

```go
type Applicant struct {
    Name       string
    Age        int
    HasLicense bool
    Claims     []Claim
}
```

আরেকটা struct `Applicant` (আবেদনকারী):

- `Name` — নাম (string)।
- `Age` — বয়স (int)।
- `HasLicense` — license আছে কি না (bool)।
- `Claims` — `[]Claim` — claim-গুলোর slice। এটা **nested struct** — একটা struct-এর ভেতরে আরেক struct-এর slice।

### Line 17

```go
func main() {
```

Program-এর entry point।

### Lines 18–26

```go
applicant := Applicant{
    Name:       "John Doe",
    Age:        25,
    HasLicense: true,
    Claims: []Claim{{
        Month:  "March",
        Amount: 40000,
    }},
}
```

একটা `Applicant` instance। লক্ষ্য করো `Claims: []Claim{{...}}` — ভেতরের `{{...}}` দিয়ে একটা single-element Claim slice। John Doe-র March মাসে 40000-এর একটা দাবি আছে।

### Line 28

```go
hasBigClaim := len(applicant.Claims) > 0 && applicant.Claims[0].Amount >= 100000
```

**`hasBigClaim`** — complex boolean expression:

- `len(applicant.Claims) > 0` — claim আছে কিনা (empty slice-এ `[0]` access crash করায় আগে check দরকার)।
- `&&` — AND; দুই শর্তই true হতে হবে।
- `applicant.Claims[0].Amount >= 100000` — প্রথম claim-টা 100000 বা তার বেশি কিনা।

John Doe-র claim amount 40000, 100000-এর কম ⇒ `hasBigClaim = false`।

### Line 29

```go
isAdult := applicant.Age >= 18
```

`Age >= 18` → 25 ⇒ `isAdult = true`।

### Line 30

```go
isEligible := isAdult && applicant.HasLicense && !hasBigClaim
```

**`isEligible`** — তিনটা শর্ত AND:

- `isAdult` (সত্য) 
- `&& applicant.HasLicense` (সত্য)
- `&& !hasBigClaim` (`!` দিয়ে উল্টানো; `hasBigClaim=false` ⇒ `!false = true`)

সব সত্য ⇒ `isEligible = true`।

### Lines 32–35

```go
fmt.Printf("Applicant:  %s\n", applicant.Name)
fmt.Printf("Adult:      %v\n", isAdult)
fmt.Printf("Big claim:  %v\n", hasBigClaim)
fmt.Printf("Eligible:   %v\n", isEligible)
```

চারটা মান `fmt.Printf` দিয়ে print:

- `%s` — string (নাম)
- `%v` — value (bool-গুলো: true/false)

### Lines 37–41

```go
if !isEligible {
    fmt.Println("Decision: referred to a human underwriter")
    return
}
fmt.Println("Decision: auto-approved, quote 9000 BDT")
```

- `if !isEligible` — eligible না হলে manual underwriter-এ পাঠানো হয়।
- নাহলে auto-approved message। এখানে eligible=true, তাই `Decision: auto-approved, quote 9000 BDT`।

---

## Expected Output

```
Applicant:  John Doe
Adult:      true
Big claim:  false
Eligible:   true
Decision: auto-approved, quote 9000 BDT
```

## মূল শিক্ষা / Key Takeaways

1. **Nested struct** — struct-এর ভেতরে struct-এর slice (`Claims []Claim`)।
2. **Boolean operators** — `&&` (AND), `!` (NOT)।
3. **Compound condition** — একাধিক শর্ত মিলিয়ে eligibility।
4. **Guard against empty slice** — `[0]` access-এর আগে `len(...) > 0` check।
5. **`%s` / `%v`** — string ও value formatting।

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

### Lines 5–8

```go
type Claim struct {
    Month  string
    Amount int
}
```

Defines a struct `Claim` — an insurance claim: `Month` (in which month) and `Amount` (the claim amount, int).

### Lines 10–15

```go
type Applicant struct {
    Name       string
    Age        int
    HasLicense bool
    Claims     []Claim
}
```

Another struct `Applicant`:

- `Name` — name (string).
- `Age` — age (int).
- `HasLicense` — whether they have a license (bool).
- `Claims` — `[]Claim` — a slice of claims. This is a **nested struct** — a struct's slice inside another struct.

### Line 17

```go
func main() {
```

Program entry point.

### Lines 18–26

```go
applicant := Applicant{
    Name:       "John Doe",
    Age:        25,
    HasLicense: true,
    Claims: []Claim{{
        Month:  "March",
        Amount: 40000,
    }},
}
```

Creates an `Applicant` instance. Note `Claims: []Claim{{...}}` — the inner `{{...}}` creates a single-element Claim slice. John Doe has one complaint of 40000 in March.

### Line 28

```go
hasBigClaim := len(applicant.Claims) > 0 && applicant.Claims[0].Amount >= 100000
```

**`hasBigClaim`** — a complex boolean expression:

- `len(applicant.Claims) > 0` — whether there's a claim (accessing `[0]` on an empty slice would crash, so this check is needed first).
- `&&` — AND; both conditions must be true.
- `applicant.Claims[0].Amount >= 100000` — whether the first claim is 100000 or more.

John Doe's claim amount is 40000, less than 100000 ⇒ `hasBigClaim = false`.

### Line 29

```go
isAdult := applicant.Age >= 18
```

`Age >= 18` → 25 ⇒ `isAdult = true`.

### Line 30

```go
isEligible := isAdult && applicant.HasLicense && !hasBigClaim
```

**`isEligible`** — three conditions ANDed:

- `isAdult` (true)
- `&& applicant.HasLicense` (true)
- `&& !hasBigClaim` (`!` inverts; `hasBigClaim=false` ⇒ `!false = true`)

All true ⇒ `isEligible = true`.

### Lines 32–35

```go
fmt.Printf("Applicant:  %s\n", applicant.Name)
fmt.Printf("Adult:      %v\n", isAdult)
fmt.Printf("Big claim:  %v\n", hasBigClaim)
fmt.Printf("Eligible:   %v\n", isEligible)
```

Prints four values with `fmt.Printf`:

- `%s` — string (name)
- `%v` — value (the bools: true/false)

### Lines 37–41

```go
if !isEligible {
    fmt.Println("Decision: referred to a human underwriter")
    return
}
fmt.Println("Decision: auto-approved, quote 9000 BDT")
```

- `if !isEligible` — if not eligible, refer to a human underwriter.
- Otherwise, the auto-approved message. Here eligible=true, so it prints `Decision: auto-approved, quote 9000 BDT`.

---

## Expected Output

```
Applicant:  John Doe
Adult:      true
Big claim:  false
Eligible:   true
Decision: auto-approved, quote 9000 BDT
```

## Key Takeaways

1. **Nested struct** — a slice of structs inside a struct (`Claims []Claim`).
2. **Boolean operators** — `&&` (AND), `!` (NOT).
3. **Compound condition** — combining multiple conditions for eligibility.
4. **Guard against empty slice** — check `len(...) > 0` before accessing `[0]`.
5. **`%s` / `%v`** — string and value formatting.
