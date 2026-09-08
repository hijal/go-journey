# data-validation-pipeline

Go-তে **function type** (`type Validator func(string) error`), **variadic function parameters** আর **error collection** (pipeline pattern) শেখার ছোট example।

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

- `fmt` — `Errorf`, `Println`।
- `strings` — `ContainsAny`।

### Line 8

```go
type Validator func(string) error
```

**Function type** — `Validator` একটা type যা এই signature-র function-কে বোঝায়: একটা `string` নেয়, `error` return করে। Go-তে functions-ও প্রথম-class citizens — variable/parameter হিসেবে pass করা যায়।

### Lines 10–19

```go
func validate(data string, validators ...Validator) []error {
	var errs []error

	for _, v := range validators {
		if err := v(data); err != nil {
			errs = append(errs, err)
		}
	}
	return errs
}
```

`validate` — validation **pipeline**:

- `data string` — validate করার ডেটা।
- `validators ...Validator` — **variadic**: একাধিক `Validator` function (একটা `[]Validator` slice)।
- Loop-এ প্রতিটা validator `v(data)` call হয়; যদি error return করে, `errs`-এ **append**।
- সব `[]error` — ফেল হওয়া সব validation-এর error collect, প্রথমটা-এ stop করে না (fail-fast নয়, fail-all)।
- **Note:** `var errs []error` — nil slice। Error না হলে `nil` slice return হয় (zero value), nil slice-এ range-এ loop চলে না।

### Lines 21–34

```go
func main() {
	isLongEnough := func(s string) error {
		if len(s) < 5 {
			return fmt.Errorf("too short")
		}
		return nil
	}

	hasNumber := func(s string) error {
		if !strings.ContainsAny(s, "1234567890") {
			return fmt.Errorf("needs a number")
		}
		return nil
	}
```

দুটো **anonymous function** (নাম-হীন closure), প্রতিটা `Validator`-এর signature match:

- `isLongEnough` — string-এর length 5-এর কম হলে `fmt.Errorf("too short")`, না হলে `nil`।
- `hasNumber` — string-এ কোনো digit না থাকলে `fmt.Errorf("needs a number")`, না হলে `nil`। `strings.ContainsAny(s, "1234567890")` — s-এ ওই characters-গুলোর একটা থাকলে true।

প্রতিটাকে variable-এ assign করা হয়েছে, তারপর validator হিসেবে pass।

### Line 36

```go
errs := validate("abc", isLongEnough, hasNumber)
```

`validate("abc", ...)`:

- `isLongEnough("abc")` — len 3 < 5 → error "too short"
- `hasNumber("abc")` — কোনো digit নেই → error "needs a number"

`errs` = দুটো error-এর slice।

### Lines 38–40

```go
for _, err := range errs {
	fmt.Println(err)
}
```

দুটো error print:

- `too short`
- `needs a number`

### Line 41

```go
}
```

Closing brace — `main` function শেষ হয়।

---

## Expected Output

```
too short
needs a number
```

## মূল শিক্ষা / Key Takeaways

1. **Function type** — `type Validator func(string) error` — signature-কে একটা নাম-বিশিষ্ট type।
2. **Variadic functions** — `...Validator` — pipeline-র মতো একাধিক validator pull করা।
3. **Fail-all pattern** — সব error collect (`[]error`), প্রথমটায় stop না করে।
4. **Anonymous functions** — inline-defined function variable।
5. **`fmt.Errorf`** — formatted error message create করা।

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

- `fmt` — for `Errorf`, `Println`.
- `strings` — for `ContainsAny`.

### Line 8

```go
type Validator func(string) error
```

**Function type** — `Validator` names the type of a function that takes a `string` and returns `error`. In Go functions are first-class citizens — passable as variables/parameters.

### Lines 10–19

```go
func validate(data string, validators ...Validator) []error {
	var errs []error

	for _, v := range validators {
		if err := v(data); err != nil {
			errs = append(errs, err)
		}
	}
	return errs
}
```

`validate` — a validation **pipeline**:

- `data string` — the data to validate.
- `validators ...Validator` — **variadic**: multiple `Validator` functions (one `[]Validator` slice).
- In the loop each validator `v(data)` is called; if it returns an error, it's **appended** to `errs`.
- Returns all of `[]error` — collects errors from every failing validation rather than stopping at the first (fail-all, not fail-fast).
- **Note:** `var errs []error` — a nil slice. If no errors, a nil slice is returned (the zero value); ranging over a nil slice runs zero iterations.

### Lines 21–34

```go
func main() {
	isLongEnough := func(s string) error {
		if len(s) < 5 {
			return fmt.Errorf("too short")
		}
		return nil
	}

	hasNumber := func(s string) error {
		if !strings.ContainsAny(s, "1234567890") {
			return fmt.Errorf("needs a number")
		}
		return nil
	}
```

Two **anonymous functions**, each matching the `Validator` signature:

- `isLongEnough` — if the string's length is under 5, returns `fmt.Errorf("too short")`, else `nil`.
- `hasNumber` — if the string has no digit, returns `fmt.Errorf("needs a number")`, else `nil`. `strings.ContainsAny(s, "1234567890")` — true if `s` contains any of those characters.

Each is assigned to a variable, then passed as a validator.

### Line 36

```go
errs := validate("abc", isLongEnough, hasNumber)
```

`validate("abc", ...)`:

- `isLongEnough("abc")` — len 3 < 5 → error "too short"
- `hasNumber("abc")` — no digit → error "needs a number"

`errs` = a slice of the two errors.

### Lines 38–40

```go
for _, err := range errs {
	fmt.Println(err)
}
```

Prints both errors:

- `too short`
- `needs a number`

### Line 41

```go
}
```

Closing brace — ends the `main` function.

---

## Expected Output

```
too short
needs a number
```

## Key Takeaways

1. **Function type** — `type Validator func(string) error` — names a signature.
2. **Variadic functions** — `...Validator` — pull in multiple validators like a pipeline.
3. **Fail-all pattern** — collect all errors (`[]error`) instead of stopping at the first.
4. **Anonymous functions** — inline-defined function variables.
5. **`fmt.Errorf`** — creates a formatted error message.