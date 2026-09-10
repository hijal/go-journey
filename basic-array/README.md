# basic-array

Go-তে **fixed-size array** (`[4]int`) শেখার ছোট example — ৪ ধরণের `range` loop, array copy semantics, 2D array, pointer-to-array, slice-as-view, আর `%v`/`%+v`/`%#v`/`%T` formatting।

**📖 ভাষা নির্বাচন করুন / Choose language:**

[🇧🇩 বাংলা](#bangla) • [🇬🇧 English](#english)

---

<a name="bangla"></a>

## 🇧🇩 বাংলা সংস্করণ

### Lines 1–5

```go
package main

import (
	"fmt"
)
```

একটা executable program (`main` package) declare করে, যা `go run` দিয়ে চালানো যায়। `fmt` — `Println`, `Printf`-এর জন্য।

### Lines 7–11

```go
type Coordinate [2]float64

type Point struct {
	x, y int
}
```

- `Coordinate` — `[2]float64`-এর একটা **named type**: array-সাইজ type-এর অংশ, তাই `[2]float64` আর `[4]float64` ভিন্ন type।
- `Point` — struct: unexported field `x, y int`।

### Lines 13–14

```go
func main() {
	a := [4]int{5, 10, 15, 20}
```

- `[4]int{...}` — **fixed-size array**: ঠিক 4টা element। Slice-এর মতো dynamic না; size type-এর অংশ।
- Element: `5, 10, 15, 20`।

### Lines 16–18

```go
	for i, v := range a {
		fmt.Println(i, v)
	}
```

১ম loop — **index + value**: `range a`-তে প্রতিটা iteration-এ `i` (index) আর `v` (value copy) → `0 5`, `1 10`, ...।

### Lines 19–23

```go
	fmt.Println("-----")
	for i := range a {
		a[i] *= 2
		fmt.Println(a[i])
	}
```

২য় loop — **index-only + in-place mutation**: `for i := range a` শুধু index; `a[i] *= 2` element double (`5→10, 10→20, 15→30, 20→40`)। Fixed-size বলে সরাসরি mutate করা যায়।

### Lines 24–27

```go
	fmt.Println("-----")
	for _, v := range a {
		fmt.Println(v)
	}
```

৩য় loop — **value-only**: `_` দিয়ে index discard, doubled values print: `10, 20, 30, 40`।

### Lines 28–31

```go
	fmt.Println("-----")
	for i := range len(a) {
		fmt.Println(i)
	}
```

৪র্থ loop — **int-range (Go 1.22+)**: `for i := range len(a)` — `0` থেকে `len(a)-1`, এখানে `0 1 2 3`।

### Lines 33–46

```go
	b := a

	b[0] = 100
	b[1] = 200
	b[2] = 300
	b[3] = 400
	fmt.Println("-----")
	for i, v := range b {
		fmt.Println(i, v)
	}
	fmt.Println("-----")
	for i, v := range a {
		fmt.Println(i, v)
	}
```

**Array copy semantics** — সবচেয়ে গুরুত্বপূর্ণ অংশ:

- `b := a` — array **assignment-এ পুরোটা copy** হয় (slice-এর মতো reference/header নয়)।
- `b`-র element বদলানোতেই → `b` এখন `100, 200, 300, 400`।
- কিন্তু `a` **অপরিবর্তিত** — এখনো `10, 20, 30, 40`। দুটো আলাদা data।
- **Slice-এর সাথে তুলনা:** slice-এ `s2 := s1` করলে একটা পিছনে ধাক্কা — data share করে। Array-তে প্রতিটা assignment-এ copy।

### Lines 47–50

```go
	fmt.Println("-----")
	var board [3][3]string
	board[1][1] = "X"
	fmt.Println(len(board), len(board[0]))
```

**2D array**: `[3][3]string` — 3×3 grid, zero values `""` দিয়ে initialize। `board[1][1] = "X"` center-এ set। `len(board)` (3) = row, `len(board[0])` (3) = column → `3 3`।

### Lines 52–55

```go
	var pickup Coordinate
	pickup[0] = 37.7749
	pickup[1] = -122.4194
	fmt.Println(pickup)
```

Named array type ব্যবহার: `Coordinate`-র zero value, তারপর index-এ lat/lng set → `[37.7749 -122.4194]`।

### Lines 57–61

```go
	arr := [2]int{1, 2}
	p := &arr
	p[0] = 10
	p[1] = 20
	fmt.Println(arr)
```

**Pointer to array**: `p := &arr` — array-র pointer; `p[0]` dereference ছাড়াই element mutate (Go auto-deref)। `arr` বদলে যায় → `[10 20]`। Pointer-এ copy হয় না — direct।

### Lines 63–68

```go
	arr1 := [4]int{1, 2, 3, 4}

	s := arr1[:]
	s[0] = 99
	fmt.Println(arr1)
	fmt.Println(s)
```

**Slice as view**: `s := arr1[:]` — array-র উপরে একটা slice (same backing memory)। `s[0] = 99` → `arr1`-তেও দেখায় → `[99 2 3 4]`, `s`-ও `[99 2 3 4]`। এখানে mutation shared — array/pointer-এর সাথে contrast।

### Lines 70–75

```go
	pt := Point{x: 1, y: 2}
	fmt.Printf("value(%v): %v\n", pt, pt)
	fmt.Printf("Plus (%+v):   %+v\n", pt, pt)
	fmt.Printf("Syntax (%#v): %#v\n", pt, pt)
	fmt.Printf("Type (%T):    %T\n", pt, pt)
```

**Formatting verbs** — struct-এ:

- `%v` — default: `{1 2}`
- `%+v` — field name সহ: `{x:1 y:2}`
- `%#v` — Go-সyntax: `main.Point{x:1, y:2}`
- `%T` — type name: `main.Point`

---

## Expected Output

```
0 5
1 10
2 15
3 20
-----
10
20
30
40
-----
10
20
30
40
-----
0
1
2
3
-----
0 100
1 200
2 300
3 400
-----
0 10
1 20
2 30
3 40
-----
3 3
[37.7749 -122.4194]
[10 20]
[99 2 3 4]
[99 2 3 4]
value({1 2}): {1 2}
Plus ({x:1 y:2}):   {x:1 y:2}
Syntax (main.Point{x:1, y:2}): main.Point{x:1, y:2}
Type (main.Point):    main.Point
```

## মূল শিক্ষা / Key Takeaways

1. **Fixed-size array** — `[4]int`, size type-এর অংশ (slice নয়)।
2. **Array = value** — `b := a` পুরো copy; `a` অপরিবর্তিত।
3. **Pointer-to-array** — `p := &arr` এড়িয়ে গেলেও mutate হয়।
4. **Slice = view** — `arr1[:]` পিছনে memory share; mutation দুই জায়গায়।
5. **2D array** — `[3][3]string`, `len` per dimension।
6. **Format verbs** — `%v`/`%+v`/`%#v`/`%T`।
7. **Int-range** — `for i := range len(a)` (Go 1.22+)।

---

---

<a name="english"></a>

## 🇬🇧 English Version

### Lines 1–5

```go
package main

import (
	"fmt"
)
```

Declares an executable program (`main` package), runnable via `go run`. `fmt` — for `Println`, `Printf`.

### Lines 7–11

```go
type Coordinate [2]float64

type Point struct {
	x, y int
}
```

- `Coordinate` — a **named type** for `[2]float64`: the array size is part of the type, so `[2]float64` and `[4]float64` are different types.
- `Point` — a struct: unexported fields `x, y int`.

### Lines 13–14

```go
func main() {
	a := [4]int{5, 10, 15, 20}
```

- `[4]int{...}` — a **fixed-size array**: exactly 4 elements. Unlike a slice, the size is part of the type.
- Elements: `5, 10, 15, 20`.

### Lines 16–18

```go
	for i, v := range a {
		fmt.Println(i, v)
	}
```

1st loop — **index + value**: `range a` gives `i` (index) and `v` (a copy of the value) each iteration → `0 5`, `1 10`, ....

### Lines 19–23

```go
	fmt.Println("-----")
	for i := range a {
		a[i] *= 2
		fmt.Println(a[i])
	}
```

2nd loop — **index-only + in-place mutation**: `for i := range a` gives only the index; `a[i] *= 2` doubles the element (`5→10, 10→20, 15→30, 20→40`). Fixed size allows direct mutation.

### Lines 24–27

```go
	fmt.Println("-----")
	for _, v := range a {
		fmt.Println(v)
	}
```

3rd loop — **value-only**: `_` discards the index, printing the doubled values: `10, 20, 30, 40`.

### Lines 28–31

```go
	fmt.Println("-----")
	for i := range len(a) {
		fmt.Println(i)
	}
```

4th loop — **int-range (Go 1.22+)**: `for i := range len(a)` — from `0` to `len(a)-1`, here `0 1 2 3`.

### Lines 33–46

```go
	b := a

	b[0] = 100
	b[1] = 200
	b[2] = 300
	b[3] = 400
	fmt.Println("-----")
	for i, v := range b {
		fmt.Println(i, v)
	}
	fmt.Println("-----")
	for i, v := range a {
		fmt.Println(i, v)
	}
```

**Array copy semantics** — the most important part:

- `b := a` — assigning an array **copies the whole thing** (not a reference/header like a slice).
- Mutating `b` → `b` becomes `100, 200, 300, 400`.
- But `a` is **unchanged** — still `10, 20, 30, 40`. Two separate data sets.
- **Compare with slices:** `s2 := s1` makes them share backing data. Array assignment always copies.

### Lines 47–50

```go
	fmt.Println("-----")
	var board [3][3]string
	board[1][1] = "X"
	fmt.Println(len(board), len(board[0]))
```

**2D array**: `[3][3]string` — a 3×3 grid, initialized with zero values (`""`). `board[1][1] = "X"` sets the center. `len(board)` (3) = rows, `len(board[0])` (3) = columns → `3 3`.

### Lines 52–55

```go
	var pickup Coordinate
	pickup[0] = 37.7749
	pickup[1] = -122.4194
	fmt.Println(pickup)
```

Using the named array type: `Coordinate`'s zero value, then index-based lat/lng set → `[37.7749 -122.4194]`.

### Lines 57–61

```go
	arr := [2]int{1, 2}
	p := &arr
	p[0] = 10
	p[1] = 20
	fmt.Println(arr)
```

**Pointer to array**: `p := &arr` — a pointer to the array; `p[0]` mutates the element without explicit dereference (Go auto-derefs). `arr` changes → `[10 20]`. No copy through a pointer — direct.

### Lines 63–68

```go
	arr1 := [4]int{1, 2, 3, 4}

	s := arr1[:]
	s[0] = 99
	fmt.Println(arr1)
	fmt.Println(s)
```

**Slice as view**: `s := arr1[:]` — a slice over the array (same backing memory). `s[0] = 99` shows up in `arr1` → `[99 2 3 4]`, and `s` is also `[99 2 3 4]`. Mutation is shared here — the contrast to arrays/pointers.

### Lines 70–75

```go
	pt := Point{x: 1, y: 2}
	fmt.Printf("value(%v): %v\n", pt, pt)
	fmt.Printf("Plus (%+v):   %+v\n", pt, pt)
	fmt.Printf("Syntax (%#v): %#v\n", pt, pt)
	fmt.Printf("Type (%T):    %T\n", pt, pt)
```

**Formatting verbs** — on a struct:

- `%v` — default: `{1 2}`
- `%+v` — with field names: `{x:1 y:2}`
- `%#v` — Go syntax: `main.Point{x:1, y:2}`
- `%T` — the type name: `main.Point`

---

## Expected Output

```
0 5
1 10
2 15
3 20
-----
10
20
30
40
-----
10
20
30
40
-----
0
1
2
3
-----
0 100
1 200
2 300
3 400
-----
0 10
1 20
2 30
3 40
-----
3 3
[37.7749 -122.4194]
[10 20]
[99 2 3 4]
[99 2 3 4]
value({1 2}): {1 2}
Plus ({x:1 y:2}):   {x:1 y:2}
Syntax (main.Point{x:1, y:2}): main.Point{x:1, y:2}
Type (main.Point):    main.Point
```

## Key Takeaways

1. **Fixed-size array** — `[4]int`, the size is part of the type (not a slice).
2. **Array = value** — `b := a` copies everything; `a` is unchanged.
3. **Pointer-to-array** — `p := &arr` mutates without a copy.
4. **Slice = view** — `arr1[:]` shares backing memory; mutations show up in both.
5. **2D array** — `[3][3]string`, `len` per dimension.
6. **Format verbs** — `%v`/`%+v`/`%#v`/`%T`.
7. **Int-range** — `for i := range len(a)` (Go 1.22+).