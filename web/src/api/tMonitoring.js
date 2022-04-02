import service from '@/utils/request'

// @Tags TMonitoring
// @Summary 创建TMonitoring
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TMonitoring true "创建TMonitoring"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /tMonitoring/createTMonitoring [post]
export const createTMonitoring = (data) => {
  return service({
    url: '/tMonitoring/createTMonitoring',
    method: 'post',
    data
  })
}

// @Tags TMonitoring
// @Summary 删除TMonitoring
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TMonitoring true "删除TMonitoring"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /tMonitoring/deleteTMonitoring [delete]
export const deleteTMonitoring = (data) => {
  return service({
    url: '/tMonitoring/deleteTMonitoring',
    method: 'delete',
    data
  })
}

// @Tags TMonitoring
// @Summary 删除TMonitoring
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除TMonitoring"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /tMonitoring/deleteTMonitoring [delete]
export const deleteTMonitoringByIds = (data) => {
  return service({
    url: '/tMonitoring/deleteTMonitoringByIds',
    method: 'delete',
    data
  })
}

// @Tags TMonitoring
// @Summary 更新TMonitoring
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TMonitoring true "更新TMonitoring"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /tMonitoring/updateTMonitoring [put]
export const updateTMonitoring = (data) => {
  return service({
    url: '/tMonitoring/updateTMonitoring',
    method: 'put',
    data
  })
}

// @Tags TMonitoring
// @Summary 用id查询TMonitoring
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.TMonitoring true "用id查询TMonitoring"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /tMonitoring/findTMonitoring [get]
export const findTMonitoring = (params) => {
  return service({
    url: '/tMonitoring/findTMonitoring',
    method: 'get',
    params
  })
}

// @Tags TMonitoring
// @Summary 分页获取TMonitoring列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取TMonitoring列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /tMonitoring/getTMonitoringList [get]
export const getTMonitoringList = (params) => {
  return service({
    url: '/tMonitoring/getTMonitoringList',
    method: 'get',
    params
  })
}
