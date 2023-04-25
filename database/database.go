package database

import (
	"fmt"
	"os"
	"sync"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var dbInstance *gorm.DB
var once sync.Once

func InitDB() *gorm.DB {
	once.Do(func() {
		db, err := gorm.Open(mysql.Open(getDSN()))

		if err != nil {
			panic(err)
		}
		dbInstance = db
	})
	return dbInstance
}

func getDSN() string {

	DSN := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"))

	fmt.Println(DSN)

	return DSN
}
