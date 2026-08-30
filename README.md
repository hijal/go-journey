# go-journey

My daily Go practice. Each folder is a small, self-contained example from my learning.

## Run an example

```bash
go run ./payment-tracker
```

## Examples

| Folder                           | Concept                                                   |
| -------------------------------- | --------------------------------------------------------- |
| `api-backend`                    | Struct + constructor method                               |
| `background-jobs`                | Loop, error handling, map lookup                          |
| `cart-stock`                     | Case-sensitive variables (`quantity` vs `Quantity`)       |
| `concurrent-webhook`             | Goroutine + channel + WaitGroup                           |
| `config-loader`                  | Builtin `len` shadowing                                   |
| `devops-config-loader`           | Package-level vars, const, funcs                          |
| `double-entry-ledger`            | Methods with pointer receiver, insufficient balance error |
| `env-variable-parsing`           | `strconv` parsing with defaults                           |
| `fingerprint-config`             | SHA-256 checksum + string/byte immutability               |
| `invoice-batch-scanner`          | `continue` / `break` in loop                              |
| `invoice-json`                   | JSON marshal + unexported field                           |
| `inventory-stock`                | Unsigned underflow + safe compare                         |
| `interget-overflow`              | Integer overflow (`int32` vs `int64`)                     |
| `ledger-entry`                   | Custom type (`Money`)                                     |
| `money-conversion-testing`       | Table-driven unit test + `math.Round`                     |
| `monitoring-agent`               | Float precision + epsilon compare                         |
| `multi-gateway-payment-proccess` | Interface + polymorphism                                  |
| `order-status-enum`              | `iota` enum + `String()` method                           |
| `payment-amount`                 | Custom type (`PaymentAmount`) + paisa-based money         |
| `payment-fee-calculator`         | Constant + function                                       |
| `payment-risk-flagging`          | `switch` with `fallthrough`                               |
| `payment-status-poll`            | `select` + channel timeout                                |
| `payment-tracker`                | Variable naming / underscores                             |
| `plan-based-rate-limit`          | Maps + `slices`/`maps` stdlib (Go 1.21+)                  |
| `rune-byte`                      | Byte vs rune, UTF-8 decoding                              |
| `settlement-reconciliation`      | `goto` retry + labeled `break`                            |
| `testing`                        | `go/token.IsIdentifier` unit test                         |
| `transaction-rollback`           | `defer` for rollback pattern                              |

## Notes

- Go version: 1.26
