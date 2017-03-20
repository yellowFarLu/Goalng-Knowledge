package chapter07

import (
	"database/sql"
    "fmt"
)

func SqlTest()  {				 // '/'后面是数据库名字
	db, err := sql.Open("mysql", "root:1234@tcp(localhost:3306)/test?charset=utf8")
    if err!=nil {
        fmt.Println(err);
    }

    // 创建数据库
    rs, err := db.Exec("create database if not exists test");
    if err!=nil {
        fmt.Println(err)
    }
    fmt.Println(rs.RowsAffected());

    // 创建表
    rs, err = db.Exec("create table if not exists Son(name text);");
    if err!=nil {
        fmt.Println(err)
    }
    fmt.Println(rs.RowsAffected());
}