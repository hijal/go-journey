# deploy-tool

Go-তে **`defer`** ব্যবহার করে function execution-time audit log করার ছোট example -- cleanup/audit code শেষে run হয়, scope শেষ হওয়ার আগে.

**ভাষা নির্বাচন করুন / Choose language:**

[বাংলা](#bangla) * [English](#english)

---

<a name="bangla"></a>

## বাংলা সংস্করণ

### Lines 1-6

```go
package main

import (
    "fmt"
    "time"
)
```

- `package main` -- একটা executable program.
- `fmt` -- output print করতে.
- `time` -- time measurement আর `time.Sleep` এর জন্য.

### Lines 8-18

```go
func deployRelease(version string) {
    start := time.Now()

    defer func() {
        fmt.Printf("[audit] deploy %s took %v\n", version, time.Since(start).Round(time.Millisecond))
    }()

    fmt.Printf("deploying version %s...\n", version)
    time.Sleep(120 * time.Millisecond)
    fmt.Printf("deployment finished\n")
}
```

- `start := time.Now()` -- function entry-তে timestamp record.
- `defer func() { ... }()` -- anonymous function defer; function return-এর আগে run হবে.
- `time.Since(start)` -- elapsed time বের করে.
- `time.Sleep(120 * time.Millisecond)` -- simulate deployment work.

### Lines 20-23

```go
func main() {
    deployRelease("v2.4.1")
    fmt.Println("main continues after deploy")
}
```

`deployRelease("v2.4.1")` কল করলে sequence:

1. "deploying version v2.4.1..." print
2. 120ms wait
3. "deployment finished" print
4. defer block run: "[audit] deploy v2.4.1 took 121ms"
5. main-এ ফিরে "main continues after deploy" print

---

## Expected Output

```
deploying version v2.4.1...
deployment finished
[audit] deploy v2.4.1 took 121ms
main continues after deploy
```

## মূল শিক্ষা / Key Takeaways

1. **`defer`** -- ফাংশন শেষে (return-এর আগে) একটা block guarantee execute হয়.
2. **Deferred closure captures variables** -- `start`, `version` captured হয়ে শেষে use হয়.
3. **Time measurement** -- `time.Since` দিয়ে elapsed duration বের করা যায়.
4. **Audit / cleanup pattern** -- defer সাধারণত logging, resource release, rollback-এ ব্যবহৃত হয়.

---

---

<a name="english"></a>

## English Version

### Lines 1-6

```go
package main

import (
    "fmt"
    "time"
)
```

- `package main` -- an executable program.
- `fmt` -- for console output.
- `time` -- for time measurement and `time.Sleep`.

### Lines 8-18

```go
func deployRelease(version string) {
    start := time.Now()

    defer func() {
        fmt.Printf("[audit] deploy %s took %v\n", version, time.Since(start).Round(time.Millisecond))
    }()

    fmt.Printf("deploying version %s...\n", version)
    time.Sleep(120 * time.Millisecond)
    fmt.Printf("deployment finished\n")
}
```

- `start := time.Now()` -- record a timestamp at function entry.
- `defer func() { ... }()` -- deferred anonymous function; runs before the function returns.
- `time.Since(start)` -- computes elapsed time.
- `time.Sleep(120 * time.Millisecond)` -- simulates the deployment work.

### Lines 20-23

```go
func main() {
    deployRelease("v2.4.1")
    fmt.Println("main continues after deploy")
}
```

Sequence on `deployRelease("v2.4.1")`:

1. "deploying version v2.4.1..." prints
2. 120ms wait
3. "deployment finished" prints
4. deferred block runs: "[audit] deploy v2.4.1 took 121ms"
5. control returns to main, "main continues after deploy" prints

---

## Expected Output

```
deploying version v2.4.1...
deployment finished
[audit] deploy v2.4.1 took 121ms
main continues after deploy
```

## Key Takeaways

1. **`defer`** -- guarantees a block runs when the function ends (before return).
2. **Deferred closure captures variables** -- `start` and `version` are captured and used later.
3. **Time measurement** -- `time.Since` returns elapsed duration.
4. **Audit / cleanup pattern** -- defer is commonly used for logging, resource release, and rollback.
