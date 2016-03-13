package main
import (
    "net"
    "os"
)

func get_internal(){
    addrs,err :=net.InterfaceAddrs()
    if err !=nil{
        os.Stderr.WriteString("Oops:"+err.Error())
        os.Exit(1)
    }
    for _,a :=range addrs{
        if ipnet,ok :=a.(*net.IPNet);ok && !ipnet.IP.IsLoopback(){
            if ipnet.IP.To4() !=nil{
                os.Stdout.WriteString(ipnet.IP.String())
            }
        }
    }
    os.Exit(0)
}

func main(){
    get_internal()
}