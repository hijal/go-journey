# go-journey

My daily Go practice. Each folder is a small, self-contained example from my learning.

## Run an example

```bash
go run ./payment-tracker
```

## Examples

| Folder | Concept | Docs |
| --- | --- | --- |
| `api-backend` | Struct + constructor method | [README](./api-backend/README.md) |
| `attendance` | `min`/`max` builtins + overtime split | [README](./attendance/README.md) |
| `background-jobs` | Loop, error handling, map lookup | [README](./background-jobs/README.md) |
| `basic-types` | Basic types + rune as numeric code point | [README](./basic-types/README.md) |
| `bill-split` | Float constants + explicit cast + `%.2f` | [README](./bill-split/README.md) |
| `byte-size-constants` | iota + bit shift for KB/MB/GB constants | [README](./byte-size-constants/README.md) |
| `cart-stock` | Case-sensitive variables (`quantity` vs `Quantity`) | [README](./cart-stock/README.md) |
| `checkout-cart` | Float arithmetic + type conversion | [README](./checkout-cart/README.md) |
| `concurrent-webhook` | Goroutine + channel + WaitGroup | [README](./concurrent-webhook/README.md) |
| `config-loader` | Builtin `len` shadowing | [README](./config-loader/README.md) |
| `devops-config-loader` | Package-level vars, const, funcs | [README](./devops-config-loader/README.md) |
| `devops-tooling` | Multiple assignment + `strconv.Atoi` | [README](./devops-tooling/README.md) |
| `double-entry-ledger` | Methods with pointer receiver, insufficient balance error | [README](./double-entry-ledger/README.md) |
| `env-variable-parsing` | `strconv` parsing with defaults | [README](./env-variable-parsing/README.md) |
| `fingerprint-config` | SHA-256 checksum + string/byte immutability | [README](./fingerprint-config/README.md) |
| `fixed-tax-discount` | Constants + `fmt.Printf` formatting | [README](./fixed-tax-discount/README.md) |
| `free-shipping` | Boolean logic (`||` `&&` `!`) | [README](./free-shipping/README.md) |
| `group-variable` | Grouped `const`/`iota` + grouped `var` | [README](./group-variable/README.md) |
| `http-status-code` | Unexported constants + `switch` for status mapping | [README](./http-status-code/README.md) |
| `invoice-batch-scanner` | `continue` / `break` in loop | [README](./invoice-batch-scanner/README.md) |
| `invoice-json` | JSON marshal + unexported field | [README](./invoice-json/README.md) |
| `inventory-stock` | Unsigned underflow + safe compare | [README](./inventory-stock/README.md) |
| `interget-overflow` | Integer overflow (`int32` vs `int64`) | [README](./interget-overflow/README.md) |
| `ip-subnet` | Bitwise ops + IP subnet membership | [README](./ip-subnet/README.md) |
| `ledger-entry` | Custom type (`Money`) | [README](./ledger-entry/README.md) |
| `leaderboard-tie-break` | `cmp.Compare` + nested switch tie-break | [README](./leaderboard-tie-break/README.md) |
| `money-conversion-testing` | Table-driven unit test + `math.Round` | [README](./money-conversion-testing/README.md) |
| `monitoring-agent` | Float precision + epsilon compare | [README](./monitoring-agent/README.md) |
| `multi-gateway-payment-proccess` | Interface + polymorphism | [README](./multi-gateway-payment-proccess/README.md) |
| `new-account` | Zero values (`string`/`int64`/`bool`) | [README](./new-account/README.md) |
| `order-status-enum` | `iota` enum + `String()` method | [README](./order-status-enum/README.md) |
| `package-local-variable` | Package-level vs local variable scope | [README](./package-local-variable/README.md) |
| `pagination` | Ceiling division + offset pagination | [README](./pagination/README.md) |
| `payment-amount` | Custom type (`PaymentAmount`) + paisa-based money | [README](./payment-amount/README.md) |
| `payment-fee-calculator` | Constant + function | [README](./payment-fee-calculator/README.md) |
| `payment-risk-flagging` | `switch` with `fallthrough` | [README](./payment-risk-flagging/README.md) |
| `payment-status-poll` | `select` + channel timeout | [README](./payment-status-poll/README.md) |
| `payment-tracker` | Variable naming / underscores | [README](./payment-tracker/README.md) |
| `plan-based-rate-limit` | Maps + `slices`/`maps` stdlib (Go 1.21+) | [README](./plan-based-rate-limit/README.md) |
| `rate-limit-config` | Exported vs unexported constants + unit suffix naming | [README](./rate-limit-config/README.md) |
| `rune-byte` | Byte vs rune, UTF-8 decoding | [README](./rune-byte/README.md) |
| `retry-backoff-timing` | Linear backoff + `time.Duration` constants | [README](./retry-backoff-timing/README.md) |
| `retry-policy` | Retry loop + backoff with `time.Sleep` | [README](./retry-policy/README.md) |
| `settlement-reconciliation` | `goto` retry + labeled `break` | [README](./settlement-reconciliation/README.md) |
| `testing` | `go/token.IsIdentifier` unit test | [README](./testing/README.md) |
| `transaction-rollback` | `defer` for rollback pattern | [README](./transaction-rollback/README.md) |
| `unix-permission-bits` | Bit flags (rwx permission bitmask) | [README](./unix-permission-bits/README.md) |

## Notes

- Go version: 1.26
