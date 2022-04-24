package twitter

import (
	"smart-admin/server/global"
	"smart-admin/server/model/common/request"
	"smart-admin/server/model/twitter"
	systemReq "smart-admin/server/model/twitter/request"
)

type TTwitterContentService struct {
}

// CreateTTwitterContent 创建TTwitterContent记录
// Author [piexlmax](https://github.com/piexlmax)
func (tTwitterContentService *TTwitterContentService) CreateTTwitterContent(tTwitterContent twitter.TTwitterContent) (err error) {
	err = global.GVA_DB.Create(&tTwitterContent).Error
	return err
}

// DeleteTTwitterContent 删除TTwitterContent记录
// Author [piexlmax](https://github.com/piexlmax)
func (tTwitterContentService *TTwitterContentService) DeleteTTwitterContent(tTwitterContent twitter.TTwitterContent) (err error) {
	err = global.GVA_DB.Delete(&tTwitterContent).Error
	return err
}

// DeleteTTwitterContentByIds 批量删除TTwitterContent记录
// Author [piexlmax](https://github.com/piexlmax)
func (tTwitterContentService *TTwitterContentService) DeleteTTwitterContentByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]twitter.TTwitterContent{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateTTwitterContent 更新TTwitterContent记录
// Author [piexlmax](https://github.com/piexlmax)
func (tTwitterContentService *TTwitterContentService) UpdateTTwitterContent(tTwitterContent twitter.TTwitterContent) (err error) {
	err = global.GVA_DB.Save(&tTwitterContent).Error
	return err
}

// GetTTwitterContent 根据id获取TTwitterContent记录
// Author [piexlmax](https://github.com/piexlmax)
func (tTwitterContentService *TTwitterContentService) GetTTwitterContent(id uint) (err error, tTwitterContent twitter.TTwitterContent) {
	err = global.GVA_DB.Where("id = ?", id).First(&tTwitterContent).Error
	return
}

// GetTTwitterContentInfoList 分页获取TTwitterContent记录
// Author [piexlmax](https://github.com/piexlmax)
func (tTwitterContentService *TTwitterContentService) GetTTwitterContentInfoList(info systemReq.TTwitterContentSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&twitter.TTwitterContent{})
	var tTwitterContents []twitter.TTwitterContent
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.UserName != "" {
		db = db.Where("`user_name` LIKE ?", "%"+info.UserName+"%")
	}

	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Order("time_parsed desc").Limit(limit).Offset(offset).Find(&tTwitterContents).Error
	return err, tTwitterContents, total
}
