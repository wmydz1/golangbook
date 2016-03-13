package main

import (
    "net/http"


    "strings"


    "fmt"
)
func myFileServer(root http.FileSystem) http.Handler {
    return &fileHandler{root}
}

type fileHandler struct {
    root http.FileSystem
}

func (f *fileHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

    w.Write([]byte("go service is running"))
    r.ParseForm()
    upath := r.URL.Path
    if !strings.HasPrefix(upath, "/") {
        upath = "/" + upath
        r.URL.Path = upath
    }
    for key,val := range r.Header{
        fmt.Println(key,val)
    }
    fmt.Println(r.Host)
    fmt.Println(r.Trailer)
    for key,val := range r.Trailer{
        fmt.Println(key,val)
    }
    fmt.Println(r.RemoteAddr)
    fmt.Println(r.Method)
    fmt.Println(r.Proto)
    fmt.Println(r.Close)
    fmt.Println(r.RequestURI)
}



func fileServer() http.Handler {
    //    handler := http.FileServer(http.Dir("./"));
    handler := myFileServer(http.Dir("./"));
    return handler
}
func main() {

    http.Handle("/", fileServer())
    //    http.Handle("/", fileserver.FileServer(http.Dir("./")));
    http.ListenAndServe(":8080", nil)
}
