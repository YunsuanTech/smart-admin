package twitter

import (
	"github.com/gin-gonic/gin"
	"smart-admin/server/api/v1"
	"smart-admin/server/middleware"
)

type TTwitterContentRouter struct {
}

// InitTTwitterContentRouter 初始化 TTwitterContent 路由信息
func (s *TTwitterContentRouter) InitTTwitterContentRouter(Router *gin.RouterGroup) {
	tTwitterContentRouter := Router.Group("tTwitterContent").Use(middleware.OperationRecord())
	tTwitterContentRouterWithoutRecord := Router.Group("tTwitterContent")
	var tTwitterContentApi = v1.ApiGroupApp.TwitterApiGroup.TTwitterContentApi
	{
		tTwitterContentRouter.POST("createTTwitterContent", tTwitterContentApi.CreateTTwitterContent)             // 新建TTwitterContent
		tTwitterContentRouter.DELETE("deleteTTwitterContent", tTwitterContentApi.DeleteTTwitterContent)           // 删除TTwitterContent
		tTwitterContentRouter.DELETE("deleteTTwitterContentByIds", tTwitterContentApi.DeleteTTwitterContentByIds) // 批量删除TTwitterContent
		tTwitterContentRouter.PUT("updateTTwitterContent", tTwitterContentApi.UpdateTTwitterContent)              // 更新TTwitterContent
	}
	{
		tTwitterContentRouterWithoutRecord.GET("findTTwitterContent", tTwitterContentApi.FindTTwitterContent)       // 根据ID获取TTwitterContent
		tTwitterContentRouterWithoutRecord.GET("getTTwitterContentList", tTwitterContentApi.GetTTwitterContentList) // 获取TTwitterContent列表
	}
}
