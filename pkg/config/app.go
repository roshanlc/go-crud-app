package config

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB

// Database credentials
var (
	username = "root"
	password = "password"
	db       = "crud_app"
)

// Connect connects to the database
func Connect() {
	connString := fmt.Sprintf("%s:%s@/%s%s", username, password, db, "?charset=utf8&parseTime=True&loc=Local")
	db, err := gorm.Open("mysql", connString)

	if err != nil {
		panic(err)
	}

	DB = db
}

func GETDB() *gorm.DB {
	return DB
}
