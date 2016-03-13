package main
import (
    "gopkg.in/redis.v3"
    "fmt"
)

func ExampleNewClient() *redis.Client {
    client := redis.NewClient(&redis.Options{
        Addr:"localhost:6379",
        Password:"",
        DB:0,
    })
    //    pong, err := client.Ping().Result()
    //    fmt.Println(pong, err)
    return client
}
func ExampleClient() {
    client := ExampleNewClient()
    err := client.Set("key", "logoocc", 0).Err()
    if err != nil {
        panic(err)
    }
    val, err := client.Get("key").Result()
    if err !=nil {
        panic(err)
    }
    fmt.Println("key", val)

    val2, err := client.Get("key2").Result()

    if err == redis.Nil {
        fmt.Println("key2 does not exists")
    }else if err !=nil {
        panic(err)
    }else {
        fmt.Println("key2", val2)
    }
}

func InsertRedis(key, value string) error {

    client := ExampleNewClient()
    err := client.Set(key, value, 0).Err()
    if err != nil {
        return err
    }
    fmt.Println("insert sucessful ")
    return nil
}
func QueryRedis(key string) string {
    client := ExampleNewClient()
    val, err := client.Get(key).Result()
    if err == redis.Nil {
        fmt.Println("this key is not exits")
    }else if err !=nil {
        panic(err)
    }else {
        return val
    }
    return ""
}
func main() {
//    err := InsertRedis("name", "samchen")
//    if err != nil {
//        fmt.Println(err.Error())
//    }

    val := QueryRedis("name")
    if val != "" {
        fmt.Println(val)
    }
}