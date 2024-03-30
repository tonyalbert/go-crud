package database

import (
    "database/sql"
	"fmt"
    _ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func InitDB() error {
    var err error
    DB, err = sql.Open("mysql", "root:root@tcp(localhost:3306)/mystore")
    if err != nil {
        return err
    }
    err = DB.Ping()
    if err != nil {
        return err
    }

	fmt.Println("Successfully connected!")
    return nil
}

func CloseDB() error {
    if DB != nil {
		fmt.Println("Closing DB connection")
        return DB.Close()
    }
    return nil
}