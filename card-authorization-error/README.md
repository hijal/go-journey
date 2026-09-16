# card-authorization-error

Go-তে **custom error type + `Error()` method + `errors.As` typed classification** শেখার ছোট example — card decline handling।

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
	"errors"
	"fmt"
)
```

- `errors` — `As`।
- `fmt` — `Println`, `Sprintf`।

### Lines 8–15

```go
type DeclineError struct {
	Code string
	Msg  string
}

func (e *DeclineError) Error() string {
	return fmt.Sprintf("declined (code %s): %s", e.Code, e.Msg)
}
```

**Custom error type** — `Error() string` method দিয়ে `error` interface satisfied। Code (`"51"`) + message।

*(`*DeclineError` pointer receiver — type assertion-এ `*DeclineError` ম্যাচ।)*

### Lines 17–22

```go
func authorize(amountCents int64) error {
	if amountCents > 500_000 {
		return &DeclineError{Code: "51", Msg: "insufficient funds"}
	}
	return nil
}
```

**Rule** — limit-এর বেশি → decline error (`500_000` — numeric underscore readability)। নাহলে `nil`।

### Lines 24–35

```go
	if err := authorize(100_000); err != nil {
		fmt.Println("unexpected:", err)
	} else {
		fmt.Println("approved, err =", err)
	}
```

**First check** — 100k approved: error nil → success path (`err = <nil>` print)।

```go
	err := authorize(750_000)
	if err == nil {
		fmt.Println("unexpected approval")
		return
	}
```

**Guard** — 750k error ঠিকই আসবে; nil হলে bug।

### Lines 37–43

```go
	var decline *DeclineError

	if errors.As(err, &decline) {
		fmt.Println("retry with another card, code:", decline.Code)
	} else {
		fmt.Println("generic failure:", err)
	}
```

**`errors.As`** — `err` chain-এ কোনো `*DeclineError` আছে কিনা দেখে; মিললে pointer-এ ঢোকে → decline-specific handling (retry with another card)।

*(`As` vs `Is`: `Is` exact/sentinel value; `As` typed value-কে extract করে।)*

---

## Expected Output

```
approved, err = <nil>
retry with another card, code: 51
```

## মূল শিক্ষা / Key Takeaways

1. **Custom error type** — `Error()` method।
2. **`errors.As`** — typed classification (decline vs generic)।
3. **Numeric underscores** — `500_000` readability।
4. **nil-check branching** — approved → nil path।
5. **Typed action** — code-based retry logic।

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
	"errors"
	"fmt"
)
```

- `errors` — for `As`.
- `fmt` — for `Println`, `Sprintf`.

### Lines 8–15

```go
type DeclineError struct {
	Code string
	Msg  string
}

func (e *DeclineError) Error() string {
	return fmt.Sprintf("declined (code %s): %s", e.Code, e.Msg)
}
```

**A custom error type** — the `Error() string` method satisfies the `error` interface. A code (`"51"`) + message.

*(The `*DeclineError` pointer receiver — matches `*DeclineError` assertions.)*

### Lines 17–22

```go
func authorize(amountCents int64) error {
	if amountCents > 500_000 {
		return &DeclineError{Code: "51", Msg: "insufficient funds"}
	}
	return nil
}
```

**The rule** — over the limit returns a decline error (`500_000` — the numeric-underscore readability). Otherwise `nil`.

### Lines 24–35

```go
	if err := authorize(100_000); err != nil {
		fmt.Println("unexpected:", err)
	} else {
		fmt.Println("approved, err =", err)
	}
```

**First check** — 100k approves: nil error → the success path (`err = <nil>` printed).

```go
	err := authorize(750_000)
	if err == nil {
		fmt.Println("unexpected approval")
		return
	}
```

**Guard** — 750k must error; nil would be a bug.

### Lines 37–43

```go
	var decline *DeclineError

	if errors.As(err, &decline) {
		fmt.Println("retry with another card, code:", decline.Code)
	} else {
		fmt.Println("generic failure:", err)
	}
```

**`errors.As`** — checks the `err` chain for a `*DeclineError`, extracts it into the pointer → decline-specific handling (retry with another card).

*(`As` vs `Is`: `Is` matches an exact/sentinel value; `As` extracts a typed value.)*

---

## Expected Output

```
approved, err = <nil>
retry with another card, code: 51
```

## Key Takeaways

1. **Custom error type** — the `Error()` method.
2. **`errors.As`** — typed classification (decline vs generic).
3. **Numeric underscores** — `500_000` readability.
4. **nil-check branching** — the approved nil path.
5. **Typed action** — code-based retry logic.