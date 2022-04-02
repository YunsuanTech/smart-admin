package request

import (
	"smart-admin/server/model/common/request"
	"smart-admin/server/model/twitter"
)

type TMonitoringSearch struct {
	twitter.TMonitoring
	request.PageInfo
}
