import service from '@/utils/request'
// @Tags Tokens
// @Summary 创建tokens表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Tokens true "创建tokens表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /tokens/createTokens [post]
export const createTokens = (data) => {
  return service({
    url: '/tokens/createTokens',
    method: 'post',
    data
  })
}

// @Tags Tokens
// @Summary 删除tokens表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Tokens true "删除tokens表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /tokens/deleteTokens [delete]
export const deleteTokens = (params) => {
  return service({
    url: '/tokens/deleteTokens',
    method: 'delete',
    params
  })
}

// @Tags Tokens
// @Summary 批量删除tokens表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除tokens表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /tokens/deleteTokens [delete]
export const deleteTokensByIds = (params) => {
  return service({
    url: '/tokens/deleteTokensByIds',
    method: 'delete',
    params
  })
}

// @Tags Tokens
// @Summary 更新tokens表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Tokens true "更新tokens表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /tokens/updateTokens [put]
export const updateTokens = (data) => {
  return service({
    url: '/tokens/updateTokens',
    method: 'put',
    data
  })
}

// @Tags Tokens
// @Summary 用id查询tokens表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Tokens true "用id查询tokens表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /tokens/findTokens [get]
export const findTokens = (params) => {
  return service({
    url: '/tokens/findTokens',
    method: 'get',
    params
  })
}

// @Tags Tokens
// @Summary 分页获取tokens表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取tokens表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /tokens/getTokensList [get]
export const getTokensList = (params) => {
  return service({
    url: '/tokens/getTokensList',
    method: 'get',
    params
  })
}

// @Tags Tokens
// @Summary 不需要鉴权的tokens表接口
// @accept application/json
// @Produce application/json
// @Param data query BinanceReq.TokensSearch true "分页获取tokens表列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /tokens/getTokensPublic [get]
export const getTokensPublic = () => {
  return service({
    url: '/tokens/getTokensPublic',
    method: 'get',
  })
}
