package main
import (

    "fmt"
    "sync"
    "runtime"
)
/*
协同多个Goroutines
同上，close channel还可以用于协同多个Goroutines，比如下面这个例子，我们创建了100个Worker Goroutine，
这些Goroutine在被创建出来后都阻塞在"<-start"上，
直到我们在main goroutine中给出开工的信号："close(start)"，这些goroutines才开始真正的并发运行起来。
*/

var wg sync.WaitGroup
func worker(start chan bool, index int) {
    defer wg.Done()//任务完成，将任务队列中的任务数量-1，其实.Done就是.Add(-1)
    <-start
    fmt.Println("The worker is ", index)
}

func main() {
    runtime.GOMAXPROCS(runtime.NumCPU())
    start := make(chan bool)
    for i := 0; i<1000; i++ {
        wg.Add(1)//每创建一个goroutine，就把任务队列中任务的数量+1
        go worker(start, i)
    }

//        for i := 0; i<1000; i++ {
//            start <- true
//        }
    close(start)
    wg.Wait()//.Wait()这里会发生阻塞，直到队列中所有的任务结束就会解除阻塞
}