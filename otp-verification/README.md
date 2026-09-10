# otp-verification

Go-তে **fixed-size array-কে `==` দিয়ে element-wise compare** শেখার ছোট example — OTP verify (payment approval)।

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

`fmt` — `Println`।

### Lines 5–6

```go
func main() {
	serverOTP := [6]int{7, 3, 0, 9, 1, 4}
```

`[6]int{7, 3, 0, 9, 1, 4}` — server-এ generate হওয়া **fixed-size array**। 6-digit OTP: `730914`।

### Line 7

```go
	userOTP := [6]int{7, 3, 0, 9, 1, 5}
```

`userOTP` — user-এর input: `730915`। **শেষ digit-টা আলাদা** (4 বনাম 5)।

### Line 8

```go
	fmt.Println("OTP length:", len(serverOTP))
```

`len(serverOTP)` — 6 → `OTP length: 6`।

### Lines 10–14

```go
	if serverOTP == userOTP {
		fmt.Println("OTP verified: payment approved")
	} else {
		fmt.Println("OTP mismatch: payment declined")
	}
```

**মূল concept — array comparison:**

- `serverOTP == userOTP` — **element-wise comparison**: প্রতিটা index-এ value তুলনা হয়। সব মিললে `true`, যেকোনো একটা পার্থক্যে `false`।
- **Array `==`-যোগ্য** যখন element type comparable (এখানে `int`)। Array-র সাইজও type-এর অংশ।
- **Slice-এর সাথে পার্থক্য:** slice-তে `s1 == s2` **compile error** — শুধু `nil` দিয়ে তুলনা allowed। Array-তে সরাসরি `==` কাজ করে।
- এখানে শেষ digit mismatch → `false` → `OTP mismatch: payment declined`।

---

## Expected Output

```
OTP length: 6
OTP mismatch: payment declined
```

> `userOTP`-র শেষ digit 4 করলে (`...1, 4`) output হবে `OTP verified: payment approved`।

## মূল শিক্ষা / Key Takeaways

1. **Array `==`** — element-wise, comparable element type-এ।
2. **Fixed-size type** — `[6]int` vs `[6]int`, সাইজ type-এর অংশ।
3. **Slice নয়** — slice-এ `==` compile-time error, array-তে ঠিক আছে।
4. **`len()`** — element-সংখ্যা।
5. **Match/decline branch** — `if/else` OTP verification flow।

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

`fmt` — for `Println`.

### Lines 5–6

```go
func main() {
	serverOTP := [6]int{7, 3, 0, 9, 1, 4}
```

`[6]int{7, 3, 0, 9, 1, 4}` — a **fixed-size array** generated on the server. The 6-digit OTP: `730914`.

### Line 7

```go
	userOTP := [6]int{7, 3, 0, 9, 1, 5}
```

`userOTP` — the user's input: `730915`. The **last digit differs** (4 vs 5).

### Line 8

```go
	fmt.Println("OTP length:", len(serverOTP))
```

`len(serverOTP)` — 6 → `OTP length: 6`.

### Lines 10–14

```go
	if serverOTP == userOTP {
		fmt.Println("OTP verified: payment approved")
	} else {
		fmt.Println("OTP mismatch: payment declined")
	}
```

**The core concept — array comparison:**

- `serverOTP == userOTP` — **element-wise comparison**: values compared at every index. All equal → `true`; any difference → `false`.
- An array is **comparable with `==`** when its element type is comparable (here `int`). The array size is also part of the type.
- **Contrast with slices:** `s1 == s2` is a **compile error** for slices — only a `nil` comparison is allowed. Arrays support direct `==`.
- Here the last digit mismatches → `false` → `OTP mismatch: payment declined`.

---

## Expected Output

```
OTP length: 6
OTP mismatch: payment declined
```

> Change the last digit of `userOTP` to 4 (`...1, 4`) and the output becomes `OTP verified: payment approved`.

## Key Takeaways

1. **Array `==`** — element-wise, with comparable element types.
2. **Fixed-size type** — `[6]int`, the size is part of the type.
3. **Not a slice** — `==` on slices is a compile-time error; on arrays it works.
4. **`len()`** — the number of elements.
5. **Match/decline branch** — the `if/else` OTP verification flow.