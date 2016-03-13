package main
import (

    "code.google.com/p/go.net/websocket"
    "fmt"
    "net/http"
    "log"
)
func Echo(ws *websocket.Conn) {
    var err error
    for {
        var reply string
        if err =websocket.Message.Receive(ws, &reply); err !=nil {
            fmt.Println("Can't recevive")
            break
        }
        fmt.Println("Received back from client:"+reply)
        msg := "Received: "+reply
        fmt.Println("Sending to client: "+msg)
        if err =websocket.Message.Send(ws, msg); err!=nil {
            fmt.Println("Can't send")
            break
        }
    }
}
func main() {
    http.Handle("/", websocket.Handler(Echo))
    if err := http.ListenAndServe(":12000", nil); err!=nil {
        log.Fatal("ListenAndServe: ", err)
    }
}