package main
import (

    "google.golang.org/grpc"
    "log"
    pd "github.com/grpc/grpc-common/go/helloworld"
    "golang.org/x/net/context"
    "fmt"
)
const (
    address = "127.0.0.1:10000"
    defaultName = "world"
)

func main() {



    conn, err := grpc.Dial(address)
    if err !=nil {
        log.Printf("did not connect: %v", err)
    }
    defer conn.Close()
   c:= pd.NewGreeterClient(conn)

    //Contact the server and print out its response
    var name string
    fmt.Scanf("%s\n",&name)

    r,err :=c.SayHello(context.Background(),&pd.HelloRequest{Name:name})
    if err !=nil{
        log.Printf("could not greet: %v ",err)
    }
    fmt.Printf("Greet: %s",r.Message)

}
