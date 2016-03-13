package main
import (

    "flag"
    "github.com/siddontang/moonmq/broker"
)

var addr = flag.String("addr", "127.0.0.1:8081", "moonmq broker listen address")
var httpAddr = flag.String("http_addr", "127.0.0.1:8080", "moonmq broker http listen addr")

func main() {
    flag.Parse()

    cfg := broker.NewDefaultConfig()
    cfg.Addr = *addr
    cfg.HttpAddr = *httpAddr

    app, err := broker.NewAppWithConfig(cfg)
    if err !=nil {
        panic(err)
    }
    app.Run()
}