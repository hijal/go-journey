# config-loader

Go-তে `len` builtin function, slicing, আর variable shadowing বুঝতে সহায়ক ছোট example।

**📖 ভাষা নির্বাচন করুন / Choose language:**

[🇧🇩 বাংলা](#bangla) • [🇬🇧 English](#english)

---

<a name="bangla"></a>

## 🇧🇩 বাংলা সংস্করণ

### Line 1

```go
package main
```

Package-টি `main` হিসেবে declare করা হয়, মানে এটা একটা executable program যা `go run` দিয়ে চালানো যায়।

### Line 3

```go
import "fmt"
```

Console-এ output print করার জন্য standard library থেকে `fmt` package import করা হয়।

### Line 5

```go
func main() {
```

Program-এর entry point। Execution এখান থেকে শুরু হয়, আর `{` দিয়ে function body শুরু হয়।

### Lines 6–10

```go
supportedRegions := []string{
    "BD",
    "IN",
    "SG",
}
```

`supportedRegions` নামে একটা variable declare করা হয়। Type হলো `[]string` — একটা **slice** (dynamic list) of strings। এখানে অ্যাপ-র তিনটা supported region code আছে: `"BD"` (বাংলাদেশ), `"IN"` (ভারত), আর `"SG"` (সিঙ্গাপুর)। এটা supported region-এর একটা config list।

### Line 12

```go
regionCount := len(supportedRegions)
```

`len()` একটা Go builtin function যা element-এর সংখ্যা return করে। তাই `len(supportedRegions)` return করে `3`। `:=` দিয়ে সেই মান একটা নতুন variable `regionCount`-এ store করি।

### Line 14

```go
fmt.Println("region count:", regionCount)
```

Label আর value print করে, output হয় `region count: 3`।

### Lines 16–20

```go
{
    len := regionCount

    fmt.Println("Shadowed len value:", len)
}
```

একটা `{ ... }` block একটা **nested scope** তৈরি করে। এই block-এর ভেতরে `len := regionCount` লিখলে একটা **নতুন local variable** `len` declare হয়। এখন এই block-এর ভেতরে `len` নামটা আর builtin function-কে বোঝায় না — বরং এই variable-কে বোঝায়। একে বলে **variable shadowing**। তাই `fmt.Println` আমাদের নতুন variable-এর value (`3`) print করে। Block শেষ হলে shadowed `len` variable-টা scope-র বাইরে চলে যায় এবং হারিয়ে যায়।

### Line 22

```go
fmt.Println("builtin len still works:", len(supportedRegions))
```

Block-এর বাইরে (আবার `main`-এর scope-এ) shadowing আর প্রযোজ্য নয়। এখানে `len` আবার builtin function-কে বোঝায়, তাই `len(supportedRegions)` `3` return করে এবং print হয়। এটা প্রমাণ করে shadowing শুধু সেই block-এর ভেতরেই সীমাবদ্ধ।

### Line 23

```go
}
```

Closing brace — `main` function body শেষ হয়, আর program শেষ।

---

## Expected Output

```
region count: 3
Shadowed len value: 3
builtin len still works: 3
```

## মূল শিক্ষা / Key Takeaways

1. **`len` builtin** — slice/array-র element সংখ্যা বের করে।
2. **Slice literal** — `[]string{"BD","IN","SG"}` দিয়ে dynamic list।
3. **Variable shadowing** — inner scope-এ নাম (যেমন `len`) নতুন variable বোঝায়; outer scope ফিরলে আবার builtin।
4. **Scope** — shadowing শুধু নিজের block-এর ভেতরে থাকে।

---

---

<a name="english"></a>

## 🇬🇧 English Version

### Line 1

```go
package main
```

Declares the package as `main`, meaning this is an executable program we can run with `go run`.

### Line 3

```go
import "fmt"
```

Imports the `fmt` package from the standard library, used for printing output to the console.

### Line 5

```go
func main() {
```

The entry point of the program. Execution starts here, and the `{` opens the function body.

### Lines 6–10

```go
supportedRegions := []string{
    "BD",
    "IN",
    "SG",
}
```

We declare a variable `supportedRegions`. The type is `[]string` — a **slice** (dynamic list) of strings. Here it holds three region codes the app supports: `"BD"` (Bangladesh), `"IN"` (India), and `"SG"` (Singapore). This represents a config list of supported regions.

### Line 12

```go
regionCount := len(supportedRegions)
```

`len()` is a Go builtin that returns the number of elements. So `len(supportedRegions)` returns `3`. We store that in a new variable `regionCount` using `:=`.

### Line 14

```go
fmt.Println("region count:", regionCount)
```

Prints the label and the value, producing `region count: 3`.

### Lines 16–20

```go
{
    len := regionCount

    fmt.Println("Shadowed len value:", len)
}
```

A `{ ... }` block creates a **nested scope**. Inside this block we write `len := regionCount`, which declares a **new local variable** named `len`. Now, within this block, the name `len` no longer refers to the builtin function — it refers to this variable. This is called **variable shadowing**. So `fmt.Println` prints the value of our new variable (`3`). When the block ends, the shadowed `len` variable goes out of scope and disappears.

### Line 22

```go
fmt.Println("builtin len still works:", len(supportedRegions))
```

Outside the block (back in `main`'s scope), the shadowing no longer applies. Here `len` again refers to the builtin function, so `len(supportedRegions)` returns `3` and gets printed. This proves shadowing is limited to the block where it was declared.

### Line 23

```go
}
```

Closing brace — ends the `main` function body, and the program finishes.

---

## Expected Output

```
region count: 3
Shadowed len value: 3
builtin len still works: 3
```

## Key Takeaways

1. **`len` builtin** — Returns the number of elements in a slice/array.
2. **Slice literal** — A `[]string` slice holds a dynamic list of strings.
3. **Variable shadowing** — Inside a block a name can shadow another; it's limited to that block.
4. **Scope** — Shadowing only lives within its own block.
