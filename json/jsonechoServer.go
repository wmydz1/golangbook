package main
import (
    "net"
    "fmt"
    "encoding/json"
)

type Person struct {
    Name  Name
    Email []Email
}
type Email struct {
    Kind    string
    Address string
}
type Name struct {
    Family   string
    Personal string
}
func (p Person) String() string {
    s := p.Name.Personal +" "+p.Name.Family
    for _, v := range p.Email {
        s +="\n"+v.Kind+":"+v.Address
    }
    return s
}
func main() {
    service := ":12000"
    tcpAddr, err := net.ResolveTCPAddr("tcp", service)
    if err != nil {
        fmt.Println(err.Error())
    }
    listener, err := net.ListenTCP("tcp", tcpAddr)
    if err !=nil {
        fmt.Println(err.Error())
    }
    for {
        conn, err := listener.Accept()
        if err !=nil {
            continue
        }
        encoder := json.NewEncoder(conn)
        decoder := json.NewDecoder(conn)
        for i := 0; i<10; i++ {
            var person Person
            decoder.Decode(&person)
            fmt.Println(person.String())
            encoder.Encode(person)
        }
        conn.Close()
    }

}

