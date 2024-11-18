package Twitter

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/Twitter"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type ScoreRecordService struct{}

// CreateScoreRecord 创建scoreRecord表记录
// Author [yourname](https://github.com/yourname)
func (scoreRecordService *ScoreRecordService) CreateScoreRecord(scoreRecord *Twitter.ScoreRecord) (err error) {
	err = global.GVA_DB.Create(scoreRecord).Error
	return err
}

// DeleteScoreRecord 删除scoreRecord表记录
// Author [yourname](https://github.com/yourname)
func (scoreRecordService *ScoreRecordService) DeleteScoreRecord(id string) (err error) {
	err = global.GVA_DB.Delete(&Twitter.ScoreRecord{}, "id = ?", id).Error
	return err
}

// DeleteScoreRecordByIds 批量删除scoreRecord表记录
// Author [yourname](https://github.com/yourname)
func (scoreRecordService *ScoreRecordService) DeleteScoreRecordByIds(ids []string) (err error) {
	err = global.GVA_DB.Delete(&[]Twitter.ScoreRecord{}, "id in ?", ids).Error
	return err
}

// UpdateScoreRecord 更新scoreRecord表记录
// Author [yourname](https://github.com/yourname)
func (scoreRecordService *ScoreRecordService) UpdateScoreRecord(scoreRecord Twitter.ScoreRecord) (err error) {
	err = global.GVA_DB.Model(&Twitter.ScoreRecord{}).Where("id = ?", scoreRecord.Id).Updates(&scoreRecord).Error
	return err
}

// GetScoreRecord 根据id获取scoreRecord表记录
// Author [yourname](https://github.com/yourname)
func (scoreRecordService *ScoreRecordService) GetScoreRecord(id string) (scoreRecord Twitter.ScoreRecord, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&scoreRecord).Error
	return
}

// GetScoreRecordInfoList 分页获取scoreRecord表记录
// Author [yourname](https://github.com/yourname)
func (scoreRecordService *ScoreRecordService) GetScoreRecordInfoList(info request.PageInfo, scoreRecord Twitter.ScoreRecord) (list []Twitter.ScoreRecord, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&Twitter.ScoreRecord{})
	var scoreRecords []Twitter.ScoreRecord
	// 如果有条件搜索 下方会自动创建搜索语句
	if scoreRecord.Query != "" {
		db = db.Where("query LIKE ?", "%"+scoreRecord.Query+"%")
	}

	err = db.Count(&total).Error
	if err != nil {
		return
	}

	// 添加倒序排序
	db = db.Order("id desc") // 按照 id 字段倒序排列
	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&scoreRecords).Error
	return scoreRecords, total, err
}
func (scoreRecordService *ScoreRecordService) GetScoreRecordPublic() {
	// 此方法为获取数据源定义的数据
	// 请自行实现
}
