package main

import (
    "fmt"

    "os"
)

//读取文件名字 (f *File).Name()这个函数是返回文件的名称，函数原型func (f *File) Name() string要文件的指针操作，返回字符串，感觉比较鸡助的方法底层实现
//f *File).Read()这个是函数的指针来操作的,属于*FIlE的method,函数原型func (f *File) Read(b []byte) (n int, err error)输入读取的字节数，返回字节的长度和error信息
func main() {
    f, _ := os.Open("docker.txt")

    defer f.Close()

    fmt.Println("filename is : " + f.Name())

    b := make([]byte, 1024) //设置读取的字节数
    n, _ := f.Read(b)
    fmt.Println(n)
    fmt.Println(string(b[:n]))
    //输出内容 为什么是n而不直接输入100呢？底层这样实现的
    /*
                n, e := f.read(b)
                   if n < 0 {
                        n = 0
                }
                if n == 0 && len(b) > 0 && e == nil {
                      return 0, io.EOF
                  }
    */
    //所以字节不足100就读取n
    //f *File).ReadAt()这个函数的原型是func (f *File) ReadAt(b []byte, off int64) (n int, err error)加入了下标，可以自定义读取多少
    fmt.Println("------------------------------------------------")
    c := make([]byte, 20)
    n2, _ := f.ReadAt(c, 15)
    fmt.Println(n2)
    fmt.Println(string(c[:n2]))
    fmt.Println("------------------------------------------------")
    //(f *File).Readdir()函数原型func (f *File) Readdir(n int) (fi []FileInfo, err error)，我们要打开一个文件夹，然后设置读取文件夹文件的个数，返回的是文件的fileinfo信息

    f2, err := os.Open("src") //打开一个目录
    if err != nil {
        fmt.Println(err)
    }
    defer f2.Close()
    nums, _ := f2.Readdir(10) //设置读取的数量 <=0是读取所有的文件 返回的[]fileinfo
    for i, num := range nums {
        fmt.Printf("filename %d : %+v\n", i, num.Name()) //我们输出文件的名称

    }
    //(f *File).Readdirnames这个函数的作用是读取目录内的文件名，其实上一个函数我们已经实现了这个函数的功能，函数的原型func (f *File) Readdirnames(n int) (names []string, err error)，跟上边一下只不过返回的是文件名 []string的slice
    fmt.Println("--------------------------------------------")
    f3, _ := os.Open("src")
    defer f3.Close()
    names, err := f3.Readdirnames(0)
    if err != nil {
        fmt.Println(err)
    }
    for i, name := range names {
        fmt.Printf("filename %d : %s\n", i, name)
    }

    fmt.Println("--------------------------------------------")
    //(f *File).Seek()这个函数大家一看就懂了，就是偏移指针的地址，函数的原型是func (f *File) Seek(offset int64, whence int) (ret int64, err error) 其中offset是文件指针的位置 whence为0时代表相对文件开始的位置，1代表相对当前位置，2代表相对文件结尾的位置 ret返回的是现在指针的位置

    e := make([]byte, 10)
    f4, _ := os.Open("docker.txt") //相当于开始位置偏移1

    defer f4.Close()
    f4.Seek(1, 0)
    n4, _ := f4.Read(e)
    fmt.Println(string(e[:n4])) //原字符package 输出ackage
    //(f *File).Stat()其中跟前边的os.Stat()一样都是返回Fileinfo所以不多讲了
    //(f *File).Truncate()这个函数跟前边的os.Truncate()函数是一样的我就不多讲了
    //(f *File) Write像文件中写入内容，函数原型func (f *File) Write(b []byte) (n int, err error)返回的是n写入的字节数

}
