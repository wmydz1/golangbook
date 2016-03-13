package main

import (
	"fmt"
)

func main() {
	m := make(map[string]int)
	m["Answer"] = 42
	fmt.Println("The Values:", m["Answer"])

	m["Answer"] = 48
	fmt.Println("The Values:", m["Answer"])

	delete(m, "Answer")
	fmt.Println("The Value:", m["Answer"])

	v, ok := m["Answer"]
	fmt.Println("The Value:", v, "Present?", ok)
}
