package main

import (
    "github.com/hydra13142/webui"
    "net/http"
    "strconv"
    "strings"
)

type Cond struct {
    last float64
    op   string
    in   bool
}

func NewCond() interface{} {
    return &Cond{0, "=", false}
}

func main() {
    w := &webui.Window{
        Width:  180,
        Height: 220,
        Sub: []webui.Object{
            &webui.Text{Common: webui.Common{"tx", "0", 10, 10, 160, 40, nil}, Readonly: true},
            &webui.Button{Common: webui.Common{"n0", "0", 10, 170, 40, 40, input0}},
            &webui.Button{Common: webui.Common{"n1", "1", 10, 130, 40, 40, input("1")}},
            &webui.Button{Common: webui.Common{"n2", "2", 50, 130, 40, 40, input("2")}},
            &webui.Button{Common: webui.Common{"n3", "3", 90, 130, 40, 40, input("3")}},
            &webui.Button{Common: webui.Common{"n4", "4", 10, 90, 40, 40, input("4")}},
            &webui.Button{Common: webui.Common{"n5", "5", 50, 90, 40, 40, input("5")}},
            &webui.Button{Common: webui.Common{"n6", "6", 90, 90, 40, 40, input("6")}},
            &webui.Button{Common: webui.Common{"n7", "7", 10, 50, 40, 40, input("7")}},
            &webui.Button{Common: webui.Common{"n8", "8", 50, 50, 40, 40, input("8")}},
            &webui.Button{Common: webui.Common{"n9", "9", 90, 50, 40, 40, input("9")}},
            &webui.Button{Common: webui.Common{"pt", ".", 50, 170, 40, 40, inputd}},
            &webui.Button{Common: webui.Common{"o1", "+", 130, 50, 40, 40, operate("+")}},
            &webui.Button{Common: webui.Common{"o2", "-", 130, 90, 40, 40, operate("-")}},
            &webui.Button{Common: webui.Common{"o3", "*", 130, 130, 40, 40, operate("*")}},
            &webui.Button{Common: webui.Common{"o4", "/", 130, 170, 40, 40, operate("/")}},
            &webui.Button{Common: webui.Common{"eq", "=", 90, 170, 40, 40, operate("=")}},
        },
    }
    http.ListenAndServe(":9999", webui.NewHandler(w, "calc.htm", NewCond))
}

func exec(a, b float64, o string) float64 {
    switch o {
        case "+":
        return a + b
        case "-":
        return a - b
        case "*":
        return a * b
        case "/":
        return a / b
        case "=":
        return b
    }
    return a
}
func input(t string) func(c *webui.Context) {
    return func(c *webui.Context) {
        hold := c.Hold.(*Cond)
        if hold.in {
            s := c.Para["tx"]
            if len(s) >= 16 {
                c.Err = "line too long"
                return
            }
            if s[0] == '0' && !strings.ContainsRune(s, '.') {
                s = s[1:]
            }
            c.Ans["tx"] = s + t
        } else {
            hold.in = true
            c.Ans["tx"] = t
        }
    }
}
func operate(o string) func(c *webui.Context) {
    return func(c *webui.Context) {
        hold := c.Hold.(*Cond)
        x, _ := strconv.ParseFloat(c.Para["tx"], 64)
        x = exec(hold.last, x, hold.op)
        hold.last, hold.op, hold.in = x, o, false
        c.Ans["tx"] = strconv.FormatFloat(x, 'g', -1, 64)
    }
}
func input0(c *webui.Context) {
    hold := c.Hold.(*Cond)
    if hold.in {
        s := c.Para["tx"]
        if len(s) >= 16 {
            c.Err = "line too long"
            return
        }
        if s != "0" {
            s = s + "0"
        }
        c.Ans["tx"] = s
    } else {
        hold.in = true
        c.Ans["tx"] = "0"
    }
}
func inputd(c *webui.Context) {
    hold := c.Hold.(*Cond)
    if hold.in {
        s := c.Para["tx"]
        if len(s) >= 16 {
            c.Err = "line too long"
            return
        }
        if !strings.ContainsRune(s, '.') {
            s = s + "."
        }
        c.Ans["tx"] = s
    } else {
        hold.in = true
        c.Ans["tx"] = "0."
    }
}
