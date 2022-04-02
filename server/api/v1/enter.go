package v1

import (
	"smart-admin/server/api/v1/autocode"
	"smart-admin/server/api/v1/example"
	"smart-admin/server/api/v1/system"
	"smart-admin/server/api/v1/twitter"
)

type ApiGroup struct {
	SystemApiGroup   system.ApiGroup
	ExampleApiGroup  example.ApiGroup
	AutoCodeApiGroup autocode.ApiGroup
	TwitterApiGroup  twitter.ApiGroup
}

var ApiGroupApp = new(ApiGroup)
