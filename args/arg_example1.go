package main
import (
    "fmt"
    "os"
)
func main() {
    fmt.Println(os.Args)
    fmt.Println(os.Args[1])
}
//可见返回了一个三个元素的数组，第０个元素是程序的名字包括路径，os.Args就第一个参数，os.Args就是第二个参数。