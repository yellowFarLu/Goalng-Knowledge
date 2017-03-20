package mysql

// 进行测试的

import (
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
)

func findByPk(pk int) int {
    var num int = 0
    db, err := sql.Open("mysql", "root:1234@tcp(127.0.0.1:3306)/test?charset=utf8")
    if err != nil {
        panic(err.Error())
    }
    defer db.Close()
    return num
}