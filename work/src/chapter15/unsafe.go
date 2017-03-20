package chapter15

import (
	"fmt"
	"unsafe"
)

/*
func Sizeof(v ArbitraryType) uintptr
Sizeof函数返回变量v占用的内存空间的字节数，该字节数不是按照变量v实际占用的内存计算，
而是按照v的“top level”内存计算。比如，在64位系统中，如果变量v是int类型，会返回8，
因为v的“top level”内存就是它的值使用的内存；如果变量v是string类型，会返回16，
因为v的“top level”内存不是存放着实际的字符串，而是该字符串的地址；如果变量v是slice类型，会返回24，
这是因为slice的描述符就占了24个字节。
*/
func SizeofTest()  {
    fmt.Println(unsafe.Sizeof("wo"))
}



/*
- （1）任何类型的指针都可以被转化为Pointer
- （2）Pointer可以被转化为任何类型的指针
- （3）uintptr可以被转化为Pointer
- （4）Pointer可以被转化为uintptr
*/
func PointerTest()  {
    i := 100
    fmt.Println(i)  // 100
    p := (*int)(unsafe.Pointer(&i))
    fmt.Println(*p) // 100
    *p = 0
    fmt.Println(i)  // 0
    fmt.Println(*p) // 0
}