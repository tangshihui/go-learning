package repo

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net/url"
)

var DB *gorm.DB

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
	//dbConn, err := sql.Open(`mysql`, dsn)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Printf("failed connect to db.")
		panic(err)
	}
	//if err := db.pi(); err != nil {
	//	panic(err)
	//}

	DB = db
}
