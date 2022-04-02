package twitter

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"smart-admin/server/global"
	"smart-admin/server/model/common/request"
	"smart-admin/server/model/common/response"
	"smart-admin/server/model/twitter"
	systemReq "smart-admin/server/model/twitter/request"
)

type TTwitterContentApi struct {
}

// CreateTTwitterContent 创建TTwitterContent
// @Tags TTwitterContent
// @Summary 创建TTwitterContent
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.TTwitterContent true "创建TTwitterContent"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /tTwitterContent/createTTwitterContent [post]
func (tTwitterContentApi *TTwitterContentApi) CreateTTwitterContent(c *gin.Context) {
	var tTwitterContent twitter.TTwitterContent
	_ = c.ShouldBindJSON(&tTwitterContent)
	if err := tTwitterContentService.CreateTTwitterContent(tTwitterContent); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteTTwitterContent 删除TTwitterContent
// @Tags TTwitterContent
// @Summary 删除TTwitterContent
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.TTwitterContent true "删除TTwitterContent"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /tTwitterContent/deleteTTwitterContent [delete]
func (tTwitterContentApi *TTwitterContentApi) DeleteTTwitterContent(c *gin.Context) {
	var tTwitterContent twitter.TTwitterContent
	_ = c.ShouldBindJSON(&tTwitterContent)
	if err := tTwitterContentService.DeleteTTwitterContent(tTwitterContent); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteTTwitterContentByIds 批量删除TTwitterContent
// @Tags TTwitterContent
// @Summary 批量删除TTwitterContent
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除TTwitterContent"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /tTwitterContent/deleteTTwitterContentByIds [delete]
func (tTwitterContentApi *TTwitterContentApi) DeleteTTwitterContentByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := tTwitterContentService.DeleteTTwitterContentByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateTTwitterContent 更新TTwitterContent
// @Tags TTwitterContent
// @Summary 更新TTwitterContent
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.TTwitterContent true "更新TTwitterContent"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /tTwitterContent/updateTTwitterContent [put]
func (tTwitterContentApi *TTwitterContentApi) UpdateTTwitterContent(c *gin.Context) {
	var tTwitterContent twitter.TTwitterContent
	_ = c.ShouldBindJSON(&tTwitterContent)
	if err := tTwitterContentService.UpdateTTwitterContent(tTwitterContent); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindTTwitterContent 用id查询TTwitterContent
// @Tags TTwitterContent
// @Summary 用id查询TTwitterContent
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query autocode.TTwitterContent true "用id查询TTwitterContent"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /tTwitterContent/findTTwitterContent [get]
func (tTwitterContentApi *TTwitterContentApi) FindTTwitterContent(c *gin.Context) {
	var tTwitterContent twitter.TTwitterContent
	_ = c.ShouldBindQuery(&tTwitterContent)
	if err, retTwitterContent := tTwitterContentService.GetTTwitterContent(tTwitterContent.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"retTwitterContent": retTwitterContent}, c)
	}
}

// GetTTwitterContentList 分页获取TTwitterContent列表
// @Tags TTwitterContent
// @Summary 分页获取TTwitterContent列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query autocodeReq.TTwitterContentSearch true "分页获取TTwitterContent列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /tTwitterContent/getTTwitterContentList [get]
func (tTwitterContentApi *TTwitterContentApi) GetTTwitterContentList(c *gin.Context) {
	var pageInfo systemReq.TTwitterContentSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := tTwitterContentService.GetTTwitterContentInfoList(pageInfo); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}
