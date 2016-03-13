package main
import (
    "strings"
    "fmt"
)

func main() {
    str := "welcome \n to \n golang.com"
    // 去除空格
    str =strings.Replace(str, " ", "", -1)
    // 去除换行符
    str =strings.Replace(str, "\n", "", -1)
    fmt.Println(str)
}
