package repo

import (
	"database/sql"
	"fmt"
	"net/url"
)

var DB *sql.DB

func init() {
	dbHost := "localhost"
	dbPort := "3306"
	dbUser := "root"
	dbPass := "root"
	dbName := "test"
	connStr := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)
	val := url.Values{}
	val.Add("parseTime", "1")
	val.Add("loc", "Asia/Shanghai")
	dsn := fmt.Sprintf("%s?%s", connStr, val.Encode())
	dbConn, err := sql.Open(`mysql`, dsn)
	if err != nil {
		panic(err)
	}
	if err := dbConn.Ping(); err != nil {
		panic(err)
	}

	DB = dbConn
}
