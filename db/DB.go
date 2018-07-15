package db

import (
	"github.com/jinzhu/gorm"
)

// DB control sql action
var DB *gorm.DB

// Init returns gorms.DB
func Init() (*gorm.DB, error) {
	db, err := gorm.Open("mysql", "root:@tcp(127.0.0.1:3306)/mysql?charset=utf8")
	//db, err := gorm.Open("mysql", "root:mysql@/wblog?charset=utf8&parseTime=True&loc=Asia/Shanghai")
	if err == nil {
		DB = db
		//db.LogMode(true)
		return db, err
	}
	return nil, err
}

// TODO
