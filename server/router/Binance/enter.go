package Binance

import api "github.com/flipped-aurora/gin-vue-admin/server/api/v1"

type RouterGroup struct {
	ArticlesRouter
	TokensRouter
}

var (
	articlesApi = api.ApiGroupApp.BinanceApiGroup.ArticlesApi
	tokensApi   = api.ApiGroupApp.BinanceApiGroup.TokensApi
)
