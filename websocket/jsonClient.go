package main
import (
    "code.google.com/p/go.net/websocket"
    "fmt"
)

type Person struct {
    Name string
    Emails []string
}
func main() {
    service := "ws://192.168.0.2:12000"
    conn, err := websocket.Dial(service, "", "http://192.168.0.2")
    if err != nil {
        fmt.Println("conn err ", err.Error())
    }
    person := Person{Name:"samchen", Emails:[]string{"home@gmail.com", "company@gmail.com"}}
    err =websocket.JSON.Send(conn, person)
    if err != nil {
        fmt.Println("can't send msg to client", err.Error())
    }

}
