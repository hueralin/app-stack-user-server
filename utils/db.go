package utils

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func DbConnect() *sql.DB {
	db, err := sql.Open("mysql", "root:123456789@(localhost:3306)/app-stack-user")
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(5)
	if err != nil {
		panic(err)
	}
	if err := db.Ping(); err != nil {
		fmt.Println("连接失败")
		panic(err.Error())
	}
	fmt.Println("数据库连接成功")
	return db
}
