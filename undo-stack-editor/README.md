# undo-stack-editor

Go-তে **snapshot-based undo stack (`[][]string`)** শেখার ছোট example — edit-এর আগে state সেভ করে undo-এর সময় পুনরুদ্ধার।

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
	"errors"
	"fmt"
	"slices"
)
```

- `errors` — `errors.New` sentinel errors।
- `fmt` — `Println`।
- `slices` — `Clone` (snapshot/defensive copy)।

### Lines 9–12

```go
type Editor struct {
	lines   []string
	history [][]string
}
```

- `lines` — current document।
- `history` — **snapshot stack**: `[][]string` — প্রতিটা undoable state-এর copy।

### Lines 14–19

```go
func newEditor() *Editor {
	return &Editor{
		lines:   []string{},
		history: [][]string{},
	}
}
```

Constructor — both fields initialized to non-nil empty slices।

### Lines 21–23

```go
func (e *Editor) snapshot() {
	e.history = append(e.history, slices.Clone(e.lines))
}
```

**Snapshot-before-mutation ডিজাইন:**

- Mutate-র **আগে** বর্তমান lines-এর copy `history`-তে push।
- `slices.Clone` — শেয়ার্ড backing-array এড়ায়; নাহলে পরে `e.lines` বদলালে history-র পুরোনো entries-ও বদলে যেত।

### Lines 25–28

```go
func (e *Editor) Insert(line string) {
	e.snapshot()
	e.lines = append(e.lines, line)
}
```

- First snapshot (পুরনো state সেভ), তারপর append।

### Lines 30–37

```go
func (e *Editor) DeleteLast() error {
	if len(e.lines) == 0 {
		return errors.New("nothing to delete")
	}
	e.snapshot()
	e.lines = e.lines[:len(e.lines)-1]
	return nil
}
```

- Empty-check → `errors.New("nothing to delete")`।
- Snapshot তারপর reslice (`[:len-1]` — শেষ line বাদ)।

### Lines 39–46

```go
func (e *Editor) Undo() error {
	if len(e.history) == 0 {
		return errors.New("nothing to undo")
	}
	e.lines = e.history[len(e.history)-1]
	e.history = e.history[:len(e.history)-1]
	return nil
}
```

**Undo — stack pop:**

- No-history → `errors.New("nothing to undo")`।
- `e.lines = history[last]` — শেষ mutation-এর **আগের** state restore।
- `history = history[:len-1]` — pop।

### Lines 48–50

```go
func (e *Editor) View() []string {
	return slices.Clone(e.lines)
}
```

**Defensive copy return** — caller যদি mutate করে, editor-এর data সেফ থাকে।

### Lines 52–67

```go
func main() {
	ed := newEditor()
	ed.Insert("line 1")
	ed.Insert("line 2")
	ed.Insert("line 3")

	fmt.Println("after inserts:", ed.View())

	if err := ed.DeleteLast(); err != nil {
		fmt.Println("delete error:", err)
	}
	fmt.Println("after delete:", ed.View())
	if err := ed.Undo(); err != nil {
		fmt.Println("undo error:", err)
	}
	fmt.Println("after undo:", ed.View())
}
```

**Trace:**

- Insert l1 → hist=`[[]]`, lines=`[l1]`
- Insert l2 → hist=`[[], [l1]]`, lines=`[l1 l2]`
- Insert l3 → hist=`[[], [l1], [l1, l2]]`, lines=`[l1 l2 l3]` → `after inserts: [line 1 line 2 line 3]`
- DeleteLast → snapshot `[l1 l2 l3]`, lines=`[l1 l2]` → `after delete`
- Undo → lines = hist last = `[l1 l2 l3]` → `after undo`

---

## Expected Output

```
after inserts: [line 1 line 2 line 3]
after delete: [line 1 line 2]
after undo: [line 1 line 2 line 3]
```

## মূল শিক্ষা / Key Takeaways

1. **Snapshot stack** — `[][]string` history।
2. **Clone before mutate** — shared-backing hazard এড়ানো।
3. **Undo = pop** — শুধু last state restore।
4. **Defensive copy** — `View()` return-এ Clone।
5. **Sentinel errors** — `errors.New` empty-case।

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
	"errors"
	"fmt"
	"slices"
)
```

