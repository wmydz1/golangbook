package main

import (
    "github.com/gorilla/mux"
    "net/http"
    "fmt"
)
func main() {
    r := mux.NewRouter().StrictSlash(false)
    r.HandleFunc("/", HomeHandler)
    // Posts collection
    posts := r.Path("/posts").Subrouter()
    posts.Methods("GET").HandlerFunc(PostIndexHandler)
    posts.Methods("POST").HandlerFunc(PostCreateHandler)
    //post singular
    post := r.Path("/posts/{id}").Subrouter()
    post.Methods("GET").Path("/edit").HandlerFunc(PostEditHandler)
    post.Methods("GET").HandlerFunc(PostShowHandler)
    post.Methods("PUT", "POST").HandlerFunc(PostUpadateHandler)
    post.Methods("DELETE").HandlerFunc(PostDeleteHandler)

    http.ListenAndServe(":10000", r)
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Home")
}
func PostIndexHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "post index")
}
func PostCreateHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "post create")
}
func PostShowHandler(w http.ResponseWriter, r *http.Request) {
    id := mux.Vars(r)["id"]
    fmt.Fprintln(w, "showing post", id)
}
func PostUpadateHandler(w http.ResponseWriter, r *http.Request) {

    fmt.Fprintln(w, "post update")
}
func PostEditHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "post edit")
}

func PostDeleteHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "post delete")
}