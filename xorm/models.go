package xorm
import (
    _ "github.com/go-sql-driver/mysql"
    "github.com/go-xorm/xorm"
    "log"
    "os"
    "errors"
    "fmt"
)
// 创建账户
type Account struct {
    Id int64
    Name string `xorm:"unique"`
    Balance float64
    Version int `xorm:"version"` //乐观锁
}
// ORM 引擎
var engine *xorm.Engine
func init() {
    // 创建 ORM 引擎和数据库
    var err error
    engine, err = xorm.NewEngine("mysql", "root:xxxx@/gomysql?charset=utf8")

    if err !=nil {
        log.Fatal("Fail to create engine:%v\n", err)
    }

    // 同步结构体和数据表
    if err := engine.Sync(new(Account)); err!=nil {
        log.Fatal("Fail to sync database: %v\n", err)
    }
    //记录日志

    //    = os.Create("sql.txt")

    f, err := os.OpenFile("sql.txt", os.O_RDWR|os.O_APPEND, 0755)
    if err!=nil {
        f, _= os.Create("sql.txt")
        f, _ = os.OpenFile("sql.txt", os.O_RDWR|os.O_APPEND, 0755)
    }

    engine.ShowSQL=true
    engine.ShowInfo=false
    engine.Logger=xorm.NewSimpleLogger(f)


}
// 创建账号
func NewAccount(name string, balance float64) error {
    _, err := engine.Insert(&Account{Name:name, Balance:balance})
    return err
}
// 获取账号数量
func GetAccountCount() (int64, error) {
    return engine.Count(new(Account))
}
//获取账号信息
func GetAccount(id int64) (*Account, error) {
    a := &Account{}
    // 直接操作 ID 的简便方法
    has, err := engine.Id(id).Get(a)
    //判断操作是否发生错误或对象是否存在
    if err !=nil {
        return nil, err
    }else if !has {
        return nil, errors.New("Account does not exist")
    }
    return a, nil
}
// 按照 ID  正序排序返回所有账号
func GetAccountAscId() (as []Account, err error) {
    // 使用 Find 方法批量获取记录
    err =engine.Find(&as)
    return as, err
}
// 删除账户
func DeleteAccount(id int64) error {
    _, err := engine.Delete(&Account{Id:id})
    return err
}
// 用户存款
func MakeDeposit(id int64, deposit float64) (*Account, error) {
    a, err := GetAccount(id)
    if err!=nil {
        return nil, err
    }
    a.Balance+=deposit
    fmt.Println(a.Balance)
    //对已有记录进行更新
    _, err =engine.Id(a.Id).Update(a)
    return a, err
}
// 用户取款
func MakeWithdraw(id int64, withdraw float64) (*Account, error) {
    a, err := GetAccount(id)
    if err!=nil {
        return nil, err
    }
    if a.Balance<withdraw {
        return nil, errors.New("Not enough balance")
    }
    a.Balance-=withdraw
    fmt.Println(a.Balance)
    _, err =engine.Id(a.Id).Update(a)
    return a, err
}
// 用户转账
func MakeTransfer(id1, id2 int64, balance float64) error {
    a1, err := GetAccount(id1)
    if err!=nil {
        return err
    }
    a2, err := GetAccount(id2)
    if err!=nil {
        return err
    }
    if a1.Balance<balance {
        return errors.New("Not enough blance")
    }

    a1.Balance-=balance
    a2.Balance+=balance
    // 创建 Session 对象，
    sess := engine.NewSession()
    defer sess.Close()
    //启动事务
    if err =sess.Begin(); err !=nil {
        return err
    }
    if _, err =sess.Id(a1.Id).Update(a1); err !=nil {
        // 发送错误时候进行回滚
        return err
    }else if _, err=sess.Id(a2.Id).Update(a2); err !=nil {
        return err
    }
    // 完成事务
    return sess.Commit()

}