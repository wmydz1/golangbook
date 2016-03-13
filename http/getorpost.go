package main
import (

    "net/http"
    "fmt"
    "os"
    "io/ioutil"
    "strings"
)
func Get(){
    client:=&http.Client{}
    reqest,err:=http.NewRequest("GET","http://www.baidu.com/",nil)
    if err !=nil{
        fmt.Println("Fatal error",err.Error())
        os.Exit(0)
    }
    //Add 头协议
    reqest.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8")
    reqest.Header.Add("Accept-Language", "ja,zh-CN;q=0.8,zh;q=0.6")
    reqest.Header.Add("Connection", "keep-alive")
    reqest.Header.Add("Cookie", "设置cookie")
    reqest.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 6.1; WOW64; rv:12.0) Gecko/20100101 Firefox/12.0")

    response,err :=client.Do(reqest)
    defer  response.Body.Close()

    //遍历cookies
    cookies:=response.Cookies()

    for _,cookie :=range cookies{
        fmt.Println("cookie: ",cookie)
    }
    body,err:=ioutil.ReadAll(response.Body)
    if err !=nil{
        fmt.Println("Read data error",err.Error())
    }
    fmt.Println(string(body))
}

func Post(){
    client :=&http.Client{}
    reqest,err :=http.NewRequest("POST","http://www.baidu.com/",strings.NewReader("name=1&age=12"))
    if err !=nil{
        fmt.Println("Request Error",err.Error())
        os.Exit(1)
    }
    reqest.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8")
    reqest.Header.Add("Accept-Language", "ja,zh-CN;q=0.8,zh;q=0.6")
    reqest.Header.Add("Connection", "keep-alive")
    reqest.Header.Add("Cookie", "设置cookie")
    reqest.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 6.1; WOW64; rv:12.0) Gecko/20100101 Firefox/12.0")

    respose,err :=client.Do(reqest)
    cookies:=respose.Cookies()

    for _,cookie := range cookies{
        fmt.Println("cookie:",cookie)
    }
    defer respose.Body.Close()

    body,err :=ioutil.ReadAll(respose.Body)
    if err!=nil{
        fmt.Println("Read Data  Error")
    }
    fmt.Println(string(body))

}

func main(){
    //    Get()
    Post()
}