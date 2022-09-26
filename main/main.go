package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func main() {
	database, err := sqlx.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/mysql")
	if err != nil {
		fmt.Println("open mysql failed,", err)
		return
	}
	fmt.Println(database)
}
