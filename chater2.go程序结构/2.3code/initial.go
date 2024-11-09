package main

import (
	"fmt"
)

func main() {
	var i, j, k int = 1, 2, 3
	fmt.Println(i, j, k)

	first, second := 1, 2
	fmt.Println(first, second)
	first, second = second, first
	fmt.Println(first, second)

}
