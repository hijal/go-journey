# cart-methods-receivers

Go-তে **Value vs Pointer receiver** শেখার ছোট example — shopping cart-এর সাধারণ operations (`AddItem`, `RemoveItem`, `TotalPrice`) যেখানে mutation-কারী method-গুলো pointer receiver এবং read-only method-টা value receiver ব্যবহার করে।

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

- `fmt` — `Printf` (formatted output)।

### Lines 5–13

```go
type Item struct {
	ID    string
	Price float64
	Qty   int
}

type Cart struct {
	Items []Item
}
```

দুটো struct:

- `Item` — একটি পণ্য: `ID`, `Price` (float64), `Qty` (int)।
- `Cart` — `Items` slice, যাতে একাধিক `Item` থাকে (slice → dynamic list)।

### Lines 15–18

```go
func (c *Cart) AddItem(item Item) {
	c.Items = append(c.Items, item)
}
```

`AddItem` — **pointer receiver** (`(c *Cart)`):

- `append(c.Items, item)` — নতুন item টা cart-এর list-এ যোগ করে।
- কেন pointer receiver? কারণ এটা cart-এর **state পরিবর্তন করছে** (`c.Items` slice-টা বদলায়)।
- Receiver pointer হলে `c.Items = ...` লেখা **original struct-এ** effect করে।

### Lines 20–26

```go
func (c *Cart) RemoveItem(itemID string) {
	for i, item := range c.Items {
		if item.ID == itemID {
			c.Items = append(c.Items[:i], c.Items[i+1:]...)
			return
		}
	}
}
```

`RemoveItem` — **pointer receiver**, item ID দিয়ে খুঁজে মুছে ফেলে:

- `for i, item := range c.Items` — প্রতিটা index (`i`) আর item-এর উপর loop।
- `item.ID == itemID` — মিল পেলে:
  - `c.Items[:i]` — index `i` পর্যন্ত অংশ।
  - `c.Items[i+1:]...` — `i`-এর পরের অংশ, **`...`** দিয়ে `append`-এ unpack।
  - মাঝের element বাদ → item মুছে যায়।
  - `return` — পাওয়া গেলে loop থামায় (আর scan-এ জরুরি নেই)।
- আবারও pointer receiver — কারণ `c.Items` **mutate** হচ্ছে।

### Lines 28–34

```go
func (c Cart) TotalPrice() float64 {
	var total float64
	for _, item := range c.Items {
		total += item.Price * float64(item.Qty)
	}
	return total
}
```

`TotalPrice` — **value receiver** (`(c Cart)`):

- `c.Items`-এর প্রতিটা item-এর `Price * Qty` (গুণ করে `float64`-এ cast, কারণ `Qty` int)।
- সব যোগ → `total`।
- **কেন value receiver?** এটা cart-এর state **পরিবর্তন করে না** — শুধু পড়ে। Value receiver-এ struct-এর **copy** নেওয়া হয়; পরিবর্তন না করলে copy-তেই কোনো ক্ষতি নেই।
- মানে: pointer receiver তখনই দরকার যখন state mutate করতে হবে; না হলে value receiver-ই ভালো (সাধারণ Go convention)।

### Lines 36–46

```go
func main() {
	cart := &Cart{}

	cart.AddItem(Item{ID: "apple", Price: 1.20, Qty: 5})
	cart.AddItem(Item{ID: "banana", Price: 0.80, Qty: 10})

	fmt.Printf("Total: $%.2f\n", cart.TotalPrice())

	cart.RemoveItem("apple")

	fmt.Printf("total after remove: $%.2f\n", cart.TotalPrice())
}
```

- `cart := &Cart{}` — empty cart-এর pointer।
- দুটো item add: apple (5 × 1.20 = 6.00), banana (10 × 0.80 = 8.00)।
- `cart.TotalPrice()` — `Total` = **14.00**। `%.2f` দিয়ে 2 দশমিক।
- `cart.RemoveItem("apple")` — apple মুছে।
- `cart.TotalPrice()` আবার → **8.00**।

**Auto-address-of:** `cart` pointer, তাই pointer receiver-এর method-গুলো সরাসরি কাজ করে। `cart` value হওয়া-ও চলত — Go নিজে `(&cart).AddItem(...)` করে দিত।

---

## Expected Output

```
Total: $14.00
total after remove: $8.00
```

## মূল শিক্ষা / Key Takeaways

