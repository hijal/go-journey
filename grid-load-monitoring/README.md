# grid-load-monitoring

Go-তে **fixed-size array + parallel array (index mapping)** শেখার ছোট example — সাপ্তাহিক পাওয়ার গ্রিড load-এর (MW) total, average, peak, low track করা।

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

`fmt` — `Printf`।

### Lines 5–6

```go
func main() {
	weeklyLoad := [7]float64{412.5, 398.2, 441.7, 460.1, 455.9, 380.4, 372.8}
	days := [7]string{"Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"}
```

দুটো **fixed-size array**, same length (7):

- `weeklyLoad` — প্রতিটা দিনের পাওয়ার load (MW)।
- `days` — দিনের নাম।

**Parallel array pattern:** প্রতিটা index-এ এক দিনের load + day name ম্যাপ করা রয়েছে (`weeklyLoad[3]` = Thu-র load 460.1)। Struct-এ না ভেঙে দুটো array — আলাদা units (float + string) একসাথে।

### Lines 8–10

```go
	total := 0.0
	peakDay, lowDay := 0, 0
```

- `total` — sum accumulator।
- `peakDay`, `lowDay` — সর্বোচ্চ/সর্বনিম্ন load-এর **index** (day)। দুটো-ই `0` থেকে শুরু — প্রথম element baseline ধরে, নিজের সাথে তুলনা false, তাই ঠিকই first-pass-এ আপডেট হয়।

### Lines 12–22

```go
	for day, load := range weeklyLoad {
		total += load

		if load > weeklyLoad[peakDay] {
			peakDay = day
		}

		if load < weeklyLoad[lowDay] {
			lowDay = day
		}
	}
```

এক-pass scan:

- `for day, load := range weeklyLoad` — index `day` + value `load`।
- `total += load` — accumulate।
- `load > weeklyLoad[peakDay]` — বর্তমান peak-এর চেয়ে বড় হলে peakDay আপডেট। **Index-ভিত্তিক tracking** — load value না, দিনের index save।
- `load < weeklyLoad[lowDay]` — তেমনি low-র জন্য।
- এক-এক scan-এ total + peak + low সব একসাথে।

### Lines 24–27

```go
	fmt.Printf("weekly total: %.1f MW\n", total)
	fmt.Printf("average: %.2f MW\n", total/float64(len(weeklyLoad)))
	fmt.Printf("peak: %s %.1f MW\n", days[peakDay], weeklyLoad[peakDay])
	fmt.Printf("low: %s %.1f MW\n", days[lowDay], weeklyLoad[lowDay])
```

- `weekly total` — `%.1f` এক দশমিক।
- `average` — `total / float64(len(weeklyLoad))` — `int → float64` cast (float/int invalid)।
- `peak`/`low` — saved index দিয়ে `days[i]` (নাম) + `weeklyLoad[i]` (মান)।

---

## Expected Output

```
weekly total: 2921.6 MW
average: 417.37 MW
peak: Thu 460.1 MW
low: Sun 372.8 MW
```

Verification: `412.5+398.2+441.7+460.1+455.9+380.4+372.8 = 2921.6`; `avg = 2921.6/7 = 417.37`। Peak = Thu (`460.1`), low = Sun (`372.8`)।

## মূল শিক্ষা / Key Takeaways

1. **Fixed-size array** — `[7]float64`, size type-এর অংশ।
2. **Parallel arrays** — same-index-ভিত্তিক দুটো array (load ↔ day name)।
3. **Index tracking** — value-এর বদলে index-save করে peak/low, পরে name+value দুটো-ই access।
4. **`float64()` cast** — `total/float64(len(...))`।
5. **One-pass scan** — total, peak, low একসাথে।

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

`fmt` — for `Printf`.

### Lines 5–6

```go
func main() {
	weeklyLoad := [7]float64{412.5, 398.2, 441.7, 460.1, 455.9, 380.4, 372.8}
	days := [7]string{"Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"}
```

Two **fixed-size arrays**, same length (7):

- `weeklyLoad` — the grid load per day (MW).
- `days` — the day names.

**Parallel array pattern:** each index pairs one day's load with its day name (`weeklyLoad[3]` is Thu's load, 460.1). Two arrays instead of a struct — separate units (float + string) side by side.

### Lines 8–10

```go
	total := 0.0
	peakDay, lowDay := 0, 0
```

- `total` — a sum accumulator.
- `peakDay`, `lowDay` — the **index** of the highest/lowest load. Both start at `0` — the first element is the baseline; comparing it with itself is false, so the updates work naturally on the first pass.

### Lines 12–22

```go
	for day, load := range weeklyLoad {
		total += load

		if load > weeklyLoad[peakDay] {
			peakDay = day
		}

		if load < weeklyLoad[lowDay] {
			lowDay = day
		}
	}
```

One-pass scan:

- `for day, load := range weeklyLoad` — index `day` + value `load`.
- `total += load` — accumulate.
- `load > weeklyLoad[peakDay]` — if larger than the current peak, update `peakDay`. **Index-based tracking** — the day's index is stored, not the load value.
- `load < weeklyLoad[lowDay]` — same for the low.
- Total, peak and low all come out of one same scan.

### Lines 24–27

```go
	fmt.Printf("weekly total: %.1f MW\n", total)
	fmt.Printf("average: %.2f MW\n", total/float64(len(weeklyLoad)))
	fmt.Printf("peak: %s %.1f MW\n", days[peakDay], weeklyLoad[peakDay])
	fmt.Printf("low: %s %.1f MW\n", days[lowDay], weeklyLoad[lowDay])
```

- `weekly total` — `%.1f`, one decimal.
- `average` — `total / float64(len(weeklyLoad))` — `int → float64` cast (float/int is invalid).
- `peak`/`low` — saved index to read `days[i]` (name) + `weeklyLoad[i]` (value).

---

## Expected Output

```
weekly total: 2921.6 MW
average: 417.37 MW
peak: Thu 460.1 MW
low: Sun 372.8 MW
```

Verification: `412.5+398.2+441.7+460.1+455.9+380.4+372.8 = 2921.6`; `avg = 2921.6/7 = 417.37`. Peak = Thu (`460.1`), low = Sun (`372.8`).

## Key Takeaways

1. **Fixed-size array** — `[7]float64`, the size is part of the type.
2. **Parallel arrays** — two arrays joined by the same index (load ↔ day name).
3. **Index tracking** — store indexes (not values) for peak/low, then read both name and value.
4. **`float64()` cast** — `total/float64(len(...))`.
5. **One-pass scan** — total, peak and low in a single loop.