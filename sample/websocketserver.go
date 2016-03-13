package main
import (
    "code.google.com/p/go.net/websocket"
    "net/http"
    "fmt"

)


func tellcho(ws *websocket.Conn) {


    msg := "welcome to websocket server..."
    for {
        if err := websocket.Message.Send(ws, msg); err !=nil {
            fmt.Println("Can't send...", err)

        }

        var reply string
        if err := websocket.Message.Receive(ws, &reply); err !=nil {
            fmt.Println("Can't receive..", err)
        }
        fmt.Println("Recive data : ", reply)
    }



}

func main() {
    http.Handle("/", websocket.Handler(tellcho))
    http.ListenAndServe(":10000", nil)
}