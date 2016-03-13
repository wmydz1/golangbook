package main

import (
	"fmt"
)

type Vertex struct {
	Lat, Long float64
}

var m = map[string]Vertex{
	"Bell Labs": Vertex{
		40.321213, -873.323232,
	},
	"Google": Vertex{
		34.23224432, -24233.32323,
	},
}

func main() {
	fmt.Println(m)
}
