package request

import (
	"smart-admin/server/model/common/request"
	"smart-admin/server/model/twitter"
)

type TSystemSearch struct {
	twitter.TSystem
	request.PageInfo
}
