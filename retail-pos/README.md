# retail-pos

Go-তে **closure with state** শেখার ছোট example -- প্রতিটা register/counter-এর নিজস্ব captured state থাকে, আলাদা counter আলাদা total track করে.

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

- `package main` -- executable program.
- `fmt` -- output print করতে.

### Lines 5-11

```go
func newRegister(name string) func(float64) float64 {
    total := 0.0
    return func(amount float64) float64 {
        total += amount
        return total
    }
}
```

`newRegister` factory function:

- `name string` parameter নেয় (এখানে print/log-এ ব্যবহার না হলেও semantic).
- `total := 0.0` -- local state, প্রতিটা নতুন register-এ zero দিয়ে শুরু.
- returned function `total` কে **capture** করে (closure).
- প্রতি call-এ `total += amount` দিয়ে accumulate হয়.

### Lines 12-22

```go
func main() {
    counterA := newRegister("Counter A")
    counterB := newRegister("Counter B")

    counterA(500)
    counterA(300)
    counterB(1000)

    fmt.Printf("counter-A total: %.2f\n", counterA(200))
    fmt.Printf("counter-B total: %.2f\n", counterB(450))
}
```

দুটো আলাদা register instance:

- `counterA` -- 500, 300, 200 = 1000.
- `counterB` -- 1000, 450 = 1450.
- প্রতিটা instance-এর নিজস্ব captured `total` থাকে, পরস্পর interfere করে না.

---

## Expected Output

```
counter-A total: 1000.00
counter-B total: 1450.00
```

## মূল শিক্ষা / Key Takeaways

1. **Closure with state** -- function-এর ভিতরের local variable captured হয়ে বারবার ব্যবহৃত হয়.
2. **Independent state per closure** -- প্রতিটা call-এ `newRegister` নতুন scope তৈরি করে, তাই state আলাদা থাকে.
3. **Factory pattern** -- `newRegister` returns a configured closure, encapsulates state.
4. **Accumulator pattern** -- `total += amount` -- running total maintain করা.

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

### Lines 5-11

```go
func newRegister(name string) func(float64) float64 {
    total := 0.0
    return func(amount float64) float64 {
        total += amount
        return total
    }
}
```

`newRegister` is a factory function:

- `name string` parameter is taken (not used for printing here, but semantically meaningful).
- `total := 0.0` -- local state, each new register starts at zero.
- the returned function **captures** `total` (closure).
- on each call `total += amount` accumulates.

### Lines 12-22

```go
func main() {
    counterA := newRegister("Counter A")
    counterB := newRegister("Counter B")

    counterA(500)
    counterA(300)
    counterB(1000)

    fmt.Printf("counter-A total: %.2f\n", counterA(200))
    fmt.Printf("counter-B total: %.2f\n", counterB(450))
}
```

Two independent register instances:

- `counterA` -- 500, 300, 200 = 1000.
- `counterB` -- 1000, 450 = 1450.
- each instance has its own captured `total`, no cross-interference.

---

## Expected Output

```
counter-A total: 1000.00
counter-B total: 1450.00
```

## Key Takeaways

1. **Closure with state** -- a function can capture and persist a local variable.
2. **Independent state per closure** -- each `newRegister` call creates a new scope, so state is isolated.
3. **Factory pattern** -- `newRegister` returns a configured closure that encapsulates state.
4. **Accumulator pattern** -- `total += amount` maintains a running total.
