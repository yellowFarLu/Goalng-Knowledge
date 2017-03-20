package chapter01

import (
    "fmt"
)

func TemDir()  {
     str := "Go"

     s := "Go";

     fmt.Printf("%p --- %p", &str, &s);
}