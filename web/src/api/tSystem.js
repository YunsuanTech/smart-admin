import service from '@/utils/request'

// @Tags TSystem
// @Summary 创建TSystem
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TSystem true "创建TSystem"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /tSystem/createTSystem [post]
export const createTSystem = (data) => {
  return service({
    url: '/tSystem/createTSystem',
    method: 'post',
    data
  })
}

// @Tags TSystem
// @Summary 删除TSystem
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TSystem true "删除TSystem"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /tSystem/deleteTSystem [delete]
export const deleteTSystem = (data) => {
  return service({
    url: '/tSystem/deleteTSystem',
    method: 'delete',
    data
  })
}

// @Tags TSystem
// @Summary 删除TSystem
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除TSystem"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /tSystem/deleteTSystem [delete]
export const deleteTSystemByIds = (data) => {
  return service({
    url: '/tSystem/deleteTSystemByIds',
    method: 'delete',
    data
  })
}

// @Tags TSystem
// @Summary 更新TSystem
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TSystem true "更新TSystem"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /tSystem/updateTSystem [put]
export const updateTSystem = (data) => {
  return service({
    url: '/tSystem/updateTSystem',
    method: 'put',
    data
  })
}

// @Tags TSystem
// @Summary 用id查询TSystem
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.TSystem true "用id查询TSystem"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /tSystem/findTSystem [get]
export const findTSystem = (params) => {
  return service({
    url: '/tSystem/findTSystem',
    method: 'get',
    params
  })
}

// @Tags TSystem
// @Summary 分页获取TSystem列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取TSystem列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /tSystem/getTSystemList [get]
export const getTSystemList = (params) => {
  return service({
    url: '/tSystem/getTSystemList',
    method: 'get',
    params
  })
}

// @Tags TSystem
// @Summary 分页获取TSystem列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query true "获取TSystem列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /tSystem/getTSystemList [get]
export const getTSystemAll = (params) => {
  return service({
    url: '/tSystem/getTSystemAll',
    method: 'get',
    params
  })
}
