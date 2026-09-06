# sent-notification

Go-তে **variadic function** + **slice spread** (`...`) শেখার ছোট example -- dynamic number of channels-এ notification পাঠানো.

**ভাষা নির্বাচন করুন / Choose language:**

[বাংলা](#bangla) * [English](#english)

---

<a name="bangla"></a>

## বাংলা সংস্করণ

### Lines 1-6

```go
package main

import (
    "fmt"
    "strings"
)
```

- `package main` -- executable program.
- `fmt` -- output format করতে.
- `strings` -- slice কে comma-separated string-এ join করতে.

### Lines 8-14

```go
func sendAlert(recipient string, channels ...string) string {
    if len(channels) == 0 {
        channels = []string{"email"}
    }

    return fmt.Sprintf("alert -> %s via %s", recipient, strings.Join(channels, ", "))
}
```

Variadic function:

- `channels ...string` -- যেকোনো সংখ্যক channel নেয়, function-এর ভিতরে `[]string` slice.
- কোনো channel না দিলে default `["email"]` set করা হয়.
- `fmt.Sprintf` দিয়ে formatted alert message return করে.

### Lines 16-22

```go
func main() {
    fmt.Println(sendAlert("oncall@team.dev", "sms", "slack"))
    fmt.Println(sendAlert("oncall@team.dev"))

    allChannels := []string{"pagerduty", "sms", "slack", "email"}
    fmt.Println(sendAlert("oncall@team.dev", allChannels...))
}
```

তিনটা call:

- 2 channels explicit: sms, slack.
- কোনো channel নেই: default email.
- slice spread (`allChannels...`): slice-কে variadic argument-এ unpack.

---

## Expected Output

```
alert -> oncall@team.dev via sms, slack
alert -> oncall@team.dev via email
alert -> oncall@team.dev via pagerduty, sms, slack, email
```

## মূল শিক্ষা / Key Takeaways

1. **Variadic + default value** -- কোনো argument না দিলে sensible default বসানো.
2. **Slice spread operator** -- `allChannels...` দিয়ে slice-কে variadic-এ unpack করা যায়.
3. **`strings.Join`** -- slice-কে delimiter দিয়ে string-এ join করা.
4. **`fmt.Sprintf` formatting** -- placeholders দিয়ে structured string তৈরি.

---

---

<a name="english"></a>

## English Version

### Lines 1-6

```go
package main

import (
    "fmt"
    "strings"
)
```

- `package main` -- an executable program.
- `fmt` -- for formatting output.
- `strings` -- to join a slice into a comma-separated string.

### Lines 8-14

```go
func sendAlert(recipient string, channels ...string) string {
    if len(channels) == 0 {
        channels = []string{"email"}
    }

    return fmt.Sprintf("alert -> %s via %s", recipient, strings.Join(channels, ", "))
}
```

Variadic function:

- `channels ...string` -- accepts any number of channels; inside the function it becomes a `[]string` slice.
- when no channel is supplied, defaults to `["email"]`.
- returns a formatted alert message via `fmt.Sprintf`.

### Lines 16-22

```go
func main() {
    fmt.Println(sendAlert("oncall@team.dev", "sms", "slack"))
    fmt.Println(sendAlert("oncall@team.dev"))

    allChannels := []string{"pagerduty", "sms", "slack", "email"}
    fmt.Println(sendAlert("oncall@team.dev", allChannels...))
}
```

Three calls:

- 2 explicit channels: sms, slack.
- no channels: defaults to email.
- slice spread (`allChannels...`): a slice is unpacked into the variadic argument.

---

## Expected Output

```
alert -> oncall@team.dev via sms, slack
alert -> oncall@team.dev via email
alert -> oncall@team.dev via pagerduty, sms, slack, email
```

## Key Takeaways

1. **Variadic + default value** -- provide a sensible default when no arguments are passed.
2. **Slice spread operator** -- `allChannels...` unpacks a slice into variadic arguments.
3. **`strings.Join`** -- joins a slice into a delimited string.
4. **`fmt.Sprintf` formatting** -- builds a structured string with placeholders.
