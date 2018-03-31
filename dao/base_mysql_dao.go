package dao

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	gorm "github.com/jinzhu/gorm"
)

var mysqlDao *gorm.DB

func GetMysqlDao() *gorm.DB {
	if mysqlDao == nil {

		connectionUrl := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=true", "root", "1234", "localhost", 3306, "demo")

		if db, err := gorm.Open("mysql", connectionUrl); err != nil {
			panic(err.Error())
		} else {
			db.DB().SetMaxIdleConns(10)
			db.DB().SetMaxOpenConns(100)
			mysqlDao = db
		}
	}
	return mysqlDao
}
