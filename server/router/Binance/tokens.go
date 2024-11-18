package Binance

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type TokensRouter struct {}

// InitTokensRouter 初始化 tokens表 路由信息
func (s *TokensRouter) InitTokensRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	tokensRouter := Router.Group("tokens").Use(middleware.OperationRecord())
	tokensRouterWithoutRecord := Router.Group("tokens")
	tokensRouterWithoutAuth := PublicRouter.Group("tokens")
	{
		tokensRouter.POST("createTokens", tokensApi.CreateTokens)   // 新建tokens表
		tokensRouter.DELETE("deleteTokens", tokensApi.DeleteTokens) // 删除tokens表
		tokensRouter.DELETE("deleteTokensByIds", tokensApi.DeleteTokensByIds) // 批量删除tokens表
		tokensRouter.PUT("updateTokens", tokensApi.UpdateTokens)    // 更新tokens表
	}
	{
		tokensRouterWithoutRecord.GET("findTokens", tokensApi.FindTokens)        // 根据ID获取tokens表
		tokensRouterWithoutRecord.GET("getTokensList", tokensApi.GetTokensList)  // 获取tokens表列表
	}
	{
	    tokensRouterWithoutAuth.GET("getTokensPublic", tokensApi.GetTokensPublic)  // tokens表开放接口
	}
}
