package compare

import (
	"smart-admin/server/service"
)

type ApiGroup struct {
	TPlatformApi
}

var (
	tPlatformService = service.ServiceGroupApp.CompareServiceGroup.TPlatformService
)
