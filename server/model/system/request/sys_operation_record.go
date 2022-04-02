package request

import (
	"smart-admin/server/model/common/request"
	"smart-admin/server/model/system"
)

type SysOperationRecordSearch struct {
	system.SysOperationRecord
	request.PageInfo
}
