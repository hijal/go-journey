# background-jobs

Go-তে **`for range` loop**, **error-handling** আর **map lookup** বুঝতে সহায়ক ছোট example।
(A small example for understanding **`for range` loops**, **error handling** and **map lookups** in Go.)

---

## Line-by-line Explanation

### Line 1

```go
package main
```

**English:** Declares an executable program (`main` package), runnable via `go run`.

**বাংলা:** একটা executable program (`main` package) declare করে, যা `go run` দিয়ে চালানো যায়।

---

### Line 3

```go
import "fmt"
```

**English:** Imports the `fmt` package for console output.

**বাংলা:** Console-এ output print করার জন্য `fmt` package import করা হয়।

---

### Lines 5–10

```go
func processJob(jobID string) (bool, error) {
    if jobID == "" {
        return false, fmt.Errorf("empty job id")
    }
    return true, nil
}
```

**English:** Defines a function `processJob` that takes a `jobID` (string) and returns two values: `bool` (success flag) and `error`. This is Go's idiomatic error-handling style — functions report problems through an `error` return value.
- If `jobID` is empty, we build an error with `fmt.Errorf("empty job id")` and return `false` plus that error.
- Otherwise, we return `true` and `nil` (meaning "no error").

**বাংলা:** `processJob` নামক function-টা define করে, যা একটা `jobID` (string) নেয় এবং দুটো value return করে: `bool` (success flag) আর `error`। এটা Go-র idiomatic error-handling style — function-গুলো সমস্যা `error` return value দিয়ে জানায়।
- `jobID` খালি হলে `fmt.Errorf("empty job id")` দিয়ে একটা error তৈরি করে `false` আর ঐ error return করি।
- নাহলে `true` আর `nil` (মানে "কোনো error নেই") return করি।

---

### Line 12

```go
func main() {
```

**English:** Program entry point; `{` opens the body.

**বাংলা:** Program-এর entry point; `{` দিয়ে body শুরু হয়।

---

### Lines 13–17

```go
jobIDs := []string{
    "job-101",
    "job-102",
    "job-103",
}
```

**English:** Declares `jobIDs`, a `[]string` slice holding three job identifiers we want to process.

**বাংলা:** `jobIDs` declare করা হয়, একটা `[]string` slice যা process করতে চাওয়া তিনটা job identifier ধারণ করে।

---

### Line 19

```go
for _, id := range jobIDs {
```

**English:** The `for ... range` loop iterates over each element of the slice. `range` yields the index and the element; here we use `_` (blank identifier) to discard the index since we don't need it, and `id` gets each element on every iteration.

**বাংলা:** `for ... range` loop slice-এর প্রতিটা element-এর উপর iterate করে। `range` index আর element এ দুটো দেয়; এখানে `_` (blank identifier) দিয়ে index discard করি যেহেতু দরকার নেই, আর প্রতি iteration-এ `id` প্রতিটা element পায়।

---

### Line 20

```go
ok, err := processJob(id)
```

**English:** Calls `processJob(id)` and captures both return values into `ok` (bool) and `err` (error) using `:=`.

**বাংলা:** `processJob(id)` call করে `:=` দিয়ে দুটো return value `ok` (bool) আর `err` (error)-এ capture করি।

---

### Lines 22–25

```go
if err != nil {
    fmt.Println("Failed:", err)
    continue
}
```

**English:** If `err` is not `nil` (meaning an error happened), we print `Failed:` with the error message, then use `continue` to skip the rest of the loop body and go to the next job. In this example no job IDs are empty, so this branch won't run.

**বাংলা:** যদি `err` `nil` না হয় (মানে error হয়েছে), তাহলে error message-সহ `Failed:` print করি, তারপর `continue` দিয়ে বাকি loop body skip করে পরের job-এ যাই। এই example-এ কোনো job ID খালি নেই, তাই এই branch চালবে না।

---

### Line 26

```go
fmt.Println("Processed:", id, " success:", ok)
```

**English:** If there was no error, we print the job id and the success flag. Output for each: `Processed: job-101 success: true` and so on.

**বাংলা:** Error না থাকলে job id আর success flag print করি। প্রতিটির output: `Processed: job-101 success: true` ইত্যাদি।

---

### Line 27

```go
}
```

**English:** Closing brace — ends the `for` loop.

**বাংলা:** Closing brace — `for` loop শেষ হয়।

---

### Lines 29–32

```go
statusCount := map[string]int{
    "done":   3,
    "failed": 0,
}
```

**English:** Declares a `map` — a key-value collection. Here the keys are strings (`"done"`, `"failed"`) and the values are ints. It says 3 jobs are done and 0 failed.

**বাংলা:** একটা `map` declare করা হয় — একটা key-value collection। এখানে key-গুলো string (`"done"`, `"failed"`) আর value-গুলো int। মানে ৩টা job done আর ০টা failed।

---

### Lines 34–36

```go
if _, exists := statusCount["retrying"]; !exists {
    fmt.Println("No jobs are currently retrying")
}
```

**English:** When you read from a map like `statusCount["retrying"]`, Go returns **two** values: the value (if any) and a bool `exists` telling you whether the key exists. Here we discard the value with `_` and keep `exists`.
This line uses Go's **`if` with an initializer** — the map lookup runs first (setting `exists`), then the condition `!exists` (meaning the key is NOT present) is checked. Since there is no `"retrying"` key, `exists` is `false`, so `!exists` is `true`, and the message is printed.
Useful because a missing key returns the zero value (`0`), which can be confused with a real value — the `exists` bool removes that ambiguity.

**বাংলা:** Map থেকে `statusCount["retrying"]` পড়লে Go **দুটো** value return করে: value (যদি থাকে) আর একটা bool `exists` যা বলে key-টা আছে কিনা। এখানে `_` দিয়ে value discard করি আর `exists` রাখি।
এই লাইনটা Go-র **`if` with an initializer** ব্যবহার করে — আগে map lookup চলে (`exists` set হয়), তারপর `!exists` (মানে key-টা **নাই**) শর্তটা পরীক্ষা করা হয়। যেহেতু `"retrying"` key-টা নাই, `exists` হলো `false`, তাই `!exists` হলো `true`, এবং বার্তাটা print হয়।
কারণটা দরকারি: key না থাকলে map zero value (`0`) return করে, যেটা real value-র সাথে গুলিয়ে যেতে পারে — `exists` bool সেই ambiguity দূর করে।

---

### Line 37

```go
}
```

**English:** Closing brace — ends the `main` function.

**বাংলা:** Closing brace — `main` function শেষ হয়।

---

## Expected Output

```
Processed: job-101  success: true
Processed: job-102  success: true
Processed: job-103  success: true
No jobs are currently retrying
```

## Key Takeaways / মূল শিক্ষা

1. **`(bool, error)` return** — function error-কে return value দিয়ে জানায়। / Functions report errors via an `error` return value.
2. **`if err != nil`** — Go-র সবচেয়ে সাধারণ error-check idiom। / The most common error-checking idiom in Go.
3. **`continue`** — বাকি loop body skip করে পরের iteration-এ যায়। / Skips the rest of the loop body and moves to the next iteration.
4. **`for ... range`** — slice/array-র প্রতিটা item-এ loop। / Loops over each item of a slice/array.
5. **Map lookup `_, exists := m[k]`** — key আছে কিনা সেটা confirms। / Confirms whether a key exists in the map.
6. **`if` with initializer** — `if x := ...; cond` — আগে initializer, তারপর শর্ত। / Runs the initializer first, then checks the condition.
