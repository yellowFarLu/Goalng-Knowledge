package chapter01

import (
    "bufio"
    "fmt"
    "strings"
)
	
func Test() {
    reader := bufio.NewReaderSize(strings.NewReader("http://studygolang.com.\t It is the home of gophers"), 14)
    line, _ := reader.Peek(14)
    fmt.Printf("%s\n", line)
}

