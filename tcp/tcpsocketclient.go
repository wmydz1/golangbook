package main
import (

    "net"
    "log"
    "os"
    "io/ioutil"
    "fmt"
)
func main() {
    service := "127.0.0.1:12000"
//    service := "180.149.132.47:80"//baidu server
    tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
    if err!=nil {
        log.Fatal("resolve error", err)
        os.Exit(1)
    }
    conn, err := net.DialTCP("tcp", nil, tcpAddr)
    if err!=nil {
        log.Fatal("dialing error", err)
        os.Exit(1)
    }
    //模拟一个基于HTTP协议的客户端请求去连接一个Web服务端
    _, err =conn.Write([]byte("HEAD / HTTP/1.0\r\n\r\n"))
    if err!=nil {
        log.Fatal("write error", err)
    }
    result, err := ioutil.ReadAll(conn)
    if err!=nil {
        log.Fatal("recivece error", err)
    }
    fmt.Println(string(result))
    os.Exit(0)

}
