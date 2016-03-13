package main
import (
    "code.google.com/p/go.net/websocket"
    "fmt"
    "net/http"

)
type Person struct {
    Name string
    Emails []string
}
func sendmsg(ws *websocket.Conn, person *Person) {
    err := websocket.JSON.Receive(ws, person)
    if err != nil {
        fmt.Println("can't recive from server", err.Error())
    }else {
        fmt.Println("Name is ", person.Name)

        for _, email := range person.Emails {
            fmt.Println("email is "+email)
        }
    }

}
func recivePerson(ws *websocket.Conn) {

    var person Person
    for {
         sendmsg(ws, &person)

    }
}

func main() {
    http.Handle("/", websocket.Handler(recivePerson))
    err := http.ListenAndServe(":12000", nil)
    if err != nil {
        fmt.Println("listen err ", err.Error())
    }
}
