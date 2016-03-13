package main

//BY: 29295842@qq.com
//这个有一定问题   如果配置信息里有中文就不行
//[Server] ;MYSQL配置
//Server=localhost   ;主机
//golang 读取 ini配置信息
//http://www.widuu.com/archives/02/961.html
import (
    "fmt"
    "github.com/widuu/goini"
//"runtime"
//"time"
)

func Read_ini_string(file_data string, KEY1 string, KEY2 string) string {
    conf := goini.SetConfig(file_data) //goini.SetConfig(filepath) 其中filepath是你ini 配置文件的所在位置
    return conf.GetValue(KEY1, KEY2)   //database是你的[section]，username是你要获取值的key名称
}
func main() {
    fmt.Println(Read_ini_string("server.ini", "Server", "Username1"))

    //conf := goini.SetConfig("server.ini")           //goini.SetConfig(filepath) 其中filepath是你ini 配置文件的所在位置
    //username := conf.GetValue("Server", "Username") //database是你的[section]，username是你要获取值的key名称
    //fmt.Println(username)
    /*for {
        time.Sleep(1 * time.Second)
        runtime.Gosched()
    }  */
}
