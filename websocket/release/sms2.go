package main
import (
    "os"
    "fmt"
    "io"
    "net/http"
    "code.google.com/p/go.net/websocket"
)

func readSingle() string {
    simpledb := "./senddb.txt"
    f, err := os.OpenFile(simpledb, os.O_CREATE|os.O_RDWR, 0644)
    if err !=nil {
        fmt.Println(err.Error())
        os.Exit(1)
    }
    defer f.Close()
    reader := io.LimitReader(f, 1024 * 100)
    p := make([]byte, 1024 * 100)
    var total int
    for {
        n, err := reader.Read(p)
        if err == io.EOF {
            break
        }
        total =total+n
    }
    return string(p[:total])
}

func writeSingle(reciveMsg string) {
    getdb := "./getdb.txt"

    fout, err := os.OpenFile(getdb, os.O_CREATE|os.O_TRUNC|os.O_RDWR, 0644)
    if err != nil {
        fmt.Println(getdb, err)
        return
    }
    defer fout.Close()
    fout.WriteString(reciveMsg)

    //    fout.Write([]byte("Just a test!\r\n"))
}

func existfile(filename string) bool {
    _, err := os.Stat(filename)
    return err == nil || os.IsExist(err)
}


func sms(ws *websocket.Conn) {

    msg := readSingle()
    err := websocket.Message.Send(ws, msg)
    if err !=nil {
        fmt.Println("can't send")
        return
    }
    var reply string
    err =websocket.Message.Receive(ws, &reply)
    if err !=nil {
        fmt.Println("Can't recive")
        return
    }
    fmt.Println("reciced back from client : "+reply)
    writeSingle(reply)


}


func main() {
    http.Handle("/", websocket.Handler(sms))
    err := http.ListenAndServe(":9000", nil)
    if err != nil {
        fmt.Println("listen port err", err.Error())
    }

}