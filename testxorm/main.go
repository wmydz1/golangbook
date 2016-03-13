package main
import (
    . "gowebprogram/xorm"
    "fmt"

)

func main() {
     err := MakeTransfer(2, 4, 100.1)
    if err !=nil {
        fmt.Println("err",err)
    }

}
