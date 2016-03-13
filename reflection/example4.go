package main
import (
    "fmt"
    "reflect"
)

type Foo struct {
    X string
    Y int
}

func (f Foo) Do() {
    fmt.Printf(" X is %s,Y is %d ", f.X, f.Y)
}
func main() {
    var i int = 123
    var f float32
    var l []string = []string{"a", "b", "c"}

    fmt.Println(reflect.ValueOf(i))
    fmt.Println(reflect.ValueOf(f))
    fmt.Println(reflect.ValueOf(l))

    var foo *Foo
    fmt.Println(reflect.TypeOf(foo))

    foo =&Foo{"abc",123}
    reflect.ValueOf(foo).MethodByName("Do").Call([]reflect.Value{})
}
