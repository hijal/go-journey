# linked-list

Go-তে **singly-linked list + pointer-traversal append + `Stringer` interface** শেখার ছোট example — manual linked list।

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

`fmt` — `Println`, `Sprintf`।

### Lines 5–8

```go
type Node struct {
	Value int
	Next  *Node
}
```

**Self-referential struct** — প্রতিটি `Node`-এর পরের node-pointer (`Next`), শেষটা `nil`।

### Lines 10–12

```go
type LinkedList struct {
	Head *Node
}
```

শুধু `Head` ধরে রাখা — গোটা chain তার মাধ্যমে।

### Lines 14–29

```go
func (l *LinkedList) Append(value int) {
	newNode := &Node{Value: value}

	if l.Head == nil {
		l.Head = newNode
		return
	}

	current := l.Head

	for current.Next != nil {
		current = current.Next
	}

	current.Next = newNode
}
```

**Tail-append:**

- Empty-list: `Head`-ই `newNode`।
- নাহলে **pointer walk**: `current.Next != nil` পর্যন্ত forward; শেষ node-র `Next`-এ নতুন node।
- `*LinkedList` receiver — মূল list mutate।

### Lines 31–38

```go
func (l *LinkedList) String() string {
	result := ""

	for current := l.Head; current != nil; current = current.Next {
		result += fmt.Sprintf("%d -> ", current.Value)
	}
	return result + "nil"
}
```

**Stringer method** — walking `Head` থেকে প্রতিটি value জোড়া দেয়, শেষে `nil`।

### Lines 40–47

```go
func main() {
	list := &LinkedList{}

	list.Append(10)
	list.Append(20)
	list.Append(30)

	fmt.Println(list)
}
```

`fmt.Println(list)` — স্বয়ংক্রিয়ভাবে `String()` call করে (`fmt.Stringer` interface — type-এ `String() string` থাকলে print-এ সেটা ব্যবহার হয়)।

---

## Expected Output

```
10 -> 20 -> 30 -> nil
```

## মূল শিক্ষা / Key Takeaways

1. **Self-referential pointer struct** — `Next *Node`।
2. **Pointer traversal** — `for current.Next != nil` tail-walk।
3. **Append-to-tail** — শেষ node-এ attach।
4. **`Stringer` interface** — `String() string` auto-used by `fmt.Print*`।
5. **Pointer receiver** — list mutate।

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

`fmt` — for `Println`, `Sprintf`.

### Lines 5–8

```go
type Node struct {
	Value int
	Next  *Node
}
```

A **self-referential struct** — each `Node` holds a pointer to the next (`Next`); the last one is `nil`.

### Lines 10–12

```go
type LinkedList struct {
	Head *Node
}
```

Only `Head` is kept; the whole chain hangs off it.

### Lines 14–29

```go
func (l *LinkedList) Append(value int) {
	newNode := &Node{Value: value}

	if l.Head == nil {
		l.Head = newNode
		return
	}

	current := l.Head

	for current.Next != nil {
		current = current.Next
	}

	current.Next = newNode
}
```

**Tail-append:**

- Empty list: `Head` becomes `newNode`.
- Otherwise a **pointer walk**: advance while `current.Next != nil`; attach the new node to the last node's `Next`.
- `*LinkedList` receiver — mutates the real list.

### Lines 31–38

```go
func (l *LinkedList) String() string {
	result := ""

	for current := l.Head; current != nil; current = current.Next {
		result += fmt.Sprintf("%d -> ", current.Value)
	}
	return result + "nil"
}
```

A **Stringer method** — walking from `Head`, pairing each value, ending with `nil`.

### Lines 40–47

```go
func main() {
	list := &LinkedList{}

	list.Append(10)
	list.Append(20)
	list.Append(30)

	fmt.Println(list)
}
```

`fmt.Println(list)` — `String()` is called automatically (the `fmt.Stringer` interface: any type with `String() string` is used when printed).

---

## Expected Output

```
10 -> 20 -> 30 -> nil
```

## Key Takeaways

1. **Self-referential pointer struct** — `Next *Node`.
2. **Pointer traversal** — the `for current.Next != nil` tail-walk.
3. **Append-to-tail** — attaching at the last node.
4. **`Stringer` interface** — `String() string` auto-used by `fmt.Print*`.
5. **Pointer receiver** — lists mutate in place.