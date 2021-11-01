package database

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

//SqlClient create a mysql connection
func SqlClient() *sql.DB {
	dbDriver := "mysql"
	dbUser := "task"
	dbPass := "12345"
	dbName := "EmployeeManagementSystem"
	dbPort := 3309
	db, err := sql.Open(dbDriver, fmt.Sprintf("%s:%s@tcp(0.0.0.0:%v)/%s", dbUser, dbPass, dbPort, dbName))
	if err != nil {
		log.Fatal("error connecting DB : ", err.Error())
	}
	return db
}