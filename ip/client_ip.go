package main
import (
    "net/http"
    "fmt"
)

func getip(w http.ResponseWriter, r *http.Request) {
    r.ParseForm()
    fmt.Println(r.RemoteAddr)
    w.Write([]byte("welocme to service.. \n your ip is " + r.RemoteAddr +"\n"))
}
func main() {
    http.HandleFunc("/", getip)
    http.ListenAndServe(":10000", nil)
}
