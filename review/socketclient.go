package main
import (

    "net"
    "fmt"
    "io/ioutil"
)

func main() {
    service := ":10000"
    tcpaddr, err := net.ResolveTCPAddr("tcp4", service)
    if err !=nil {
        fmt.Println("can't resolve tcp addr ", err)
    }
    conn, err := net.DialTCP("tcp", nil, tcpaddr)
    if err!=nil {
        fmt.Println("can't dial addr", err)
    }
    _, err =conn.Write([]byte("HEAD / HTTP/1.0\r\n\r\n"))

    result, err := ioutil.ReadAll(conn)
    if err!=nil {
        fmt.Println("read conn err ", err)
    }
    fmt.Println(string(result))

}
