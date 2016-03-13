package main
import (

    "net"
    "log"
    "os"
    "time"
    "strconv"
)
//这个服务端没有处理客户端实际请求的内容。如果我们需要通过从客户端发送不同的请求来获取不同的时间格式，而且需要一个长连接，该怎么做呢

func main() {
    service := ":12000"
    tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
    if err!=nil {
        log.Fatal("can't resolve tcp addr ", err)
        os.Exit(1)
    }
    listener, err := net.ListenTCP("tcp", tcpAddr)
    if err!=nil {
        log.Fatal("listening error ", err)
        os.Exit(1)
    }
    for {
        conn, err := listener.Accept()
        if err!=nil {
            log.Fatal("conn error", err)
            continue
        }
        go handleClientMsg(conn)
    }
}
func handleClientMsg(conn net.Conn) {

    conn.SetReadDeadline(time.Now().Add(2 * time.Minute)) // set 2 minutes time out
    request := make([]byte, 128)//set maxium request length to 128kb prevent flood attack
    defer conn.Close() //close connection before exit
    for {
        read_len, err := conn.Read(request)
        if err !=nil {
            log.Fatal("read error", err)
            break
        }

        if read_len==0 {
            break //connection already closed by client
        }else if string(request) =="timestamp" {
            daytime := strconv.FormatInt(time.Now().Unix(), 10)
            conn.Write([]byte(daytime))
        }else {
            daytime := time.Now().String()
            conn.Write([]byte(daytime))
        }

        request=make([]byte, 128) //clear last read content

    }
}