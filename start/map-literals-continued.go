package main

import (
	"fmt"
)

type Vertex struct {
	Lat, Long float64
}

var m = map[string]Vertex{
	"Bell Labs": {34.2324, 32.4242},
	"Google":    {42.2322, 242.4222424},
}

func main() {
	fmt.Println(m)
}
