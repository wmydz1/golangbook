package main
import (

    "github.com/logoocc/httprounter"
    "net/http"
    "fmt"
    "github.com/qiniu/log"
)
func main() {
    r := httprounter.New()
    r.GET("/", HomeHandler)

    // Posts clooection
    r.GET("/posts", PostIndexHandler)
    r.POST("/posts", PostCreateHandler)


    // Post singular
    r.GET("/posts/:id", PostShowHandler)
    r.PUT("/posts/:id", PostUpadateHandler)
    r.GET("/posts/:id/edit", PostEditHandler)


    fmt.Println("service is running")
    http.ListenAndServe(":12000", r)

}

func HomeHandler(w http.ResponseWriter, r *http.Request, p httprounter.Params) {
    fmt.Fprintln(w, "Home")
}
func PostIndexHandler(w http.ResponseWriter, r *http.Request, p httprounter.Params) {
    fmt.Fprintln(w, "post index")
}
func PostCreateHandler(w http.ResponseWriter, r *http.Request, p httprounter.Params) {
    fmt.Fprintln(w, "post create")
}
func PostShowHandler(w http.ResponseWriter, r *http.Request, p httprounter.Params) {
    id := p.ByName("id")
    fmt.Fprintln(w, "showing post", id)
}
func PostUpadateHandler(w http.ResponseWriter, r *http.Request, p httprounter.Params) {

    fmt.Fprintln(w, "post update")
}
func PostEditHandler(w http.ResponseWriter, r *http.Request, p httprounter.Params) {
    fmt.Fprintln(w, "post edit")
}