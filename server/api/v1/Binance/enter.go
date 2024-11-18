package Binance

import "github.com/flipped-aurora/gin-vue-admin/server/service"

type ApiGroup struct {
	ArticlesApi
	TokensApi
}

var (
	articlesService = service.ServiceGroupApp.BinanceServiceGroup.ArticlesService
	tokensService   = service.ServiceGroupApp.BinanceServiceGroup.TokensService
)
