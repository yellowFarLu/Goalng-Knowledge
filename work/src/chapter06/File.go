package chapter06

import (
    "os"
	"fmt"
	"log"
    "syscall"
    "time"
)

func GetLastVisitTime()  {
    // Stat返回文件的相关信息
    fileInfo, err := os.Stat("./chapter06/test.txt")
    if err != nil {
        log.Fatal(err)
    }

    // Sys()获取底层数据来源
    sys := fileInfo.Sys()
    // 判断是不是syscall.Stat_t的类型，是的话，返回这个类型的变量
    stat := sys.(*syscall.Stat_t)
    fmt.Println(time.Unix(stat.Atimespec.Unix()))
}

func FileTruncate()  {
    file, _ := os.Open("./chapter06/test.txt");
    file.Truncate(1024);
    arr := make([]byte, 1024)
    file.Read(arr);
    // fmt.Println(string(arr));
    fmt.Println(arr);
}

func FileReaderTest()  {
    file, error := os.Open("./chapter06/test.txt");
    if error!=nil {
        // fmt.Println(error)
        // return;
    }

    arr := make([]byte, 1024)
    n,e := file.ReadAt(arr, 0);
    if e!=nil {
        fmt.Println(e);
        // return;
    }
    fmt.Println(n, arr)
}