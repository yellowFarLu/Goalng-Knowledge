package chapter04

import (
    // "syscall"
    "time"
    "fmt"
)

// 设置环境变量 `TZ`
func setTZ()  {
    // tz, ok := syscall.Getenv("TZ")
    // switch {
    // case !ok:
    //     z, err := loadZoneFile("", "/etc/localtime")
    //     if err == nil {
    //         localLoc = *z
    //         localLoc.name = "Local"
    //         return
    //     }
    // case tz != "" && tz != "UTC":
    //     if z, err := loadLocation(tz); err == nil {
    //         localLoc = *z
    //         return
    //     }
    // }
}

func ParseTimeTest()  {
    // 获取当前时间
    now := time.Now().Unix();
    fmt.Println(now);

    // 将时间戳变成需要的类型
    // 1、将时间戳转化成time类型
    tm := time.Unix(now, 0);
    str := tm.Format("2006-01-02 03:04:05 PM");
    fmt.Println(str);

    // 第一个参数指定时间格式，函数把第二个参数的时间转化成时间戳
    // "2006-01-02 15:04:05" 固定写法
    t, _ := time.Parse("2006-01-02 15:04:05", "2016-06-13 09:14:00")
    fmt.Println(time.Now().Sub(t).Hours())
}