package main

import (
    "encoding/xml"
    "fmt"
)

type Person struct {
    //这个表示这个xml的根节点是Person
    XMLName xml.Name `xml:"Person"`
    //这里的name是属性，所以使用attr标记
    Name           string         `xml:"name,attr"`
    Age            int            `xml:"age,attr"`
    FavoriteFruits FavoriteFruits `xml:"FavoriteFruits"`
}

//对于包含多个相同结构子节点的节点，定义为子节点对象切片
type FavoriteFruits struct {
    Fruits []Fruit `xml:"Fruit"`
}

type Fruit struct {
    //如果没有任何标记，那么这个节点是叶子节点
    Name  string `xml:"Name"`
    Count int    `xml:"Count"`
}

func main() {
    //列出我喜欢的水果，和一次饕餮数量
    var fruits = []Fruit{
        Fruit{"Apple", 2},
        Fruit{"Orange", 1},
        Fruit{"Grape", 10},
        Fruit{"Banana", 3},
    }

    var favoriteFruits = FavoriteFruits{
        Fruits: fruits,
    }

    //创建一个我的对象
    var person = Person{
        Name:           "duokr",
        Age:            25,
        FavoriteFruits: favoriteFruits,
    }
    //好了，用xml的函数MarshalIndent来生成xml
    if output, err := xml.MarshalIndent(person, "", "\t"); err != nil {
        fmt.Println(err)
    } else {
        fmt.Print(xml.Header)
        fmt.Println(string(output))
    }
    fmt.Println()
    //你也可以不用Indent，生成一个压缩的xml字符串
//    if output, err := xml.Marshal(person); err != nil {
//        fmt.Println(err)
//    } else {
//        fmt.Print(xml.Header)
//        fmt.Println(string(output))
//    }
}
