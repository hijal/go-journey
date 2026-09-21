# error-http-status-mapping

Go-তে **error → HTTP status mapping (`errors.Is` + `errors.As` + custom error types)** শেখার ছোট example — account service error handling।

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
	"net/http"
)
```

- `errors` — `Is`, `As`, `New`।
- `net/http` — status codes।
- `fmt` — `Errorf`, `Sprintf`, `Printf`।

### Lines 9–29

```go
type NotFoundError struct {
	Resource string
	ID       string
}

func (e *NotFoundError) Error() string {
	return fmt.Sprintf("%s %q not found", e.Resource, e.ID)
}

type ValidationError struct {
	Field string
	Err   error
}

func (e *ValidationError) Error() string {
	return e.Field + ": " + e.Err.Error()
}

func (e *ValidationError) Unwrap() error {
	return e.Err
}
```

**Custom error types**:

- `NotFoundError` — structured fields (Resource + ID) + `Error()`।
- `ValidationError` — wraps inner error + **`Unwrap()`** (cause-chain-এ অংশ)।

`Unwrap` মানেই `errors.Is`/`As` **এক layer ভেদ করতে পারে**।

### Lines 31–34

```go
var (
	ErrUnauthorized = errors.New("unauthorized")
	ErrNegative     = errors.New("must not be negative")
)
```

Sentinel errors।

### Lines 36–52

```go
func GetAccount(id, token string) error {
	if token == "" {
		return ErrUnauthorized
	}

	if id != "acc-1" {
		return fmt.Errorf("account service: %w", &NotFoundError{Resource: "account", ID: id})
	}
	return nil
}

func Withdraw(amount int64) error {
	if amount < 0 {
		return fmt.Errorf("withdraw: %w", &ValidationError{Field: "amount", Err: ErrNegative})
	}
	return nil
}
```

`%w` দিয়ে custom error wrap — outer context + cause-chain intact।

### Lines 54–70

```go
func StatusFor(err error) int {
	var nf *NotFoundError
	var ve *ValidationError

	switch {
	case err == nil:
		return http.StatusOK
	case errors.Is(err, ErrUnauthorized):
		return http.StatusUnauthorized
	case errors.As(err, &nf):
		return http.StatusNotFound
	case errors.As(err, &ve):
		return http.StatusUnprocessableEntity
	default:
		return http.StatusInternalServerError
	}
}
```

**Mapping switch** (order matters):

- nil → `200`।
- `errors.Is` sentinel → `401`।
- `errors.As` → `404` (typed error পাওয়া গেলে)।
- `errors.As` ValidationError → `422`।
- otherwise → `500`।

### Lines 72–89

```go
	cases := []error{
		GetAccount("acc-1", "tok"),
		GetAccount("acc-1", ""),
		GetAccount("acc-9", "tok"),
		Withdraw(-50),
		errors.New("disk full"),
	}

	for _, err := range cases {
		fmt.Printf("%d %v\n", StatusFor(err), err)
	}

	var nf *NotFoundError
	if err := GetAccount("acc-9", "tok"); errors.As(err, &nf) {
		fmt.Println("missing resource:", nf.Resource, nf.ID)
	}
	fmt.Println("is negative:", errors.Is(Withdraw(-1), ErrNegative))
}
```

- ৫টা representative error → mapped status।
- `errors.As` — **typed field extract** (Resource, ID)।
- `errors.Is` — `Unwrap` chain ভেদ করে sentinel মেলে।

---

## Expected Output

```
200 <nil>
401 unauthorized
404 account service: account "acc-9" not found
422 withdraw: amount: must not be negative
500 disk full
missing resource: account acc-9
is negative: true
```

## মূল শিক্ষা / Key Takeaways

1. **`errors.Is`** — sentinel identity (wrappers ভেদ করে)।
2. **`errors.As`** — typed errors-কে extract + fields-এ access।
3. **Custom error + `Unwrap()`** — cause-chain অংশ।
4. **`%w` wrap** — outer context + inner cause।
5. **Mapping-এ order** — নির্দিষ্ট → generic (Is → As → default)।

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
	"net/http"
)
```

