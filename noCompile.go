package main

import "fmt"

func init() {
	panic("not allowed to import common by itself, import individual packages")
}

func main() {
	fmt.Println("common > noCompile.go > main(). This should panic")
}

// type noCompile struct
