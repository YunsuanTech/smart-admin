package twitter

import (
	"smart-admin/server/global"
	"smart-admin/server/model/common/request"
	"smart-admin/server/model/twitter"
	systemReq "smart-admin/server/model/twitter/request"
)

type TMonitoringService struct {
}

// CreateTMonitoring 创建TMonitoring记录
// Author [piexlmax](https://github.com/piexlmax)
func (tMonitoringService *TMonitoringService) CreateTMonitoring(tMonitoring twitter.TMonitoring) (err error) {
	err = global.GVA_DB.Create(&tMonitoring).Error
	return err
}

// DeleteTMonitoring 删除TMonitoring记录
// Author [piexlmax](https://github.com/piexlmax)
func (tMonitoringService *TMonitoringService) DeleteTMonitoring(tMonitoring twitter.TMonitoring) (err error) {
	err = global.GVA_DB.Delete(&tMonitoring).Error
	return err
}

// DeleteTMonitoringByIds 批量删除TMonitoring记录
// Author [piexlmax](https://github.com/piexlmax)
func (tMonitoringService *TMonitoringService) DeleteTMonitoringByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]twitter.TMonitoring{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateTMonitoring 更新TMonitoring记录
// Author [piexlmax](https://github.com/piexlmax)
func (tMonitoringService *TMonitoringService) UpdateTMonitoring(tMonitoring twitter.TMonitoring) (err error) {
	err = global.GVA_DB.Save(&tMonitoring).Error
	return err
}

// GetTMonitoring 根据id获取TMonitoring记录
// Author [piexlmax](https://github.com/piexlmax)
func (tMonitoringService *TMonitoringService) GetTMonitoring(id uint) (err error, tMonitoring twitter.TMonitoring) {
	err = global.GVA_DB.Where("id = ?", id).First(&tMonitoring).Error
	return
}

// GetTMonitoringInfoList 分页获取TMonitoring记录
// Author [piexlmax](https://github.com/piexlmax)
func (tMonitoringService *TMonitoringService) GetTMonitoringInfoList(info systemReq.TMonitoringSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&twitter.TMonitoring{})
	var tMonitorings []twitter.TMonitoring
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.TiwtterName != "" {
		db = db.Where("`tiwtter_name` LIKE ?", "%"+info.TiwtterName+"%")
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Find(&tMonitorings).Error
	return err, tMonitorings, total
}
