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

type TTwitterUserApi struct {
}

// CreateTTwitterUser 创建TTwitterUser
// @Tags TTwitterUser
// @Summary 创建TTwitterUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.TTwitterUser true "创建TTwitterUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /tTwitterUser/createTTwitterUser [post]
func (tTwitterUserApi *TTwitterUserApi) CreateTTwitterUser(c *gin.Context) {
	var tTwitterUser twitter.TTwitterUser
	_ = c.ShouldBindJSON(&tTwitterUser)
	if err := tTwitterUserService.CreateTTwitterUser(tTwitterUser); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteTTwitterUser 删除TTwitterUser
// @Tags TTwitterUser
// @Summary 删除TTwitterUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.TTwitterUser true "删除TTwitterUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /tTwitterUser/deleteTTwitterUser [delete]
func (tTwitterUserApi *TTwitterUserApi) DeleteTTwitterUser(c *gin.Context) {
	var tTwitterUser twitter.TTwitterUser
	_ = c.ShouldBindJSON(&tTwitterUser)
	if err := tTwitterUserService.DeleteTTwitterUser(tTwitterUser); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteTTwitterUserByIds 批量删除TTwitterUser
// @Tags TTwitterUser
// @Summary 批量删除TTwitterUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除TTwitterUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /tTwitterUser/deleteTTwitterUserByIds [delete]
func (tTwitterUserApi *TTwitterUserApi) DeleteTTwitterUserByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := tTwitterUserService.DeleteTTwitterUserByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateTTwitterUser 更新TTwitterUser
// @Tags TTwitterUser
// @Summary 更新TTwitterUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.TTwitterUser true "更新TTwitterUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /tTwitterUser/updateTTwitterUser [put]
func (tTwitterUserApi *TTwitterUserApi) UpdateTTwitterUser(c *gin.Context) {
	var tTwitterUser twitter.TTwitterUser
	_ = c.ShouldBindJSON(&tTwitterUser)
	if err := tTwitterUserService.UpdateTTwitterUser(tTwitterUser); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindTTwitterUser 用id查询TTwitterUser
// @Tags TTwitterUser
// @Summary 用id查询TTwitterUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query autocode.TTwitterUser true "用id查询TTwitterUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /tTwitterUser/findTTwitterUser [get]
func (tTwitterUserApi *TTwitterUserApi) FindTTwitterUser(c *gin.Context) {
	var tTwitterUser twitter.TTwitterUser
	_ = c.ShouldBindQuery(&tTwitterUser)
	if err, retTwitterUser := tTwitterUserService.GetTTwitterUser(tTwitterUser.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"retTwitterUser": retTwitterUser}, c)
	}
}

// GetTTwitterUserList 分页获取TTwitterUser列表
// @Tags TTwitterUser
// @Summary 分页获取TTwitterUser列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query autocodeReq.TTwitterUserSearch true "分页获取TTwitterUser列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /tTwitterUser/getTTwitterUserList [get]
func (tTwitterUserApi *TTwitterUserApi) GetTTwitterUserList(c *gin.Context) {
	var pageInfo twitterReq.TTwitterUserSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := tTwitterUserService.GetTTwitterUserInfoList(pageInfo); err != nil {
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
