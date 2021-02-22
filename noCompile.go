package main

import "fmt"

func init() {
	fmt.Println("noCompile.go > init(). Should panic...")
	panic("not allowed to import common by itself, import individual packages")
}

func main() {
	fmt.Println("common > noCompile.go > main(). This should panic")
}

type NoCompile struct {}
