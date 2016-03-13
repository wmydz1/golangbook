package main

import (

    "os"
    "io"
    "fmt"
)

var all string

func readSingle() {
    f, err := os.Open("./io/test.txt")
    if err !=nil {
        fmt.Println(err.Error())
        os.Exit(1)
    }
    defer f.Close()
    reader := io.LimitReader(f, 1024)
    p := make([]byte, 1024)
    var total int
    for {
        n, err := reader.Read(p)
        if err == io.EOF {
            break
        }
        total =total+n
    }
    fmt.Println(string(p[:total]))
}
func readMulti() {
    f1, err := os.Open("./io/test.txt")
    f2, err2 := os.Open("./io/spring.txt")
    if err !=nil {
        fmt.Println(err.Error())
        os.Exit(1)
    }
    if err2 !=nil {
        fmt.Println(err2.Error())
        os.Exit(1)
    }
    reader := io.MultiReader(f1, f2)
    p := make([]byte, 1024 *5)
    var total int
    var data string
    for {
        n, err := reader.Read(p)
        if err == io.EOF {
            fmt.Println("read over ", total)
            break
        }
        total =total+n
        data =data+string(p[:n])
    }
    fmt.Println("value is ", data)
    fmt.Println("count is", total)

    all =data

}
func writeMuti() {
    f1, _ := os.Create("./io/all1.txt")
    f2, _ := os.Create("./io/all2.txt")
    writer := io.MultiWriter(f1, f2)
    writer.Write([]byte(all))
}
func main() {
    //    readSingle()
    readMulti()
    writeMuti()
}
