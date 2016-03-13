package main
import (

    "net/http"
    "regexp"
    "html/template"
    "fmt"
    "path/filepath"
    "strconv"
    "time"
    "os"
    "io"

)
var mux map[string]func(http.ResponseWriter, *http.Request)
type MyHandler struct {

}
type home struct {
    Title string
}
const (
    Template_Dir = "./fileserver/view/"
    Upload_Dir = "./fileserver/upload/"
)
func (hanler *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    if h, ok := mux[r.URL.String()]; ok {
        h(w, r)
        return
    }
    if ok, _ := regexp.MatchString("/css/", r.URL.String()); ok {
        http.StripPrefix("/css/", http.FileServer(http.Dir("./fileserver/css/"))).ServeHTTP(w, r)
    }else {
       http.StripPrefix("/", http.FileServer(http.Dir("./fileserver/upload/"))).ServeHTTP(w, r)
    }
}
func upload(w http.ResponseWriter, r *http.Request) {
    if r.Method =="GET" {
        t, _ := template.ParseFiles(Template_Dir+"upload.html")
        t.Execute(w, "上传文件")
    }else {
        r.ParseMultipartForm(32 <<20)
        file, handler, err := r.FormFile("uploadfile")
        if err != nil {
            fmt.Fprintf(w, "%v", "上传错误")
            return
        }
        fileExt := filepath.Ext(handler.Filename)
        if check(fileExt) == false {
            fmt.Fprintf(w, "%v", "不允许上传该文件类型")
            return
        }
        filename := strconv.FormatInt(time.Now().Unix(), 10)+fileExt
        f, _ := os.OpenFile(Upload_Dir+filename, os.O_CREATE|os.O_WRONLY, 0660)
        _, err =io.Copy(f, file)
        if err !=nil {
            fmt.Fprintf(w, "%v", "上传失败")
            return
        }
        filedir, _ := filepath.Abs(Upload_Dir+filename)
        fmt.Fprintf(w, "%v", filename+"上传完成，服务器地址"+filedir)

    }
}

func index(w http.ResponseWriter, r *http.Request) {
    title := home{Title:"首页"}
    t, _ := template.ParseFiles(Template_Dir+"index.html")
    t.Execute(w, title)
}
func StaticServer(w http.ResponseWriter, r *http.Request) {
        http.StripPrefix("/file", http.FileServer(http.Dir(Upload_Dir))).ServeHTTP(w, r)
//    http.FileServer(http.Dir(Upload_Dir))
}
func check(name string) bool {
    ext := []string{".exe", ".js", ".css"}
    for _, v := range ext {
        if v == name {
            return false
        }
    }
    return true
}

func main() {
    server := http.Server{
        Addr:":10000",
        Handler:&MyHandler{},
        ReadTimeout:10 * time.Second,
    }
    mux =make(map[string]func(http.ResponseWriter, *http.Request))
    mux["/"] =index
    mux["/upload"] = upload
    mux["/file"] =StaticServer
    server.ListenAndServe()
}