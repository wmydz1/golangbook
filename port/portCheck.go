package main
import (

    "os"
    "fmt"
)
func main(){
    port :=os.Getenv("PORT")
    if port ==""{
        port ="12000"
    }
    fmt.Println(port)

}
