package Binance

import (
	
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/model/Binance"
    BinanceReq "github.com/flipped-aurora/gin-vue-admin/server/model/Binance/request"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

type TokensApi struct {}



// CreateTokens 创建tokens表
// @Tags Tokens
// @Summary 创建tokens表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body Binance.Tokens true "创建tokens表"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /tokens/createTokens [post]
func (tokensApi *TokensApi) CreateTokens(c *gin.Context) {
	var tokens Binance.Tokens
	err := c.ShouldBindJSON(&tokens)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = tokensService.CreateTokens(&tokens)
	if err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:" + err.Error(), c)
		return
	}
    response.OkWithMessage("创建成功", c)
}

// DeleteTokens 删除tokens表
// @Tags Tokens
// @Summary 删除tokens表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body Binance.Tokens true "删除tokens表"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /tokens/deleteTokens [delete]
func (tokensApi *TokensApi) DeleteTokens(c *gin.Context) {
	id := c.Query("id")
	err := tokensService.DeleteTokens(id)
	if err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteTokensByIds 批量删除tokens表
// @Tags Tokens
// @Summary 批量删除tokens表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /tokens/deleteTokensByIds [delete]
func (tokensApi *TokensApi) DeleteTokensByIds(c *gin.Context) {
	ids := c.QueryArray("ids[]")
	err := tokensService.DeleteTokensByIds(ids)
	if err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateTokens 更新tokens表
// @Tags Tokens
// @Summary 更新tokens表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body Binance.Tokens true "更新tokens表"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /tokens/updateTokens [put]
func (tokensApi *TokensApi) UpdateTokens(c *gin.Context) {
	var tokens Binance.Tokens
	err := c.ShouldBindJSON(&tokens)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = tokensService.UpdateTokens(tokens)
	if err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindTokens 用id查询tokens表
// @Tags Tokens
// @Summary 用id查询tokens表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query Binance.Tokens true "用id查询tokens表"
// @Success 200 {object} response.Response{data=Binance.Tokens,msg=string} "查询成功"
// @Router /tokens/findTokens [get]
func (tokensApi *TokensApi) FindTokens(c *gin.Context) {
	id := c.Query("id")
	retokens, err := tokensService.GetTokens(id)
	if err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:" + err.Error(), c)
		return
	}
	response.OkWithData(retokens, c)
}

// GetTokensList 分页获取tokens表列表
// @Tags Tokens
// @Summary 分页获取tokens表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query BinanceReq.TokensSearch true "分页获取tokens表列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /tokens/getTokensList [get]
func (tokensApi *TokensApi) GetTokensList(c *gin.Context) {
	var pageInfo BinanceReq.TokensSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := tokensService.GetTokensInfoList(pageInfo)
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

// GetTokensPublic 不需要鉴权的tokens表接口
// @Tags Tokens
// @Summary 不需要鉴权的tokens表接口
// @accept application/json
// @Produce application/json
// @Param data query BinanceReq.TokensSearch true "分页获取tokens表列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /tokens/getTokensPublic [get]
func (tokensApi *TokensApi) GetTokensPublic(c *gin.Context) {
    // 此接口不需要鉴权
    // 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
    tokensService.GetTokensPublic()
    response.OkWithDetailed(gin.H{
       "info": "不需要鉴权的tokens表接口信息",
    }, "获取成功", c)
}
