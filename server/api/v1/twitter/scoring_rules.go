package Twitter

import (
	
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/model/Twitter"
    TwitterReq "github.com/flipped-aurora/gin-vue-admin/server/model/Twitter/request"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

type ScoringRulesApi struct {}



// CreateScoringRules 创建scoringRules表
// @Tags ScoringRules
// @Summary 创建scoringRules表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body Twitter.ScoringRules true "创建scoringRules表"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /scoringRules/createScoringRules [post]
func (scoringRulesApi *ScoringRulesApi) CreateScoringRules(c *gin.Context) {
	var scoringRules Twitter.ScoringRules
	err := c.ShouldBindJSON(&scoringRules)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = scoringRulesService.CreateScoringRules(&scoringRules)
	if err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:" + err.Error(), c)
		return
	}
    response.OkWithMessage("创建成功", c)
}

// DeleteScoringRules 删除scoringRules表
// @Tags ScoringRules
// @Summary 删除scoringRules表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body Twitter.ScoringRules true "删除scoringRules表"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /scoringRules/deleteScoringRules [delete]
func (scoringRulesApi *ScoringRulesApi) DeleteScoringRules(c *gin.Context) {
	id := c.Query("id")
	err := scoringRulesService.DeleteScoringRules(id)
	if err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteScoringRulesByIds 批量删除scoringRules表
// @Tags ScoringRules
// @Summary 批量删除scoringRules表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /scoringRules/deleteScoringRulesByIds [delete]
func (scoringRulesApi *ScoringRulesApi) DeleteScoringRulesByIds(c *gin.Context) {
	ids := c.QueryArray("ids[]")
	err := scoringRulesService.DeleteScoringRulesByIds(ids)
	if err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateScoringRules 更新scoringRules表
// @Tags ScoringRules
// @Summary 更新scoringRules表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body Twitter.ScoringRules true "更新scoringRules表"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /scoringRules/updateScoringRules [put]
func (scoringRulesApi *ScoringRulesApi) UpdateScoringRules(c *gin.Context) {
	var scoringRules Twitter.ScoringRules
	err := c.ShouldBindJSON(&scoringRules)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = scoringRulesService.UpdateScoringRules(scoringRules)
	if err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindScoringRules 用id查询scoringRules表
// @Tags ScoringRules
// @Summary 用id查询scoringRules表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query Twitter.ScoringRules true "用id查询scoringRules表"
// @Success 200 {object} response.Response{data=Twitter.ScoringRules,msg=string} "查询成功"
// @Router /scoringRules/findScoringRules [get]
func (scoringRulesApi *ScoringRulesApi) FindScoringRules(c *gin.Context) {
	id := c.Query("id")
	rescoringRules, err := scoringRulesService.GetScoringRules(id)
	if err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:" + err.Error(), c)
		return
	}
	response.OkWithData(rescoringRules, c)
}

// GetScoringRulesList 分页获取scoringRules表列表
// @Tags ScoringRules
// @Summary 分页获取scoringRules表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query TwitterReq.ScoringRulesSearch true "分页获取scoringRules表列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /scoringRules/getScoringRulesList [get]
func (scoringRulesApi *ScoringRulesApi) GetScoringRulesList(c *gin.Context) {
	var pageInfo TwitterReq.ScoringRulesSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := scoringRulesService.GetScoringRulesInfoList(pageInfo)
	if err != nil {
	    global.GVA_LOG.Error("获取失败!", zap.Error(err))
        response.FailWithMessage("获取失败:" + err.Error(), c)
        return
    }
    response.OkWithDetailed(response.PageResult{
        List:     list,
        Total:    total,
        Page:     pageInfo.Page,
        PageSize: pageInfo.PageSize,
    }, "获取成功", c)
}

// GetScoringRulesPublic 不需要鉴权的scoringRules表接口
// @Tags ScoringRules
// @Summary 不需要鉴权的scoringRules表接口
// @accept application/json
// @Produce application/json
// @Param data query TwitterReq.ScoringRulesSearch true "分页获取scoringRules表列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /scoringRules/getScoringRulesPublic [get]
func (scoringRulesApi *ScoringRulesApi) GetScoringRulesPublic(c *gin.Context) {
    // 此接口不需要鉴权
    // 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
    scoringRulesService.GetScoringRulesPublic()
    response.OkWithDetailed(gin.H{
       "info": "不需要鉴权的scoringRules表接口信息",
    }, "获取成功", c)
}
