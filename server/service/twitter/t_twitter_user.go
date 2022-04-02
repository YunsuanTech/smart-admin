package twitter

import (
	"smart-admin/server/global"
	"smart-admin/server/model/common/request"
	"smart-admin/server/model/twitter"
	systemReq "smart-admin/server/model/twitter/request"
)

type TTwitterUserService struct {
}

// CreateTTwitterUser 创建TTwitterUser记录
// Author [piexlmax](https://github.com/piexlmax)
func (tTwitterUserService *TTwitterUserService) CreateTTwitterUser(tTwitterUser twitter.TTwitterUser) (err error) {
	err = global.GVA_DB.Create(&tTwitterUser).Error
	return err
}

// DeleteTTwitterUser 删除TTwitterUser记录
// Author [piexlmax](https://github.com/piexlmax)
func (tTwitterUserService *TTwitterUserService) DeleteTTwitterUser(tTwitterUser twitter.TTwitterUser) (err error) {
	err = global.GVA_DB.Delete(&tTwitterUser).Error
	return err
}

// DeleteTTwitterUserByIds 批量删除TTwitterUser记录
// Author [piexlmax](https://github.com/piexlmax)
func (tTwitterUserService *TTwitterUserService) DeleteTTwitterUserByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]twitter.TTwitterUser{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateTTwitterUser 更新TTwitterUser记录
// Author [piexlmax](https://github.com/piexlmax)
func (tTwitterUserService *TTwitterUserService) UpdateTTwitterUser(tTwitterUser twitter.TTwitterUser) (err error) {
	err = global.GVA_DB.Save(&tTwitterUser).Error
	return err
}

// GetTTwitterUser 根据id获取TTwitterUser记录
// Author [piexlmax](https://github.com/piexlmax)
func (tTwitterUserService *TTwitterUserService) GetTTwitterUser(id uint) (err error, tTwitterUser twitter.TTwitterUser) {
	err = global.GVA_DB.Where("id = ?", id).First(&tTwitterUser).Error
	return
}

// GetTTwitterUserInfoList 分页获取TTwitterUser记录
// Author [piexlmax](https://github.com/piexlmax)
func (tTwitterUserService *TTwitterUserService) GetTTwitterUserInfoList(info systemReq.TTwitterUserSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&twitter.TTwitterUser{})
	var tTwitterUsers []twitter.TTwitterUser
	if info.UserName != "" {
		db = db.Where("`user_name` LIKE ?", "%"+info.UserName+"%")
	}
	// 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Find(&tTwitterUsers).Error
	return err, tTwitterUsers, total
}
