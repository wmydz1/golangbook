package main

import (
    "encoding/json"
    "net/http"
)

type Book struct {
    Title string `json:"title"`
    Author string `json:"author"`
}
func main() {
    http.HandleFunc("/", ShowBook)
    http.ListenAndServe(":12000", nil)
}
func ShowBook(w http.ResponseWriter, r *http.Request) {
    book := Book{"what's god", "samchen"}
    js, err := json.Marshal(book)

    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    w.Header().Set("Content-Type", "application/json")
    w.Write(js)
}
