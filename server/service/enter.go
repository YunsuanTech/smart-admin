package service

import (
	"smart-admin/server/service/autocode"
	"smart-admin/server/service/example"
	"smart-admin/server/service/system"
	"smart-admin/server/service/twitter"
)

type ServiceGroup struct {
	SystemServiceGroup   system.ServiceGroup
	ExampleServiceGroup  example.ServiceGroup
	AutoCodeServiceGroup autocode.ServiceGroup
	TwitterServiceGroup  twitter.ServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)
