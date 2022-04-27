// 自动生成模板TPlatform
package compare

import (
	"smart-admin/server/global"
)

// TPlatform 结构体
// 如果含有time.Time 请自行import time包
type TPlatform struct {
	global.GVA_MODEL
	Name         string `json:"name" form:"name" gorm:"column:name;comment:名称;"`
	Type         *int   `json:"type" form:"type" gorm:"column:type;comment:1 平台 2 交易所;size:64;"`
	ChainId      string `json:"chainId" form:"chainId" gorm:"column:chain_id;comment:链id;"`
	RpcUrl       string `json:"rpcUrl" form:"rpcUrl" gorm:"column:rpc_url;comment:rpc_url;"`
	Symbol       string `json:"symbol" form:"symbol" gorm:"column:symbol;comment:符号;"`
	BlockBrowser string `json:"blockBrowser" form:"blockBrowser" gorm:"column:block_browser;comment:区块浏览器;"`
}

// TableName TPlatform 表名
func (TPlatform) TableName() string {
	return "t_platform"
}
