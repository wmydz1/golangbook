package main
import (

    "fmt"
    "time"
)
var flag = make(chan bool)

func do() {
    fmt.Println("do something.....")
    flag <- true
//    close(flag)
}
func main() {
    go do()
//    <-flag
    time.Sleep(4*time.Second)
}
