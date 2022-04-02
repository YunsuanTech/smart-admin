import service from '@/utils/request'

// @Tags TTwitterContent
// @Summary 创建TTwitterContent
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TTwitterContent true "创建TTwitterContent"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /tTwitterContent/createTTwitterContent [post]
export const createTTwitterContent = (data) => {
  return service({
    url: '/tTwitterContent/createTTwitterContent',
    method: 'post',
    data
  })
}

// @Tags TTwitterContent
// @Summary 删除TTwitterContent
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TTwitterContent true "删除TTwitterContent"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /tTwitterContent/deleteTTwitterContent [delete]
export const deleteTTwitterContent = (data) => {
  return service({
    url: '/tTwitterContent/deleteTTwitterContent',
    method: 'delete',
    data
  })
}

// @Tags TTwitterContent
// @Summary 删除TTwitterContent
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除TTwitterContent"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /tTwitterContent/deleteTTwitterContent [delete]
export const deleteTTwitterContentByIds = (data) => {
  return service({
    url: '/tTwitterContent/deleteTTwitterContentByIds',
    method: 'delete',
    data
  })
}

// @Tags TTwitterContent
// @Summary 更新TTwitterContent
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TTwitterContent true "更新TTwitterContent"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /tTwitterContent/updateTTwitterContent [put]
export const updateTTwitterContent = (data) => {
  return service({
    url: '/tTwitterContent/updateTTwitterContent',
    method: 'put',
    data
  })
}

// @Tags TTwitterContent
// @Summary 用id查询TTwitterContent
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.TTwitterContent true "用id查询TTwitterContent"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /tTwitterContent/findTTwitterContent [get]
export const findTTwitterContent = (params) => {
  return service({
    url: '/tTwitterContent/findTTwitterContent',
    method: 'get',
    params
  })
}

// @Tags TTwitterContent
// @Summary 分页获取TTwitterContent列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取TTwitterContent列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /tTwitterContent/getTTwitterContentList [get]
export const getTTwitterContentList = (params) => {
  return service({
    url: '/tTwitterContent/getTTwitterContentList',
    method: 'get',
    params
  })
}
