package main

import (
    "flag"
    "fmt"
    "log"
    "os"
    "runtime"
)

var (
    logFileName = flag.String("log", "cServer.log", "Log file name")
)

func main() {
    runtime.GOMAXPROCS(runtime.NumCPU())
    flag.Parse()

    //set logfile Stdout
    logFile, logErr := os.OpenFile(*logFileName, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
    if logErr != nil {
        fmt.Println("Fail to find", *logFile, "cServer start Failed")
        os.Exit(1)
    }
    log.SetOutput(logFile)
    log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

    //write log
    log.Printf("Server abort! Cause:%v \n", "test log file")

}
/*
一、在主程序启动之前初始化log的格式后，以后用log.xxx 都将记录在初始的cServer.log文件中

二、上面的初始化的格式在log文件显示为：

2013/03/19 10:44:26 main.go:29: Server abort! Cause:test log file
2013/03/19 10:44:27 main.go:29: Server abort! Cause:test log file

三、可以通过log.SetFlags()自定议你想要表达的格式

四、设置输出目的地log.SetOutput()

五、os.OpenFile(*logFileName, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)这是创建log文件.

5.1如果log文件不存在，创建一个新的文件os.O_CREATE

5.2打开文件的读写os.O_RDWR

5.3将log信息写入到log文件，是继承到当前log文件，不是覆盖os.O_APPEND

5.3log文件的权限位0666（即所有用户可读写）
*/
