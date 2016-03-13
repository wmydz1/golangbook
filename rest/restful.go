package main
import (
    "github.com/drone/routes"
    "net/http"

    "fmt"

)

func getuser(w http.ResponseWriter, r *http.Request) {
    params := r.URL.Query()

    uid := params.Get(":uid")

    fmt.Fprintf(w, "you are get user %s by get", uid)//http://127.0.0.1:12000/user/sam
}
func modifyuser(w http.ResponseWriter, r *http.Request) {
    params := r.URL.Query()
    uid := params.Get(":uid")
    fmt.Fprintf(w, "you are get user %s by post", uid)
}
func deleteuser(w http.ResponseWriter, r *http.Request) {
    params := r.URL.Query()
    uid := params.Get(":uid")
    fmt.Fprintf(w, "you are delete user %s by delete", uid)
}
func adduser(w http.ResponseWriter, r *http.Request) {
    params := r.URL.Query()
    uid := params.Get(":uid")
    fmt.Fprintf(w, "you are add user %s by put", uid)
}
func main() {
    mux := routes.New()
    mux.Get("/user/:uid", getuser)
    mux.Post("/user/:uid", modifyuser)
    mux.Del("/user/:uid", deleteuser)
    mux.Put("/user", adduser)
    http.Handle("/", mux)
    http.ListenAndServe(":12000", nil)

}
