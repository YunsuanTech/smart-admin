package twitter

import (
	"smart-admin/server/global"
	"smart-admin/server/model/common/request"
	"smart-admin/server/model/twitter"
	systemReq "smart-admin/server/model/twitter/request"
)

type TSystemService struct {
}

// CreateTSystem 创建TSystem记录
// Author [piexlmax](https://github.com/piexlmax)
func (tSystemService *TSystemService) CreateTSystem(tSystem twitter.TSystem) (err error) {
	err = global.GVA_DB.Create(&tSystem).Error
	return err
}

// DeleteTSystem 删除TSystem记录
// Author [piexlmax](https://github.com/piexlmax)
func (tSystemService *TSystemService) DeleteTSystem(tSystem twitter.TSystem) (err error) {
	err = global.GVA_DB.Delete(&tSystem).Error
	return err
}

// DeleteTSystemByIds 批量删除TSystem记录
// Author [piexlmax](https://github.com/piexlmax)
func (tSystemService *TSystemService) DeleteTSystemByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]twitter.TSystem{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateTSystem 更新TSystem记录
// Author [piexlmax](https://github.com/piexlmax)
func (tSystemService *TSystemService) UpdateTSystem(tSystem twitter.TSystem) (err error) {
	err = global.GVA_DB.Save(&tSystem).Error
	return err
}

// GetTSystem 根据id获取TSystem记录
// Author [piexlmax](https://github.com/piexlmax)
func (tSystemService *TSystemService) GetTSystem(id uint) (err error, tSystem twitter.TSystem) {
	err = global.GVA_DB.Where("id = ?", id).First(&tSystem).Error
	return
}

// GetTSystemInfoList 分页获取TSystem记录
// Author [piexlmax](https://github.com/piexlmax)
func (tSystemService *TSystemService) GetTSystemInfoList(info systemReq.TSystemSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&twitter.TSystem{})
	var tSystems []twitter.TSystem
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.Key != "" {
		db = db.Where("`key` LIKE ?", "%"+info.Key+"%")
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Find(&tSystems).Error
	return err, tSystems, total
}

func (TSystemService *TSystemService) GetTSystemAll() (err error, list interface{}) {
	// 创建db
	db := global.GVA_DB.Model(&twitter.TSystem{})
	var tSystems []twitter.TSystem
	err = db.Find(&tSystems).Order("created_at desc").Error
	if err != nil {
		return
	}
	return err, tSystems
}
