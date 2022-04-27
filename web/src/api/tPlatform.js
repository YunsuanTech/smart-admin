import service from '@/utils/request'

// @Tags TPlatform
// @Summary 创建TPlatform
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TPlatform true "创建TPlatform"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /tPlatform/createTPlatform [post]
export const createTPlatform = (data) => {
  return service({
    url: '/tPlatform/createTPlatform',
    method: 'post',
    data
  })
}

// @Tags TPlatform
// @Summary 删除TPlatform
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TPlatform true "删除TPlatform"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /tPlatform/deleteTPlatform [delete]
export const deleteTPlatform = (data) => {
  return service({
    url: '/tPlatform/deleteTPlatform',
    method: 'delete',
    data
  })
}

// @Tags TPlatform
// @Summary 删除TPlatform
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除TPlatform"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /tPlatform/deleteTPlatform [delete]
export const deleteTPlatformByIds = (data) => {
  return service({
    url: '/tPlatform/deleteTPlatformByIds',
    method: 'delete',
    data
  })
}

// @Tags TPlatform
// @Summary 更新TPlatform
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TPlatform true "更新TPlatform"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /tPlatform/updateTPlatform [put]
export const updateTPlatform = (data) => {
  return service({
    url: '/tPlatform/updateTPlatform',
    method: 'put',
    data
  })
}

// @Tags TPlatform
// @Summary 用id查询TPlatform
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.TPlatform true "用id查询TPlatform"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /tPlatform/findTPlatform [get]
export const findTPlatform = (params) => {
  return service({
    url: '/tPlatform/findTPlatform',
    method: 'get',
    params
  })
}

// @Tags TPlatform
// @Summary 分页获取TPlatform列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取TPlatform列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /tPlatform/getTPlatformList [get]
export const getTPlatformList = (params) => {
  return service({
    url: '/tPlatform/getTPlatformList',
    method: 'get',
    params
  })
}
