# invoice-batch-scanner

Go-তে **`for range` loop**, **`continue`**, আর **`break`** দিয়ে invoice-এর batch process (scan) শেখার ছোট example।

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

### Line 6

```go
invoiceAmounts := []float64{1200, 0, 850, -50, 3000, 400}
```

`invoiceAmounts` — একটা `[]float64` slice যাতে ৬টা invoice-এর পরিমাণ আছে: `1200, 0, 850, -50, 3000, 400`। লক্ষ্য করো এদের মধ্যে `0` (বৈধ কিন্তু শূন্য) আর `-50` (invalid/negative) আছে — process করার সময় এগুলোর সাথে বিশেষ আচরণ করা হবে।

### Line 7

```go
var total float64
```

`var total float64` — একটা accumulator variable, যাতে process করা invoice-গুলোর যোগফল জমা হবে। শুরুতে `0` (float zero value)।

### Line 9

```go
for i, amount := range invoiceAmounts {
```

`for ... range` loop slice-এর প্রতিটা element-এ যায়। প্রতিটিতে `i` = index, `amount` = পরিমাণ।

### Lines 10–12

```go
if amount == 0 {
    continue
}
```

যদি `amount == 0` হয়, তাহলে **`continue`** — বাকি loop body skip করে পরের element-এ যাই। অর্থাৎ শূন্য-পরিমাণের invoice-গুলো skipp করা হয় (তারা total-এ যোগ হয় না, print-ও হয় না)।

### Lines 13–16

```go
if amount < 0 {
    fmt.Println("invalid invoice at index", i, "- stopping batch")
    break
}
```

যদি `amount < 0` হয় (নেতিবাচক/অবৈধ), তাহলে message print করে **`break`** — পুরো loop ছেড়ে বেরিয়ে যায়। মানে নেতিবাচক invoice পেলে batch process **বন্ধ** হয়ে যায়।

### Lines 17–18

```go
total += amount
fmt.Println("processed invoice", i, ":", amount)
```

নেতিবাচক বা শূন্য না হলে, amount-টা `total`-এ যোগ হয় এবং "processed invoice" message print হয়।

### Line 19

```go
}
```

Closing brace — `for` loop শেষ।

### Line 21

```go
fmt.Println("Total processed:", total)
```

সব end-এ মোট processed পরিমাণ print করে।

**Flow (invoiceAmounts = [1200, 0, 850, -50, 3000, 400]):**

1. `i=0, amount=1200` → 0 না, negative না → total=1200, print "processed invoice 0 : 1200"
2. `i=1, amount=0` → `continue` → skip
3. `i=2, amount=850` → total=2050, print "processed invoice 2 : 850"
4. `i=3, amount=-50` → negative → print "invalid invoice at index 3 - stopping batch" → `break`
5. বাকিগুলো (3000, 400) আর process হয় না
6. Final: `Total processed: 2050`

---

## Expected Output

```
processed invoice 0 : 1200
processed invoice 2 : 850
invalid invoice at index 3 - stopping batch
Total processed: 2050
```

## মূল শিক্ষা / Key Takeaways

1. **`for ... range`** — slice-এর index (`i`) আর element (`amount`) এ দুটো accessibility।
2. **`continue`** — বর্তমান iteration skip করে পরেরটায় যায় (এখানে শূন্য invoice skip)।
3. **`break`** — পুরো loop ছেড়ে বেরিয়ে যায় (এখানে invalid invoice-এ batch stop)।
4. **`float64` slice** — decimal পরিমাণের dynamic list।
5. **Accumulator** — `total += amount` দিয়ে যোগফল জমা।

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

### Line 6

```go
invoiceAmounts := []float64{1200, 0, 850, -50, 3000, 400}
```

`invoiceAmounts` — a `[]float64` slice holding six invoice amounts: `1200, 0, 850, -50, 3000, 400`. Note it contains a `0` (valid but zero) and a `-50` (invalid/negative) — these get special handling during processing.

### Line 7

```go
var total float64
```

`var total float64` — an accumulator variable that will sum up the processed invoices. It starts at `0` (the float zero value).

### Line 9

```go
for i, amount := range invoiceAmounts {
```

The `for ... range` loop goes through each element of the slice. On each iteration, `i` is the index and `amount` is the amount.

### Lines 10–12

```go
if amount == 0 {
    continue
}
```

If `amount == 0`, **`continue`** skips the rest of the loop body and moves to the next element. So zero-amount invoices are skipped (not added to the total, not printed).

### Lines 13–16

```go
if amount < 0 {
    fmt.Println("invalid invoice at index", i, "- stopping batch")
    break
}
```

If `amount < 0` (negative/invalid), print a message and **`break`** to exit the entire loop. Meaning: on encountering a negative invoice, batch processing **stops**.

### Lines 17–18

```go
total += amount
fmt.Println("processed invoice", i, ":", amount)
```

If the amount is neither zero nor negative, it's added to `total` and a "processed invoice" message is printed.

### Line 19

```go
}
```

Closing brace — ends the `for` loop.

### Line 21

```go
fmt.Println("Total processed:", total)
```

At the end, prints the total processed amount.

**Flow (invoiceAmounts = [1200, 0, 850, -50, 3000, 400]):**

1. `i=0, amount=1200` → not 0, not negative → total=1200, print "processed invoice 0 : 1200"
2. `i=1, amount=0` → `continue` → skip
3. `i=2, amount=850` → total=2050, print "processed invoice 2 : 850"
4. `i=3, amount=-50` → negative → print "invalid invoice at index 3 - stopping batch" → `break`
5. The rest (3000, 400) are not processed
6. Final: `Total processed: 2050`

---

## Expected Output

```
processed invoice 0 : 1200
processed invoice 2 : 850
invalid invoice at index 3 - stopping batch
Total processed: 2050
```

## Key Takeaways

1. **`for ... range`** — access both the index (`i`) and element (`amount`) of a slice.
2. **`continue`** — skips the current iteration and moves to the next (here, skipping zero invoices).
3. **`break`** — exits the whole loop (here, stopping the batch on an invalid invoice).
4. **`float64` slice** — a dynamic list of decimal amounts.
5. **Accumulator** — accumulates a sum with `total += amount`.
