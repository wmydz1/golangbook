package main

import (

    "code.google.com/p/go.net/websocket"
    "fmt"
    "net/http"
)
func send(ws *websocket.Conn) {
    for i := 0; i<10; i++ {
        msg := "Hello "+string(i+48)
        err := websocket.Message.Send(ws, msg)
        if err !=nil {
            fmt.Println("can't send")
            break
        }
        var reply string
        err =websocket.Message.Receive(ws, &reply)
        if err !=nil {
            fmt.Println("Can't recive")
            break
        }
        fmt.Println("reciced back from client : "+reply)

    }
}
func main(){
    http.Handle("/",websocket.Handler(send))
    err :=http.ListenAndServe(":12000",nil)
    if err != nil {
        fmt.Println("listen port err",err.Error())
    }
}

