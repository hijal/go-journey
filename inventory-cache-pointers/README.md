# inventory-cache-pointers

Go-তে **map storing pointers (`map[string]*T`) + pointer dereference mutation + validation** শেখার ছোট example — in-place stock reduce।

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

`fmt` — `Printf`, `Println`, `Errorf`।

### Lines 5–9

```go
type Product struct {
	SKU   string
	Name  string
	Stock int
}
```

Product model।

### Lines 11–13

```go
type Inventory struct {
	items map[string]*Product
}
```

**Pointer-value map** — `SKU → *Product` (value নয়, মূল object pointer)।

### Lines 15–17

```go
func NewInventory() *Inventory {
	return &Inventory{items: make(map[string]*Product)}
}
```

Constructor — map initialize (nil-map-এ write panic)।

### Lines 19–21

```go
func (i *Inventory) Add(p *Product) {
	i.items[p.SKU] = p
}
```

Pointer-same store — পাস করা object-টা make।

### Lines 23–38

```go
func (i *Inventory) ReduceStock(sku string, quantity int) error {
	product, ok := i.items[sku]

	if !ok {
		return fmt.Errorf("product %s not found", sku)
	}

	if product.Stock < quantity {
		return fmt.Errorf("insufficient stock for %s: have %d, want %d",
			sku, product.Stock, quantity)
	}

	product.Stock -= quantity

	return nil
}
```

**Pointer-mutation essence:**

1. Comma-ok lookup — না-পাওয়া product error।
2. Stock-guard — `product.Stock < quantity` reject।
3. `product.Stock -= quantity` — **dereference** ব্যবহার করে map-এর object mutate।

*(Map-value হতো copy; pointer-নির্ভর map-এ same object-কে mutate করে — outer access-এ ফেরত visible।)*

### Lines 40–43

```go
	inv := NewInventory()

	inv.Add(&Product{SKU: "SKU-1", Name: "Mechanical Keyboard", Stock: 15})
```

15 stock-সহ product add।

### Lines 45–48

```go
	if err := inv.ReduceStock("SKU-1", 3); err != nil {
		fmt.Println("order error:", err)
	}
	fmt.Printf("Remaining stock: %d\n", inv.items["SKU-1"].Stock)
```

3 কমায় → map-এর product-এ `12`।

---

## Expected Output

```
Remaining stock: 12
```

## মূল শিক্ষা / Key Takeaways

1. **`map[string]*T`** — মূল object store।
2. **Pointer dereference mutate** — map-entry-ই বদলায়।
3. **Comma-ok guard** — missing-product error।
4. **Quantity guard** — insufficient stock।
5. **Value-vs-pointer map** — copy নয়, shared object।

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

`fmt` — for `Printf`, `Println`, `Errorf`.

### Lines 5–9

```go
type Product struct {
	SKU   string
	Name  string
	Stock int
}
```

The product model.

### Lines 11–13

```go
type Inventory struct {
	items map[string]*Product
}
```

A **pointer-value map** — `SKU → *Product` (pointers to the actual objects, not copies).

### Lines 15–17

```go
func NewInventory() *Inventory {
	return &Inventory{items: make(map[string]*Product)}
}
```

The constructor — the map is initialized (writing to a nil map panics).

### Lines 19–21

```go
func (i *Inventory) Add(p *Product) {
	i.items[p.SKU] = p
}
```

Stores the pointer — keeps the very object that was passed.

### Lines 23–38

```go
func (i *Inventory) ReduceStock(sku string, quantity int) error {
	product, ok := i.items[sku]

	if !ok {
		return fmt.Errorf("product %s not found", sku)
	}

	if product.Stock < quantity {
		return fmt.Errorf("insufficient stock for %s: have %d, want %d",
			sku, product.Stock, quantity)
	}

	product.Stock -= quantity

	return nil
}
```

**The essence of pointer mutation:**

1. A comma-ok lookup — a missing product errors.
2. A stock guard — `product.Stock < quantity` rejects.
3. `product.Stock -= quantity` — **dereferences** to mutate the object inside the map.

*(A value map would mutate a copy; a pointer map mutates the same object shown through outer access.)*

### Lines 40–43

```go
	inv := NewInventory()

	inv.Add(&Product{SKU: "SKU-1", Name: "Mechanical Keyboard", Stock: 15})
```

Adds a product with 15 stock.

### Lines 45–48

```go
	if err := inv.ReduceStock("SKU-1", 3); err != nil {
		fmt.Println("order error:", err)
	}
	fmt.Printf("Remaining stock: %d\n", inv.items["SKU-1"].Stock)
```

Reduces by 3 → `12` on the map's product.

---

## Expected Output

```
Remaining stock: 12
```

## Key Takeaways

1. **`map[string]*T`** — storing the actual objects.
2. **Pointer dereference mutation** — the map entry itself changes.
3. **Comma-ok guard** — missing-product errors.
4. **Quantity guard** — insufficient stock protection.
5. **Value-vs-pointer map** — shared object, not a copy.