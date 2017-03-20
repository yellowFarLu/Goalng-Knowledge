package chapter15

import (
    "unsafe"
	"fmt"
)

type Datas struct{
    c0 byte
    c1 int
    c2 string
    c3 int
}

func OffsetTest()  {
    var d Datas;
    d.c3 = 13
    // 生成指向结构体内存起始位置的指针
    p := unsafe.Pointer(&d)
    offset := unsafe.Offsetof(d.c3)
    q := (*int)(unsafe.Pointer(uintptr(p) + offset))
    fmt.Println(*q)

    *q = 10000
    fmt.Println(d.c3)
}