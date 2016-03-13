package main

import (

    "fmt"
    "time"
    "sync"
)

var w sync.WaitGroup
func do(c chan int, i int) {
    defer w.Done()
    <-c
    fmt.Println("Goroutine....  ", i)
}

func main() {
    fmt.Println(time.Now())
    c := make(chan int)
    for i := 0; i<10000; i++ {
        w.Add(1)
        go do(c, i)
    }
    for i := 0; i<10000; i++ {
        c <- 1
    }

    w.Wait()
    fmt.Println(time.Now())
}
