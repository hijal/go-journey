# customer-profile-nested

Go-তে **nested struct + struct reuse + slice field** শেখার ছোট example — `Customer`-এর ভেতরে দুটো `Address` struct, nested field access/mutation।

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

`fmt` — `Println`, `Printf`।

### Lines 5–9

```go
type Address struct {
	City    string
	Street  string
	ZipCode string
}
```

`Address` struct — City/Street/ZipCode। এটা **reusable** — একাধিক জায়গায় একই struct type ব্যবহার করা যায়।

### Lines 11–16

```go
type Customer struct {
	Name        string
	Home        Address
	Office      Address
	PhoneNumber []string
}
```

`Customer` — composition:

- `Home`/`Office` — **দুটো আলাদা `Address`** — struct-এর ভেতরে struct (nested)।
- `PhoneNumber []string` — slice field — একাধিক নম্বর।

Nested struct-টা **grouping + type safety** দেয় (City vs PhoneNumber mismash-এর field-এ মিশে যায় না)।

### Lines 18–23

```go
func main() {
	c := Customer{
		Name: "John Doe",
		Home: Address{
			City:    "New York",
			Street:  "123 Main St",
			ZipCode: "10001",
		},
```

**Nested composite literal** — `Home: Address{...}` — parent literal-এ child struct literal embed।

### Lines 26–32

```go
		Office: Address{
			City:    "New York",
			Street:  "456 Office Rd",
			ZipCode: "10002",
		},
		PhoneNumber: []string{"123-456-7890", "987-654-3210"},
	}
```

- দ্বিতীয় `Address` (Office) — একই struct type, আলাদা data।
- `PhoneNumber: []string{...}` — slice literal।

### Lines 34–37

```go
	fmt.Println("Home city:", c.Home.City)
	fmt.Println("Office city:", c.Office.City)
	c.Home.City = "Auckland"
	fmt.Printf("Updated home: %+v\n", c.Home)
```

**Nested access**:

- `c.Home.City` — dot-chain: struct → nested struct → field।
- `c.Home.City = "Auckland"` — deep field mutate — dot-chain+assign।
- `%+v` — `c.Home`-এর field name সহ print।

---

## Expected Output

```
Home city: New York
Office city: New York
Updated home: {City:Auckland Street:123 Main St ZipCode:10001}
```

## মূল শিক্ষা / Key Takeaways

1. **Nested struct** — struct field-এ struct (`Home Address`)।
2. **Struct reuse** — same `Address` type, ভিন্ন instance।
3. **Composite literal nesting** — `Home: Address{...}` embed।
4. **Dot-chain access/mutation** — `c.Home.City = ...`।
5. **Slice field** — `[]string` struct-এর ভিতরে।

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

`fmt` — for `Println`, `Printf`.

### Lines 5–9

```go
type Address struct {
	City    string
	Street  string
	ZipCode string
}
```

`Address` struct — City/Street/ZipCode. It's **reusable** — the same struct type can be used in multiple places.

### Lines 11–16

```go
type Customer struct {
	Name        string
	Home        Address
	Office      Address
	PhoneNumber []string
}
```

`Customer` — composition:

- `Home`/`Office` — **two separate `Address` values** — nesting a struct inside a struct.
- `PhoneNumber []string` — a slice field — multiple numbers.

Nested structs provide **grouping + type safety** (City won't get mixed with PhoneNumber fields).

### Lines 18–23

```go
func main() {
	c := Customer{
		Name: "John Doe",
		Home: Address{
			City:    "New York",
			Street:  "123 Main St",
			ZipCode: "10001",
		},
```

**Nested composite literal** — `Home: Address{...}` — a child struct literal embedded inside the parent literal.

### Lines 26–32

```go
		Office: Address{
			City:    "New York",
			Street:  "456 Office Rd",
			ZipCode: "10002",
		},
		PhoneNumber: []string{"123-456-7890", "987-654-3210"},
	}
```

- The second `Address` (Office) — the same struct type, different data.
- `PhoneNumber: []string{...}` — a slice literal.

### Lines 34–37

```go
	fmt.Println("Home city:", c.Home.City)
	fmt.Println("Office city:", c.Office.City)
	c.Home.City = "Auckland"
	fmt.Printf("Updated home: %+v\n", c.Home)
```

**Nested access**:

- `c.Home.City` — a dot-chain: struct → nested struct → field.
- `c.Home.City = "Auckland"` — deep field mutation — dot-chain + assignment.
- `%+v` — prints `c.Home` with field names.

---

## Expected Output

```
Home city: New York
Office city: New York
Updated home: {City:Auckland Street:123 Main St ZipCode:10001}
```

## Key Takeaways

1. **Nested struct** — a struct as a struct field (`Home Address`).
2. **Struct reuse** — the same `Address` type, different instances.
3. **Composite literal nesting** — embedding `Home: Address{...}`.
4. **Dot-chain access/mutation** — `c.Home.City = ...`.
5. **Slice field** — a `[]string` inside a struct.