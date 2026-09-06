# closure-fn

Go-তে **closure** শেখার ছোট example -- function যেটা outer variable capture করে, পরে captured value নিয়ে কাজ করে.

**ভাষা নির্বাচন করুন / Choose language:**

[বাংলা](#bangla) * [English](#english)

---

<a name="bangla"></a>

## বাংলা সংস্করণ

### Lines 1-3

```go
package main

import "fmt"
```

- `package main` -- একটা executable program.
- `fmt` -- output print করতে.

### Lines 5-9

```go
func makeAdder(x int) func(int) int {
    return func(y int) int {
        return x + y
    }
}
```

Closure তৈরির function:

- `makeAdder(x int)` -- একটা `int` parameter নেয়.
- `func(int) int` return করে -- মানে একটা function যা `int` নেয় আর `int` দেয়.
- ভিতরের anonymous function `x` কে **capture** করে (closure).
- `return x + y` -- captured `x` আর পরের parameter `y` যোগ করে.

### Lines 11-14

```go
func main() {
    add5 := makeAdder(5)
    fmt.Println(add5(10))
}
```

- `add5 := makeAdder(5)` -- `x = 5` capture হয়ে একটা function পাওয়া যায়.
- `add5(10)` -- `x=5`, `y=10`, result = `15`.

---

## Expected Output

```
15
```

## মূল শিক্ষা / Key Takeaways

1. **Closure** -- function তার lexical scope-এর variables capture করে রাখে.
2. **Function returning function** -- `makeAdder` returns `func(int) int`.
3. **Captured value persists** -- `x = 5` বারবার ব্যবহার হতে পারে returned function থেকে.
4. **Higher-order function** -- function কে value হিসেবে pass/return করার pattern.

---

---

<a name="english"></a>

## English Version

### Lines 1-3

```go
package main

import "fmt"
```

- `package main` -- an executable program.
- `fmt` -- for console output.

### Lines 5-9

```go
func makeAdder(x int) func(int) int {
    return func(y int) int {
        return x + y
    }
}
```

A closure-producing function:

- `makeAdder(x int)` -- takes one `int` parameter.
- returns `func(int) int` -- a function that takes an `int` and returns an `int`.
- the inner anonymous function **captures `x`** (closure).
- `return x + y` -- adds the captured `x` and the later parameter `y`.

### Lines 11-14

```go
func main() {
    add5 := makeAdder(5)
    fmt.Println(add5(10))
}
```

- `add5 := makeAdder(5)` -- `x = 5` is captured and a function is returned.
- `add5(10)` -- `x=5`, `y=10`, result = `15`.

---

## Expected Output

```
15
```

## Key Takeaways

1. **Closure** -- a function captures variables from its lexical scope.
2. **Function returning function** -- `makeAdder` returns `func(int) int`.
3. **Captured value persists** -- the captured `x = 5` can be reused by the returned function.
4. **Higher-order function** -- the pattern of passing/returning functions as values.
