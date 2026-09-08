# paginated-fetching

Go-তে **`for` + `break` দিয়ে paginated fetch**, **empty page-এ থামা**, আর **slice spread (`...`) দিয়ে `append`** শেখার ছোট example.

**ভাষা নির্বাচন করুন / Choose language:**

[বাংলা](#bangla) * [English](#english)

---

<a name="bangla"></a>

## বাংলা সংস্করণ

### Line 1

```go
package main
```

একটা executable program (`main` package) declare করে, যা `go run` দিয়ে চালানো যায়.

### Line 3

```go
import "fmt"
```

Console-এ output print করার জন্য `fmt` package import করা হয়.

### Lines 5-13

```go
func fetchPage(page int) []string {
	if page > 2 {
		return nil
	}
	return []string{
		fmt.Sprintf("order-%d-a", page),
		fmt.Sprintf("order-%d-b", page),
	}
}
```

`fetchPage` -- একটা page-এর order simulate করে:

- `if page > 2 { return nil }` -- 2-এর বেশি page চাইলে খালি (`nil` slice) return করে. এটাই "আর data নেই" signal.
- `fmt.Sprintf("order-%d-a", page)` -- `%d`-তে page number বসিয়ে string বানায় (যেমন page 1 -> `"order-1-a"`). `Sprintf` print না করে string return করে.
- প্রতিটা valid page-এ 2টা order return করে.

### Line 15

```go
func main() {
```

Program-এর entry point.

### Lines 16-17

```go
var allOrders []string
page := 1
```

- `var allOrders []string` -- সব page-এর order জমানোর slice, শুরুতে `nil` (খালি).
- `page := 1` -- page counter, 1 থেকে শুরু (API-তে page সাধারণত 1-based).

### Lines 20-29

```go
for {
	orders := fetchPage(page)
	if len(orders) == 0 {
		break
	}

	allOrders = append(allOrders, orders...)
	fmt.Printf("page %d: fetched %d orders\n", page, len(orders))
	page++
}
```

**Infinite `for` + `break` pattern** -- কতটা page আছে আগে জানা নেই, তাই loop চালিয়ে যাওয়া হয় যতক্ষণ data আসে:

- `orders := fetchPage(page)` -- current page fetch করে.
- `if len(orders) == 0 { break }` -- খালি page মানে সব data শেষ, loop থেকে বেরিয়ে যাওয়া. (`nil` slice-এর `len` 0, তাই এটা কাজ করে.)
- `allOrders = append(allOrders, orders...)` -- নতুন order-গুলো আগেরগুলোর সাথে যোগ করে. `orders...` (**spread**) মানে slice-এর প্রতিটা element আলাদা argument হিসাবে পাঠানো -- না হলে পুরো slice একটা element হিসাবে ঢুকতে চাইতো (compile error).
- `fmt.Printf("page %d: fetched %d orders\n", page, len(orders))` -- progress log.
- `page++` -- পরের page-এ যাওয়া.

**Flow:**

1. page=1 -> `["order-1-a" "order-1-b"]` -> allOrders-এ যোগ (মোট 2) -> prints "page 1: fetched 2 orders" -> page=2
2. page=2 -> `["order-2-a" "order-2-b"]` -> allOrders-এ যোগ (মোট 4) -> prints "page 2: fetched 2 orders" -> page=3
3. page=3 -> `nil` -> `len == 0` -> `break`

### Line 31

```go
fmt.Println("total orders:", len(allOrders))
```

`len(allOrders)` = 4 -- মোট কয়টা order পাওয়া গেছে print করে.

### Line 32

```go
}
```

Closing brace -- `main` function শেষ হয়.

---

## Expected Output

```
page 1: fetched 2 orders
page 2: fetched 2 orders
total orders: 4
```

## মূল শিক্ষা / Key Takeaways

1. **`for { ... break }`** -- কত iteration লাগবে জানা না থাকলে infinite loop + condition-এ break.
2. **Empty page = stop signal** -- `len(orders) == 0` মানে আর data নেই (`nil`-এর `len` 0).
3. **Slice spread `...`** -- `append(dst, src...)` দিয়ে এক slice-এর সব element আরেকটায় যোগ.
4. **`Sprintf`** -- format করে string বানায় (print করে না).
5. **`var s []string`** -- শুরুতে `nil` slice; `append` দিয়ে বাড়ানো যায়.

---

---

<a name="english"></a>

##  English Version

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

### Lines 5-13

```go
func fetchPage(page int) []string {
	if page > 2 {
		return nil
	}
	return []string{
		fmt.Sprintf("order-%d-a", page),
		fmt.Sprintf("order-%d-b", page),
	}
}
```

`fetchPage` -- simulates one page of orders:

- `if page > 2 { return nil }` -- returns an empty (`nil`) slice past page 2. This is the "no more data" signal.
- `fmt.Sprintf("order-%d-a", page)` -- builds a string with the page number (e.g. page 1 -> `"order-1-a"`). `Sprintf` returns a string instead of printing.
- Each valid page returns 2 orders.

### Line 15

```go
func main() {
```

Program entry point.

### Lines 16-17

```go
var allOrders []string
page := 1
```

- `var allOrders []string` -- accumulates orders from all pages, starts as `nil` (empty).
- `page := 1` -- page counter starting at 1 (APIs usually use 1-based pages).

### Lines 20-29

```go
for {
	orders := fetchPage(page)
	if len(orders) == 0 {
		break
	}

	allOrders = append(allOrders, orders...)
	fmt.Printf("page %d: fetched %d orders\n", page, len(orders))
	page++
}
```

The **infinite `for` + `break` pattern** -- we don't know how many pages exist, so we keep looping while data arrives:

- `orders := fetchPage(page)` -- fetches the current page.
- `if len(orders) == 0 { break }` -- an empty page means all data is done, exit the loop. (A `nil` slice has `len` 0, so this works.)
- `allOrders = append(allOrders, orders...)` -- merges the new orders into the accumulator. `orders...` (**spread**) passes each element as a separate argument -- without it the whole slice would be treated as one element (compile error).
- `fmt.Printf("page %d: fetched %d orders\n", page, len(orders))` -- progress log.
- `page++` -- moves to the next page.

**Flow:**

1. page=1 -> `["order-1-a" "order-1-b"]` -> appended (total 2) -> prints "page 1: fetched 2 orders" -> page=2
2. page=2 -> `["order-2-a" "order-2-b"]` -> appended (total 4) -> prints "page 2: fetched 2 orders" -> page=3
3. page=3 -> `nil` -> `len == 0` -> `break`

### Line 31

```go
fmt.Println("total orders:", len(allOrders))
```

`len(allOrders)` = 4 -- prints how many orders were collected.

### Line 32

```go
}
```

Closing brace -- ends the `main` function.

---

## Expected Output

```
page 1: fetched 2 orders
page 2: fetched 2 orders
total orders: 4
```

## Key Takeaways

1. **`for { ... break }`** -- infinite loop + conditional break when the iteration count is unknown.
2. **Empty page = stop signal** -- `len(orders) == 0` means no more data (`nil` has `len` 0).
3. **Slice spread `...`** -- `append(dst, src...)` merges all elements of one slice into another.
4. **`Sprintf`** -- builds a formatted string (does not print).
5. **`var s []string`** -- starts as a `nil` slice; grows via `append`.
