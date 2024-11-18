package Twitter

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type ScoringRulesRouter struct {}

// InitScoringRulesRouter 初始化 scoringRules表 路由信息
func (s *ScoringRulesRouter) InitScoringRulesRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	scoringRulesRouter := Router.Group("scoringRules").Use(middleware.OperationRecord())
	scoringRulesRouterWithoutRecord := Router.Group("scoringRules")
	scoringRulesRouterWithoutAuth := PublicRouter.Group("scoringRules")
	{
		scoringRulesRouter.POST("createScoringRules", scoringRulesApi.CreateScoringRules)   // 新建scoringRules表
		scoringRulesRouter.DELETE("deleteScoringRules", scoringRulesApi.DeleteScoringRules) // 删除scoringRules表
		scoringRulesRouter.DELETE("deleteScoringRulesByIds", scoringRulesApi.DeleteScoringRulesByIds) // 批量删除scoringRules表
		scoringRulesRouter.PUT("updateScoringRules", scoringRulesApi.UpdateScoringRules)    // 更新scoringRules表
	}
	{
		scoringRulesRouterWithoutRecord.GET("findScoringRules", scoringRulesApi.FindScoringRules)        // 根据ID获取scoringRules表
		scoringRulesRouterWithoutRecord.GET("getScoringRulesList", scoringRulesApi.GetScoringRulesList)  // 获取scoringRules表列表
	}
	{
	    scoringRulesRouterWithoutAuth.GET("getScoringRulesPublic", scoringRulesApi.GetScoringRulesPublic)  // scoringRules表开放接口
	}
}
