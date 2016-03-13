package main
import (

    "reflect"
    "fmt"
)
type Foo struct {
    FirstName string `tag_name:"tag 1"`
    LastName string `tag_name:"tag 2"`
    Age int `tag_name:"tag 3"`
}
func (f *Foo) reflect() {
    val := reflect.ValueOf(f).Elem()

    for i := 0; i <val.NumField(); i++ {
        valieField := val.Field(i)
        typeField := val.Type().Field(i)
        tag := typeField.Tag

        fmt.Printf("Field Name: %s \t Filed Value: %v ,\t Tag Value: %s\n", typeField.Name, valieField.Interface(), tag.Get("tag_name"))
    }
}

func main() {
    f := &Foo{
        FirstName:"Drew",
        LastName:"Olson",
        Age:100,
    }
    f.reflect()
}