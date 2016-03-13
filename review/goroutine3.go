package main
import (

    "fmt"
    "runtime"
)

func say(s string) {
    for i := 0; i<100; i++ {
        runtime.Gosched()
        fmt.Println(s, i)
    }
}

func main() {
    go say("golang ")
    say("hello ")
}
