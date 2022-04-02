package twitter

import (
	"smart-admin/server/global"
)

type TMonitoring struct {
	global.GVA_MODEL
	IsEnable    *bool  `json:"isEnable" form:"isEnable" gorm:"column:is_enable;comment:是否启用;"`
	TiwtterName string `json:"tiwtterName" form:"tiwtterName" gorm:"column:tiwtter_name;comment:推特用户;size:100;"`
	UserTypeId  string `json:"userTypeId" form:"userTypeId" gorm:"column:user_type_id;comment:外键id;size:50;"`
}

// TableName TMonitoring 表名
func (TMonitoring) TableName() string {
	return "t_monitoring"
}
