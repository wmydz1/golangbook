package main
import (

    "golang.org/x/net/context"
    pd "github.com/grpc/grpc-common/go/helloworld"
    "log"
    "net"
    "google.golang.org/grpc"
)
const (
    port = ":10000"
)
type server struct {

}

func (s *server) SayHello(ctx context.Context, in *pd.HelloRequest) (*pd.HelloReply, error) {
    log.Println(in.Name)
    return &pd.HelloReply{Message:"Hello "+in.Name}, nil
}
func main() {
    lis, err := net.Listen("tcp", port)
    if err != nil {
        log.Printf("failed to listen: %v", err)
    }
    s := grpc.NewServer()
    pd.RegisterGreeterServer(s, &server{})
    s.Serve(lis)

}