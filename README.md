SQL-Explain
===========

```go

query := "SELECT * FROM user_account where email = ?"
args := []interface{}{"food@example.com"}
queryDebug := Explain(query, args...)
```

### Benchmark

```
goos: darwin
goarch: amd64
pkg: sql-explain
BenchmarkExplainSQL-8                           10310656               114 ns/op              80 B/op          2 allocs/op
BenchmarkExplainSQLWithArg-8                     7353726               164 ns/op             128 B/op          2 allocs/op
BenchmarkExplainSQLWith2Arg-8                    5246095               226 ns/op             176 B/op          2 allocs/op
BenchmarkExplainSQLWithManyManyArgs-8            2439931               500 ns/op             368 B/op          2 allocs/op
PASS
ok      sql-explain     6.103s
```