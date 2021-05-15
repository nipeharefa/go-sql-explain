package main

import (
	"fmt"
	sqlexplain "sql-explain"
)

func main() {
	a := sqlexplain.Explain("SELECT * FROM users where email = ? and username = ?", "nipe@nias.dev", 1)
	fmt.Println(a)
}
