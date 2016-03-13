package main

import (


    "net"
    "log"
    "os"
    "time"
)

func main() {

    service := ":12000"
    //    service := "127.0.0.1:12000"
    tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
    if err !=nil {
        log.Fatal("resolve error", err)
        os.Exit(1)
    }
    listener, err := net.ListenTCP("tcp", tcpAddr)
    if err!=nil {
        log.Fatal("listen error", err)
    }

    for {
        conn, err := listener.Accept()
        if err!=nil {
            log.Fatal("conn error", err)
            continue
        }
        daytime := time.Now().String()
        conn.Write([]byte(daytime))
        conn.Close()

    }


}

