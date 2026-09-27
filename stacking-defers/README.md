# stacking-defers

Go-তে **deferred calls LIFO stacking** — function-এর শেষে সব `defer` উল্টো-ক্রমে চালায়।

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

`fmt` (print)।

### Lines 5–12

```go
func main() {
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}
```

**Stacked defers** — প্রতিটি `defer` LAST-IN-FIRST-OUT:

- "counting" → 10টা defer জমা (i=0..9) → "done"।
- Function শেষে defer-গুলো **উল্টো ক্রমে**: 9, 8, 7, ... 0।

---

## Expected Output

```
counting
done
9
8
7
6
5
4
3
2
1
0
```

## মূল শিক্ষা / Key Takeaways

1. **LIFO order** — শেষ-নিবন্ধিত defer আগে চলে।
2. **Defer execution now** — args function-এ যাওয়ার মুহূর্তে evaluate হয়।
3. **Use-case** — resource cleanup-এর সাধারণ আদল।
4. **No guarantee of runtime** — defer শেষে-ই, কোথাও না।

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

`fmt` (print).

### Lines 5–12

```go
func main() {
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}
```

**Stacked defers** — each `defer` is last-in-first-out:

- "counting" → 10 defers accumulate (i=0..9) → "done".
- At function exit the defers run **in reverse**: 9, 8, 7, ... 0.

---

## Expected Output

```
counting
done
9
8
7
6
5
4
3
2
1
0
```

## Key Takeaways

1. **LIFO order** — the last-registered defer runs first.
2. **Arguments evaluate now** — at the moment the `defer` statement is reached.
3. **Use-case** — the typical shape of resource cleanup.
4. **Timing** — defers always run at the end of the function.