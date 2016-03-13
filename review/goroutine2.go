package main
import (

    "fmt"
)

var c = make(chan int)

func worker(i int) {
    fmt.Println("worker is doing ... ", i)
    c <- 1
}
func main() {

    for i := 0; i<30000; i++ {
        go worker(i)
    }

    for i := 0; i<30000; i++ {
        <-c
    }


}
