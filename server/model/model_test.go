package model

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"testing"
)

func Testing(t *testing.T) {
	dsn := "root:123456@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(Tag{})
}
