package request

import (
	"smart-admin/server/model/common/request"
	"smart-admin/server/model/twitter"
)

type TTwitterUserSearch struct {
	twitter.TTwitterUser
	request.PageInfo
}
