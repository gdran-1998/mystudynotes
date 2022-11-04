package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

// initMySQL 初始化MySQL连接
func initMySQL() (err error) {
	dsn := "root:123456@tcp(127.0.0.1:3306)/sql_test"
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		fmt.Printf("connect to db failed, err:%v\n", err)
		return
	}
	db.SetMaxOpenConns(100) // 最大连接数
	db.SetMaxIdleConns(10)  // 最大空闲连接数
	return
}

type user struct {
	id   int
	age  int
	name string
}

func main() {
	if err := initMySQL(); err != nil {
		fmt.Printf("connect to db failed, err:%v\n", err)
	}
	// Close() 用来释放掉数据库连接相关的资源
	defer db.Close() // 注意这行代码要写在上面err判断的下面
	fmt.Println("connect to db success")
}
