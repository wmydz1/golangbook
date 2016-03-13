package main
import (

    "fmt"
    . "gowebprogram/xorm"
)
const prompt = `Please enter number of operation:
1. Create new account
2. Show detail of account
3. Deposit
4. Withdraw
5. Make transfer
6. List exist accounts by Id
7. List exist accounts by balance
8. Delete account
9. Exit`

func main() {
    fmt.Println("Welcome bank of xorm!")
    for {
        fmt.Println(prompt)
        //获取用户输入并转换成数字
        var num int
        fmt.Scanf("%d\n", &num)
        //判断用户选择
        switch num{
            case 1:
            fmt.Println("Please enter <name> <balance>:")
            var name string
            var balance float64
            fmt.Scanf("%s %f\n", &name, &balance)
            if err := NewAccount(name, balance); err !=nil {
                fmt.Println("Fail to create new account:", err)
            }else {
                fmt.Println("New account has been created")
            }
            case 2:
            fmt.Println("Please enter <id>:")
            var id int64
            fmt.Scanf("%d\n", &id)
            a, err := GetAccount(id)
            if err !=nil {
                fmt.Println("Fail to get account:", err)
            }else {
                fmt.Printf("%#v\n", a)
            }
            case 3:
            fmt.Println("Please enter <id> <deposit>:")
            var id int64
            var deposit float64
            fmt.Scanf("%d %f\n", &id, &deposit)
            a, err := MakeDeposit(id, deposit)
            if err!=nil {
                fmt.Println("Fail to deposit:", err)
            }else {
                fmt.Printf("%#v\n", a)
            }
            case 4:

            fmt.Println("Please enter <id>")
            case 5:

            fmt.Println("Please enter <id> <balance> <id>:")
            var id1, id2 int64
            var balance float64
            fmt.Scanf("%d %f %d\n", &id1, &balance, &id2)
            fmt.Println(id1, id2, balance)
            if err := MakeTransfer(id1, id2, balance); err !=nil {
                fmt.Println("Fail to transfer :", err)
            }else {
                fmt.Println("Transfer has been made")
            }
            case 6:

            as, err := GetAccountsAscId()
            if err!=nil {
                fmt.Println("Fail to get accounts:", err)
            }else {
                for i, a := range as {
                    fmt.Printf("%d: %#v\n", i, a)
                }
            }

        }
    }
}
