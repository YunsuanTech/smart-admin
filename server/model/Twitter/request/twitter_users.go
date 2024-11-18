package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/Twitter"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type TwitterUsersSearch struct {
	request.PageInfo
	Twitter.TwitterUsers
}
