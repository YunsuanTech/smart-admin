package Twitter

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/Twitter"
	TwitterReq "github.com/flipped-aurora/gin-vue-admin/server/model/Twitter/request"
)

type ScoringRulesService struct{}

// CreateScoringRules 创建scoringRules表记录
// Author [yourname](https://github.com/yourname)
func (scoringRulesService *ScoringRulesService) CreateScoringRules(scoringRules *Twitter.ScoringRules) (err error) {
	err = global.GVA_DB.Create(scoringRules).Error
	return err
}

// DeleteScoringRules 删除scoringRules表记录
// Author [yourname](https://github.com/yourname)
func (scoringRulesService *ScoringRulesService) DeleteScoringRules(id string) (err error) {
	err = global.GVA_DB.Delete(&Twitter.ScoringRules{}, "id = ?", id).Error
	return err
}

// DeleteScoringRulesByIds 批量删除scoringRules表记录
// Author [yourname](https://github.com/yourname)
func (scoringRulesService *ScoringRulesService) DeleteScoringRulesByIds(ids []string) (err error) {
	err = global.GVA_DB.Delete(&[]Twitter.ScoringRules{}, "id in ?", ids).Error
	return err
}

// UpdateScoringRules 更新scoringRules表记录
// Author [yourname](https://github.com/yourname)
func (scoringRulesService *ScoringRulesService) UpdateScoringRules(scoringRules Twitter.ScoringRules) (err error) {
	err = global.GVA_DB.Model(&Twitter.ScoringRules{}).Where("id = ?", scoringRules.Id).Updates(&scoringRules).Error
	return err
}

// GetScoringRules 根据id获取scoringRules表记录
// Author [yourname](https://github.com/yourname)
func (scoringRulesService *ScoringRulesService) GetScoringRules(id string) (scoringRules Twitter.ScoringRules, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&scoringRules).Error
	return
}

// GetScoringRulesInfoList 分页获取scoringRules表记录
// Author [yourname](https://github.com/yourname)
func (scoringRulesService *ScoringRulesService) GetScoringRulesInfoList(info TwitterReq.ScoringRulesSearch) (list []Twitter.ScoringRules, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&Twitter.ScoringRules{})
	var scoringRuless []Twitter.ScoringRules
	// 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	// 添加倒序排序
	db = db.Order("id desc") // 按照 id 字段倒序排列

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&scoringRuless).Error
	return scoringRuless, total, err
}
func (scoringRulesService *ScoringRulesService) GetScoringRulesPublic() {
	// 此方法为获取数据源定义的数据
	// 请自行实现
}
