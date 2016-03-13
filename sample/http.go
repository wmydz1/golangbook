package main
import (

    "net/http"
)
func sayHello(w http.ResponseWriter, r *http.Request) {

    w.Write([]byte("Hello,Welcome to golang server..."))
}
func main() {
    http.HandleFunc("/", sayHello)
    http.ListenAndServe(":10000", nil)

}
