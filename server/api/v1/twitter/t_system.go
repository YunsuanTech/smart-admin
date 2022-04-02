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

type TSystemApi struct {
}

// CreateTSystem 创建TSystem
// @Tags TSystem
// @Summary 创建TSystem
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.TSystem true "创建TSystem"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /tSystem/createTSystem [post]
func (tSystemApi *TSystemApi) CreateTSystem(c *gin.Context) {
	var tSystem twitter.TSystem
	_ = c.ShouldBindJSON(&tSystem)
	if err := tSystemService.CreateTSystem(tSystem); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteTSystem 删除TSystem
// @Tags TSystem
// @Summary 删除TSystem
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.TSystem true "删除TSystem"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /tSystem/deleteTSystem [delete]
func (tSystemApi *TSystemApi) DeleteTSystem(c *gin.Context) {
	var tSystem twitter.TSystem
	_ = c.ShouldBindJSON(&tSystem)
	if err := tSystemService.DeleteTSystem(tSystem); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteTSystemByIds 批量删除TSystem
// @Tags TSystem
// @Summary 批量删除TSystem
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除TSystem"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /tSystem/deleteTSystemByIds [delete]
func (tSystemApi *TSystemApi) DeleteTSystemByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := tSystemService.DeleteTSystemByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateTSystem 更新TSystem
// @Tags TSystem
// @Summary 更新TSystem
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.TSystem true "更新TSystem"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /tSystem/updateTSystem [put]
func (tSystemApi *TSystemApi) UpdateTSystem(c *gin.Context) {
	var tSystem twitter.TSystem
	_ = c.ShouldBindJSON(&tSystem)
	if err := tSystemService.UpdateTSystem(tSystem); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindTSystem 用id查询TSystem
// @Tags TSystem
// @Summary 用id查询TSystem
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query autocode.TSystem true "用id查询TSystem"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /tSystem/findTSystem [get]
func (tSystemApi *TSystemApi) FindTSystem(c *gin.Context) {
	var tSystem twitter.TSystem
	_ = c.ShouldBindQuery(&tSystem)
	if err, retSystem := tSystemService.GetTSystem(tSystem.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"retSystem": retSystem}, c)
	}
}

// GetTSystemList 分页获取TSystem列表
// @Tags TSystem
// @Summary 分页获取TSystem列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query systemReq.TSystemSearch true "分页获取TSystem列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /tSystem/getTSystemList [get]
func (tSystemApi *TSystemApi) GetTSystemList(c *gin.Context) {
	var pageInfo twitterReq.TSystemSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := tSystemService.GetTSystemInfoList(pageInfo); err != nil {
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

// GetTSystemList 获取TSystem列表
// @Tags TSystem
// @Summary 获取TSystem列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query  true "获取TSystem列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /tSystem/getTSystemAll [get]
func (tSystemApi *TSystemApi) GetTSystemAll(c *gin.Context) {
	err, list := tSystemService.GetTSystemAll()
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(list, c)
	}
}
