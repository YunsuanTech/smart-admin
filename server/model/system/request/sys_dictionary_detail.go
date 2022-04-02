package request

import (
	"smart-admin/server/model/common/request"
	"smart-admin/server/model/system"
)

type SysDictionaryDetailSearch struct {
	system.SysDictionaryDetail
	request.PageInfo
}
