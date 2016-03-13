package main
import (
    "reflect"
    "fmt"
)
func repeat(x interface{}, n int) interface{} {
    v := reflect.ValueOf(x)
    ys := reflect.MakeSlice(reflect.SliceOf(v.Type()), n, n)
    for i := 0; i<n; i++ {
        ys.Index(i).Set(v)
    }
    return ys.Interface()
}
type MyInt int

func main() {
    var xs []int = repeat(3, 10).([]int)
    fmt.Println(reflect.ValueOf(xs).Type(), xs)

    var ss []string = repeat("go", 3).([]string)
    fmt.Println(reflect.ValueOf(ss).Type(), ss)


    var ms []MyInt = repeat(MyInt(5), 6).([]MyInt)
    fmt.Println(reflect.ValueOf(ms).Type(), ms)
}