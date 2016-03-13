package main
import (

    "net"
    "fmt"
    "time"
    "strconv"
)

func main() {
    service := ":10000"

    tcpaddr, err := net.ResolveTCPAddr("tcp4", service)
    if err!=nil {
        fmt.Println("can't resolve rcp addr ", err)
    }
    listener, err := net.ListenTCP("tcp", tcpaddr)
    if err !=nil {
        fmt.Println("listen error", err)
    }
    for {
        conn, err := listener.Accept()
        if err !=nil {
            continue
        }
        go handlerequset(conn)

    }
}

func handlerequset(conn net.Conn) {
    // set  2 minute time out
    conn.SetDeadline(time.Now().Add(2 * time.Minute))
    // set maxium request length to 128kb to prevent floor attack
    request := make([]byte, 128)

    defer conn.Close()
    for {
        read_len, err := conn.Read(request)
        if err !=nil {
            fmt.Println(err)
            break
        }

        if read_len ==0 {
            break //connection already closed by client
        }else if string(request) =="timestamp" {
            daytime := strconv.FormatInt(time.Now().Unix(), 10)
            conn.Write([]byte(daytime))
        }else {
            daytime := time.Now().String()
            conn.Write([]byte(daytime))
        }

        request=make([]byte, 128) //clean last read content


    }
}
