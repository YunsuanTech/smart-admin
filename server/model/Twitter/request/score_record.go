package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/Twitter"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type ScoreRecordSearch struct {
	request.PageInfo
	Twitter.ScoreRecord
}
