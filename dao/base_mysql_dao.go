package dao

import (
	"fmt"
	. "product-query/global"

	_ "github.com/go-sql-driver/mysql"
	gorm "github.com/jinzhu/gorm"
)

var mysqlDao *gorm.DB

func GetMysqlDao() *gorm.DB {
	if mysqlDao == nil {

		connectionUrl := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=true", Config.Database.Demo.User, Config.Database.Demo.Password, Config.Database.Demo.Host, Config.Database.Demo.Port, Config.Database.Demo.Dbname)

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
