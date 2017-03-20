package chapter02

import (
    // "strconv"
    "fmt"
    "strings"
)

func String()  {
    fmt.Println(strings.ContainsAny("team", "i"))
	fmt.Println(strings.ContainsAny("failure", "u & i"))
	fmt.Println(strings.ContainsAny("in failure", "s g"))
	fmt.Println(strings.ContainsAny("foo", ""))
	fmt.Println(strings.ContainsAny("", ""))
}

func JoinTest()  {
    fmt.Println(strings.Join([]string{"name=xxx", "age=xx"}, "&"))
}

func ReplaceTest()  {
    r := strings.NewReplacer("<", "&lt;", ">", "&gt;")
	fmt.Println(r.Replace("This is <b>HTML</b>!"))
}