1. **Pointer receiver (`*Cart`)** — যখন method struct-এর state mutate করে (`AddItem`, `RemoveItem`)।
2. **Value receiver (`Cart`)** — যখন method শুধু পড়ে (`TotalPrice`) — state অপরিবর্তনীয়, copy-তেই কাজ চলে।
3. **Rule of thumb** — mutate করতে হলে pointer; না হলে value receiver (কিংবা consistency-র জন্য pointer)।
4. **Auto-address-of** — value-তে pointer-receiver method call করলে Go নিজে `&` নেয়।
5. **Slice splice** — `append(s[:i], s[i+1:]...)` দিয়ে মাঝের element delete।

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

- `fmt` — for `Printf` (formatted output).

### Lines 5–13

```go
type Item struct {
	ID    string
	Price float64
	Qty   int
}

type Cart struct {
	Items []Item
}
```

Two structs:

- `Item` — a product: `ID`, `Price` (float64), `Qty` (int).
- `Cart` — an `Items` slice holding multiple `Item`s (slice → dynamic list).

### Lines 15–18

```go
func (c *Cart) AddItem(item Item) {
	c.Items = append(c.Items, item)
}
```

`AddItem` — **pointer receiver** (`(c *Cart)`):

- `append(c.Items, item)` — adds the new item to the cart.
- Why a pointer receiver? Because it's **mutating the cart's state** (`c.Items` changes).
- With a pointer receiver, `c.Items = ...` affects the **original struct**.

### Lines 20–26

```go
func (c *Cart) RemoveItem(itemID string) {
	for i, item := range c.Items {
		if item.ID == itemID {
			c.Items = append(c.Items[:i], c.Items[i+1:]...)
			return
		}
	}
}
```

`RemoveItem` — **pointer receiver**, finds by ID and removes it:

- `for i, item := range c.Items` — loops over each index (`i`) and item.
- On match (`item.ID == itemID`):
  - `c.Items[:i]` — the items up to index `i`.
  - `c.Items[i+1:]...` — the items after `i`, unpacked with `...` for `append`.
  - The middle element is dropped → the item is gone.
  - `return` — stops the loop once found.
- Again a pointer receiver — because `c.Items` **mutates**.

### Lines 28–34

```go
func (c Cart) TotalPrice() float64 {
	var total float64
	for _, item := range c.Items {
		total += item.Price * float64(item.Qty)
	}
	return total
}
```

`TotalPrice` — **value receiver** (`(c Cart)`):

- For each item, `Price * Qty` (cast to `float64` because `Qty` is int).
- Sums everything into `total`.
- **Why a value receiver?** It doesn't **change** the cart's state — it only reads. A value receiver takes a **copy**; that's harmless when nothing is modified.
- Meaning: use a pointer receiver only when you must mutate state; otherwise a value receiver is the idiomatic Go choice.

### Lines 36–46

```go
func main() {
	cart := &Cart{}

	cart.AddItem(Item{ID: "apple", Price: 1.20, Qty: 5})
	cart.AddItem(Item{ID: "banana", Price: 0.80, Qty: 10})

	fmt.Printf("Total: $%.2f\n", cart.TotalPrice())

	cart.RemoveItem("apple")

	fmt.Printf("total after remove: $%.2f\n", cart.TotalPrice())
}
```

- `cart := &Cart{}` — a pointer to an empty cart.
- Adds two items: apple (5 × 1.20 = 6.00), banana (10 × 0.80 = 8.00).
- `cart.TotalPrice()` — `Total` = **14.00**. `%.2f` prints 2 decimals.
- `cart.RemoveItem("apple")` — removes the apple.
- `cart.TotalPrice()` again → **8.00**.

**Auto-address-of:** `cart` is a pointer, so the pointer-receiver methods just work. It would also work if `cart` were a value — Go automatically inserts `(&cart).AddItem(...)`.

---

## Expected Output

```
Total: $14.00
total after remove: $8.00
```

## Key Takeaways

1. **Pointer receiver (`*Cart`)** — when the method mutates the struct's state (`AddItem`, `RemoveItem`).
2. **Value receiver (`Cart`)** — when the method only reads (`TotalPrice`) — state immutable, works on a copy.
3. **Rule of thumb** — pointer to mutate; value receiver (or pointer for consistency) otherwise.
4. **Auto-address-of** — Go inserts `&` when calling a pointer-receiver method on a value.
5. **Slice splice** — delete a middle element with `append(s[:i], s[i+1:]...)`.