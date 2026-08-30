# plan-based-rate-limit

Go-তে **map**, **map update/delete**, আর Go 1.21+ এর **`slices`/`maps`** standard library (sorted iteration) দিয়ে plan-based rate-limit config শেখার ছোট example।

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

### Lines 3–7

```go
import (
    "fmt"
    "maps"
    "slices"
)
```

তিনটা package import করা হয়:

- `fmt` — console-এ output print করার জন্য।
- `maps` — Go 1.21+ standard library। এখানে `maps.Keys` ব্যবহার করা হয়, যা map-এর সব key-র একটা slice দেয় (order-হীন, তারপর sort করা হবে)।
- `slices` — Go 1.21+ standard library। এখানে `slices.Sorted` ব্যবহার করা হয়, যা একটা **sorted** sequence return করে।

> **গুরুত্বপূর্ণ:** `maps` আর `slices` package-গুলো **Go 1.21** এ standard library-তে যোগ হয়েছিল। আগে এগুলো `golang.org/x/exp` মডিউল-এ ছিল।

### Line 9

```go
func main() {
```

Program-এর entry point।

### Lines 10–14

```go
rateLimits := map[string]int{
    "free":       100,
    "pro":        1000,
    "enterprise": 10000,
}
```

`rateLimits` — একটা **map**: `map[string]int`, অর্থাৎ key-টা string (plan-এর নাম) আর value-টা int (requests/hour)। তিনটা plan:

- `"free"` → 100 requests/hour
- `"pro"` → 1000 requests/hour
- `"enterprise"` → 10000 requests/hour

### Line 16

```go
requestedPlan := "pro"
```

`requestedPlan` — client-টা কোন plan-এ আছে, তা সেট করা হয় `"pro"`।

### Line 18

```go
limit, exist := rateLimits[requestedPlan]
```

Map থেকে lookup: `rateLimits["pro"]`। Go-র map lookup **দুটো** value return করে:

- `limit` — value (1000)।
- `exist` — একটা bool: key টা map-এ **আছে কিনা**।

এটা দরকারি কারণ key না থাকলে map value-টার **zero value** (`0`) return করে, যা আসল value-র সাথে গুলিয়ে যেতে পারে — `exist` সেই ambiguity সরিয়ে দেয়।

### Lines 20–23

```go
if !exist {
    fmt.Println("unknown plan:", requestedPlan)
    return
}
```

যদি key-টা না থাকে (`!exist`), তাহলে "unknown plan" print করে `return` দিয়ে program শেষ করে। (এখানে `"pro"` আছে, তাই এটা চলে না।)

### Line 24

```go
fmt.Println(requestedPlan, "plan allows", limit, "requests/hour")
```

প্রশ্ন করা plan-র limit print করে: `pro plan allows 1000 requests/hour`।

### Lines 26–27

```go
rateLimits["pro"] = 1500
delete(rateLimits, "free")
```

Map-টা update হয়:

- `rateLimits["pro"] = 1500` — existing key-এর value পরিবর্তন করে; `"pro"` এখন 1500।
- `delete(rateLimits, "free")` — builtin `delete` দিয়ে `"free"` key-টা (এবং তার value) map থেকে সরিয়ে দেয়। Map-এ এখন সুধু ২টা entry: `pro` আর `enterprise`।

### Lines 29–31

```go
for _, plan := range slices.Sorted(maps.Keys(rateLimits)) {
    fmt.Println(plan, "->", rateLimits[plan])
}
```

Sorted order-এ সব plan print করা হয়:

- `maps.Keys(rateLimits)` — map-এর সব key-র একটা **orderless slice** দেয়: `["enterprise", "pro"]` (কোনো নির্দিষ্ট order নেই)।
- `slices.Sorted(...)` — ওই keys-টাকে **alphabetically sort** করে: `["enterprise", "pro"]`।
- ফলাফলের উপর `range` — sorted order-এ প্রতিটা plan print হয় value-সহ।

**কেন sorted iteration দরকারি:** native Go map-এর iteration order randomized/nondeterministic — sort করতে গেলে প্রতিবার ভিন্ন order আসতে পারে। `slices.Sorted` দিয়ে output **predictable/stabil** হয়।

### Line 32

```go
}
```

Closing brace — `main` function শেষ হয়।

---

## Expected Output

```
pro plan allows 1000 requests/hour
enterprise -> 10000
pro -> 1500
```

## মূল শিক্ষা / Key Takeaways

