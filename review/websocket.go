package main
import (

    "net/http"
    "golang.org/x/net/websocket"
    "fmt"
)
func handle(ws *websocket.Conn) {

    var err error

    for {
        var reply string

        if err =websocket.Message.Receive(ws, &reply); err !=nil {
            fmt.Println("can't receive...")
            break
        }
        fmt.Println("recive data from client is ", reply)


        msg := "welcome to golang websocket server !!"

        if err =websocket.Message.Send(ws, msg); err !=nil {
            fmt.Println("can't send data")
            break
        }
    }


}

func main() {
    http.Handle("/", websocket.Handler(handle))
    http.ListenAndServe(":10000", nil)
}