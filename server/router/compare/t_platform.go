package compare

import (
	"github.com/gin-gonic/gin"
	"smart-admin/server/api/v1"
	"smart-admin/server/middleware"
)

type TPlatformRouter struct {
}

// InitTPlatformRouter 初始化 TPlatform 路由信息
func (s *TPlatformRouter) InitTPlatformRouter(Router *gin.RouterGroup) {
	tPlatformRouter := Router.Group("tPlatform").Use(middleware.OperationRecord())
	tPlatformRouterWithoutRecord := Router.Group("tPlatform")
	var tPlatformApi = v1.ApiGroupApp.CompareGroup.TPlatformApi
	{
		tPlatformRouter.POST("createTPlatform", tPlatformApi.CreateTPlatform)             // 新建TPlatform
		tPlatformRouter.DELETE("deleteTPlatform", tPlatformApi.DeleteTPlatform)           // 删除TPlatform
		tPlatformRouter.DELETE("deleteTPlatformByIds", tPlatformApi.DeleteTPlatformByIds) // 批量删除TPlatform
		tPlatformRouter.PUT("updateTPlatform", tPlatformApi.UpdateTPlatform)              // 更新TPlatform
	}
	{
		tPlatformRouterWithoutRecord.GET("findTPlatform", tPlatformApi.FindTPlatform)       // 根据ID获取TPlatform
		tPlatformRouterWithoutRecord.GET("getTPlatformList", tPlatformApi.GetTPlatformList) // 获取TPlatform列表
	}
}
