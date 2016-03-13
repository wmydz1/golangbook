package main
import (
    "net/http"
    "gowebprogram/fileserver"
)

func main() {


    http.Handle("/", fileserver.FileServer(http.Dir("./")));
    http.ListenAndServe(":9000", nil)
}

