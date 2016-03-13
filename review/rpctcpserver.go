package main
import (

    "errors"
    "net/rpc"
    "net"
    "fmt"
)
type Args struct {
    A, B int
}
type Quotient struct {
    Quo, Rem int
}
type Arith int

func (t *Arith) Multiply(args *Args, reply *int) error {
    *reply =args.A *args.B
    return nil
}

func (t *Arith) Divide(args *Args, quo *Quotient) error {
    if args.B ==0 {
        return errors.New("divide by zero")
    }
    quo.Quo =args.A / args.B
    quo.Rem =args.A % args.B

    return nil
}

func main() {
    arith := new(Arith)
    rpc.Register(arith)
    service := ":10000"
    tcpaddr, err := net.ResolveTCPAddr("tcp", service)
    if err !=nil {
        fmt.Println(err)
    }

    listener, err := net.ListenTCP("tcp", tcpaddr)

    if err !=nil {
        fmt.Println("listen err", err)
    }
    for {
        conn, err := listener.Accept()
        if err!=nil {
            continue
        }
        go rpc.ServeConn(conn)
    }

}
