package main
import (

    "net"
    "log"
    "os"
    "fmt"
)
func main() {
    service := "127.0.0.1:12000"
    //    service := "180.149.132.47:80"//baidu server
    udpAddr, err := net.ResolveUDPAddr("udp4", service)
    if err !=nil {
        log.Fatal("can't resolve udp addr", err)
        os.Exit(1)
    }
    conn, err := net.DialUDP("udp", nil, udpAddr)
    if err !=nil {
        log.Fatal("dialing error", err)
        os.Exit(1)
    }
    _, err=conn.Write([]byte("anything"))
    if err !=nil {
        log.Fatal("writing error", err)
        os.Exit(1)
    }

    var buf [512]byte
    n, err := conn.Read(buf[0:])
    if err!=nil {
        log.Fatal("read error", err)
        os.Exit(1)
    }
    fmt.Println(string(buf[0:n]))
    os.Exit(0)



}
