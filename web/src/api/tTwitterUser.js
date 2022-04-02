import service from '@/utils/request'

// @Tags TTwitterUser
// @Summary 创建TTwitterUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TTwitterUser true "创建TTwitterUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /tTwitterUser/createTTwitterUser [post]
export const createTTwitterUser = (data) => {
  return service({
    url: '/tTwitterUser/createTTwitterUser',
    method: 'post',
    data
  })
}

// @Tags TTwitterUser
// @Summary 删除TTwitterUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TTwitterUser true "删除TTwitterUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /tTwitterUser/deleteTTwitterUser [delete]
export const deleteTTwitterUser = (data) => {
  return service({
    url: '/tTwitterUser/deleteTTwitterUser',
    method: 'delete',
    data
  })
}

// @Tags TTwitterUser
// @Summary 删除TTwitterUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除TTwitterUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /tTwitterUser/deleteTTwitterUser [delete]
export const deleteTTwitterUserByIds = (data) => {
  return service({
    url: '/tTwitterUser/deleteTTwitterUserByIds',
    method: 'delete',
    data
  })
}

// @Tags TTwitterUser
// @Summary 更新TTwitterUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TTwitterUser true "更新TTwitterUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /tTwitterUser/updateTTwitterUser [put]
export const updateTTwitterUser = (data) => {
  return service({
    url: '/tTwitterUser/updateTTwitterUser',
    method: 'put',
    data
  })
}

// @Tags TTwitterUser
// @Summary 用id查询TTwitterUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.TTwitterUser true "用id查询TTwitterUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /tTwitterUser/findTTwitterUser [get]
export const findTTwitterUser = (params) => {
  return service({
    url: '/tTwitterUser/findTTwitterUser',
    method: 'get',
    params
  })
}

// @Tags TTwitterUser
// @Summary 分页获取TTwitterUser列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取TTwitterUser列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /tTwitterUser/getTTwitterUserList [get]
export const getTTwitterUserList = (params) => {
  return service({
    url: '/tTwitterUser/getTTwitterUserList',
    method: 'get',
    params
  })
}
