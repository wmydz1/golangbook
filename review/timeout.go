package main
import (

    "fmt"
    "time"
)

func timeout(c chan int, o chan bool) {

    for {
        select {
        case v := <-c:
            fmt.Println(v)
        case <-time.After(5 * time.Second):
            fmt.Println("timeout...")
            o <- true
            break

        }
    }
}

func main() {
    c := make(chan int)
    o := make(chan bool)

    go timeout(c, o)

//    c <- 10000
    <-o
}
