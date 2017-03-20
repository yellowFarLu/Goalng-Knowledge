package main

import (
    // "github.com/astaxie/beego"
	"fmt"
    "encoding/json"
)

type Student struct {
    Name    string
    Age     int
    Guake   bool
    Classes []string
    Price   float32
}


func main() {
    st := &Student {
    "Xiao Ming",
    16,
    true,
    []string{"Math", "English", "Chinese"},
     9.99,
    }

    b, err := json.Marshal(st)
    if err!=nil {
        fmt.Println(err)
    }
    fmt.Println(string(b))

    fmt.Println("-----------------")

    p := &Student{}
    json.Unmarshal(b, p)
    fmt.Println(p)
}