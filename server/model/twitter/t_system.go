package twitter

import (
	"smart-admin/server/global"
)

// TSystem 结构体
// 如果含有time.Time 请自行import time包
type TSystem struct {
	global.GVA_MODEL
	Key    string `json:"key" form:"key" gorm:"column:key;comment:键;size:50;"`
	Value  string `json:"value" form:"value" gorm:"column:value;comment:值;size:50;"`
	Remark string `json:"remark" form:"remark" gorm:"column:remark;comment:备注;size:50;"`
}

// TableName TSystem 表名
func (TSystem) TableName() string {
	return "t_system"
}
