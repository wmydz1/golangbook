package main
import (
    "reflect"
    "fmt"
)

type MyStruct struct {
    N int
}
func main(){
    n :=MyStruct{100}

    // get
    immutable :=reflect.ValueOf(n)
    val :=immutable.FieldByName("N").Int()
    fmt.Printf("N=%d\n",val)

    // set
    mutable :=reflect.ValueOf(&n).Elem()
    mutable.FieldByName("N").SetInt(10000)
    fmt.Printf("N=%d\n",n.N)
}
