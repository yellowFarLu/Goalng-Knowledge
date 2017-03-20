package chapter02

import (
    "unicode/utf8"
	"fmt"
)

func UTFTest()  {
    // 一个rune只是对应一个unicode编码，一个unicode编码映射着一个字符。
    arr := make([]byte, 5);
    n := utf8.EncodeRune(arr, '哈');
    fmt.Println(arr, n);

    hui, n := utf8.DecodeRune(arr);
    fmt.Println(string(hui), n);
}