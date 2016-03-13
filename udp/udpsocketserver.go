package main
import (

    "net"
    "log"
    "os"
    "time"
)
func main() {
    service := ":12000"
    udpAddr, err := net.ResolveUDPAddr("udp4", service)
    if err !=nil {
        log.Fatal("can't resolve udp addr", err)
        os.Exit(1)
    }
    conn, err := net.ListenUDP("udp", udpAddr)
    if err !=nil {
        log.Fatal("listening error", err)
        os.Exit(1)
    }
    for {
        handleUdpReq(conn)
    }
}
func handleUdpReq(conn *net.UDPConn) {
    var buf [512]byte
    _, addr, err := conn.ReadFromUDP(buf[0:])
    if err!=nil {
        return
    }
    daytime := time.Now().String()
    conn.WriteToUDP([]byte(daytime), addr)

}