package main
import (
    "html/template"
    "net/http"
)


type data struct {
    Content string
}

func handler(w http.ResponseWriter, r *http.Request) {
    t, _ := template.ParseFiles("./template/view/show.html")

    c := data{Content:"samchen"}
    t.Execute(w, c)
}
func main() {
    http.HandleFunc("/", handler)
    http.ListenAndServe(":12000", nil)

}
