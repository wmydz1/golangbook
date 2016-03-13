package main
import (
    "gowebprogram/redis/helper"
    "fmt"
    "os"
    "io"
    "net/http"
    "strings"
    "html/template"
)
func query(name string) string {
    val := helper.QueryRedis(name)
    if val !="" {
        return val
    }
    return ""
}

func readSingle() string {
    f, err := os.Open("./redis/qq/QQ说说8.txt")
    if err !=nil {
        fmt.Println(err.Error())
        os.Exit(1)
    }
    defer f.Close()
    reader := io.LimitReader(f, 1024 * 100)
    p := make([]byte, 1024 * 100)
    var total int
    for {
        n, err := reader.Read(p)
        if err == io.EOF {
            break
        }
        total =total+n
    }
    return string(p[:total])
}



func save() {
    val := readSingle()
    err := helper.InsertRedis("qq8", val)
    if err!=nil {
        fmt.Println(err.Error())
    }
}

type data struct {
    Content template.HTML
}
func echo(w http.ResponseWriter, r *http.Request) {
    // Content-Type:text/html;charset=utf-8
    w.Header().Set("Content-Type", "text/html;charset=utf-8")
    str := strings.Replace(query("qq1"), "\n", "<br>", -1)

    t, _ := template.ParseFiles("./redis/view/show.html")
    str2 := template.HTML(strings.Join(strings.Split(str,"<br>"),"<br>"))

    c := data{Content:str2}
    t.Execute(w, c)

    //    fmt.Fprintln(w, str)
}

func echo2(w http.ResponseWriter, r *http.Request) {
    // Content-Type:text/html;charset=utf-8
    w.Header().Set("Content-Type", "text/html;charset=utf-8")
    str := strings.Replace(query("qq2"), "\n", "<br>", -1)

    t, _ := template.ParseFiles("./redis/view/show.html")
    str2 := template.HTML(strings.Join(strings.Split(str,"<br>"),"<br>"))

    c := data{Content:str2}
    t.Execute(w, c)

    //    fmt.Fprintln(w, str)
}
func echo3(w http.ResponseWriter, r *http.Request) {
    // Content-Type:text/html;charset=utf-8
    w.Header().Set("Content-Type", "text/html;charset=utf-8")
    str := strings.Replace(query("qq3"), "\n", "<br>", -1)

    t, _ := template.ParseFiles("./redis/view/show.html")
    str2 := template.HTML(strings.Join(strings.Split(str,"<br>"),"<br>"))

    c := data{Content:str2}
    t.Execute(w, c)

    //    fmt.Fprintln(w, str)
}
func echo4(w http.ResponseWriter, r *http.Request) {
    // Content-Type:text/html;charset=utf-8
    w.Header().Set("Content-Type", "text/html;charset=utf-8")
    str := strings.Replace(query("qq4"), "\n", "<br>", -1)

    t, _ := template.ParseFiles("./redis/view/show.html")
    str2 := template.HTML(strings.Join(strings.Split(str,"<br>"),"<br>"))

    c := data{Content:str2}
    t.Execute(w, c)

    //    fmt.Fprintln(w, str)
}
func echo5(w http.ResponseWriter, r *http.Request) {
    // Content-Type:text/html;charset=utf-8
    w.Header().Set("Content-Type", "text/html;charset=utf-8")
    str := strings.Replace(query("qq5"), "\n", "<br>", -1)

    t, _ := template.ParseFiles("./redis/view/show.html")
    str2 := template.HTML(strings.Join(strings.Split(str,"<br>"),"<br>"))

    c := data{Content:str2}
    t.Execute(w, c)

    //    fmt.Fprintln(w, str)
}
func echo6(w http.ResponseWriter, r *http.Request) {
    // Content-Type:text/html;charset=utf-8
    w.Header().Set("Content-Type", "text/html;charset=utf-8")
    str := strings.Replace(query("qq6"), "\n", "<br>", -1)

    t, _ := template.ParseFiles("./redis/view/show.html")
    str2 := template.HTML(strings.Join(strings.Split(str,"<br>"),"<br>"))

    c := data{Content:str2}
    t.Execute(w, c)

    //    fmt.Fprintln(w, str)
}
func echo7(w http.ResponseWriter, r *http.Request) {
    // Content-Type:text/html;charset=utf-8
    w.Header().Set("Content-Type", "text/html;charset=utf-8")
    str := strings.Replace(query("qq7"), "\n", "<br>", -1)

    t, _ := template.ParseFiles("./redis/view/show.html")
    str2 := template.HTML(strings.Join(strings.Split(str,"<br>"),"<br>"))

    c := data{Content:str2}
    t.Execute(w, c)

    //    fmt.Fprintln(w, str)
}
func echo8(w http.ResponseWriter, r *http.Request) {
    // Content-Type:text/html;charset=utf-8
    w.Header().Set("Content-Type", "text/html;charset=utf-8")
    str := strings.Replace(query("qq8"), "\n", "<br>", -1)

    t, _ := template.ParseFiles("./redis/view/show.html")
    str2 := template.HTML(strings.Join(strings.Split(str,"<br>"),"<br>"))

    c := data{Content:str2}
    t.Execute(w, c)

    //    fmt.Fprintln(w, str)
}

func main() {

    http.HandleFunc("/", echo)
    http.HandleFunc("/2", echo2)
    http.HandleFunc("/3", echo3)
    http.HandleFunc("/4", echo4)
    http.HandleFunc("/5", echo5)
    http.HandleFunc("/6", echo6)
    http.HandleFunc("/7", echo7)
    http.HandleFunc("/8", echo8)

    http.ListenAndServe(":10000", nil)
}

