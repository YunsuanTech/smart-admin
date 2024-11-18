import service from '@/utils/request'
// @Tags TwitterUsers
// @Summary 创建twitterUsers表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TwitterUsers true "创建twitterUsers表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /twitterUsers/createTwitterUsers [post]
export const createTwitterUsers = (data) => {
  return service({
    url: '/twitterUsers/createTwitterUsers',
    method: 'post',
    data
  })
}

// @Tags TwitterUsers
// @Summary 删除twitterUsers表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TwitterUsers true "删除twitterUsers表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /twitterUsers/deleteTwitterUsers [delete]
export const deleteTwitterUsers = (params) => {
  return service({
    url: '/twitterUsers/deleteTwitterUsers',
    method: 'delete',
    params
  })
}

// @Tags TwitterUsers
// @Summary 批量删除twitterUsers表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除twitterUsers表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /twitterUsers/deleteTwitterUsers [delete]
export const deleteTwitterUsersByIds = (params) => {
  return service({
    url: '/twitterUsers/deleteTwitterUsersByIds',
    method: 'delete',
    params
  })
}

// @Tags TwitterUsers
// @Summary 更新twitterUsers表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TwitterUsers true "更新twitterUsers表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /twitterUsers/updateTwitterUsers [put]
export const updateTwitterUsers = (data) => {
  return service({
    url: '/twitterUsers/updateTwitterUsers',
    method: 'put',
    data
  })
}

// @Tags TwitterUsers
// @Summary 用id查询twitterUsers表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.TwitterUsers true "用id查询twitterUsers表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /twitterUsers/findTwitterUsers [get]
export const findTwitterUsers = (params) => {
  return service({
    url: '/twitterUsers/findTwitterUsers',
    method: 'get',
    params
  })
}

// @Tags TwitterUsers
// @Summary 分页获取twitterUsers表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取twitterUsers表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /twitterUsers/getTwitterUsersList [get]
export const getTwitterUsersList = (params) => {
  return service({
    url: '/twitterUsers/getTwitterUsersList',
    method: 'get',
    params
  })
}

// @Tags TwitterUsers
// @Summary 不需要鉴权的twitterUsers表接口
// @accept application/json
// @Produce application/json
// @Param data query TwitterReq.TwitterUsersSearch true "分页获取twitterUsers表列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /twitterUsers/getTwitterUsersPublic [get]
export const getTwitterUsersPublic = () => {
  return service({
    url: '/twitterUsers/getTwitterUsersPublic',
    method: 'get',
  })
}
