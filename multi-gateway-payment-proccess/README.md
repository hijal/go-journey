# multi-gateway-payment-proccess

Go-তে **interface**, **implicit implementation**, আর **polymorphism** দিয়ে একাধিক payment gateway (processor) manage করা শেখার ছোট example।

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

### Lines 5–8

```go
type PaymentProcessor interface {
    Charge(amountCents int) error
    Name() string
}
```

একটা **interface** `PaymentProcessor` define করে। Interface হলো একটা **contract** — এটা বলে "যে কোনো type-কে `PaymentProcessor` হিসেবে ব্যবহার করতে হলে এই method-গুলো থাকতে হবে":

- `Charge(amountCents int) error` — একটা method যা cent-এ amount নেয় এবং `error` return করে।
- `Name() string` — একটা method যা processor-এর নাম string হিসেবে return করে।

Interface-টা নিজে কোনো implementation ধারণ করে না — শুধু "কী করা উচিত" define করে।

### Line 10

```go
type StripeStyleProcessor struct{}
```

`StripeStyleProcessor` নামে একটা **empty struct** type declare করে। এতে কোনো field নেই — আমরা শুধু behavior-এর জন্য type টা ব্যবহার করব। Name টা প্রতীকী: card network-এর মতো process করা (Stripe-style)।

### Lines 12–16

```go
func (StripeStyleProcessor) Charge(amountCents int) error {
    fmt.Println("charging", amountCents, "cents via card network")
    return nil
}
```

`StripeStyleProcessor`-এর উপর **method** `Charge` define করে। Receiver-টা `(StripeStyleProcessor)` — কোনো receiver variable-এর নাম নেই, কারণ আমাদের কিছুর দরকার নেই (empty struct)। Method-টা একটা message print করে এবং `nil` (no error) return করে।

### Lines 18–20

```go
func (StripeStyleProcessor) Name() string {
    return "card"
}
```

`Name` method define করে যা `"card"` return করে।

এখন `StripeStyleProcessor`-এর interface-এর দুটো method-ই আছে → এটা automatic-এ `PaymentProcessor` interface-কে **satisfy** করে। Go-তে interface implement করার জন্য `implements` keyword-এর দরকার নেই — **structural typing**: method set-টা ঠিক থাকলেই হয়।

### Line 22

```go
type MobileWalletProcessor struct{}
```

আরেকটা empty struct type `MobileWalletProcessor` — মোবাইল wallet (যেমন bKash/Nagad/GPay-এর মতো) gateway-কে represent করে।

### Lines 24–27

```go
func (MobileWalletProcessor) Charge(amountCents int) error {
    fmt.Println("charging", amountCents, "cents via mobile wallet")
    return nil
}
```

`MobileWalletProcessor`-এর উপর `Charge` method — ভিন্ন message print করে, কিন্তু একই signature (`int -> error`), তাই interface-এর সাথে match করে।

### Lines 29–31

```go
func (MobileWalletProcessor) Name() string {
    return "mobile-wallet"
}
```

`Name` method — `"mobile-wallet"` return করে। দুটো type এখন দুটোই `PaymentProcessor` interface-কে satisfy করে।

### Lines 33–38

```go
func processPayment(p PaymentProcessor, amountCents int) {
    fmt.Println("using processor:", p.Name())
    if err := p.Charge(amountCents); err != nil {
        fmt.Println("Payment failed", err)
    }
}
```

আলাদা function `processPayment` যা একজন `PaymentProcessor` (`p`) আর একটা amount নেয়। **Polymorphism-এর মজা এখানেই:** function টা শুধু interface-কে জানে — `StripeStyleProcessor` নাকি `MobileWalletProcessor` সেটা জানার দরকার নেই।

- `p.Name()` — কোন processor ব্যবহার হচ্ছে, print করে।
- `p.Charge(amountCents)` — charge করে; result error `err`-এ। `if err := ...; err != nil` pattern দিয়ে check — error হলে "Payment failed" print করে।
- Caller হিসেবে আমরা **যেকোনো** `PaymentProcessor` দিতে পারি। নতুন gateway যোগ করলে function-টা অপরিবর্তিত থাকে — এটা open/closed-এর মতো principle-এর একটা সাধারণ উদাহরণ।

### Line 40

```go
func main() {
```

Program-এর entry point।

### Lines 41–44

```go
processors := []PaymentProcessor{
    StripeStyleProcessor{},
    MobileWalletProcessor{},
}
```

`processors` — একটা `[]PaymentProcessor` slice। গুরুত্বপূর্ণ: element-গুলোর type হলো **interface** `PaymentProcessor`, কিন্তু আমরা দুটো concrete type-এর value দিচ্ছি (`StripeStyleProcessor{}` আর `MobileWalletProcessor{}`)। যেহেতু ওরা interface-কে satisfy করে, ওরা-কে interface-তে assign করা যায়।

### Lines 46–48

```go
for _, p := range processors {
    processPayment(p, 25000)
}
```

slice-এর প্রতিটা `p`-তে `processPayment(p, 25000)` call করে — amount `25000` cents (250 টাকা/ডলার) দিয়ে। `p` code-এ interface হলোেও, runtime-এ ঠিক আসল concrete type-এর method-ই চলে (dynamic dispatch)।

### Line 49

```go
}
```

