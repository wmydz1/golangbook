package main
import (


    "fmt"
)
func sum(a []int, c chan int) {
    total := 0
    for _, v := range a {
        total+=v
    }
    c <- total
//    close(c)
}

func main() {
    a := []int{1, 2, 3, 4, 5, 6, 7, 8}
    c := make(chan int)
    //    go sum(a[:len(a)/2], c)
    go sum(a[len(a)/2:], c)

    //    for v := range c {
    //        fmt.Println(v)
    //    }
    fmt.Println(<-c)
}
/*
for i := range c能够不断的读取channel里面的数据，直到该channel被显式的关闭
*/