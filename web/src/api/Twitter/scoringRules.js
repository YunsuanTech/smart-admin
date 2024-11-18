import service from '@/utils/request'
// @Tags ScoringRules
// @Summary 创建scoringRules表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ScoringRules true "创建scoringRules表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /scoringRules/createScoringRules [post]
export const createScoringRules = (data) => {
  return service({
    url: '/scoringRules/createScoringRules',
    method: 'post',
    data
  })
}

// @Tags ScoringRules
// @Summary 删除scoringRules表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ScoringRules true "删除scoringRules表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /scoringRules/deleteScoringRules [delete]
export const deleteScoringRules = (params) => {
  return service({
    url: '/scoringRules/deleteScoringRules',
    method: 'delete',
    params
  })
}

// @Tags ScoringRules
// @Summary 批量删除scoringRules表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除scoringRules表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /scoringRules/deleteScoringRules [delete]
export const deleteScoringRulesByIds = (params) => {
  return service({
    url: '/scoringRules/deleteScoringRulesByIds',
    method: 'delete',
    params
  })
}

// @Tags ScoringRules
// @Summary 更新scoringRules表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ScoringRules true "更新scoringRules表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /scoringRules/updateScoringRules [put]
export const updateScoringRules = (data) => {
  return service({
    url: '/scoringRules/updateScoringRules',
    method: 'put',
    data
  })
}

// @Tags ScoringRules
// @Summary 用id查询scoringRules表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.ScoringRules true "用id查询scoringRules表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /scoringRules/findScoringRules [get]
export const findScoringRules = (params) => {
  return service({
    url: '/scoringRules/findScoringRules',
    method: 'get',
    params
  })
}

// @Tags ScoringRules
// @Summary 分页获取scoringRules表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取scoringRules表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /scoringRules/getScoringRulesList [get]
export const getScoringRulesList = (params) => {
  return service({
    url: '/scoringRules/getScoringRulesList',
    method: 'get',
    params
  })
}

// @Tags ScoringRules
// @Summary 不需要鉴权的scoringRules表接口
// @accept application/json
// @Produce application/json
// @Param data query TwitterReq.ScoringRulesSearch true "分页获取scoringRules表列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /scoringRules/getScoringRulesPublic [get]
export const getScoringRulesPublic = () => {
  return service({
    url: '/scoringRules/getScoringRulesPublic',
    method: 'get',
  })
}
