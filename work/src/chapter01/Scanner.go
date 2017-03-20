package chapter01

import (
    "bufio"
	"os"
	"fmt"
    "strings"
)

func Scanner()  {
    // os.Stdin相当于java中的System.in，这里通过scanner和Stdin联系起来，scanner能从控制条读入信息
    scanner := bufio.NewScanner(os.Stdin)
    // 相当于定义分割函数，一定要在输入前定义
    scanner.Split(bufio.ScanWords);

    // Scan() :读取一次token(装有分割后的一份字符串的byte[])
	if scanner.Scan() {
	    fmt.Println(scanner.Text()) 
	}
}

func Scanner2()  {
    const input = "This is The Golang Standard Library.\nWelcome you!"
	scanner := bufio.NewScanner(strings.NewReader(input))
	scanner.Split(bufio.ScanWords)
	count := 0
	for scanner.Scan() {
        fmt.Println(scanner.Text())
		count++
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading input:", err)
	}
	fmt.Println(count)
}