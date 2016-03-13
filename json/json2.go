package main

import (
    "encoding/json"
    "fmt"
)

var jsonBlob = []byte(`[
    {"Name": "Platypus", "Order": "Monotremata"},
    {"Name": "Quoll",    "Order": "Dasyuromorphia"}
]`)

type Animal struct {
    Name  string
    Order string
}

var animals []Animal

func main() {
    err := json.Unmarshal(jsonBlob, &animals)
    if err != nil {
        fmt.Println("error:", err)
    }
    fmt.Printf("%+v", animals)
    // fmt.Println(animals)

}
