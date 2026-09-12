# product-value-pointer

Go-তে **value semantics vs pointer semantics** শেখার ছোট example — value-এ struct-টা copy হয়, pointer-এ same struct mutate হয়; struct copy-ও দেখায়।

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

`fmt` — `Println`।

### Lines 5–9

```go
type Product struct {
	SKU   string
	Name  string
	Price float64
}
```

`Product` struct — SKU, Name, Price।

### Lines 11–13

```go
func applyDiscountBad(p Product) {
	p.Price *= 0.9
}
```

`applyDiscountBad` — **value receiver** (নামে "bad" — ইচ্ছাকৃত)।

- `p Product` — struct-টা **copy** হয়ে আসে।
- `p.Price *= 0.9` — **copy-তে** 10% discount apply।
- Function শেষে copy dispose — **caller-এর struct অপরিবর্তিত**। এটাই value semantics-এর trap।

### Lines 15–17

```go
func applyDiscount(p *Product, pct float64) {
	p.Price *= 1 - pct
}
```

`applyDiscount` — **pointer receiver**, discount percentage param সহ।

- `p *Product` — original struct-এর **pointer** (address)।
- `p.Price *= 1 - pct` — `pct=0.1` হলে `× 0.9`। Direct mutate — caller-এ দেখা যায়।
- এটাই pointer semantics: function টা মূল struct-কে বদলায়।

### Lines 19–24

```go
func main() {
	laptop := Product{
		SKU:   "LAP123",
		Name:  "Laptop Pro 14",
		Price: 90000.0,
	}
```

`laptop` — composite literal-এ initialize, `Price: 90000`.

### Lines 26–30

```go
	applyDiscountBad(laptop)
	fmt.Println("After bad discount:", laptop)

	applyDiscount(&laptop, 0.1)
	fmt.Println("After good discount:", laptop)
```

- `applyDiscountBad(laptop)` — value pass → Price অপরিবর্তিত `90000`।
- `applyDiscount(&laptop, 0.1)` — **`&laptop`** (address) pass → `90000 × 0.9 = 81000` ✓

**মূল পয়েন্ট:** `%v` (direct) print করলে struct value-টা field-order-এ আসে।

### Lines 32–35

```go
	backup := laptop
	backup.Price = 0
	fmt.Println("Backup:", backup)
	fmt.Println("Laptop untouched:", laptop)
```

**Struct copy:**

- `backup := laptop` — পুরো struct-টা **copy** (সব field)।
- `backup.Price = 0` — copy-তে update।
- `laptop` untouched — `81000`-ই থাকে। দুটো আলাদা data।

---

## Expected Output

```
After bad discount: {LAP123 Laptop Pro 14 90000}
After good discount: {LAP123 Laptop Pro 14 81000}
Backup: {LAP123 Laptop Pro 14 0}
Laptop untouched: {LAP123 Laptop Pro 14 81000}
```

## মূল শিক্ষা / Key Takeaways

1. **Value semantics** — struct pass the name copy; mutations এখানে-ওখানে হারায়।
2. **Pointer semantics** — `*Product` দিয়ে same struct mutate।
3. **`&` operator** — struct-এর address নেওয়া।
4. **Struct copy** — `backup := laptop` সব field copy; independent।
5. **Mutate-lাগলে pointer, read-লে value** — receiver select-র মূল নিয়ম।

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

`fmt` — for `Println`.

### Lines 5–9

```go
type Product struct {
	SKU   string
	Name  string
	Price float64
}
```

`Product` struct — SKU, Name, Price.

### Lines 11–13

```go
func applyDiscountBad(p Product) {
	p.Price *= 0.9
}
```

`applyDiscountBad` — **value receiver** (the name says "bad" — intentional).

- `p Product` — the struct arrives as a **copy**.
- `p.Price *= 0.9` — applies a 10% discount to the **copy**.
- The copy is discarded when the function returns — **the caller's struct is unchanged**. That's the value-semantics trap.

### Lines 15–17

```go
func applyDiscount(p *Product, pct float64) {
	p.Price *= 1 - pct
}
```

`applyDiscount` — **pointer receiver**, with a percentage param.

- `p *Product` — a **pointer** (address) to the original struct.
- `p.Price *= 1 - pct` — with `pct=0.1` that's `× 0.9`. A direct mutation, visible to the caller.
- This is pointer semantics: the function mutates the actual struct.

### Lines 19–24

```go
func main() {
	laptop := Product{
		SKU:   "LAP123",
		Name:  "Laptop Pro 14",
		Price: 90000.0,
	}
```

`laptop` — initialized with a composite literal, `Price: 90000`.

### Lines 26–30

```go
	applyDiscountBad(laptop)
	fmt.Println("After bad discount:", laptop)

	applyDiscount(&laptop, 0.1)
	fmt.Println("After good discount:", laptop)
```

- `applyDiscountBad(laptop)` — passed by value → Price unchanged at `90000`.
- `applyDiscount(&laptop, 0.1)` — passes **`&laptop`** (address) → `90000 × 0.9 = 81000` ✓

**Key point:** printing the struct directly with `%v` shows the value in field order.

### Lines 32–35

```go
	backup := laptop
	backup.Price = 0
	fmt.Println("Backup:", backup)
	fmt.Println("Laptop untouched:", laptop)
```

**Struct copy:**

- `backup := laptop` — the whole struct is **copied** (all fields).
- `backup.Price = 0` — updates the copy.
- `laptop` stays untouched — still `81000`. Two separate data sets.

---

## Expected Output

```
After bad discount: {LAP123 Laptop Pro 14 90000}
After good discount: {LAP123 Laptop Pro 14 81000}
Backup: {LAP123 Laptop Pro 14 0}
Laptop untouched: {LAP123 Laptop Pro 14 81000}
```

## Key Takeaways

1. **Value semantics** — a struct passed by value is copied; mutations are lost.
2. **Pointer semantics** — mutate the same struct with `*Product`.
3. **`&` operator** — taking the address of a struct.
4. **Struct copy** — `backup := laptop` copies all fields; fully independent.
5. **Pointer when mutating, value when reading** — the core receiver rule.