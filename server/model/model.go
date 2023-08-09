package model

import "gorm.io/gorm"

type Model struct {
	gorm.Model
	CreateBy string `json:"create_by;comment:创建者"`
	UpdateBy string `json:"update_by;comment:删除者"`
}

//func NewDBEngine(databasesSetting *setting.DatabasesSetting) (gorm.Model, error) {
//
//}