- `errors` — `Is`, `As`, `New`.
- `net/http` — status codes.
- `fmt` — `Errorf`, `Sprintf`, `Printf`.

### Lines 9–29

```go
type NotFoundError struct {
	Resource string
	ID       string
}

func (e *NotFoundError) Error() string {
	return fmt.Sprintf("%s %q not found", e.Resource, e.ID)
}

type ValidationError struct {
	Field string
	Err   error
}

func (e *ValidationError) Error() string {
	return e.Field + ": " + e.Err.Error()
}

func (e *ValidationError) Unwrap() error {
	return e.Err
}
```

**Custom error types**:

- `NotFoundError` — structured fields (Resource + ID) + `Error()`.
- `ValidationError` — wraps an inner error + **`Unwrap()`** (part of the cause chain).

`Unwrap` is what lets `errors.Is`/`As` peel through one layer.

### Lines 31–34

```go
var (
	ErrUnauthorized = errors.New("unauthorized")
	ErrNegative     = errors.New("must not be negative")
)
```

Sentinel errors.

### Lines 36–52

```go
func GetAccount(id, token string) error {
	if token == "" {
		return ErrUnauthorized
	}

	if id != "acc-1" {
		return fmt.Errorf("account service: %w", &NotFoundError{Resource: "account", ID: id})
	}
	return nil
}

func Withdraw(amount int64) error {
	if amount < 0 {
		return fmt.Errorf("withdraw: %w", &ValidationError{Field: "amount", Err: ErrNegative})
	}
	return nil
}
```

`%w` wraps the custom errors — outer context plus an intact cause chain.

### Lines 54–70

```go
func StatusFor(err error) int {
	var nf *NotFoundError
	var ve *ValidationError

	switch {
	case err == nil:
		return http.StatusOK
	case errors.Is(err, ErrUnauthorized):
		return http.StatusUnauthorized
	case errors.As(err, &nf):
		return http.StatusNotFound
	case errors.As(err, &ve):
		return http.StatusUnprocessableEntity
	default:
		return http.StatusInternalServerError
	}
}
```

**The mapping switch** (order matters):

- nil → `200`.
- `errors.Is` sentinel → `401`.
- `errors.As` → `404` (a typed error found).
- `errors.As` for ValidationError → `422`.
- otherwise → `500`.

### Lines 72–89

```go
	cases := []error{
		GetAccount("acc-1", "tok"),
		GetAccount("acc-1", ""),
		GetAccount("acc-9", "tok"),
		Withdraw(-50),
		errors.New("disk full"),
	}

	for _, err := range cases {
		fmt.Printf("%d %v\n", StatusFor(err), err)
	}

	var nf *NotFoundError
	if err := GetAccount("acc-9", "tok"); errors.As(err, &nf) {
		fmt.Println("missing resource:", nf.Resource, nf.ID)
	}
	fmt.Println("is negative:", errors.Is(Withdraw(-1), ErrNegative))
}
```

- 5 representative errors → mapped statuses.
- `errors.As` — **typed-field extraction** (Resource, ID).
- `errors.Is` — an `Unwrap` chain reaches the sentinel.

---

## Expected Output

```
200 <nil>
401 unauthorized
404 account service: account "acc-9" not found
422 withdraw: amount: must not be negative
500 disk full
missing resource: account acc-9
is negative: true
```

## Key Takeaways

1. **`errors.Is`** — sentinel identity (through wrappers).
2. **`errors.As`** — extract typed errors + access their fields.
3. **Custom error + `Unwrap()`** — part of the cause chain.
4. **`%w` wrapping** — outer context + inner cause.
5. **Mapping order** — specific → generic (Is → As → default).