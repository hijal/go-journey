# testing

Go-তে **table-driven unit testing** এবং `go/token.IsIdentifier` দিয়ে Go identifier-এর validity rule test করা শেখার ছোট example। এতে একটা `main.go` আর একটা `identifier_test.go` আছে।

**চালানো:** `go test -v` (এই folder-এ)

**📖 ভাষা নির্বাচন করুন / Choose language:**

[🇧🇩 বাংলা](#bangla) • [🇬🇧 English](#english)

---

<a name="bangla"></a>

## 🇧🇩 বাংলা সংস্করণ

### `main.go`

#### Line 1

```go
package main
```

একটা executable program (`main` package) declare করে। Test file-টাও একই `package main`-এ আছে, যাতে একসাথে compile হয়।

#### Line 3

```go
import "fmt"
```

Console-এ output print করার জন্য `fmt` package import করা হয়।

#### Lines 5–7

```go
func main() {
    fmt.Println("Run 'go test -v' in this folder to check identifier validity rules.")
}
```

`main` function-টা শুধু একটা নির্দেশনা message print করে — এটা দিয়ে user-কে বোঝানো হয় যে বাস্তব চেকটা `go test` দিয়ে চলে।

---

### `identifier_test.go`

#### Line 1

```go
package main
```

Test file-টা `main` package-এ, কারণ এটা `main.go`-র সাথে একই package-এ test করে।

#### Lines 3–6

```go
import (
    "go/token"
    "testing"
)
```

দুটো standard library package import:

- `go/token` — এটা Go language-এর lexical scanner/token package। এখানে `token.IsIdentifier` function আছে, যা check করে একটা string valid Go identifier কিনা। (Go নিজে নিজের source-এর identifier validate করতে এটা use করে।)
- `testing` — Go-র builtin testing framework। এটা দিয়ে `TestXxx` function define করা হয়।

#### Line 8

```go
func TestIsValidIdentifier(t *testing.T) {
```

একটা **test function**। নিয়ম: নাম `Test` দিয়ে শুরু, argument `*testing.T`। `go test` এই function-টা চালায় এবং report করে pass/fail।

#### Lines 9–13

```go
cases := []struct {
    name     string
    input    string
    expected bool
}{
```

এটা একটা **anonymous struct slice** — table-driven testing-এর "table"। প্রতিটা struct-এ তিনটা field:

- `name` — test case-এর নাম (diagnostic-এর জন্য)।
- `input` — যে string-টা আমরা validate করব।
- `expected` — আশা করা result (valid হলে `true`, না হলে `false`).

#### Lines 14–24

```go
{"simple lowercase", "amount", true},
{"simple exported", "Amount", true},
{"with underscore", "amount_cents", true},
{"starts with underscore", "_temp", true},
{"blank identifier", "_", true},
{"starts with digit", "1amount", false},
{"contains space", "amount cents", false},
{"contains hyphen", "amount-cents", false},
{"empty string", "", false},
{"go keyword", "func", false},
```

**Table-এর প্রতিটা row** একটা test case:

| input | expected | reason |
|-------|----------|--------|
| `"amount"` | `true` | সাধারণ lowercase identifier — valid |
| `"Amount"` | `true` | বড় হাতের দিয়ে শুরু exported identifier — valid |
| `"amount_cents"` | `true` | underscore থাকাটা legal — valid |
| `"_temp"` | `true` | underscore দিয়ে শুরু legal — valid |
| `"_"` | `true` | প্লেইন `_` হলো **blank identifier** — valid |
| `"1amount"` | `false` | digit দিয়ে শুরু **হতে পারে না** — invalid |
| `"amount cents"` | `false` | space আছে, যা identifier-এ নিষিদ্ধ — invalid |
| `"amount-cents"` | `false` | hyphen (`-`) identifier-এ অক্ষর নয় — invalid |
| `""` | `false` | খালি string identifier নয় — invalid |
| `"func"` | `false` | Go-র **keyword** identifier হিসেবে ব্যবহার করা যায় না — invalid |

#### Line 26

```go
for _, tc := range cases {
```

`range` দিয়ে table-এর প্রতিটা case-এ loop হয়। `_` index discard, `tc` হয় প্রতিটা struct।

#### Line 27

```go
t.Run(tc.name, func(t *testing.T) {
```

`t.Run(name, fn)` একটা **subtest** চালায় — প্রতিটা case আলাদা named test হিসেবে show হয় (`go test -v`-এ)। এতে fail হলে কোন case ভেঙেছে সহজে বোঝা যায়।

#### Line 28

```go
got := token.IsIdentifier(tc.input)
```

`token.IsIdentifier(tc.input)` — এই function-টা input-টা valid identifier কিনা check করে এবং result `got`-এ রাখে। এটাই আসল testing logic।

#### Lines 29–31

```go
if got != tc.expected {
    t.Errorf("IsIdentifier(%q) = %v, want %v", tc.input, got, tc.expected)
}
```

যদি `got` আর `expected` মিলে না যায়, তাহলে `t.Errorf` দিয়ে failure report করে — input, actual result (`got`), আর expected value print করে। `%q` string-কে quote সহ, `%v` value print করে। মিলে গেলে কিছুই হয় না (নীরবে pass)।

#### Line 32

```go
})
```

Subtest-এর closure শেষ।

#### Line 33

```go
}
```

`for` loop শেষ।

#### Line 34

```go
}
```

Test function শেষ।

---

## Expected Output

`go test -v`:

```
=== RUN   TestIsValidIdentifier
=== RUN   TestIsValidIdentifier/simple_lowercase
=== RUN   TestIsValidIdentifier/simple_exported
... (সব subtest)
--- PASS: TestIsValidIdentifier (0.00s)
PASS
ok  	go-journey/testing	0.810s
```

`go run .`:

```
Run 'go test -v' in this folder to check identifier validity rules.
```

## মূল শিক্ষা / Key Takeaways

1. **Table-driven testing** — struct slice দিয়ে অনেকগুলো input/expected test case-কে সংগঠিত করা।
2. **`go/token.IsIdentifier`** — string valid Go identifier কিনা দেখার standard library function।
3. **`t.Run` subtests** — প্রতিটা case আলাদা named test হিসেবে; fail হলে location স্পষ্ট।
4. **`t.Errorf`** — test fail হলে diagnostic message report করা।
5. **Test conventions** — `_test.go` file, `TestXxx` function, `*testing.T` argument, `go test` দিয়ে চালানো।

---

---

<a name="english"></a>

## 🇬🇧 English Version

### `main.go`

#### Line 1

```go
package main
```

Declares an executable program (`main` package). The test file is in the same `package main`, so they compile together.

#### Line 3

```go
import "fmt"
```

Imports the `fmt` package for console output.

#### Lines 5–7

```go
func main() {
    fmt.Println("Run 'go test -v' in this folder to check identifier validity rules.")
}
```

The `main` function just prints an instruction message — telling the user that the real checking happens via `go test`.

---

### `identifier_test.go`

#### Line 1

```go
package main
```

The test file is in the `main` package, since it tests code in the same package as `main.go`.

#### Lines 3–6

```go
import (
    "go/token"
    "testing"
)
```

Two standard library packages are imported:

- `go/token` — Go's lexical scanner/token package. It provides `token.IsIdentifier`, a function that checks whether a string is a valid Go identifier (Go itself uses this on its own source code).
- `testing` — Go's built-in testing framework, used to define `TestXxx` functions.

#### Line 8

```go
func TestIsValidIdentifier(t *testing.T) {
```

A **test function**. By convention it starts with `Test`, takes `*testing.T`, and `go test` runs it and reports pass/fail.

#### Lines 9–13

```go
cases := []struct {
    name     string
    input    string
    expected bool
}{
```

An **anonymous struct slice** — the "table" of table-driven testing. Each struct has three fields:

- `name` — a label for the test case (for diagnostics).
- `input` — the string we want to validate.
- `expected` — the expected result (`true` if valid, `false` otherwise).

#### Lines 14–24

```go
{"simple lowercase", "amount", true},
{"simple exported", "Amount", true},
{"with underscore", "amount_cents", true},
{"starts with underscore", "_temp", true},
{"blank identifier", "_", true},
{"starts with digit", "1amount", false},
{"contains space", "amount cents", false},
{"contains hyphen", "amount-cents", false},
{"empty string", "", false},
{"go keyword", "func", false},
```

**Each row of the table** is one test case:

| input | expected | reason |
|-------|----------|--------|
| `"amount"` | `true` | a normal lowercase identifier — valid |
| `"Amount"` | `true` | exported identifier starting uppercase — valid |
| `"amount_cents"` | `true` | underscores are legal — valid |
| `"_temp"` | `true` | starting with underscore is legal — valid |
| `"_"` | `true` | plain `_` is the **blank identifier** — valid |
| `"1amount"` | `false` | cannot start with a digit — invalid |
| `"amount cents"` | `false` | contains a space, which is forbidden — invalid |
| `"amount-cents"` | `false` | hyphen (`-`) is not a valid identifier character — invalid |
| `""` | `false` | an empty string is not an identifier — invalid |
| `"func"` | `false` | a Go **keyword** can't be used as an identifier — invalid |

#### Line 26

```go
for _, tc := range cases {
```

Loops over each case in the table with `range`. `_` discards the index, and `tc` is each struct.

#### Line 27

```go
t.Run(tc.name, func(t *testing.T) {
```

`t.Run(name, fn)` runs a **subtest** — each case appears as its own named test in `go test -v`. If one fails, it's easy to see exactly which case broke.

#### Line 28

```go
got := token.IsIdentifier(tc.input)
```

`token.IsIdentifier(tc.input)` checks whether the input is a valid identifier and stores the result in `got`. This is the actual testing logic.

#### Lines 29–31

```go
if got != tc.expected {
    t.Errorf("IsIdentifier(%q) = %v, want %v", tc.input, got, tc.expected)
}
```

If `got` doesn't match `expected`, `t.Errorf` reports the failure with the input, the actual result (`got`), and the expected value. `%q` prints the string quoted, `%v` prints values. If it matches, nothing happens (silently passes).

#### Line 32

```go
})
```

Ends the subtest closure.

#### Line 33

```go
}
```

Ends the `for` loop.

#### Line 34

```go
}
```

Ends the test function.

---

## Expected Output

`go test -v`:

```
=== RUN   TestIsValidIdentifier
=== RUN   TestIsValidIdentifier/simple_lowercase
=== RUN   TestIsValidIdentifier/simple_exported
... (all subtests)
--- PASS: TestIsValidIdentifier (0.00s)
PASS
ok  	go-journey/testing	0.810s
```

`go run .`:

```
Run 'go test -v' in this folder to check identifier validity rules.
```

## Key Takeaways

1. **Table-driven testing** — organizing many input/expected cases in a struct slice.
2. **`go/token.IsIdentifier`** — standard library function to check if a string is a valid Go identifier.
3. **`t.Run` subtests** — each case as a named test; fail location is clear.
4. **`t.Errorf`** — reporting a failure with a diagnostic message.
5. **Test conventions** — `_test.go` file, `TestXxx` function, `*testing.T` argument, run with `go test`.
