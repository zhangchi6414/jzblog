package model

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"jzblog/global"
	"jzblog/pkg/setting"
)

type Model struct {
	gorm.Model
	CreateBy string `json:"create_by;comment:创建者"`
	UpdateBy string `json:"update_by;comment:删除者"`
}

func NewDBEngine(databasesSetting *setting.DBSettings) (*gorm.DB, error) {
	s := "%s:%s@tcp(%s)/%s?charset=%s&parseTime=%t&loc=Local"
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN: fmt.Sprintf(
			s,
			databasesSetting.Username,
			databasesSetting.Password,
			databasesSetting.Host,
			databasesSetting.DBName,
			databasesSetting.Charset,
			databasesSetting.ParseTime)}))
	if err != nil {
		return nil, err
	}

	if global.ServerSetting.RunMode == "debug" {
		db.Debug()
	}
	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(databasesSetting.MaxIdleConns)
	sqlDB.SetMaxOpenConns(databasesSetting.MaxOpenConns)

	return db, nil
}
