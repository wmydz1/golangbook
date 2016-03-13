package main
import (

    "fmt"
    "time"
    "os"
)
/*
select是Go语言特有的操作，使用select我们可以同时在多个channel上进行发送/接收操作。下面是select的基本操作。
*/
func main() {
    ch := make(chan string,1)
    ch <- "Go"
    go func() {
        select {
        case x, ok := <-ch:
            if !ok {
                fmt.Println("error")
            }
            fmt.Println(x)
        case <-time.After(3*time.Second):
            fmt.Println("time out....")
            os.Exit(2)
        }
    }()

    time.Sleep(5*time.Second)
    os.Exit(1)


}
