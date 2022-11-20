package models

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

// SetupDB: initalizing mysql database
func SetupDB() *gorm.DB {
	USER := "root"
	PASSWORD := "$BUll34.01"
	HOST := "localhost"
	PORT := "3306"
	DBNAME := "test_db_for_test"
	URL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", USER, PASSWORD, HOST, PORT, DBNAME)

	db, err := gorm.Open("mysql", URL)

	if err != nil {
		// stop current gorountine
		panic(err.Error())
	}

	return db
}