1. **Map** — `map[string]int` key-value collection; lookup `m[k]`।
2. **`limit, exist := m[k]`** — key আছে কিনা এমন দু-মান lookout; zero-value গুলিয়া এড়াতে।
3. **Map mutation** — `m[k] = v` দিয়ে update, `delete(m, k)` দিয়ে remove।
4. **`maps.Keys` + `slices.Sorted`** — map-এর key-গুলো sorted order-এ iterate করা (Go 1.21+)।
5. **Deterministic output** — sorted iteration map-এর random order-কে নির্ভরযোগ্য বানায়।

---

---

<a name="english"></a>

## 🇬🇧 English Version

### Line 1

```go
package main
```

Declares an executable program (`main` package), runnable via `go run`.

### Lines 3–7

```go
import (
    "fmt"
    "maps"
    "slices"
)
```

Three packages are imported:

- `fmt` — for printing output to the console.
- `maps` — Go 1.21+ standard library. Here we use `maps.Keys`, which returns a slice of all the map's keys (unordered — to be sorted next).
- `slices` — Go 1.21+ standard library. Here we use `slices.Sorted`, which returns a **sorted** sequence.

> **Important:** The `maps` and `slices` packages joined the standard library in **Go 1.21**. Before that they lived in the `golang.org/x/exp` module.

### Line 9

```go
func main() {
```

Program entry point.

### Lines 10–14

```go
rateLimits := map[string]int{
    "free":       100,
    "pro":        1000,
    "enterprise": 10000,
}
```

`rateLimits` — a **map**: `map[string]int` — string keys (plan names) mapping to int values (requests/hour). Three plans:

- `"free"` → 100 requests/hour
- `"pro"` → 1000 requests/hour
- `"enterprise"` → 10000 requests/hour

### Line 16

```go
requestedPlan := "pro"
```

`requestedPlan` — which plan the client is on, set to `"pro"`.

### Line 18

```go
limit, exist := rateLimits[requestedPlan]
```

A map lookup: `rateLimits["pro"]`. A Go map lookup returns **two** values:

- `limit` — the value (1000).
- `exist` — a bool telling you whether the key **exists** in the map.

This matters because absent keys return the map value's **zero value** (`0`), which could be confused with a real value — `exist` removes that ambiguity.

### Lines 20–23

```go
if !exist {
    fmt.Println("unknown plan:", requestedPlan)
    return
}
```

If the key doesn't exist (`!exist`), print "unknown plan" and `return` to end the program. (Here `"pro"` exists, so this doesn't run.)

### Line 24

```go
fmt.Println(requestedPlan, "plan allows", limit, "requests/hour")
```

Prints the limit for the queried plan: `pro plan allows 1000 requests/hour`.

### Lines 26–27

```go
rateLimits["pro"] = 1500
delete(rateLimits, "free")
```

The map is mutated:

- `rateLimits["pro"] = 1500` — changes the value of an existing key; `"pro"` is now 1500.
- `delete(rateLimits, "free")` — removes the `"free"` key (and its value) using the builtin `delete`. The map now holds only two entries: `pro` and `enterprise`.

### Lines 29–31

```go
for _, plan := range slices.Sorted(maps.Keys(rateLimits)) {
    fmt.Println(plan, "->", rateLimits[plan])
}
```

Prints all plans in sorted order:

- `maps.Keys(rateLimits)` — returns an **unordered slice** of all the map's keys: `["enterprise", "pro"]` (no specific order).
- `slices.Sorted(...)` — **sorts** those keys alphabetically: `["enterprise", "pro"]`.
- `range` over the result prints each plan with its value in sorted order.

**Why sorted iteration matters:** native Go map iteration order is randomized/nondeterministic — iterating directly could give different order every run. `slices.Sorted` makes the output **predictable/stable**.

### Line 32

```go
}
```

Closing brace — ends the `main` function.

---

## Expected Output

```
pro plan allows 1000 requests/hour
enterprise -> 10000
pro -> 1500
```

## Key Takeaways

1. **Map** — `map[string]int` key-value collection; lookup with `m[k]`.
2. **`limit, exist := m[k]`** — two-value lookup confirming key existence; avoids zero-value confusion.
3. **Map mutation** — update with `m[k] = v`, remove with `delete(m, k)`.
4. **`maps.Keys` + `slices.Sorted`** — iterate a map's keys in sorted order (Go 1.21+).
5. **Deterministic output** — sorted iteration makes the map's random order reliable.