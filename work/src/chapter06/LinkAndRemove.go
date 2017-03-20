package chapter06

import (
	"os"
	"fmt"
)


func MakeDirTest()  {
    // `perm` 参数指定了新目录的权限
    error := os.Mkdir("super.hah", os.ModeType);
    fmt.Println(error)
}


func LinkTest()  {
      error:= os.Link("studygolang.txt", "super.txt");
      fmt.Println(error)
}

func RemoveLinkTest()  {
    e:= os.Remove("super.txt");
    fmt.Println(e)
}

func ReadLinkTest()  {
    str, e := os.Readlink("super.hah");
    fmt.Println(str, e);
}
