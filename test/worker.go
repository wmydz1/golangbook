package test
import "fmt"
func Work() {
    fmt.Println("Hello Golang")
}
func Add(a, b int) {
    fmt.Println(a+b)
}
//要外部才能调，自己包下调不了