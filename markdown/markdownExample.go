package main
import (
    "net/http"
    "github.com/slene/blackfriday"
)

func main() {
    http.HandleFunc("/markdown", GenerateMarkdown)
    http.Handle("/", http.FileServer(http.Dir("./markdown/public")))
    http.ListenAndServe(":12000", nil)
}

func GenerateMarkdown(w http.ResponseWriter, r *http.Request) {
    markdown := blackfriday.MarkdownCommon([]byte(r.FormValue("body")))
    w.Write(markdown)
}