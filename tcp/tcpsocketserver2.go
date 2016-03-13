package main
import (

    "net"
    "log"
    "os"
    "time"
)
func main() {
    service := ":12000"

    tcpAddr, err := net.ResolveTCPAddr("tcp4", service)

    if err !=nil {
        log.Fatal("Can't resolve tcp addr", err)
        os.Exit(1)
    }
    listener, err := net.ListenTCP("tcp", tcpAddr)
    if err!=nil {
        log.Fatal("listening error", err)
        os.Exit(1)
    }

    for {
        conn, err := listener.Accept()
        if err!=nil {
            continue
        }
        go handleClient(conn)
    }

}
func handleClient(conn net.Conn) {
    defer conn.Close()
    daytime := time.Now().String()
    conn.Write([]byte(daytime))
}