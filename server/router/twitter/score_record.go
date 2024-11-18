package Twitter

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type ScoreRecordRouter struct {}

// InitScoreRecordRouter 初始化 scoreRecord表 路由信息
func (s *ScoreRecordRouter) InitScoreRecordRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	scoreRecordRouter := Router.Group("scoreRecord").Use(middleware.OperationRecord())
	scoreRecordRouterWithoutRecord := Router.Group("scoreRecord")
	scoreRecordRouterWithoutAuth := PublicRouter.Group("scoreRecord")
	{
		scoreRecordRouter.POST("createScoreRecord", scoreRecordApi.CreateScoreRecord)   // 新建scoreRecord表
		scoreRecordRouter.DELETE("deleteScoreRecord", scoreRecordApi.DeleteScoreRecord) // 删除scoreRecord表
		scoreRecordRouter.DELETE("deleteScoreRecordByIds", scoreRecordApi.DeleteScoreRecordByIds) // 批量删除scoreRecord表
		scoreRecordRouter.PUT("updateScoreRecord", scoreRecordApi.UpdateScoreRecord)    // 更新scoreRecord表
	}
	{
		scoreRecordRouterWithoutRecord.GET("findScoreRecord", scoreRecordApi.FindScoreRecord)        // 根据ID获取scoreRecord表
		scoreRecordRouterWithoutRecord.GET("getScoreRecordList", scoreRecordApi.GetScoreRecordList)  // 获取scoreRecord表列表
	}
	{
	    scoreRecordRouterWithoutAuth.GET("getScoreRecordPublic", scoreRecordApi.GetScoreRecordPublic)  // scoreRecord表开放接口
	}
}
