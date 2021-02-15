package db

import (
	"fmt"
	"github.com/rahulkhairwar/go-test/sql"
)

func TestImportSql() {
	sql.TestForSQLV0_0_3()
}

func SomeFuncForV0_0_2() {
	fmt.Println("db > util.go > SomeFuncForV0_0_2()")
}
