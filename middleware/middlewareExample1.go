package main
import (
    "github.com/codegangsta/negroni"
    "net/http"
    "log"
    "fmt"
)

func main() {
    // middleware stack
    n := negroni.New(
    negroni.NewRecovery(),
    negroni.HandlerFunc(MyMiddleware),
    negroni.NewLogger(),
    negroni.NewStatic(http.Dir("./middleware/public")),
    )
    n.Run(":12000")
}

func MyMiddleware(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
    log.Println("logging on the way there...")

    // http://127.0.0.1:12000/?password=123456
    if r.URL.Query().Get("password") =="123456" {
        fmt.Println("login success")
        next(w, r)
    }else {
        http.Error(w, "NOT AUTHORIZED", http.StatusUnauthorized)
    }
    log.Println("logging on the way back...")
}
