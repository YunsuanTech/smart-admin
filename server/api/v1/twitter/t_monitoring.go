package twitter

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"smart-admin/server/global"
	"smart-admin/server/model/common/request"
	"smart-admin/server/model/common/response"
	"smart-admin/server/model/twitter"
	twitterReq "smart-admin/server/model/twitter/request"
)

type TMonitoringApi struct {
}

// CreateTMonitoring 创建TMonitoring
// @Tags TMonitoring
// @Summary 创建TMonitoring
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.TMonitoring true "创建TMonitoring"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /tMonitoring/createTMonitoring [post]
func (tMonitoringApi *TMonitoringApi) CreateTMonitoring(c *gin.Context) {
	var tMonitoring twitter.TMonitoring
	_ = c.ShouldBindJSON(&tMonitoring)
	if err := tMonitoringService.CreateTMonitoring(tMonitoring); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteTMonitoring 删除TMonitoring
// @Tags TMonitoring
// @Summary 删除TMonitoring
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.TMonitoring true "删除TMonitoring"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /tMonitoring/deleteTMonitoring [delete]
func (tMonitoringApi *TMonitoringApi) DeleteTMonitoring(c *gin.Context) {
	var tMonitoring twitter.TMonitoring
	_ = c.ShouldBindJSON(&tMonitoring)
	if err := tMonitoringService.DeleteTMonitoring(tMonitoring); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteTMonitoringByIds 批量删除TMonitoring
// @Tags TMonitoring
// @Summary 批量删除TMonitoring
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除TMonitoring"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /tMonitoring/deleteTMonitoringByIds [delete]
func (tMonitoringApi *TMonitoringApi) DeleteTMonitoringByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := tMonitoringService.DeleteTMonitoringByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateTMonitoring 更新TMonitoring
// @Tags TMonitoring
// @Summary 更新TMonitoring
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.TMonitoring true "更新TMonitoring"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /tMonitoring/updateTMonitoring [put]
func (tMonitoringApi *TMonitoringApi) UpdateTMonitoring(c *gin.Context) {
	var tMonitoring twitter.TMonitoring
	_ = c.ShouldBindJSON(&tMonitoring)
	if err := tMonitoringService.UpdateTMonitoring(tMonitoring); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindTMonitoring 用id查询TMonitoring
// @Tags TMonitoring
// @Summary 用id查询TMonitoring
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query autocode.TMonitoring true "用id查询TMonitoring"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /tMonitoring/findTMonitoring [get]
func (tMonitoringApi *TMonitoringApi) FindTMonitoring(c *gin.Context) {
	var tMonitoring twitter.TMonitoring
	_ = c.ShouldBindQuery(&tMonitoring)
	if err, retMonitoring := tMonitoringService.GetTMonitoring(tMonitoring.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"retMonitoring": retMonitoring}, c)
	}
}

// GetTMonitoringList 分页获取TMonitoring列表
// @Tags TMonitoring
// @Summary 分页获取TMonitoring列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query autocodeReq.TMonitoringSearch true "分页获取TMonitoring列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /tMonitoring/getTMonitoringList [get]
func (tMonitoringApi *TMonitoringApi) GetTMonitoringList(c *gin.Context) {
	var pageInfo twitterReq.TMonitoringSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := tMonitoringService.GetTMonitoringInfoList(pageInfo); err != nil {
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
