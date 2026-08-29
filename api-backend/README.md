# api-backend

Go-তে **struct**, constructor-like **factory function** আর **receiver method** বুঝতে সহায়ক ছোট example।

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
type PaymentAPIClient struct {
    BaseURL    string
    APIKey     string
    httpClient string
}
```

`type PaymentAPIClient struct { ... }` দিয়ে `PaymentAPIClient` নামক কাস্টম data type declare করা হয়। একটা struct সম্পর্কিত data field-গুলো একসাথে group করে। এর তিনটা field আছে:

- `BaseURL` (বড় হাতের) — payment API-র base URL।
- `APIKey` (বড় হাতের) — authentication-এর জন্য secret key।
- `httpClient` (ছোট হাতের) — একটা internal detail। ছোট হাতের letter দিয়ে শুরু হওয়ায় এটা **unexported** (এই package-র বাইরে visible নয়)। এখানে string হিসেবে রাখা, কিন্তু real app-এ এখানে `*http.Client` থাকত।

### Lines 11–17

```go
func NewPaymentAPIClient(baseURL, apiKey string) *PaymentAPIClient {
    return &PaymentAPIClient{
        BaseURL:    baseURL,
        APIKey:     apiKey,
        httpClient: "default-http-client",
    }
}
```

এটা একটা **factory function** — Go-তে constructor নেই, কিন্তু convention অনুযায়ী `New...` নামের function constructor-এর মতো কাজ করে। এটা `baseURL` আর `apiKey` (দুটোই `string`) নেয় এবং `*PaymentAPIClient` (একটি `PaymentAPIClient`-এর pointer) return করে।

- `&PaymentAPIClient{ ... }` নতুন struct-এর একটা pointer তৈরি করে এবং named field initialization দিয়ে field-গুলো সেট করে।
- `httpClient`-এর ভেতরে default value `"default-http-client"` সেট করা হয়।

### Lines 19–21

```go
func (c *PaymentAPIClient) ID() string {
    return c.APIKey
}
```

এটা একটা **method** — type-এর সাথে attached function। `(c *PaymentAPIClient)` অংশটা হলো **receiver**, তাই যেকোনো `PaymentAPIClient` (বা তার pointer) এর উপর `ID()` call করা যায়। ভেতরে `c` receiver-কে বোঝায়। `ID()` `c.APIKey` return করে — অর্থাৎ client-এর API key। Method দিয়েই Go-তে custom type-এর behavior define করা হয়।

### Lines 23–27

```go
func main() {
    client := NewPaymentAPIClient(
        "https://api.pay.example.com",
        "sk_xxxxxxxxxxxxxxxxxx",
    )
```

`main`-এ factory function-কে একটা real base URL আর ফেক API key দিয়ে call করি এবং return করা pointer-টা `client`-এ store করি।

### Lines 29–30

```go
fmt.Println("Base URL:", client.BaseURL)
fmt.Println("Client ID:", client.ID())
```

প্রথমে `client.BaseURL` দিয়ে field-টা সরাসরি print করি (struct field access করার জন্য dot notation)। তারপর `client.ID()` দিয়ে method call করি। Output:

```
Base URL: https://api.pay.example.com
Client ID: sk_xxxxxxxxxxxxxxxxxx
```

### Line 31

```go
}
```

Closing brace — `main` function শেষ হয়।

---

## Expected Output

```
Base URL: https://api.pay.example.com
Client ID: sk_xxxxxxxxxxxxxxxxxx
```

## মূল শিক্ষা / Key Takeaways

1. **Struct** — সম্পর্কিত data field একসাথে group করার custom type।
2. **Factory function (`New...`)** — constructor-এর মতো pattern; struct initialize করে pointer return করে।
3. **Receiver method** — `func (c *Type) Method()` দিয়ে type-এ behavior যোগ করা।
4. **Pointer `&`** — struct-এর memory address return করতে `&` ব্যবহার।
5. **Exported vs unexported** — বড় হাতের field outside visible, ছোট হাতের না।

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
type PaymentAPIClient struct {
    BaseURL    string
    APIKey     string
    httpClient string
}
```

`type PaymentAPIClient struct { ... }` declares a custom data type called `PaymentAPIClient`. A struct groups related data fields together. It has three fields:

- `BaseURL` (uppercase) — the base URL of the payment API.
- `APIKey` (uppercase) — the secret key used for authentication.
- `httpClient` (lowercase) — an internal detail. Because it starts lowercase it is **unexported** (not visible outside this package). Here it's stored as a string, but in a real app this would hold an `*http.Client`.

### Lines 11–17

```go
func NewPaymentAPIClient(baseURL, apiKey string) *PaymentAPIClient {
    return &PaymentAPIClient{
        BaseURL:    baseURL,
        APIKey:     apiKey,
        httpClient: "default-http-client",
    }
}
```

This is a **factory function** — Go doesn't have constructors, but by convention a function named `New...` acts like one. It takes `baseURL` and `apiKey` (both `string`), and returns `*PaymentAPIClient` (a pointer to a `PaymentAPIClient`).

- `&PaymentAPIClient{ ... }` creates a pointer to a new struct and sets its fields using named field initialization.
- `httpClient` gets the default value `"default-http-client"` internally.

### Lines 19–21

```go
func (c *PaymentAPIClient) ID() string {
    return c.APIKey
}
```

This is a **method** — a function attached to a type. The part `(c *PaymentAPIClient)` is the **receiver**, so `ID()` can be called on any `PaymentAPIClient` (or pointer to it). Inside, `c` refers to the receiver. `ID()` returns `c.APIKey`, i.e. the client's API key. Methods are how you give behavior to your custom types in Go.

### Lines 23–27

```go
func main() {
    client := NewPaymentAPIClient(
        "https://api.pay.example.com",
        "sk_xxxxxxxxxxxxxxxxxx",
    )
```

In `main` we call the factory function with a real base URL and a fake API key, and store the returned pointer in `client`.

### Lines 29–30

```go
fmt.Println("Base URL:", client.BaseURL)
fmt.Println("Client ID:", client.ID())
```

First we print the field directly with `client.BaseURL` (dot notation to access a struct field). Then we call the method with `client.ID()`. Output:

```
Base URL: https://api.pay.example.com
Client ID: sk_xxxxxxxxxxxxxxxxxx
```

### Line 31

```go
}
```

Closing brace — ends the `main` function.

---

## Expected Output

```
Base URL: https://api.pay.example.com
Client ID: sk_xxxxxxxxxxxxxxxxxx
```

## Key Takeaways

1. **Struct** — A custom type that groups related data fields.
2. **Factory function (`New...`)** — Constructor-like pattern that initializes a struct and returns a pointer.
3. **Receiver method** — Adding behavior to a type with `func (c *Type) Method()`.
4. **Pointer `&`** — Use `&` to return the memory address (pointer) of a struct.
5. **Exported vs unexported** — Uppercase fields are visible externally, lowercase are not.
