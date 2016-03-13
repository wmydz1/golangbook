package main
import (

    "fmt"
    "io"
    "code.google.com/p/go.net/websocket"
    "time"

)

func main() {
    service := "ws://192.168.0.2:12000"
    conn, err := websocket.Dial(service, "", "http://192.168.0.2")
    if err != nil {
        fmt.Println("conn err",err.Error())
    }
    var msg string
    for {
        err := websocket.Message.Receive(conn, &msg)
        if err !=nil {
            if err ==io.EOF {
                break
            }
            fmt.Println("can't recive from server")
            break
        }
        fmt.Println("recive from server "+msg)
        err =websocket.Message.Send(conn, msg)
        if err!=nil {
            fmt.Println("can't send  msg to server ")
            break
        }

    }
    time.Sleep(5 * time.Second)


}
