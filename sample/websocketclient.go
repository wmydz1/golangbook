package main
import (
    "code.google.com/p/go.net/websocket"

    "fmt"
)
func main() {
    //    service := ":10000"
    origin := "http://localhost:10000"
    url := "ws://localhost:10000/ws"
    ws, err := websocket.Dial(url, "", origin)
    if err!=nil {
        fmt.Println(err)
    }

    err =websocket.Message.Send(ws, "Hello,server")
    if err!=nil {
        fmt.Println(err)
    }

    var reply string
    err =websocket.Message.Receive(ws,&reply)
    if err!=nil{
        fmt.Println(err)
    }
    fmt.Println(reply)

}
