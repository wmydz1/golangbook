package main
import (

    "net/http"
    "log"

    "fmt"
    "strings"
)
/*
http://localhost:10000
http://localhost:10000/?url_long=111&url_long=222
*/

func worker(w http.ResponseWriter, r *http.Request) {
    //解析参数, 默认是不会解析的
    r.ParseForm()
    //这些是服务器端的打印信息
    fmt.Println(r.Form)
    fmt.Println("path", r.URL.Path)
    fmt.Println(r.URL.Scheme)
    fmt.Println(r.Form["url_long"])

    for k, v := range r.Form {
        fmt.Println("key ", k)
        fmt.Println("val ", strings.Join(v, ""))
    }
    fmt.Fprintf(w, "welcome to golang server")

}
func main() {
    http.HandleFunc("/", worker)
    err := http.ListenAndServe(":10000", nil)
    if err!=nil {
        log.Fatal("ListenAndServe", err)
    }
}
Status API Training Shop Blog About Pricing
