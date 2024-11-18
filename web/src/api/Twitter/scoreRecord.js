import service from '@/utils/request'
// @Tags ScoreRecord
// @Summary 创建scoreRecord表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ScoreRecord true "创建scoreRecord表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /scoreRecord/createScoreRecord [post]
export const createScoreRecord = (data) => {
  return service({
    url: '/scoreRecord/createScoreRecord',
    method: 'post',
    data
  })
}

// @Tags ScoreRecord
// @Summary 删除scoreRecord表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ScoreRecord true "删除scoreRecord表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /scoreRecord/deleteScoreRecord [delete]
export const deleteScoreRecord = (params) => {
  return service({
    url: '/scoreRecord/deleteScoreRecord',
    method: 'delete',
    params
  })
}

// @Tags ScoreRecord
// @Summary 批量删除scoreRecord表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除scoreRecord表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /scoreRecord/deleteScoreRecord [delete]
export const deleteScoreRecordByIds = (params) => {
  return service({
    url: '/scoreRecord/deleteScoreRecordByIds',
    method: 'delete',
    params
  })
}

// @Tags ScoreRecord
// @Summary 更新scoreRecord表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ScoreRecord true "更新scoreRecord表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /scoreRecord/updateScoreRecord [put]
export const updateScoreRecord = (data) => {
  return service({
    url: '/scoreRecord/updateScoreRecord',
    method: 'put',
    data
  })
}

// @Tags ScoreRecord
// @Summary 用id查询scoreRecord表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.ScoreRecord true "用id查询scoreRecord表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /scoreRecord/findScoreRecord [get]
export const findScoreRecord = (params) => {
  return service({
    url: '/scoreRecord/findScoreRecord',
    method: 'get',
    params
  })
}

// @Tags ScoreRecord
// @Summary 分页获取scoreRecord表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取scoreRecord表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /scoreRecord/getScoreRecordList [get]
export const getScoreRecordList = (params) => {
  return service({
    url: '/scoreRecord/getScoreRecordList',
    method: 'get',
    params
  })
}

// @Tags ScoreRecord
// @Summary 不需要鉴权的scoreRecord表接口
// @accept application/json
// @Produce application/json
// @Param data query TwitterReq.ScoreRecordSearch true "分页获取scoreRecord表列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /scoreRecord/getScoreRecordPublic [get]
export const getScoreRecordPublic = () => {
  return service({
    url: '/scoreRecord/getScoreRecordPublic',
    method: 'get',
  })
}
