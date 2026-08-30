# money-conversion-testing

Go-তে **table-driven unit testing** শেখার ছোট example — `math.Round` দিয়ে BDT→paisa conversion-এ float error সামলানো এবং `_test.go` test ফাইল। এখানে কোন `main` function নেই; শুধু `go test` দিয়ে চলে।

**চালানো:** `go test -v` (এই folder-এ)

**📖 ভাষা নির্বাচন করুন / Choose language:**

[🇧🇩 বাংলা](#bangla) • [🇬🇧 English](#english)

---

<a name="bangla"></a>

## 🇧🇩 বাংলা সংস্করণ

### `money.go`

#### Line 1

```go
package main
```

Executable program নয় (কোনো `main` function নেই), তবুও `package main` — test file-ও একই package-এ থাকে, যাতে একসাথে compile হয়।

#### Line 3

```go
import "math"
```

`math` package — নিচে `math.Round` ব্যবহারের জন্য।

#### Lines 5–7

```go
func BDTToPaisa(bdt float64) int64 {
	return int64(math.Round(bdt * 100))
}
```

`BDTToPaisa` — function যা টাকার amount-কে (`float64` হিসেবে, যেমন `1499.5`) paisa-তে (`int64`, যেমন `149950`) convert করে:

- `bdt * 100` — টাকাকে paisa-তে: `1499.5 × 100 = 149950`। কিন্তু float-এ `1499.5 × 100` মাঝে মাঝে `149949.99999...`-এর মতো আসতে পারে।
- `math.Round(...)` — নিকটতম পূর্ণসংখ্যায় **round** করে। এটা float-এর মাঝেমধ্যে 0.000...1 drift-টা সামলায়, `149950`-তে ঠিকঠাক আনে।
- `int64(...)` — round-করা float-কে `int64`-তে convert করে (fraction নাই এখন)।

> **কেন Round দরকারি:** `float64`-এ টাকার হিসাব exact নয় (যেমন `0.1+0.2` != `0.3`)। paisa-তে convert করার সময় truncate না করে round করলে binary float-এর সামান্য ত্রুটি হাতে যায় না।

### `money_test.go`

#### Line 1

```go
package main
```

Test file-টা `money.go`-র সাথে একই package-এ, যাতে unexported-ও test করা যায় এবং একসাথে compile হয়।

#### Line 3

```go
import "testing"
```

Go-র builtin testing framework।

#### Line 5

```go
func TestBDTToPaisa(t *testing.T) {
```

**Test function** — নাম `Test` দিয়ে শুরু, argument `*testing.T`, `go test` চালায়।

#### Lines 6–17

```go
tests := []struct {
	name string
	bdt  float64
	want int64
}{
	{"whole number", 100.0, 10000},
	{"two decimal places", 1499.50, 149950},
	{"rounds up", 19.995, 2000},
	{"rounds down", 19.994, 1999},
	{"zero", 0.0, 0},
	{"negative (refund)", -50.25, -5025},
}
```

**Table** — anonymous struct slice, প্রতিটা case-এ ৩টা field:

- `name` — case-এর নাম (diagnostic)।
- `bdt` — input: টাকার পরিমাণ।
- `want` — expected paisa value।

6টা case cover করে:

| case              | input (`bdt`) | `want`   | মন্তব্য                                       |
| ----------------- | ------------- | -------- | --------------------------------------------- |
| whole number      | `100.0`       | `10000`  | সরল                                           |
| two decimals      | `1499.50`     | `149950` | normal case                                   |
| rounds up         | `19.995`      | `2000`   | float drift থাকলেও round-এ `19.995` → `20.00` |
| rounds down       | `19.994`      | `1999`   | নিচে round                                    |
| zero              | `0.0`         | `0`      | edge                                          |
| negative (refund) | `-50.25`      | `-5025`  | refund-এর জন্য negative                       |

#### Lines 19–27

```go
for _, tt := range tests {
	t.Run(tt.name, func(t *testing.T) {
		got := BDTToPaisa(tt.bdt)

		if got != tt.want {
			t.Errorf("BDTToPaisa(%v) = %d; want %d", tt.bdt, got, tt.want)
		}
	})
}
```

- `for _, tt := range tests` — প্রতিটা case-এ loop।
- `t.Run(tt.name, ...)` — প্রতিটা case আলাদা **subtest**; `go test -v`-এ "BDTToPaisa/rounds_up"-এর মতো দেখায়।
- `got := BDTToPaisa(tt.bdt)` — আসল function call।
- `if got != tt.want` — mismatch হলে `t.Errorf` diagnostic message।

---

## Expected Output

`go test -v`:

```
=== RUN   TestBDTToPaisa
=== RUN   TestBDTToPaisa/whole_number
=== RUN   TestBDTToPaisa/two_decimal_places
=== RUN   TestBDTToPaisa/rounds_up
=== RUN   TestBDTToPaisa/rounds_down
=== RUN   TestBDTToPaisa/zero
=== RUN   TestBDTToPaisa/negative_(refund)
--- PASS: TestBDTToPaisa (0.00s)
    --- PASS: TestBDTToPaisa/whole_number (0.00s)
    --- PASS: TestBDTToPaisa/two_decimal_places (0.00s)
    --- PASS: TestBDTToPaisa/rounds_up (0.00s)
    --- PASS: TestBDTToPaisa/rounds_down (0.00s)
    --- PASS: TestBDTToPaisa/zero (0.00s)
    --- PASS: TestBDTToPaisa/negative_(refund) (0.00s)
PASS
ok  	go-journey/money-conversion-testing	0.002s
```

## মূল শিক্ষা / Key Takeaways

1. **Table-driven testing** — struct slice-এ অনেক input/expected case-গুলো সংগঠিত করা।
2. **`math.Round`** — float→int convert করার সময় drift সামলে নিকটতম পূর্ণসংখ্যা।
3. **Subtests** — `t.Run(name, ...)` থাকে প্রতিটা case named test হিসেবে।
4. **`t.Errorf`** — fail হলে input, got, want report করা।
5. **Money in cents/paisa** — float-এ exact না; integer-unit-এ convert করে store করা।

---

---

<a name="english"></a>

## 🇬🇧 English Version

### `money.go`

#### Line 1

```go
package main
```

Not an executable (there's no `main` function), but still `package main` — so the test file lives in the same package.

#### Line 3

```go
import "math"
```

For `math.Round` below.

#### Lines 5–7

```go
func BDTToPaisa(bdt float64) int64 {
	return int64(math.Round(bdt * 100))
}
```

`BDTToPaisa` — converts a taka amount (as `float64`, e.g. `1499.5`) into paisa (an `int64`, e.g. `149950`):

- `bdt * 100` — taka to paisa: `1499.5 × 100 = 149950`. But in float, `1499.5 × 100` can come out like `149949.99999...`.
- `math.Round(...)` — **rounds** to the nearest whole number. This absorbs the occasional 0.000...1 float drift and lands on `149950`.
- `int64(...)` — converts the rounded float to `int64` (no fraction left).

> **Why Round matters:** money math in `float64` isn't exact (e.g. `0.1+0.2` != `0.3`). When converting to paisa, rounding instead of truncating avoids shipping the binary float's tiny errors.

### `money_test.go`

#### Line 1

```go
package main
```

The test file is in the same package as `money.go`, so they compile together.

#### Line 3

```go
import "testing"
```

Go's built-in testing framework.

#### Line 5

```go
func TestBDTToPaisa(t *testing.T) {
```

**Test function** — starts with `Test`, takes `*testing.T`, run by `go test`.

#### Lines 6–17

```go
tests := []struct {
	name string
	bdt  float64
	want int64
}{
	{"whole number", 100.0, 10000},
	{"two decimal places", 1499.50, 149950},
	{"rounds up", 19.995, 2000},
	{"rounds down", 19.994, 1999},
	{"zero", 0.0, 0},
	{"negative (refund)", -50.25, -5025},
}
```

**The table** — an anonymous struct slice; each case has 3 fields:

- `name` — a label for the case (diagnostics).
- `bdt` — the input amount in taka.
- `want` — the expected paisa value.

Six cases cover:

| case              | input (`bdt`) | `want`   | comment                                               |
| ----------------- | ------------- | -------- | ----------------------------------------------------- |
| whole number      | `100.0`       | `10000`  | simple                                                |
| two decimals      | `1499.50`     | `149950` | the normal case                                       |
| rounds up         | `19.995`      | `2000`   | even with float drift, round sends `19.995` → `20.00` |
| rounds down       | `19.994`      | `1999`   | rounds down                                           |
| zero              | `0.0`         | `0`      | edge case                                             |
| negative (refund) | `-50.25`      | `-5025`  | negative for refunds                                  |

#### Lines 19–27

```go
for _, tt := range tests {
	t.Run(tt.name, func(t *testing.T) {
		got := BDTToPaisa(tt.bdt)

		if got != tt.want {
			t.Errorf("BDTToPaisa(%v) = %d; want %d", tt.bdt, got, tt.want)
		}
	})
}
```

- `for _, tt := range tests` — loops over every case.
- `t.Run(tt.name, ...)` — each case runs as its own **subtest**; `go test -v` shows names like "BDTToPaisa/rounds_up".
- `got := BDTToPaisa(tt.bdt)` — the actual function call.
- `if got != tt.want` — a mismatch triggers `t.Errorf` with a diagnostic message.

---

## Expected Output

`go test -v`:

```
=== RUN   TestBDTToPaisa
=== RUN   TestBDTToPaisa/whole_number
=== RUN   TestBDTToPaisa/two_decimal_places
=== RUN   TestBDTToPaisa/rounds_up
=== RUN   TestBDTToPaisa/rounds_down
=== RUN   TestBDTToPaisa/zero
=== RUN   TestBDTToPaisa/negative_(refund)
--- PASS: TestBDTToPaisa (0.00s)
    --- PASS: TestBDTToPaisa/whole_number (0.00s)
    --- PASS: TestBDTToPaisa/two_decimal_places (0.00s)
    --- PASS: TestBDTToPaisa/rounds_up (0.00s)
    --- PASS: TestBDTToPaisa/rounds_down (0.00s)
    --- PASS: TestBDTToPaisa/zero (0.00s)
    --- PASS: TestBDTToPaisa/negative_(refund) (0.00s)
PASS
ok  	go-journey/money-conversion-testing	0.002s
```

## Key Takeaways

1. **Table-driven testing** — organize many input/expected cases in a struct slice.
2. **`math.Round`** — nearest-integer rounding absorbs float drift when converting to int.
3. **Subtests** — `t.Run(name, ...)` makes each case a named test.
4. **`t.Errorf`** — reports input, got, and want on failure.
5. **Money in cents/paisa** — floats aren't exact; convert and store in integer units.