Closing brace — `main` function শেষ হয়।

---

## Expected Output

```
using processor: card
charging 25000 cents via card network
using processor: mobile-wallet
charging 25000 cents via mobile wallet
```

## মূল শিক্ষা / Key Takeaways

1. **Interface** — method-এর contract; type-কে কী কী করতে হবে define করে।
2. **Implicit implementation** — `implements` keyword নাই; method set-টা match করলেই interface satisfy হয়।
3. **Polymorphism** — `[]PaymentProcessor` slice-এ বিভিন্ন concrete type রাখা যায়; একই function সব চালায়।
4. **Dynamic dispatch** — interface ভেরিয়েবল দিয়ে method call করলে runtime-এ আসল type-এর method চলে।
5. **Extensibility** — নতুন gateway যোগ করতে হলে শুধু interface-এর method-গুলো implement করলেই চলবে।

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

### Lines 5–8

```go
type PaymentProcessor interface {
    Charge(amountCents int) error
    Name() string
}
```

Defines an **interface** `PaymentProcessor`. An interface is a **contract** — it declares which methods any type must have to be used as a `PaymentProcessor`:

- `Charge(amountCents int) error` — a method taking an amount in cents and returning an `error`.
- `Name() string` — a method returning the processor's name as a string.

An interface holds no implementation itself — it only defines "what must be possible".

### Line 10

```go
type StripeStyleProcessor struct{}
```

Declares an **empty struct** type `StripeStyleProcessor`. It has no fields — we only use the type for its behavior. The name is symbolic: processing like a card-network (Stripe-style).

### Lines 12–16

```go
func (StripeStyleProcessor) Charge(amountCents int) error {
    fmt.Println("charging", amountCents, "cents via card network")
    return nil
}
```

Defines the **method** `Charge` on `StripeStyleProcessor`. The receiver is `(StripeStyleProcessor)` — no receiver variable name, because we don't need one (empty struct). The method prints a message and returns `nil` (no error).

### Lines 18–20

```go
func (StripeStyleProcessor) Name() string {
    return "card"
}
```

Defines the `Name` method, returning `"card"`.

Now `StripeStyleProcessor` has both methods of the interface → it automatically **satisfies** `PaymentProcessor`. Go needs no `implements` keyword — it's **structural typing**: if the method set matches, it satisfies the interface.

### Line 22

```go
type MobileWalletProcessor struct{}
```

Another empty struct type `MobileWalletProcessor` — representing a mobile-wallet gateway (like bKash/Nagad/GPay).

### Lines 24–27

```go
func (MobileWalletProcessor) Charge(amountCents int) error {
    fmt.Println("charging", amountCents, "cents via mobile wallet")
    return nil
}
```

`Charge` on `MobileWalletProcessor` — prints a different message but keeps the same signature (`int -> error`), so it matches the interface.

### Lines 29–31

```go
func (MobileWalletProcessor) Name() string {
    return "mobile-wallet"
}
```

`Name` method returning `"mobile-wallet"`. Both types now satisfy the `PaymentProcessor` interface.

### Lines 33–38

```go
func processPayment(p PaymentProcessor, amountCents int) {
    fmt.Println("using processor:", p.Name())
    if err := p.Charge(amountCents); err != nil {
        fmt.Println("Payment failed", err)
    }
}
```

A separate function `processPayment` that takes a `PaymentProcessor` (`p`) and an amount. **This is the payoff of polymorphism:** the function only knows the interface — it doesn't need to know `StripeStyleProcessor` vs `MobileWalletProcessor`.

- `p.Name()` — prints which processor is being used.
- `p.Charge(amountCents)` — performs the charge; the resulting error goes into `err`. Checked with the `if err := ...; err != nil` pattern — on error it prints "Payment failed".
- As a caller we can pass **any** `PaymentProcessor`. Adding a new gateway leaves this function untouched — a simple take on the open/closed principle.

### Line 40

```go
func main() {
```

Program entry point.

### Lines 41–44

```go
processors := []PaymentProcessor{
    StripeStyleProcessor{},
    MobileWalletProcessor{},
}
```

`processors` — a `[]PaymentProcessor` slice. Key point: the element type is the **interface** `PaymentProcessor`, but we store concrete values of both types (`StripeStyleProcessor{}` and `MobileWalletProcessor{}`). Because they satisfy the interface, they can be assigned to it.

### Lines 46–48

```go
for _, p := range processors {
    processPayment(p, 25000)
}
```

For each `p` in the slice, calls `processPayment(p, 25000)` with `25000` cents (250 taka/dollars). Although `p` is declared as the interface, at runtime the actual concrete type's method executes (dynamic dispatch).

### Line 49

```go
}
```

Closing brace — ends the `main` function.

---

## Expected Output

```
using processor: card
charging 25000 cents via card network
using processor: mobile-wallet
charging 25000 cents via mobile wallet
```

## Key Takeaways

1. **Interface** — a contract of methods; defines what a type must be able to do.
2. **Implicit implementation** — no `implements` keyword; a matching method set satisfies the interface.
3. **Polymorphism** — different concrete types can live in a `[]PaymentProcessor` slice and run through one function.
4. **Dynamic dispatch** — calling a method through an interface variable dispatches to the real type at runtime.
5. **Extensibility** — adding a new gateway only requires implementing the interface's methods.