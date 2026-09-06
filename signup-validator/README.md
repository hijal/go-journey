# signup-validator

Go-তে **function type + variadic validators** শেখার ছোট example -- reusable validation functions chain করে user input check করা.

**ভাষা নির্বাচন করুন / Choose language:**

[বাংলা](#bangla) * [English](#english)

---

<a name="bangla"></a>

## বাংলা সংস্করণ

### Lines 1-7

```go
package main

import (
    "errors"
    "fmt"
    "strings"
)
```

- `package main` -- executable program.
- `errors`, `fmt`, `strings` -- error, formatting, string-utility packages.

### Line 9

```go
type Validator func(string) error
```

Custom function type `Validator` -- signature `func(string) error`.

### Lines 11-16

```go
func notEmpty(value string) error {
    if strings.TrimSpace(value) == "" {
        return errors.New("value is empty")
    }
    return nil
}
```

প্রথম validator -- value empty কিনা (spaces সহ) check করে.

### Lines 18-25

```go
func minLength(n int) Validator {
    return func(value string) error {
        if len(value) < n {
            return fmt.Errorf("need at least %d characters", n)
        }
        return nil
    }
}
```

Validator factory -- `n` capture করে একটা custom validator return করে.
এখানে closure pattern ব্যবহৃত -- `n` captured parameter.

### Lines 27-36

```go
func runValidators(field string, value string, checks ...Validator) []error {
    var errs []error

    for _, check := range checks {
        if err := check(value); err != nil {
            errs = append(errs, fmt.Errorf("%s: %w", field, err))
        }
    }
    return errs
}
```

Variadic validators একসাথে run করে সব error collect করে:

- `checks ...Validator` -- variable number of validators.
- `%w` wrap করে field name সহ error decorate করা হয়.
- return type `[]error` -- সব error list.

### Lines 38-52

```go
func main() {
    password := "hi"
    errs := runValidators("password", password, notEmpty, minLength(8))

    if len(errs) == 0 {
        fmt.Println("password accepted")
    } else {
        for _, err := range errs {
            fmt.Println("rejected:", err)
        }
    }

    var check Validator = minLength(4)
    fmt.Println("username check:", check("raf"))
}
```

- `password := "hi"` -- "hi" 2 chars, 8 chars minimum পূরণ করে না.
- দুটো validator run: `notEmpty`, `minLength(8)`.
- `"hi"` not empty, কিন্তু 8 chars-এর কম -> 1টা error.
- `minLength(4)` factory ব্যবহার করে `check("raf")` run করলে 4 chars minimum পূরণ হয় না -> error.

---

## Expected Output

```
rejected: password: need at least 8 characters
username check: need at least 4 characters
```

## মূল শিক্ষা / Key Takeaways

1. **Function type alias** -- `type Validator func(string) error` reusable type তৈরি করে.
2. **Higher-order function** -- function যা function return করে (`minLength`).
3. **Variadic validators** -- any number of checks chain করা যায়.
4. **Error wrapping with `%w`** -- context যোগ করা, `errors.Is/Unwrap` friendly.
5. **Composition** -- validators আলাদা লিখে একসাথে apply করা.

---

---

<a name="english"></a>

## English Version

### Lines 1-7

```go
package main

import (
    "errors"
    "fmt"
    "strings"
)
```

- `package main` -- an executable program.
- `errors`, `fmt`, `strings` -- error, formatting, and string-utility packages.

### Line 9

```go
type Validator func(string) error
```

Custom function type `Validator` -- signature `func(string) error`.

### Lines 11-16

```go
func notEmpty(value string) error {
    if strings.TrimSpace(value) == "" {
        return errors.New("value is empty")
    }
    return nil
}
```

First validator -- checks if the value is empty (including whitespace-only).

### Lines 18-25

```go
func minLength(n int) Validator {
    return func(value string) error {
        if len(value) < n {
            return fmt.Errorf("need at least %d characters", n)
        }
        return nil
    }
}
```

Validator factory -- captures `n` and returns a custom validator.
This is a closure pattern: `n` is captured from the enclosing scope.

### Lines 27-36

```go
func runValidators(field string, value string, checks ...Validator) []error {
    var errs []error

    for _, check := range checks {
        if err := check(value); err != nil {
            errs = append(errs, fmt.Errorf("%s: %w", field, err))
        }
    }
    return errs
}
```

Runs a variadic list of validators and collects every error:

- `checks ...Validator` -- variable number of validators.
- `%w` wraps the original error and decorates it with the field name.
- return type `[]error` -- all errors collected.

### Lines 38-52

```go
func main() {
    password := "hi"
    errs := runValidators("password", password, notEmpty, minLength(8))

    if len(errs) == 0 {
        fmt.Println("password accepted")
    } else {
        for _, err := range errs {
            fmt.Println("rejected:", err)
        }
    }

    var check Validator = minLength(4)
    fmt.Println("username check:", check("raf"))
}
```

- `password := "hi"` -- "hi" is 2 chars, less than the required 8.
- two validators run: `notEmpty`, `minLength(8)`.
- `"hi"` is not empty, but shorter than 8 -> 1 error.
- `minLength(4)` factory used as `check("raf")` -- "raf" has 3 chars, less than 4 -> error.

---

## Expected Output

```
rejected: password: need at least 8 characters
username check: need at least 4 characters
```

## Key Takeaways

1. **Function type alias** -- `type Validator func(string) error` creates a reusable type.
2. **Higher-order function** -- a function that returns a function (`minLength`).
3. **Variadic validators** -- chain any number of checks together.
4. **Error wrapping with `%w`** -- adds context, friendly to `errors.Is/Unwrap`.
5. **Composition** -- validators are written separately and applied together.
