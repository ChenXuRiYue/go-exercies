package main

import (
	"fmt"
)

func testPointer() {
	var a int = 10
	var b *int
	b = &a
	fmt.Println("a =", a)
	fmt.Println("b =", b)
	fmt.Println("*b =", *b)
	fmt.Println("&a =", &a)
	fmt.Println("&b =", &b)
}