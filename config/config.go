package config

import (
	"prakerja2/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	dsn := "root:d3KK45kDTvFgoykXjnTp@tcp(containers-us-west-126.railway.app:5600)/railway?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("connection DB")
	}
	migration()
}

func migration() {
	DB.AutoMigrate(&model.User{})
}
