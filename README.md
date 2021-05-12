SQL-Explain
===========

```go

query := "SELECT * FROM user_account where email = ?"
args := []interface{}{"food@example.com"}
queryDebug := Explain(query, args...)
```