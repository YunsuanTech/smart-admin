package request

import (
	"smart-admin/server/model/common/request"
	"smart-admin/server/model/twitter"
)

type TTwitterContentSearch struct {
	twitter.TTwitterContent
	request.PageInfo
}
