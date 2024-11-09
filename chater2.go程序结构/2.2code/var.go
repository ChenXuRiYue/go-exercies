package main

import (
	"fmt"
)

const demo1 = "demo1"

func funcDemo(s string, t string) string {
	return s + t
}

func main() {
	var demo2 string = "demo2"
	fmt.Println("hello go", demo1, demo2)
	fmt.Println(funcDemo(demo1, demo2))
}
