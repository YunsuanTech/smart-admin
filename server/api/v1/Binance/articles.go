package Binance

import (
	
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/model/Binance"
    BinanceReq "github.com/flipped-aurora/gin-vue-admin/server/model/Binance/request"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

type ArticlesApi struct {}



// CreateArticles 创建articles表
// @Tags Articles
// @Summary 创建articles表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body Binance.Articles true "创建articles表"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /articles/createArticles [post]
func (articlesApi *ArticlesApi) CreateArticles(c *gin.Context) {
	var articles Binance.Articles
	err := c.ShouldBindJSON(&articles)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = articlesService.CreateArticles(&articles)
	if err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:" + err.Error(), c)
		return
	}
    response.OkWithMessage("创建成功", c)
}

// DeleteArticles 删除articles表
// @Tags Articles
// @Summary 删除articles表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body Binance.Articles true "删除articles表"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /articles/deleteArticles [delete]
func (articlesApi *ArticlesApi) DeleteArticles(c *gin.Context) {
	id := c.Query("id")
	err := articlesService.DeleteArticles(id)
	if err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteArticlesByIds 批量删除articles表
// @Tags Articles
// @Summary 批量删除articles表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /articles/deleteArticlesByIds [delete]
func (articlesApi *ArticlesApi) DeleteArticlesByIds(c *gin.Context) {
	ids := c.QueryArray("ids[]")
	err := articlesService.DeleteArticlesByIds(ids)
	if err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateArticles 更新articles表
// @Tags Articles
// @Summary 更新articles表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body Binance.Articles true "更新articles表"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /articles/updateArticles [put]
func (articlesApi *ArticlesApi) UpdateArticles(c *gin.Context) {
	var articles Binance.Articles
	err := c.ShouldBindJSON(&articles)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = articlesService.UpdateArticles(articles)
	if err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindArticles 用id查询articles表
// @Tags Articles
// @Summary 用id查询articles表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query Binance.Articles true "用id查询articles表"
// @Success 200 {object} response.Response{data=Binance.Articles,msg=string} "查询成功"
// @Router /articles/findArticles [get]
func (articlesApi *ArticlesApi) FindArticles(c *gin.Context) {
	id := c.Query("id")
	rearticles, err := articlesService.GetArticles(id)
	if err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:" + err.Error(), c)
		return
	}
	response.OkWithData(rearticles, c)
}

// GetArticlesList 分页获取articles表列表
// @Tags Articles
// @Summary 分页获取articles表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query BinanceReq.ArticlesSearch true "分页获取articles表列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /articles/getArticlesList [get]
func (articlesApi *ArticlesApi) GetArticlesList(c *gin.Context) {
	var pageInfo BinanceReq.ArticlesSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := articlesService.GetArticlesInfoList(pageInfo)
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

// GetArticlesPublic 不需要鉴权的articles表接口
// @Tags Articles
// @Summary 不需要鉴权的articles表接口
// @accept application/json
// @Produce application/json
// @Param data query BinanceReq.ArticlesSearch true "分页获取articles表列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /articles/getArticlesPublic [get]
func (articlesApi *ArticlesApi) GetArticlesPublic(c *gin.Context) {
    // 此接口不需要鉴权
    // 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
    articlesService.GetArticlesPublic()
    response.OkWithDetailed(gin.H{
       "info": "不需要鉴权的articles表接口信息",
    }, "获取成功", c)
}
