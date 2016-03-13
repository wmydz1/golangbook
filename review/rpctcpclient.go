package main
import (

    "net/rpc"
    "fmt"
)
type Args struct {

    A, B int

}



type Quotient struct {

    Quo, Rem int

}
func main() {
    service := ":10000"
    client, err := rpc.Dial("tcp", service)
    if err !=nil {
        fmt.Println("dial tcp error ", err)
    }
    args := Args{17, 8}
    var reply int
    err =client.Call("Arith.Multiply", args, &reply)
    if err !=nil {
        fmt.Println("call error ", err)
    }
    fmt.Printf("%d * %d = %d\n ", args.A, args.B, reply)

    var quot Quotient

    err =client.Call("Arith.Divide", args, &quot)
    if err!=nil {
        fmt.Println("call divide error", err)
    }

    fmt.Printf("%d / %d = %d remainder %d\n", args.A, args.B, quot.Quo, quot.Rem)
}
