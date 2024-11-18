package Twitter

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/Twitter"
	TwitterReq "github.com/flipped-aurora/gin-vue-admin/server/model/Twitter/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type TwitterUsersApi struct{}

// CreateTwitterUsers 创建twitterUsers表
// @Tags TwitterUsers
// @Summary 创建twitterUsers表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body Twitter.TwitterUsers true "创建twitterUsers表"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /twitterUsers/createTwitterUsers [post]
func (twitterUsersApi *TwitterUsersApi) CreateTwitterUsers(c *gin.Context) {
	var twitterUsers Twitter.TwitterUsers
	err := c.ShouldBindJSON(&twitterUsers)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = twitterUsersService.CreateTwitterUsers(&twitterUsers)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

// DeleteTwitterUsers 删除twitterUsers表
// @Tags TwitterUsers
// @Summary 删除twitterUsers表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body Twitter.TwitterUsers true "删除twitterUsers表"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /twitterUsers/deleteTwitterUsers [delete]
func (twitterUsersApi *TwitterUsersApi) DeleteTwitterUsers(c *gin.Context) {
	id := c.Query("id")
	err := twitterUsersService.DeleteTwitterUsers(id)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteTwitterUsersByIds 批量删除twitterUsers表
// @Tags TwitterUsers
// @Summary 批量删除twitterUsers表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /twitterUsers/deleteTwitterUsersByIds [delete]
func (twitterUsersApi *TwitterUsersApi) DeleteTwitterUsersByIds(c *gin.Context) {
	ids := c.QueryArray("ids[]")
	err := twitterUsersService.DeleteTwitterUsersByIds(ids)
	if err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateTwitterUsers 更新twitterUsers表
// @Tags TwitterUsers
// @Summary 更新twitterUsers表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body Twitter.TwitterUsers true "更新twitterUsers表"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /twitterUsers/updateTwitterUsers [put]
func (twitterUsersApi *TwitterUsersApi) UpdateTwitterUsers(c *gin.Context) {
	var twitterUsers Twitter.TwitterUsers
	err := c.ShouldBindJSON(&twitterUsers)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = twitterUsersService.UpdateTwitterUsers(twitterUsers)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindTwitterUsers 用id查询twitterUsers表
// @Tags TwitterUsers
// @Summary 用id查询twitterUsers表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query Twitter.TwitterUsers true "用id查询twitterUsers表"
// @Success 200 {object} response.Response{data=Twitter.TwitterUsers,msg=string} "查询成功"
// @Router /twitterUsers/findTwitterUsers [get]
func (twitterUsersApi *TwitterUsersApi) FindTwitterUsers(c *gin.Context) {
	id := c.Query("id")
	retwitterUsers, err := twitterUsersService.GetTwitterUsers(id)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:"+err.Error(), c)
		return
	}
	response.OkWithData(retwitterUsers, c)
}

// GetTwitterUsersList 分页获取twitterUsers表列表
// @Tags TwitterUsers
// @Summary 分页获取twitterUsers表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query TwitterReq.TwitterUsersSearch true "分页获取twitterUsers表列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /twitterUsers/getTwitterUsersList [get]
func (twitterUsersApi *TwitterUsersApi) GetTwitterUsersList(c *gin.Context) {
	var pageInfo TwitterReq.TwitterUsersSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := twitterUsersService.GetTwitterUsersInfoList(pageInfo.PageInfo, pageInfo.TwitterUsers)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败:"+err.Error(), c)
		return
	}
	response.OkWithDetailed(response.PageResult{
		List:     list,
		Total:    total,
		Page:     pageInfo.Page,
		PageSize: pageInfo.PageSize,
	}, "获取成功", c)
}

// GetTwitterUsersPublic 不需要鉴权的twitterUsers表接口
// @Tags TwitterUsers
// @Summary 不需要鉴权的twitterUsers表接口
// @accept application/json
// @Produce application/json
// @Param data query TwitterReq.TwitterUsersSearch true "分页获取twitterUsers表列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /twitterUsers/getTwitterUsersPublic [get]
func (twitterUsersApi *TwitterUsersApi) GetTwitterUsersPublic(c *gin.Context) {
	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	twitterUsersService.GetTwitterUsersPublic()
	response.OkWithDetailed(gin.H{
		"info": "不需要鉴权的twitterUsers表接口信息",
	}, "获取成功", c)
}
