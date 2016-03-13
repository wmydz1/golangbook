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
    person := Person{
        Name:Name{Family:"chen", Personal:"samchen"},
        Email:[]Email{Email{Kind:"home", Address:"home@gmail"}, Email{Kind:"work", Address:"work@gmail"}}}
    service := ":12000"

    conn, err := net.Dial("tcp", service)
    if err !=nil {
        fmt.Println(err.Error())
    }
    encoder := json.NewEncoder(conn)
    decoder := json.NewDecoder(conn)
    for i := 0; i<10; i++ {
        encoder.Encode(person)
        var newPerson Person
        decoder.Decode(&newPerson)
        fmt.Println(newPerson.String())
    }

}