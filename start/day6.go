package main

import (
	"fmt"
)

func add(x int, y int) int {
	return x + y
}
func main() {
	fmt.Println(add(10000, 2999))
}