# invoice-json

Go-তে struct-কে JSON-এ convert (`json.Marshal`) এবং **exported vs unexported** field-এর ভূমিকা বুঝতে সহায়ক ছোট example।

**📖 ভাষা নির্বাচন করুন / Choose language:**

[🇧🇩 বাংলা](#bangla) • [🇬🇧 English](#english)

---

<a name="bangla"></a>

## 🇧🇩 বাংলা সংস্করণ

### Lines 1–6

```go
package main

import (
    "encoding/json"
    "fmt"
)
```

- `package main` — এটা একটা executable program।
- আমরা দুটো package import করি: `encoding/json` (JSON-এর সাথে কাজ করতে) আর `fmt` (console-এ output print করতে)।

### Lines 8–12

```go
type Invoice struct {
    ID           string
    CustomerName string
    amountCents  int
}
```

`Invoice` নামক একটা struct type declare করে এর তিনটা field:

- `ID` — একটা string। **বড় হাতের** letter দিয়ে শুরু, তাই এটা **exported** (package-র বাইরে visible এবং JSON-এ অন্তর্ভুক্ত)।
- `CustomerName` — একটা string। বড় হাতের letter দিয়ে শুরু → exported → JSON-এ অন্তর্ভুক্ত।
- `amountCents` — একটা int। **ছোট হাতের** letter দিয়ে শুরু, তাই এটা **unexported** (শুধু এই package-এর ভেতরে visible এবং **JSON থেকে বাদ**)।

### Line 14

```go
func main() {
```

Program-এর entry point।

### Lines 15–19

```go
inv := Invoice{
    ID:           "INV-1000",
    CustomerName: "John Doe",
    amountCents:  5600,
}
```

Named field initialization দিয়ে `inv` নামক একটা `Invoice` instance তৈরি করা হয়। সব field সেট করা হয়। যেহেতু আমরা একই package-এ আছি, এমনকি unexported field `amountCents`-ও সেট করতে পারি।

### Line 21

```go
data, err := json.Marshal(inv)
```

`json.Marshal(inv)` struct-কে JSON bytes-এ convert করে। এটা দুটো value return করে: `data` (JSON টা `[]byte` হিসেবে) আর `err` (হতে পারে এমন যেকোনো error)।

### Lines 23–26

```go
if err != nil {
    fmt.Println("Marshal error", err)
    return
}
```

Error check করা হয়। হলে error print করে `return` দিয়ে program বন্ধ করে দেয়।

### Line 28

```go
fmt.Println(string(data))
```

JSON bytes `data`-কে string-এ convert করে print করে।

`json.Marshal` শুধু **exported (বড় হাতের)** field-গুলোই অন্তর্ভুক্ত করে, তাই output হলো:

`{"ID":"INV-1000","CustomerName":"John Doe"}`

লক্ষ্য করো: `amountCents` **নেই** — কারণ এটা ছোট হাতের/unexported, `encoding/json` package (একটি ভিন্ন package) এটা access করতে পারে না, তাই JSON-এ এটা বাদ থাকে।

### Line 29

```go
fmt.Println("Internal amount (only visible inside this package):", inv.amountCents)
```

নিজের package-এর ভেতর থেকে `inv.amountCents` সরাসরি পড়তে পারি এবং `5600` হিসেবে print করতে পারি। Value-টা আসলে আছে — শুধু JSON-এ expose হয় না।

### Line 30

```go
}
```

Closing brace — `main` function শেষ হয়।

---

## Expected Output

```
{"ID":"INV-1000","CustomerName":"John Doe"}
Internal amount (only visible inside this package): 5600
```

## মূল শিক্ষা / Key Takeaways

1. **`json.Marshal`** — struct-কে JSON bytes-এ convert করে।
2. **Exported fields only** — JSON-এ শুধু বড় হাতের (exported) field-ই যায়।
3. **Unexported = hidden** — ছোট হাতের field JSON-এ leak হয় না — sensitive data-র জন্য ভালো।
4. **Same package access** — নিজের package-এ unexported field-ও readable।
5. **Tip** — চাইলে `json:"amount_cents"` tag দিয়ে exported field-ও rename করা যায়।

---

---

<a name="english"></a>

## 🇬🇧 English Version

### Lines 1–6

```go
package main

import (
    "encoding/json"
    "fmt"
)
```

- `package main` — this is an executable program.
- We import two packages: `encoding/json` (to work with JSON) and `fmt` (for console output).

### Lines 8–12

```go
type Invoice struct {
    ID           string
    CustomerName string
    amountCents  int
}
```

Declares a struct type `Invoice` with three fields:

- `ID` — a string. Starts **uppercase**, so it is **exported** (visible outside the package and included in JSON).
- `CustomerName` — a string. Also uppercase → exported → included in JSON.
- `amountCents` — an int. Starts **lowercase**, so it is **unexported** (only visible inside this package and **excluded from JSON**).

### Line 14

```go
func main() {
```

Program entry point.

### Lines 15–19

```go
inv := Invoice{
    ID:           "INV-1000",
    CustomerName: "John Doe",
    amountCents:  5600,
}
```

Creates an `Invoice` instance called `inv` using named field initialization. All fields are set. Because we are in the same package, we can even set the unexported field `amountCents`.

### Line 21

```go
data, err := json.Marshal(inv)
```

`json.Marshal(inv)` converts the struct to JSON bytes. It returns two values: `data` (the JSON as `[]byte`) and `err` (any error that occurred).

### Lines 23–26

```go
if err != nil {
    fmt.Println("Marshal error", err)
    return
}
```

Checks for an error. If one occurred, print the error and `return` to stop the program.

### Line 28

```go
fmt.Println(string(data))
```

Converts the JSON bytes `data` to a string and prints it.

`json.Marshal` only includes **exported (uppercase)** fields, so the output is:

`{"ID":"INV-1000","CustomerName":"John Doe"}`

Notice: `amountCents` is **missing** — because it's lowercase/unexported, the `encoding/json` package (a different package) cannot access it, so it's left out of the JSON.

### Line 29

```go
fmt.Println("Internal amount (only visible inside this package):", inv.amountCents)
```

From within our own package, we can still read `inv.amountCents` directly and print it as `5600`. The value exists — it is just not exposed in the JSON.

### Line 30

```go
}
```

Closing brace — ends the `main` function.

---

## Expected Output

```
{"ID":"INV-1000","CustomerName":"John Doe"}
Internal amount (only visible inside this package): 5600
```

## Key Takeaways

1. **`json.Marshal`** — Converts a struct to JSON bytes.
2. **Exported fields only** — Only uppercase (exported) fields appear in JSON.
3. **Unexported = hidden** — Lowercase fields aren't leaked into JSON — good for sensitive data.
4. **Same package access** — Within the same package, unexported fields are still readable.
5. **Tip** — You can rename/exclude fields with struct tags like `json:"..."`.
