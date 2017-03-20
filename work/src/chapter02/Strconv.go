package chapter02

import (
    "strconv"
    "fmt"
)

func PaseIntTest()  {
    // 字符串 
    // 进制
    // 参数 *bitSize* 表示的是整数取值范围，或者说整数的具体类型。
    // 取值 0、8、16、32 和 64 分别代表 int、int8、int16、int32 和 int64
    i, _ := strconv.ParseInt("5", 10, 0);
    fmt.Println(i);
}