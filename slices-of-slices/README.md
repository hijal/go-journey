# slices-of-slices

Go-তে **nested slices** — Tic-Tac-Toe board, `[][]string` matrix।

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
	"fmt"
	"strings"
)
```

`strings` (Join)।

### Lines 8–23

```go
func main() {
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}
}
```

- `[][]string` — slice-এর slice: প্রতিটি element-ই আরেকটা slice।
- Double-index assign → `board[row][col]`।
- `strings.Join` → সারিগুলো space-যুক্ত এক লাইন।

---

## Expected Output

```
X _ X
O _ X
_ _ O
```

## মূল শিক্ষা / Key Takeaways

1. **`[][]T`** — nested structure।
2. **Two-index access** — `board[1][2]`।
3. **`strings.Join`** — tokens → string।
4. **Row-by-row print** — board rendering।

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
	"fmt"
	"strings"
)
```

`strings` (Join).

### Lines 8–23

```go
func main() {
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}
}
```

- `[][]string` — a slice of slices: each element is itself a slice.
- Double-index assignment → `board[row][col]`.
- `strings.Join` → joins each row's tokens into one line.

---

## Expected Output

```
X _ X
O _ X
_ _ O
```

## Key Takeaways

1. **`[][]T`** — a nested structure.
2. **Two-index access** — `board[1][2]`.
3. **`strings.Join`** — tokens → a string.
4. **Row-by-row print** — renders the board.