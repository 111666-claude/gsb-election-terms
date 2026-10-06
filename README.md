# election-terms

简化选举：按任期收票、过半当选、高任期接替低任期。只用标准库。

```
go test ./...
go vet ./...
go run ./cmd/election --sample=dupvote
go run ./cmd/election --sample=majority
go run ./cmd/election --sample=stale
go run ./cmd/election --sample=reject
go run ./cmd/election --sample=work
```

## 口径

- **一票**：同一个投票者在同一个任期最多记一票；重复投同一任期不再计数。
- **任期单调**：投票者见过更高任期后，低任期的投票请求一律不接受。
- **过半**：得票数严格大于投票者总数的一半才当选；恰好一半不算。
- **接替**：更高任期产生的当选者替换当前当选者。

## 不变量

- 每个投票者在每个任期最多贡献一票。
- 当选者的任期不降低。
- `scanned` 不随票数平方级放大：两千张票的 `scanned` 不超过 6000。

## 输出契约

场景打印一行 JSON，键随场景而定：`tally` 是某任期的得票映射，`leader` 是当前当选者，`accepted` 是最近一次投票是否被接受。
`--sample=work` 打印 `{"votes": N, "scanned": N}`。
