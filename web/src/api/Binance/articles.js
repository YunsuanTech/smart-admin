import service from '@/utils/request'
// @Tags Articles
// @Summary 创建articles表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Articles true "创建articles表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /articles/createArticles [post]
export const createArticles = (data) => {
  return service({
    url: '/articles/createArticles',
    method: 'post',
    data
  })
}

// @Tags Articles
// @Summary 删除articles表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Articles true "删除articles表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /articles/deleteArticles [delete]
export const deleteArticles = (params) => {
  return service({
    url: '/articles/deleteArticles',
    method: 'delete',
    params
  })
}

// @Tags Articles
// @Summary 批量删除articles表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除articles表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /articles/deleteArticles [delete]
export const deleteArticlesByIds = (params) => {
  return service({
    url: '/articles/deleteArticlesByIds',
    method: 'delete',
    params
  })
}

// @Tags Articles
// @Summary 更新articles表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Articles true "更新articles表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /articles/updateArticles [put]
export const updateArticles = (data) => {
  return service({
    url: '/articles/updateArticles',
    method: 'put',
    data
  })
}

// @Tags Articles
// @Summary 用id查询articles表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Articles true "用id查询articles表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /articles/findArticles [get]
export const findArticles = (params) => {
  return service({
    url: '/articles/findArticles',
    method: 'get',
    params
  })
}

// @Tags Articles
// @Summary 分页获取articles表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取articles表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /articles/getArticlesList [get]
export const getArticlesList = (params) => {
  return service({
    url: '/articles/getArticlesList',
    method: 'get',
    params
  })
}

// @Tags Articles
// @Summary 不需要鉴权的articles表接口
// @accept application/json
// @Produce application/json
// @Param data query BinanceReq.ArticlesSearch true "分页获取articles表列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /articles/getArticlesPublic [get]
export const getArticlesPublic = () => {
  return service({
    url: '/articles/getArticlesPublic',
    method: 'get',
  })
}
