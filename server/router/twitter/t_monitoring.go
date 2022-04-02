package twitter

import (
	"github.com/gin-gonic/gin"
	"smart-admin/server/api/v1"
	"smart-admin/server/middleware"
)

type TMonitoringRouter struct {
}

// InitTMonitoringRouter 初始化 TMonitoring 路由信息
func (s *TMonitoringRouter) InitTMonitoringRouter(Router *gin.RouterGroup) {
	tMonitoringRouter := Router.Group("tMonitoring").Use(middleware.OperationRecord())
	tMonitoringRouterWithoutRecord := Router.Group("tMonitoring")
	var tMonitoringApi = v1.ApiGroupApp.TwitterApiGroup.TMonitoringApi
	{
		tMonitoringRouter.POST("createTMonitoring", tMonitoringApi.CreateTMonitoring)             // 新建TMonitoring
		tMonitoringRouter.DELETE("deleteTMonitoring", tMonitoringApi.DeleteTMonitoring)           // 删除TMonitoring
		tMonitoringRouter.DELETE("deleteTMonitoringByIds", tMonitoringApi.DeleteTMonitoringByIds) // 批量删除TMonitoring
		tMonitoringRouter.PUT("updateTMonitoring", tMonitoringApi.UpdateTMonitoring)              // 更新TMonitoring
	}
	{
		tMonitoringRouterWithoutRecord.GET("findTMonitoring", tMonitoringApi.FindTMonitoring)       // 根据ID获取TMonitoring
		tMonitoringRouterWithoutRecord.GET("getTMonitoringList", tMonitoringApi.GetTMonitoringList) // 获取TMonitoring列表
	}
}
