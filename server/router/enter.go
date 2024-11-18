package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/router/Binance"
	"github.com/flipped-aurora/gin-vue-admin/server/router/Twitter"
	"github.com/flipped-aurora/gin-vue-admin/server/router/example"
	"github.com/flipped-aurora/gin-vue-admin/server/router/system"
)

var RouterGroupApp = new(RouterGroup)

type RouterGroup struct {
	System  system.RouterGroup
	Example example.RouterGroup
	Twitter Twitter.RouterGroup
	Binance Binance.RouterGroup
}
