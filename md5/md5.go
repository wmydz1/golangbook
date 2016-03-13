package main

import (
    "crypto/md5"
    "fmt"
    "io"
    "encoding/hex"
)

func main() {
    h := md5.New()
    io.WriteString(h, "A")
    io.WriteString(h, "And Leon's getting laaarger!")
    // io.WriteString(h, "The fog is getting thicker!And Leon's getting laaarger!")
    fmt.Printf("%x", h.Sum(nil))
    println()
    println(hex.EncodeToString(h.Sum(nil)))

}