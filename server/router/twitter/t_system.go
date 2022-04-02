package twitter

import (
	"github.com/gin-gonic/gin"
	v1 "smart-admin/server/api/v1"
	"smart-admin/server/middleware"
)

type TSystemRouter struct {
}

// InitTSystemRouter 初始化 TSystem 路由信息
func (s *TSystemRouter) InitTSystemRouter(Router *gin.RouterGroup) {
	tSystemRouter := Router.Group("tSystem").Use(middleware.OperationRecord())
	tSystemRouterWithoutRecord := Router.Group("tSystem")
	var tSystemApi = v1.ApiGroupApp.TwitterApiGroup.TSystemApi
	{
		tSystemRouter.POST("createTSystem", tSystemApi.CreateTSystem)             // 新建TSystem
		tSystemRouter.DELETE("deleteTSystem", tSystemApi.DeleteTSystem)           // 删除TSystem
		tSystemRouter.DELETE("deleteTSystemByIds", tSystemApi.DeleteTSystemByIds) // 批量删除TSystem
		tSystemRouter.PUT("updateTSystem", tSystemApi.UpdateTSystem)              // 更新TSystem
	}
	{
		tSystemRouterWithoutRecord.GET("findTSystem", tSystemApi.FindTSystem)       // 根据ID获取TSystem
		tSystemRouterWithoutRecord.GET("getTSystemList", tSystemApi.GetTSystemList) // 获取TSystem列表
		tSystemRouterWithoutRecord.GET("getTSystemAll", tSystemApi.GetTSystemAll)
	}
}
