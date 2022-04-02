package twitter

import (
	"smart-admin/server/service"
)

type ApiGroup struct {
	TSystemApi
	TMonitoringApi
	TTwitterUserApi
	TTwitterContentApi
}

var (
	tSystemService         = service.ServiceGroupApp.TwitterServiceGroup.TSystemService
	tMonitoringService     = service.ServiceGroupApp.TwitterServiceGroup.TMonitoringService
	tTwitterUserService    = service.ServiceGroupApp.TwitterServiceGroup.TTwitterUserService
	tTwitterContentService = service.ServiceGroupApp.TwitterServiceGroup.TTwitterContentService
)
