# multi-error-aggregation

Go-তে **custom error type implementing `Error()` interface**, **nil-error filtering** আর multiple error combine শেখার ছোট example।

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

### Lines 3–7

```go
import (
	"errors"
	"fmt"
	"strings"
)
```

- `errors` — `errors.New`।
- `fmt` — `Println`।
- `strings` — `Join`।

### Lines 9–11

```go
type MultiError struct {
	Errors []error
}
```

`MultiError` — একটা struct যা একাধিক `error` ধরে (একটা `[]error` field)।

### Lines 13–21

```go
func (m *MultiError) Error() string {
	msgs := make([]string, len(m.Errors))
	for i, err := range m.Errors {
		msgs[i] = err.Error()
	}
	return strings.Join(msgs, ";")
}
```

**`Error() string` method** — receiver pointer-এ (`*MultiError`)। এই method-টা থাকলে `MultiError` **`error` interface implement করে** — একজন valid error হয়ে যায়। Implementation:

- `msgs := make([]string, len(m.Errors))` — same-length string slice।
- Loop-এ প্রতিটা error-র message-টা (`err.Error()`) extract।
- `strings.Join(msgs, ";")` — `";"` দিয়ে join — সব error-র message এক string-এ।

**Why `;`:** আলাদা error-গুলো এক string-এ readable-এর জন্য semicolon separator।

### Lines 23–35

```go
func CombineErrors(errs ...error) error {
	var mErr MultiError
	for _, err := range errs {
		if err != nil {
			mErr.Errors = append(mErr.Errors, err)
		}
	}
	if len(mErr.Errors) == 0 {
		return nil
	}
	return &mErr
}
```

`CombineErrors` — variadic errors-কে সংগ্রহ:

- `errs ...error` — **variadic**: multiple error arguments।
- Loop-এ প্রতিটা `err`:
  - `if err != nil` — **nil-filter**: শুধু non-nil error-গুলো সংগ্রহ করা হয় (nil = "no error", সংগ্রহ করা ইতিবাচক)। nil-কে append করলে `MultiError`-তে useless nil-entry থাকত।
- `if len(mErr.Errors) == 0` — কোনো error না হলে **`nil` return** (কোনো error নেই — সফল)। এটা important: খালি `MultiError` return না করে `nil`।
- শেষে `return &mErr` — pointer (address) return — যেহেতু `Error()`-এর receiver-pointer, তাই `*MultiError`-কে error হিসেবে pass করা correct (value-র receiver-pointer method-টা ঠিকঠাক satisfy করে)।

### Lines 37–43

```go
func main() {
	err := CombineErrors(nil, errors.New("DB timeout"), errors.New("invalid input"))

	if err != nil {
		fmt.Println(err)
	}
}
```

- `CombineErrors(nil, errors.New("DB timeout"), errors.New("invalid input"))` — ৩টা argument: একটা nil, ২টা real error। nil filter হবে, বাকি ২টা যাবে।
- `if err != nil` — combined error নয় → print।
- `fmt.Println(err)` — `MultiError.Error()` call হয় (ইন্টারফেস dispatch মাধ্যমে): `"DB timeout;invalid input"`।

### Line 43

```go
}
```

Closing brace — `main` function শেষ হয়।

---

## Expected Output

```
DB timeout;invalid input
```

## মূল শিক্ষা / Key Takeaways

1. **Custom error type** — struct + `Error() string` method = `error` interface implement।
2. **`Error()` method** — error-র readable representation।
3. **Nil-error filtering** — `if err != nil` দিয়ে শুধু real errors সংগ্রহ।
4. **`nil` on empty** — কোনো error নেই = nil (খালি container নয়)।
5. **Pointer receiver** — `*MultiError` — method-টা pointer-এ।

---

---

<a name="english"></a>

## 🇬🇧 English Version

### Line 1

```go
package main
```

Declares an executable program (`main` package), runnable via `go run`.

### Lines 3–7

```go
import (
	"errors"
	"fmt"
	"strings"
)
```

- `errors` — for `errors.New`.
- `fmt` — for `Println`.
- `strings` — for `Join`.

### Lines 9–11

```go
type MultiError struct {
	Errors []error
}
```

`MultiError` — a struct holding multiple errors (a `[]error` field).

### Lines 13–21

```go
func (m *MultiError) Error() string {
	msgs := make([]string, len(m.Errors))
	for i, err := range m.Errors {
		msgs[i] = err.Error()
	}
	return strings.Join(msgs, ";")
}
```

**`Error() string` method** — with a pointer receiver (`*MultiError`). Having this method makes `MultiError` implement the **`error` interface** — it becomes a valid error. Implementation:

- `msgs := make([]string, len(m.Errors))` — a string slice of the same length.
- The loop extracts each error's message (`err.Error()`).
- `strings.Join(msgs, ";")` — joins with `";"` — all messages in one string.

**Why `;`:** a semicolon separator keeps the distinct errors readable in a single string.

### Lines 23–35

```go
func CombineErrors(errs ...error) error {
	var mErr MultiError
	for _, err := range errs {
		if err != nil {
			mErr.Errors = append(mErr.Errors, err)
		}
	}
	if len(mErr.Errors) == 0 {
		return nil
	}
	return &mErr
}
```

`CombineErrors` — collects variadic errors:

- `errs ...error` — **variadic**: multiple error arguments.
- The loop over each `err`:
  - `if err != nil` — **nil-filter**: only non-nil errors are collected (a nil error means "no error", so it's not worth collecting). Appending nil would put a useless nil entry in `MultiError`.
- `if len(mErr.Errors) == 0` — if there are no errors, return **`nil`** (no error — success). This matters: return `nil` rather than an empty `MultiError`.
- Finally `return &mErr` — returns the pointer (address). Since `Error()` has a pointer receiver, returning `*MultiError` makes it a proper error value (the method set of `*MultiError` includes `Error`).

### Lines 37–43

```go
func main() {
	err := CombineErrors(nil, errors.New("DB timeout"), errors.New("invalid input"))

	if err != nil {
		fmt.Println(err)
	}
}
```

- `CombineErrors(nil, errors.New("DB timeout"), errors.New("invalid input"))` — 3 arguments: one nil, two real errors. The nil is filtered, the other two go in.
- `if err != nil` — the error isn't nil → print.
- `fmt.Println(err)` — calls `MultiError.Error()` (through interface dispatch): `"DB timeout;invalid input"`.

### Line 43

```go
}
```

Closing brace — ends the `main` function.

---

## Expected Output

```
DB timeout;invalid input
```

## Key Takeaways

1. **Custom error type** — a struct with an `Error() string` method implements the `error` interface.
2. **`Error()` method** — readable representation of the error.
3. **Nil-error filtering** — `if err != nil` to gather only real errors.
4. **`nil` on empty** — no errors means `nil` (not an empty container).
5. **Pointer receiver** — `*MultiError` — the method lives on the pointer.