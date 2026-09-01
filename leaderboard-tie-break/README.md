# leaderboard-tie-break

Go-তে **`cmp.Compare`** (`cmp` package), **nested `switch`** আর **`switch {}` (expression-less)** statement দিয়ে leaderboard tie-break শেখার ছোট example।

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

### Lines 3–6

```go
import (
	"cmp"
	"fmt"
)
```

- `cmp` — `cmp.Compare` (Go 1.21+) দিয়ে generic ordering।
- `fmt` — output print।

### Line 8

```go
func main() {
```

Program-এর entry point।

### Lines 9–10

```go
yourScore, yourLevel := 1500, 9
theirScore, theirLevel := 1500, 7
```

দুটো player-এর score/level:

- আপনি: score 1500, level 9
- প্রতিপক্ষ: score 1500, level 7

Score **equal** (1500 == 1500) — তাই tie-break-এ level কথা বলবে।

### Line 12

```go
rank := cmp.Compare(yourScore, theirScore)
```

**`cmp.Compare(a, b)`** — generic ordering function, return করে:
- `-1` যদি `a < b`
- `0` যদি `a == b`
- `+1` যদি `a > b`

`cmp.Compare(1500, 1500)` = `0` — score equal।

### Lines 14–29

```go
switch {
case rank < 0:
	fmt.Println("they are ahead on score")
case rank > 0:
	fmt.Println("you are ahead on score")
default:
	levelRank := cmp.Compare(yourLevel, theirLevel)
	switch {
	case levelRank < 0:
		fmt.Println("they win on level tie-break")
	case levelRank > 0:
		fmt.Println("you win on level tie-break")
	default:
		fmt.Println("perfect tie - share first place")
	}
}
```

**`switch {}` (expression-less)** — কোন expression-এর উপর নয়, প্রতিটা `case` একটা **condition/boolean expression** evaluates করে; প্রথমটা যে true হয় সেটাই চলে:

- `rank < 0`? rank=0, তাই না।
- `rank > 0`? rank=0, তাই না।
- `default` — score equal-এর **tie-break** branch।

Tie-break-এর ভেতরে **nested `switch {}`** আবার level-এ:

- `levelRank := cmp.Compare(yourLevel, theirLevel)` = `cmp.Compare(9, 7)` = `+1` (আপনার level বেশি)
- `levelRank < 0`? না।
- `levelRank > 0`? **হ্যাঁ** → `you win on level tie-break`
- `default` — উভয় level equal হলে `perfect tie`.

**কেন `cmp.Compare`:** `<` / `>` operator চেয়ে readable — একটা function দিয়ে compare pattern; বিশেষ করে generic/compareable-type-এ।

### Line 30

```go
}
```

Closing brace — `main` function শেষ হয়।

---

## Expected Output

```
you win on level tie-break
```

## মূল শিক্ষা / Key Takeaways

1. **`cmp.Compare`** — `a<b` → -1, `a==b` → 0, `a>b` → 1।
2. **`switch {}`** — expression ছাড়া switch; প্রতিটা case boolean condition।
3. **Nested switch** — একটা switch-এর `default`-এ আরেকটা switch।
4. **Tie-break logic** — primary key (score) equal হলে secondary key (level) decide করে।
5. **`case` order** — প্রথম matching case-ই চলে; `default` সবশেষে।

---

---

<a name="english"></a>

## 🇬🇧 English Version

### Line 1

```go
package main
```

Declares an executable program (`main` package), runnable via `go run`.

### Lines 3–6

```go
import (
	"cmp"
	"fmt"
)
```

- `cmp` — `cmp.Compare` (Go 1.21+) for generic ordering.
- `fmt` — for output.

### Line 8

```go
func main() {
```

Program entry point.

### Lines 9–10

```go
yourScore, yourLevel := 1500, 9
theirScore, theirLevel := 1500, 7
```

Two players' score/level:

- You: score 1500, level 9
- Opponent: score 1500, level 7

The scores are **equal** (1500 == 1500), so level will decide the tie-break.

### Line 12

```go
rank := cmp.Compare(yourScore, theirScore)
```

**`cmp.Compare(a, b)`** — a generic ordering function returning:
- `-1` if `a < b`
- `0` if `a == b`
- `+1` if `a > b`

`cmp.Compare(1500, 1500)` = `0` — the scores are equal.

### Lines 14–29

```go
switch {
case rank < 0:
	fmt.Println("they are ahead on score")
case rank > 0:
	fmt.Println("you are ahead on score")
default:
	levelRank := cmp.Compare(yourLevel, theirLevel)
	switch {
	case levelRank < 0:
		fmt.Println("they win on level tie-break")
	case levelRank > 0:
		fmt.Println("you win on level tie-break")
	default:
		fmt.Println("perfect tie - share first place")
	}
}
```

**`switch {}` (expression-less)** — not switching on a value; each `case` evaluates a **condition (boolean expression)**, and the first true one runs:

- `rank < 0`? rank=0, no.
- `rank > 0`? rank=0, no.
- `default` — the **tie-break** branch when scores are equal.

Inside the tie-break, a **nested `switch {}`** again evaluates on level:

- `levelRank := cmp.Compare(yourLevel, theirLevel)` = `cmp.Compare(9, 7)` = `+1` (your level is higher)
- `levelRank < 0`? no.
- `levelRank > 0`? **yes** → `you win on level tie-break`
- `default` — if both levels are equal, `perfect tie`.

**Why `cmp.Compare`:** more readable than a raw `<` / `>` comparison — a single function expresses the compare pattern, especially for generic/comparable types.

### Line 30

```go
}
```

Closing brace — ends the `main` function.

---

## Expected Output

```
you win on level tie-break
```

## Key Takeaways

1. **`cmp.Compare`** — `a<b` → -1, `a==b` → 0, `a>b` → 1.
2. **`switch {}`** — a switch without an expression; each case is a boolean condition.
3. **Nested switch** — a second switch inside the `default` of another.
4. **Tie-break logic** — when the primary key (score) is equal, a secondary key (level) decides.
5. **`case` order** — the first matching case runs; `default` is last.