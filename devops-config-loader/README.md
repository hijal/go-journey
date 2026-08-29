# devops-config-loader

Go-তে **declaration order vs execution order**, global variable, `const`, আর function-এর order গুরুত্বপূর্ণ নয় — এটা বুঝতে সহায়ক ছোট example।

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

### Line 3

```go
import "fmt"
```

Console-এ output print করার জন্য `fmt` package import করা হয়।

### Lines 5–9

```go
func main() {
    fmt.Println("environment:", currentEnvironment)
    fmt.Println("max retries:", maxRetries)
    printBanner()
}
```

`main()` হলো execution-এর **আসল শুরু** — আর এই function-টা file-এর উপরে লেখা হলেও, এটাই **প্রথমে execute** হয়। ভেতরে:

- `fmt.Println("environment:", currentEnvironment)` — `currentEnvironment`-এর মান print করে (যা `main`-এর পরে source-এ declare হলেও `main` চালানোর সময় তা initialize হয়ে গেছে)।
- `fmt.Println("max retries:", maxRetries)` — constant `maxRetries`-এর মান print করে।
- `printBanner()` — নিচে define করা function call করে।

**গুরুত্বপূর্ণ ধারণা:** Go-তে file-এর **উপরে/নিচে লেখা** কোনও ব্যাপার না — `main` function আলাদা, আর `const`/`var`/function সবগুলো compile-এর আগে **সব প্রোগ্রাম-ভেদেই পরিচিত (visible)**। `main` শেষ পর্যন্ত চালানো হয়।

### Line 11

```go
var currentEnvironment = detectEnvironment()
```

একটা **package-level variable** `currentEnvironment` declare করে, যার মান আসে `detectEnvironment()` function call থেকে। এটা `main`-এর **নিচে** লেখা, কিন্তু Go-তে এটা ঠিক আছে — variable initialization `main` চালানোর **আগেই** হয়ে যায়।

### Line 13

```go
const maxRetries = 5
```

একটা **constant** `maxRetries` যার মান `5`, যা program-জুড়ে অপরিবর্তনীয়। Const-ও `main`-এর নিচে লেখা, তবুও `main`-এ ব্যবহারযোগ্য।

### Lines 15–17

```go
func detectEnvironment() string {
    return "development"
}
```

একটা function `detectEnvironment` যা `"development"` string return করে। এটা `main`-এর পরে define, কিন্তু `main`-এ নাম দিয়ে reference করা যায়। (Real app-এ এখানে পরিবেশ/env check করা হতো, যেমন `dev`/`prod`।)

### Lines 19–21

```go
func printBanner() {
    fmt.Println("---deployment config loaded---")
}
```

একটা function `printBanner` যা একটা banner/message print করে।

**মূল ধারণা:** Go-তে function, variable, constant-এর declaration **order গুরুত্বপূর্ণ নয়** — সব package-স্তরের identifier program-এর যেকোনো জায়গা থেকে ব্যবহারযোগ্য, যতক্ষণ compile-এ একটা trigger order (dependency) মেটে। এখানে `var currentEnvironment = detectEnvironment()` `detectEnvironment`-কে আগে-called মনে করলেও, compile-এ self-contained cycle নেই তাই ঠিক কাজ করে।

---

## Expected Output

```
environment: development
max retries: 5
---deployment config loaded---
```

## মূল শিক্ষা / Key Takeaways

1. **Execution order ≠ declaration order** — source-এ যেখানেই লেখা হোক, `main`-ই শেষে চলে; declarations আগে resolve হয়।
2. **Package-level visibility** — `const`/`var`/function সব জায়গা থেকে used হওয়া যায়, order নির্বিশেষে।
3. **Variable initialization** — package-level `var ... = detectEnvironment()`-এর মতো initialization `main`-এর আগে চলে।
4. **`const`** — compile-time-এ fixed, অপরিবর্তনীয় মান।
5. **Function definition order** — loop/circular dependency না থাকলে নিচে define করা function উপরে call করা যায়।

---

---

<a name="english"></a>

## 🇬🇧 English Version

### Line 1

```go
package main
```

Declares an executable program (`main` package), runnable via `go run`.

### Line 3

```go
import "fmt"
```

Imports the `fmt` package for console output.

### Lines 5–9

```go
func main() {
    fmt.Println("environment:", currentEnvironment)
    fmt.Println("max retries:", maxRetries)
    printBanner()
}
```

`main()` is the **actual start** of execution — even though this function is written near the top of the file, this is what **runs first**. Inside it:

- `fmt.Println("environment:", currentEnvironment)` — prints the value of `currentEnvironment` (which is already initialized by the time `main` runs, even though it's declared later in the source).
- `fmt.Println("max retries:", maxRetries)` — prints the value of the constant `maxRetries`.
- `printBanner()` — calls a function defined further down.

**Key idea:** In Go, where things are **written** (top vs bottom of the file) doesn't matter — `main` is special, and all `const`/`var`/functions are known **across the whole program** before anything runs. `main` simply runs at the end.

### Line 11

```go
var currentEnvironment = detectEnvironment()
```

Declares a **package-level variable** `currentEnvironment` whose value comes from calling `detectEnvironment()`. It's written **below** `main`, but that's fine — variable initialization happens **before** `main` runs.

### Line 13

```go
const maxRetries = 5
```

A **constant** `maxRetries` with the value `5`, immutable for the whole program. It's also written below `main`, yet usable inside `main`.

### Lines 15–17

```go
func detectEnvironment() string {
    return "development"
}
```

A function `detectEnvironment` that returns the string `"development"`. It's defined after `main`, but `main` can still reference it by name. (In a real app this would check an environment flag like `dev`/`prod`.)

### Lines 19–21

```go
func printBanner() {
    fmt.Println("---deployment config loaded---")
}
```

A function `printBanner` that prints a banner/message.

**Key idea:** In Go, the **order of declarations** doesn't matter — all package-level identifiers (functions, variables, constants) are usable from anywhere in the program as long as there's no compile-time dependency cycle. Here `var currentEnvironment = detectEnvironment()` references `detectEnvironment` even though it's defined later — that's allowed since there's no circular dependency.

---

## Expected Output

```
environment: development
max retries: 5
---deployment config loaded---
```

## Key Takeaways

1. **Execution order ≠ declaration order** — `main` runs last regardless of where it's written; declarations resolve first.
2. **Package-level visibility** — `const`/`var`/functions are usable from anywhere, regardless of order.
3. **Variable initialization** — package-level `var ... = detectEnvironment()`-style initialization runs before `main`.
4. **`const`** — a fixed, immutable value at compile time.
5. **Function definition order** — functions defined below can be called from above (as long as there's no cycle).
