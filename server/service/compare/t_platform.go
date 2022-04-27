package compare

import (
	"smart-admin/server/global"
	"smart-admin/server/model/common/request"
	"smart-admin/server/model/compare"
	compareReq "smart-admin/server/model/compare/request"
)

type TPlatformService struct {
}

// CreateTPlatform 创建TPlatform记录
// Author [piexlmax](https://github.com/piexlmax)
func (tPlatformService *TPlatformService) CreateTPlatform(tPlatform compare.TPlatform) (err error) {
	err = global.GVA_PGSQL_DB.Create(&tPlatform).Error
	return err
}

// DeleteTPlatform 删除TPlatform记录
// Author [piexlmax](https://github.com/piexlmax)
func (tPlatformService *TPlatformService) DeleteTPlatform(tPlatform compare.TPlatform) (err error) {
	err = global.GVA_PGSQL_DB.Delete(&tPlatform).Error
	return err
}

// DeleteTPlatformByIds 批量删除TPlatform记录
// Author [piexlmax](https://github.com/piexlmax)
func (tPlatformService *TPlatformService) DeleteTPlatformByIds(ids request.IdsReq) (err error) {
	err = global.GVA_PGSQL_DB.Delete(&[]compare.TPlatform{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateTPlatform 更新TPlatform记录
// Author [piexlmax](https://github.com/piexlmax)
func (tPlatformService *TPlatformService) UpdateTPlatform(tPlatform compare.TPlatform) (err error) {
	err = global.GVA_PGSQL_DB.Save(&tPlatform).Error
	return err
}

// GetTPlatform 根据id获取TPlatform记录
// Author [piexlmax](https://github.com/piexlmax)
func (tPlatformService *TPlatformService) GetTPlatform(id uint) (err error, tPlatform compare.TPlatform) {
	err = global.GVA_PGSQL_DB.Where("id = ?", id).First(&tPlatform).Error
	return
}

// GetTPlatformInfoList 分页获取TPlatform记录
// Author [piexlmax](https://github.com/piexlmax)
func (tPlatformService *TPlatformService) GetTPlatformInfoList(info compareReq.TPlatformSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_PGSQL_DB.Model(&compare.TPlatform{})
	var tPlatforms []compare.TPlatform
	// 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Find(&tPlatforms).Error
	return err, tPlatforms, total
}
