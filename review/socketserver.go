package main
import (

    "net"
    "fmt"
    "time"
)


func main() {

    service := ":10000"

    tcpaddr, err := net.ResolveTCPAddr("tcp4", service)

    if err !=nil {
        fmt.Println("can't resolve tcp addr ", err)
    }

    listener, err := net.ListenTCP("tcp", tcpaddr)

    if err !=nil {
        fmt.Println("can't listen.... ", err)
    }

    for {
        conn, err := listener.Accept()
        if err!=nil {
            continue
        }
        go handleclient(conn)

    }
}
func handleclient(conn net.Conn) {
    defer conn.Close()
    daytime := time.Now().String()
    conn.Write([]byte(daytime))
}