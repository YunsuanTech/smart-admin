package twitter

import (
	"github.com/gin-gonic/gin"
	"smart-admin/server/api/v1"
	"smart-admin/server/middleware"
)

type TTwitterUserRouter struct {
}

// InitTTwitterUserRouter 初始化 TTwitterUser 路由信息
func (s *TTwitterUserRouter) InitTTwitterUserRouter(Router *gin.RouterGroup) {
	tTwitterUserRouter := Router.Group("tTwitterUser").Use(middleware.OperationRecord())
	tTwitterUserRouterWithoutRecord := Router.Group("tTwitterUser")
	var tTwitterUserApi = v1.ApiGroupApp.TwitterApiGroup.TTwitterUserApi
	{
		tTwitterUserRouter.POST("createTTwitterUser", tTwitterUserApi.CreateTTwitterUser)             // 新建TTwitterUser
		tTwitterUserRouter.DELETE("deleteTTwitterUser", tTwitterUserApi.DeleteTTwitterUser)           // 删除TTwitterUser
		tTwitterUserRouter.DELETE("deleteTTwitterUserByIds", tTwitterUserApi.DeleteTTwitterUserByIds) // 批量删除TTwitterUser
		tTwitterUserRouter.PUT("updateTTwitterUser", tTwitterUserApi.UpdateTTwitterUser)              // 更新TTwitterUser
	}
	{
		tTwitterUserRouterWithoutRecord.GET("findTTwitterUser", tTwitterUserApi.FindTTwitterUser)       // 根据ID获取TTwitterUser
		tTwitterUserRouterWithoutRecord.GET("getTTwitterUserList", tTwitterUserApi.GetTTwitterUserList) // 获取TTwitterUser列表
	}
}
