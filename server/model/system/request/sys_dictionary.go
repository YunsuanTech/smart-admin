package request

import (
	"smart-admin/server/model/common/request"
	"smart-admin/server/model/system"
)

type SysDictionarySearch struct {
	system.SysDictionary
	request.PageInfo
}
