# find-max-num

Go-তে **variadic parameter** (`...int`) আর explicit int-এ loop দিয়ে সর্বোচ্চ সংখ্যা খুঁজে শেখার ছোট example — `<` comparison + accumulation।

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

### Lines 5–16

```go
func FindMaxNum(nums ...int) int {
	if len(nums) == 0 {
		return 0
	}
	max := nums[0]
	for _, num := range nums {
		if num > max {
			max = num
		}
	}
	return max
}
```

`FindMaxNum` — variadic int list-এর মধ্যে সর্বোচ্চ:

- `nums ...int` — **variadic**: একাধিক int argument এক `[]int`-এ।
- `if len(nums) == 0` — **guard**: খালি slice হলে `0` return (না হলে `nums[0]` panic করত — index out of range)।
- `max := nums[0]` — প্রাথমিক maximum প্রথম element।
- Loop-এ প্রতিটা `num` যদি `num > max` হয়, `max`-কে update।
- পরে `return max`।

**Note:** `max := nums[0]` (variable-name `max`) — এখানে scalar int-এ সিলেকশন। (`min`/`max` builtin-তেও ব্যবহার-যোগ্য, কিন্তু এখানে explicit loop দেখানো হয়েছে।)

### Line 18

```go
func main() {
```

Program-এর entry point।

### Line 19

```go
fmt.Println(FindMaxNum(10, 50, 30, 90, 20))
```

`FindMaxNum(10, 50, 30, 90, 20)` — সবচেয়ে বড়টা 90। Result print হয়: `90`।

### Line 20

```go
}
```

Closing brace — `main` function শেষ হয়।

---

## Expected Output

```
90
```

## মূল শিক্ষা / Key Takeaways

1. **Variadic `...int`** — একাধিক int argument।
2. **Guard clause** — `len(nums) == 0` → 0 (panic এড়ানো)।
3. **Accumulation pattern** — `max`-কে track করে প্রতিটা element-এর সাথে compare।
4. **`var` shadowing** — local variable-এর নাম (`max`)।

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

### Lines 5–16

```go
func FindMaxNum(nums ...int) int {
	if len(nums) == 0 {
		return 0
	}
	max := nums[0]
	for _, num := range nums {
		if num > max {
			max = num
		}
	}
	return max
}
```

`FindMaxNum` — finds the largest among a variadic int list:

- `nums ...int` — **variadic**: multiple int arguments as one `[]int`.
- `if len(nums) == 0` — **guard**: for an empty slice return `0` (otherwise `nums[0]` would panic with index out of range).
- `max := nums[0]` — the initial maximum is the first element.
- In the loop, each `num` that's `> max` updates `max`.
- Then `return max`.

**Note:** the local variable is named `max` (a scalar selection). (`min`/`max` builtins could be used, but here an explicit loop is shown.)

### Line 18

```go
func main() {
```

Program entry point.

### Line 19

```go
fmt.Println(FindMaxNum(10, 50, 30, 90, 20))
```

`FindMaxNum(10, 50, 30, 90, 20)` — the largest is `90`. It prints `90`.

### Line 20

```go
}
```

Closing brace — ends the `main` function.

---

## Expected Output

```
90
```

## Key Takeaways

1. **Variadic `...int`** — multiple int arguments.
2. **Guard clause** — `len(nums) == 0` → 0 (avoids panic).
3. **Accumulation pattern** — track `max`, comparing each element.
4. **Local naming** — a local variable named `max`.