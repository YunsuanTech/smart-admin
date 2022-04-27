package compare

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"smart-admin/server/global"
	"smart-admin/server/model/common/request"
	"smart-admin/server/model/common/response"
	"smart-admin/server/model/compare"
	compareReq "smart-admin/server/model/compare/request"
)

type TPlatformApi struct {
}

// CreateTPlatform 创建TPlatform
// @Tags TPlatform
// @Summary 创建TPlatform
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.TPlatform true "创建TPlatform"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /tPlatform/createTPlatform [post]
func (tPlatformApi *TPlatformApi) CreateTPlatform(c *gin.Context) {
	var tPlatform compare.TPlatform
	_ = c.ShouldBindJSON(&tPlatform)
	if err := tPlatformService.CreateTPlatform(tPlatform); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteTPlatform 删除TPlatform
// @Tags TPlatform
// @Summary 删除TPlatform
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.TPlatform true "删除TPlatform"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /tPlatform/deleteTPlatform [delete]
func (tPlatformApi *TPlatformApi) DeleteTPlatform(c *gin.Context) {
	var tPlatform compare.TPlatform
	_ = c.ShouldBindJSON(&tPlatform)
	if err := tPlatformService.DeleteTPlatform(tPlatform); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteTPlatformByIds 批量删除TPlatform
// @Tags TPlatform
// @Summary 批量删除TPlatform
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除TPlatform"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /tPlatform/deleteTPlatformByIds [delete]
func (tPlatformApi *TPlatformApi) DeleteTPlatformByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := tPlatformService.DeleteTPlatformByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateTPlatform 更新TPlatform
// @Tags TPlatform
// @Summary 更新TPlatform
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.TPlatform true "更新TPlatform"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /tPlatform/updateTPlatform [put]
func (tPlatformApi *TPlatformApi) UpdateTPlatform(c *gin.Context) {
	var tPlatform compare.TPlatform
	_ = c.ShouldBindJSON(&tPlatform)
	if err := tPlatformService.UpdateTPlatform(tPlatform); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindTPlatform 用id查询TPlatform
// @Tags TPlatform
// @Summary 用id查询TPlatform
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query autocode.TPlatform true "用id查询TPlatform"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /tPlatform/findTPlatform [get]
func (tPlatformApi *TPlatformApi) FindTPlatform(c *gin.Context) {
	var tPlatform compare.TPlatform
	_ = c.ShouldBindQuery(&tPlatform)
	if err, retPlatform := tPlatformService.GetTPlatform(tPlatform.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"retPlatform": retPlatform}, c)
	}
}

// GetTPlatformList 分页获取TPlatform列表
// @Tags TPlatform
// @Summary 分页获取TPlatform列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query autocodeReq.TPlatformSearch true "分页获取TPlatform列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /tPlatform/getTPlatformList [get]
func (tPlatformApi *TPlatformApi) GetTPlatformList(c *gin.Context) {
	var pageInfo compareReq.TPlatformSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := tPlatformService.GetTPlatformInfoList(pageInfo); err != nil {
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
