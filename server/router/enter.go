package router

import (
	"smart-admin/server/router/autocode"
	"smart-admin/server/router/compare"
	"smart-admin/server/router/example"
	"smart-admin/server/router/system"
	"smart-admin/server/router/twitter"
)

type RouterGroup struct {
	System   system.RouterGroup
	Example  example.RouterGroup
	Autocode autocode.RouterGroup
	Twitter  twitter.RouterGroup
	Compare  compare.RouterGroup
}

var RouterGroupApp = new(RouterGroup)