- `errors` — for `errors.New`.
- `fmt` — for `Println`.
- `slices` — for `Clone` (snapshots / defensive copies).

### Lines 9–12

```go
type Editor struct {
	lines   []string
	history [][]string
}
```

- `lines` — the current document.
- `history` — a **snapshot stack**: `[][]string` — one copy per undoable state.

### Lines 14–19

```go
func newEditor() *Editor {
	return &Editor{
		lines:   []string{},
		history: [][]string{},
	}
}
```

A constructor — both fields start as non-nil empty slices.

### Lines 21–23

```go
func (e *Editor) snapshot() {
	e.history = append(e.history, slices.Clone(e.lines))
}
```

**Snapshot-before-mutation design:**

- Pushes a copy of the current lines onto `history` **before** a mutation.
- `slices.Clone` — avoids the shared-backing-array hazard; otherwise later changes to `e.lines` would corrupt the already-stored history entries.

### Lines 25–28

```go
func (e *Editor) Insert(line string) {
	e.snapshot()
	e.lines = append(e.lines, line)
}
```

- Snapshot first (save the old state), then append.

### Lines 30–37

```go
func (e *Editor) DeleteLast() error {
	if len(e.lines) == 0 {
		return errors.New("nothing to delete")
	}
	e.snapshot()
	e.lines = e.lines[:len(e.lines)-1]
	return nil
}
```

- Empty-check → `errors.New("nothing to delete")`.
- Snapshot, then reslice (`[:len-1]` — drop the last line).

### Lines 39–46

```go
func (e *Editor) Undo() error {
	if len(e.history) == 0 {
		return errors.New("nothing to undo")
	}
	e.lines = e.history[len(e.history)-1]
	e.history = e.history[:len(e.history)-1]
	return nil
}
```

**Undo — a stack pop:**

- No history → `errors.New("nothing to undo")`.
- `e.lines = history[last]` — restores the state **before** the last mutation.
- `history = history[:len-1]` — pops the entry.

### Lines 48–50

```go
func (e *Editor) View() []string {
	return slices.Clone(e.lines)
}
```

**Defensive return** — if the caller mutates it, the editor's data stays safe.

### Lines 52–67

```go
func main() {
	ed := newEditor()
	ed.Insert("line 1")
	ed.Insert("line 2")
	ed.Insert("line 3")

	fmt.Println("after inserts:", ed.View())

	if err := ed.DeleteLast(); err != nil {
		fmt.Println("delete error:", err)
	}
	fmt.Println("after delete:", ed.View())
	if err := ed.Undo(); err != nil {
		fmt.Println("undo error:", err)
	}
	fmt.Println("after undo:", ed.View())
}
```

**Trace:**

- Insert l1 → hist=`[[]]`, lines=`[l1]`
- Insert l2 → hist=`[[], [l1]]`, lines=`[l1 l2]`
- Insert l3 → hist=`[[], [l1], [l1, l2]]`, lines=`[l1 l2 l3]` → `after inserts: [line 1 line 2 line 3]`
- DeleteLast → snapshot `[l1 l2 l3]`, lines=`[l1 l2]` → `after delete`
- Undo → lines = hist last = `[l1 l2 l3]` → `after undo`

---

## Expected Output

```
after inserts: [line 1 line 2 line 3]
after delete: [line 1 line 2]
after undo: [line 1 line 2 line 3]
```

## Key Takeaways

1. **Snapshot stack** — the `[][]string` history.
2. **Clone before mutate** — avoiding the shared-backing hazard.
3. **Undo = pop** — just restore the latest state.
4. **Defensive copy** — `Clone` on the `View()` return.
5. **Sentinel errors** — `errors.New` for the empty cases.