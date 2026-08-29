# payment-tracker

Go-তে variable declaration, naming convention আর basic type শেখার জন্য ছোট example।

**📖 ভাষা নির্বাচন করুন / Choose language:**

[🇧🇩 বাংলা](#bangla) • [🇬🇧 English](#english)

---

<a name="bangla"></a>

## 🇧🇩 বাংলা সংস্করণ

### Line 1

```go
package main
```

প্রতিটি Go file-এর শুরুতে `package` declaration থাকে। `package main` মানে এটা একটা executable program (reusable library নয়), তাই এটা `go run` দিয়ে চালানো যায়।

### Line 3

```go
import "fmt"
```

Go-র standard library থেকে `fmt` (format) package import করি। এর মাধ্যমে console-এ output print করার জন্য `fmt.Println` ব্যবহার করতে পারি।

### Line 5

```go
func main() {
```

`main()` হলো program-এর entry point — এখান থেকে Go code execute করা শুরু করে। Opening curly brace `{` দিয়ে function body শুরু হয়।

### Line 6

```go
transactionID := "TXN-2026-0091"
```

`:=` (short variable declaration) অপারেটর দিয়ে একটা **নতুন** variable এক ধাপে declare ও initialize করা হয়। এখানে `transactionID` নামে variable বানিয়ে string `"TXN-2026-0091"` দিয়েছি। Go একটা **statically typed** language, তাই compiler নিজে থেকেই type যাচাই করে নেয় — value-টা string হওয়ায় `transactionID`-এর type হয় `string`। নামটা **camelCase** আর ছোট হাতের letter দিয়ে শুরু, যেটা Go-তে বোঝায় এটা একটা **unexported** (private) variable।

### Line 7

```go
amountInCents := 150075
```

আবার `:=` দিয়ে নতুন variable declare করা হয়। Value `150075` একটা পূর্ণসংখ্যা (whole number), তাই compiler type-টা `int` হিসাবে নেয়। নামটা descriptive — `amountInCents` অর্থ পরিমাণটা cents-এ রাখা হয়েছে (যেমন 150075 cents = মূল currency-তে 1500.75)। Meaningful নাম ব্যবহার করলে কোড নিজে থেকেই বুঝা যায়।

### Line 8

```go
_isRefunded := false
```

`:=` দিয়ে `_isRefunded` নামে আরেকটা নতুন variable বানানো হয় যার value `false`, তাই এর type `bool`।

- নামটা **underscore** `_` দিয়ে শুরু। এটা Go-তে legal কিন্তু uncommon — সাধারণত "special/internal" variable বোঝাতে ব্যবহার হয়, তবে normal code-এ এড়িয়ে চলা ভালো।
- এটা একটা boolean, যেটা বোঝায় এই transaction refunded হয়েছে কি না।

### Line 10

```go
fmt.Println("Transaction:", transactionID)
```

`fmt.Println` তার arguments-গুলো console-এ print করে, প্রতিটির মাঝে space আর শেষে নতুন line যোগ করে। তাই এটি print করবে: `Transaction: TXN-2026-0091`।

### Line 11

```go
fmt.Println("Amount (cents):", amountInCents)
```

`Amount (cents):` label-এর সাথে `amountInCents`-এর value (150075) print করে। Output: `Amount (cents): 150075`।

### Line 12

```go
fmt.Println("Refunded:", _isRefunded)
```

`Refunded:` এবং তারপর boolean value `false` print করে। Output: `Refunded: false`।

### Line 13

```go
}
```

Closing curly brace — `main` function body-র শেষ বোঝায়। এখানেই program শেষ হয়।

---

## Expected Output

```
Transaction: TXN-2026-0091
Amount (cents): 150075
Refunded: false
```

## মূল শিক্ষা / Key Takeaways

1. **`:=` short declaration** — এক ধাপে নতুন variable declare + initialize করে।
2. **Type inference** — Go compiler value দেখে নিজে type বের করে নেয় (string, int, bool)।
3. **camelCase naming** — ছোট হাতের letter দিয়ে শুরু = unexported/private variable।
4. **Underscore prefix** — legal কিন্তু uncommon; সাধারণত এড়িয়ে চলা উচিত।
5. **`fmt.Println`** — একাধিক argument-কে print করে, space দিয়ে আলাদা করে।

---

---

<a name="english"></a>

## 🇬🇧 English Version

### Line 1

```go
package main
```

Every Go file starts with a `package` declaration. `package main` tells Go that this is an executable program (rather than a reusable library), so we can run it with `go run`.

### Line 3

```go
import "fmt"
```

We import the `fmt` (format) package from Go's standard library. It gives us `fmt.Println` to print output to the console.

### Line 5

```go
func main() {
```

`main()` is the entry point of the program — this is where Go starts executing the code. Opening curly brace `{` begins the function body.

### Line 6

```go
transactionID := "TXN-2026-0091"
```

The `:=` (short variable declaration) operator declares a **new** variable and initializes it in one step. Here we create `transactionID` and set it to the string `"TXN-2026-0091"`. Go is **statically typed**, so the compiler infers the type automatically — since the value is a string literal, `transactionID` becomes type `string`. The name uses **camelCase** starting with a lowercase letter, which in Go means this is an **unexported** (private/package-level) variable.

### Line 7

```go
amountInCents := 150075
```

Again `:=` declares a new variable. The value `150075` is a whole number, so the compiler infers the type as `int`. Note the descriptive name `amountInCents` — the amount is stored in cents (i.e. 150075 cents = 1500.75 in the main currency unit). Using meaningful names makes the code self-documenting.

### Line 8

```go
_isRefunded := false
```

`:=` creates another new variable named `_isRefunded` with the boolean value `false`, so its type is `bool`. Two things stand out:

- The name starts with an **underscore** `_`. This is legal in Go but uncommon — it's often used to signal a "special/internal" variable, though generally it should be avoided in normal code.
- It's a boolean, representing whether this transaction has been refunded or not.

### Line 10

```go
fmt.Println("Transaction:", transactionID)
```

`fmt.Println` prints its arguments to the console, adding a space between them and a newline at the end. So this prints: `Transaction: TXN-2026-0091`.

### Line 11

```go
fmt.Println("Amount (cents):", amountInCents)
```

Prints the label `Amount (cents):` followed by the value of `amountInCents` (150075). Output: `Amount (cents): 150075`.

### Line 12

```go
fmt.Println("Refunded:", _isRefunded)
```

Prints `Refunded:` followed by the boolean value `false`. Output: `Refunded: false`.

### Line 13

```go
}
```

Closing curly brace — marks the end of the `main` function body. The program finishes here.

---

## Expected Output

```
Transaction: TXN-2026-0091
Amount (cents): 150075
Refunded: false
```

## Key Takeaways

1. **`:=` short declaration** — Declares and initializes a new variable in one step.
2. **Type inference** — Go infers the type from the value automatically.
3. **camelCase naming** — Lowercase-start camelCase = unexported variable.
4. **Underscore prefix** — Legal but uncommon; usually avoid it.
5. **`fmt.Println`** — Prints multiple arguments separated by spaces.
