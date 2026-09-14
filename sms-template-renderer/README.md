# sms-template-renderer

Go-তে **`strings.NewReplacer` + multi-placeholder template rendering** শেখার ছোট example — bilingual SMS template।

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

### Lines 3–6

```go
import (
	"fmt"
	"strings"
)
```

- `fmt` — `Println`।
- `strings` — `NewReplacer`।

### Lines 8–10

```go
const orderTemplate = `নমস্কার {{name}},
আপনার অর্ডার {{order_id}} গ্রহণ করা হয়েছে (মোট {{total}} টাকা)।
ট্র্যাকিং লিংক: {{tracking_url}}`
```

৪টা placeholder-সহ বাংলা টেমপ্লেট (raw-string backtick — multi-line)।

### Lines 12–20

```go
func renderOrderSMS(name, orderID, total, trackingURL string) string {
	r := strings.NewReplacer(
		"{{name}}", name,
		"{{order_id}}", orderID,
		"{{total}}", total,
		"{{tracking_url}}", trackingURL,
	)
	return r.Replace(orderTemplate)
}
```

**Multi-placeholder replace:**

- `strings.NewReplacer(old→new pairs)` — এক-বারে parallel replacement।
- `r.Replace(orderTemplate)` — প্রতিটি `{{...}}`-কে মান বসিয়ে দেয়।
- বারবার `strings.Replace` চেইন করার চেয়ে clean — এক-pass।

*(`Replacer` pre-built object; বড় input-এর জন্য optimized।)*

### Lines 22–26

```go
func main() {
	fmt.Println(renderOrderSMS("রাফি", "ORD-9012", "1450", "https://example.com/t/9012"))
	fmt.Println("---")
	fmt.Println(renderOrderSMS("মিতু", "ORD-9013", "899", "https://example.com/t/9013"))
}
```

২টা order — একই টেমপ্লেট, ভিন্ন ডেটা:

- রাফি / ORD-9012 / 1450 / t/9012
- মিতু / ORD-9013 / 899 / t/9013

---

## Expected Output

```
নমস্কার রাফি,
আপনার অর্ডার ORD-9012 গ্রহণ করা হয়েছে (মোট 1450 টাকা)।
ট্র্যাকিং লিংক: https://example.com/t/9012
---
নমস্কার মিতু,
আপনার অর্ডার ORD-9013 গ্রহণ করা হয়েছে (মোট 899 টাকা)।
ট্র্যাকিং লিংক: https://example.com/t/9013
```

## মূল শিক্ষা / Key Takeaways

1. **`fmt.Sprintf`**-এর বদলে **`strings.Replacer`** — named placeholder।
2. **Old→new pairs** — এক-pass-এ সব replace।
3. **Raw-string backtick** — multi-line template।
4. **Reusability** — এক template, ভিন্ন ডেটা।

---

---

<a name="english"></a>

## 🇬🇧 English Version

### Line 1

```go
package main
```

Declares an executable program (`main` package), runnable via `go run`.

### Lines 3–6

```go
import (
	"fmt"
	"strings"
)
```

- `fmt` — for `Println`.
- `strings` — for `NewReplacer`.

### Lines 8–10

```go
const orderTemplate = `নমস্কার {{name}},
আপনার অর্ডার {{order_id}} গ্রহণ করা হয়েছে (মোট {{total}} টাকা)।
ট্র্যাকিং লিংক: {{tracking_url}}`
```

A Bengali template with 4 placeholders (raw-string backticks — multiline).

### Lines 12–20

```go
func renderOrderSMS(name, orderID, total, trackingURL string) string {
	r := strings.NewReplacer(
		"{{name}}", name,
		"{{order_id}}", orderID,
		"{{total}}", total,
		"{{tracking_url}}", trackingURL,
	)
	return r.Replace(orderTemplate)
}
```

**Multi-placeholder replacement:**

- `strings.NewReplacer(old→new pairs)` — parallel replacement in one go.
- `r.Replace(orderTemplate)` — fills each `{{...}}` with its value.
- Cleaner than chaining multiple `strings.Replace` calls — one pass.

*(`Replacer` is a pre-built object; optimized for large inputs.)*

### Lines 22–26

```go
func main() {
	fmt.Println(renderOrderSMS("রাফি", "ORD-9012", "1450", "https://example.com/t/9012"))
	fmt.Println("---")
	fmt.Println(renderOrderSMS("মিতু", "ORD-9013", "899", "https://example.com/t/9013"))
}
```

2 orders — same template, different data:

- রাফি / ORD-9012 / 1450 / t/9012
- মিতু / ORD-9013 / 899 / t/9013

---

## Expected Output

```
নমস্কার রাফি,
আপনার অর্ডার ORD-9012 গ্রহণ করা হয়েছে (মোট 1450 টাকা)।
ট্র্যাকিং লিংক: https://example.com/t/9012
---
নমস্কার মিতু,
আপনার অর্ডার ORD-9013 গ্রহণ করা হয়েছে (মোট 899 টাকা)।
ট্র্যাকিং লিংক: https://example.com/t/9013
```

## Key Takeaways

1. **`strings.Replacer`** over `fmt.Sprintf` — named placeholders.
2. **Old→new pairs** — everything replaced in a single pass.
3. **Raw-string backticks** — multiline templates.
4. **Reusability** — one template, many data sets.