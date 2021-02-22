package db

import (
	"fmt"
	// gotest "github.com/rahulkhairwar/go-test"
	"github.com/rahulkhairwar/go-test/sql"
	gotestutil "github.com/rahulkhairwar/go-test/util"
)

func TestImportSql() {
	sql.TestForSQLV0_0_3()
	sql.Test2()
}

func SomeFuncForV0_0_2() {
	fmt.Println("db > util.go > SomeFuncForV0_0_2()")
}

func SomeFuncForV0_0_3() {
	fmt.Println("db > util.go > SomeFuncForV0_0_3()")
}

func SomeFuncForV0_0_5() {
	fmt.Println("db > util.go > SomeFuncForV0_0_5()")
}

func TestNoCompile() {
	fmt.Println("TestNoCompile() ")
	// gotest.NoCompile{}
	gotestutil.SomeUtilFunc1()
}
