package request

import (
	"smart-admin/server/model/common/request"
	"smart-admin/server/model/compare"
)

type TPlatformSearch struct {
	compare.TPlatform
	request.PageInfo
}
