package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// DB control sql action
var DB *gorm.DB

// Init returns gorms.DB
func Init() (*gorm.DB, error) {
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return "prefix_" + defaultTableName
	}
	db, err := gorm.Open("mysql", "root:@tcp(127.0.0.1:3306)/mysql?charset=utf8&parseTime=True&loc=Asia/Shanghai")
	//db, err := gorm.Open("mysql", "root:mysql@/wblog?charset=utf8&parseTime=True&loc=Asia/Shanghai")
	if err == nil {
		DB = db
		//db.LogMode(true)
		db.DB().SetMaxIdleConns(10)
		db.DB().SetMaxOpenConns(100)
		db.CreateTable(&Ticker{})
		return db, err
	}
	return nil, err
}

// TODO
