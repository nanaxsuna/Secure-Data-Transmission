package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var database *sqlx.DB
var userName = "root"
var userPasswd = "123456"
var sqlAddr = "127.0.0.1:3306"

//init
/*
 * @author 	nana
 * @description init函数sql连接
 */
func init() {
	var err error
	dataSourceName := userName + ":" + userPasswd + "@tcp(" + sqlAddr + ")/mysql"

	database, err = sqlx.Open("mysql", dataSourceName)
	if err != nil {
		fmt.Println("open mysql failed,", err)
		return
	}
	fmt.Println("open mysql succeed,")
}

func main() {
	fmt.Println(database)
}
