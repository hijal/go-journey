# attendance

Go-তে **`min` / `max` builtins** (`Go 1.21+`) দিয়ে regular ও overtime work-time ভাগ করা শেখার ছোট example — integer arithmetic + modulo।

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

### Line 5

```go
func main() {
```

Program-এর entry point।

### Line 6

```go
const regularMinutes = 8 * 60 // full work day
```

`regularMinutes` — একটা constant: পুরো work day = 8 ঘণ্টা × 60 মিনিট = 480 মিনিট। Constant কারণ এটা fix।

### Line 8

```go
workedMinutes := 537
```

`workedMinutes` — চলমান value: আজ কত মিনিট কাজ হয়েছে (8 ঘণ্টা 57 মিনিট = 537)।

### Line 10

```go
regular := min(workedMinutes, regularMinutes)
```

**`min` builtin** — দুটো argument-এর মধ্যে ছোটটা return করে:

- `min(537, 480)` = `480` — regular = পুরো work day (কারণ regular 480-এর বেশি হতে পারে না)

মানে: regular minutes = 480 (কে। 537 480-এর বেশি, কিন্তু regular শুধু 480 পর্যন্ত)।

### Line 11

```go
overtime := max(0, workedMinutes-regularMinutes)
```

**`max` builtin** — দুটো argument-এর মধ্যে বড়টা return করে:

- `workedMinutes - regularMinutes` = 537 − 480 = 57
- `max(0, 57)` = `57` — negative না হলে। কাজ কম করলে (যেমন 400) difference negative হত, `max(0, negative)` = 0 — overtime 0।

মানে: overtime = 57 মিনিট।

### Lines 13–14

```go
hours, minutes := regular/60, regular%60
otHours, otMinutes := overtime/60, overtime%60
```

**Integer division + modulo** দিয়ে ঘণ্টা-vs-মিনিট ভাগ:

- `regular/60` = 480/60 = 8 (hours), `regular%60` = 0 (minutes) → `8h 0m`
- `overtime/60` = 57/60 = 0 (hours), `overtime%60` = 57 (minutes) → `0h 57m`

`/` দেয় integer quotient, `%` দেয় remainder।

### Lines 16–17

```go
fmt.Printf("regular: %dh %dm\n", hours, minutes)
fmt.Printf("overtime: %dh %dm\n", otHours, otMinutes)
```

Formatted output:

- `regular: 8h 0m`
- `overtime: 0h 57m`

### Line 18

```go
fmt.Println("overtime pay multiplier applies:", overtime > 0)
```

`overtime > 0` (`57 > 0`) = `true` — overtime হয়েছে তাই multiplier প্রযোজ্য। Output: `overtime pay multiplier applies: true`।

### Line 19

```go
}
```

Closing brace — `main` function শেষ হয়।

---

## Expected Output

```
regular: 8h 0m
overtime: 0h 57m
overtime pay multiplier applies: true
```

## মূল শিক্ষা / Key Takeaways

1. **`min` / `max` builtins** (Go 1.21+) — দুটো মানের ছোট/বড় খুঁজে আনা।
2. **`/` integer division + `%` modulo** — ঘণ্টা/মিনিট, পুরো-page/quota ভাগ।
3. **`max(0, diff)`** — negative-কে 0-এ clamp করার ট্রিক (underflow safeguard)।
4. **Constant vs variable** — fix time-এর জন্য `const`, চলমান মানের জন্য `var`/`:=`।

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

### Line 5

```go
func main() {
```

Program entry point.

### Line 6

```go
const regularMinutes = 8 * 60 // full work day
```

`regularMinutes` — a constant: full work day = 8 hours × 60 minutes = 480 minutes. Constant because it's fixed.

### Line 8

```go
workedMinutes := 537
```

`workedMinutes` — the live value: how many minutes were worked today (8h 57m = 537).

### Line 10

```go
regular := min(workedMinutes, regularMinutes)
```

**`min` builtin** — returns the smaller of two arguments:

- `min(537, 480)` = `480` — regular = the full work day

So regular minutes = 480.

### Line 11

```go
overtime := max(0, workedMinutes-regularMinutes)
```

**`max` builtin** — returns the larger of two arguments:

- `workedMinutes - regularMinutes` = 537 − 480 = 57
- `max(0, 57)` = `57` — only if not negative. If you worked fewer hours (e.g. 400), the difference would be negative and `max(0, negative)` = 0 — no overtime.

So overtime = 57 minutes.

### Lines 13–14

```go
hours, minutes := regular/60, regular%60
otHours, otMinutes := overtime/60, overtime%60
```

Split hours vs minutes with **integer division + modulo**:

- `regular/60` = 480/60 = 8 (hours), `regular%60` = 0 (minutes) → `8h 0m`
- `overtime/60` = 57/60 = 0 (hours), `overtime%60` = 57 (minutes) → `0h 57m`

`/` gives the integer quotient, `%` gives the remainder.

### Lines 16–17

```go
fmt.Printf("regular: %dh %dm\n", hours, minutes)
fmt.Printf("overtime: %dh %dm\n", otHours, otMinutes)
```

Formatted output:

- `regular: 8h 0m`
- `overtime: 0h 57m`

### Line 18

```go
fmt.Println("overtime pay multiplier applies:", overtime > 0)
```

`overtime > 0` (`57 > 0`) = `true` — there was overtime, so the multiplier applies. Output: `overtime pay multiplier applies: true`.

### Line 19

```go
}
```

Closing brace — ends the `main` function.

---

## Expected Output

```
regular: 8h 0m
overtime: 0h 57m
overtime pay multiplier applies: true
```

## Key Takeaways

1. **`min` / `max` builtins** (Go 1.21+) — the smaller/larger of two values.
2. **`/` integer division + `%` modulo** — split hours/minutes, pages/quotas.
3. **`max(0, diff)`** — clamp negatives to 0 (underflow safeguard).
4. **Constant vs variable** — `const` for fixed time, `var`/`:=` for live values.