package chapter01

import (
    "bufio"
    "strings"
    "fmt"
)

func ReadSliceAndReadBytes() {
    reader := bufio.NewReader(strings.NewReader("http://studygolang.com\nIt is the home of gophers\ndasdadad\n"))
    // 让line指向buffer
	// line, error := reader.ReadSlice('\n')
    line, error := reader.ReadBytes('\n')
	fmt.Printf("the line:%s", line)
    fmt.Println(error)
    fmt.Println("------------------")

	// 让n指向buffer
	// n, error := reader.ReadSlice('\n')
    n, error := reader.ReadBytes('\n')
	fmt.Printf("the line:%s\n", line)
	fmt.Println(string(n))
    fmt.Println(error)
    fmt.Println("-----------------")

    // v, error:= reader.ReadSlice('\n')
    v, error := reader.ReadBytes('\n');
    fmt.Printf("the line:%s\n", line)
    fmt.Print(string(n))
    fmt.Println(string(v))
    fmt.Println(error)
}

func ReadString()  {
    // func (b *Reader) ReadString(delim byte) (line string, err error) {
	// 	bytes, err := b.ReadBytes(delim)
	// 	return string(bytes), err
	// }
}