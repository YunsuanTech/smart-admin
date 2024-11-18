package Twitter

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/Twitter"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type TwitterUsersService struct{}

// CreateTwitterUsers 创建twitterUsers表记录
// Author [yourname](https://github.com/yourname)
func (twitterUsersService *TwitterUsersService) CreateTwitterUsers(twitterUsers *Twitter.TwitterUsers) (err error) {
	err = global.GVA_DB.Create(twitterUsers).Error
	return err
}

// DeleteTwitterUsers 删除twitterUsers表记录
// Author [yourname](https://github.com/yourname)
func (twitterUsersService *TwitterUsersService) DeleteTwitterUsers(id string) (err error) {
	err = global.GVA_DB.Delete(&Twitter.TwitterUsers{}, "id = ?", id).Error
	return err
}

// DeleteTwitterUsersByIds 批量删除twitterUsers表记录
// Author [yourname](https://github.com/yourname)
func (twitterUsersService *TwitterUsersService) DeleteTwitterUsersByIds(ids []string) (err error) {
	err = global.GVA_DB.Delete(&[]Twitter.TwitterUsers{}, "id in ?", ids).Error
	return err
}

// UpdateTwitterUsers 更新twitterUsers表记录
// Author [yourname](https://github.com/yourname)
func (twitterUsersService *TwitterUsersService) UpdateTwitterUsers(twitterUsers Twitter.TwitterUsers) (err error) {
	err = global.GVA_DB.Model(&Twitter.TwitterUsers{}).Where("id = ?", twitterUsers.Id).Updates(&twitterUsers).Error
	return err
}

// GetTwitterUsers 根据id获取twitterUsers表记录
// Author [yourname](https://github.com/yourname)
func (twitterUsersService *TwitterUsersService) GetTwitterUsers(id string) (twitterUsers Twitter.TwitterUsers, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&twitterUsers).Error
	return
}

// GetTwitterUsersInfoList 分页获取twitterUsers表记录
// Author [yourname](https://github.com/yourname)
func (twitterUsersService *TwitterUsersService) GetTwitterUsersInfoList(info request.PageInfo, Users Twitter.TwitterUsers) (list []Twitter.TwitterUsers, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&Twitter.TwitterUsers{})
	var twitterUserss []Twitter.TwitterUsers
	// 如果有条件搜索 下方会自动创建搜索语句
	if Users.Username != "" {
		db = db.Where("username LIKE ?", "%"+Users.Username+"%")
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&twitterUserss).Error
	return twitterUserss, total, err
}
func (twitterUsersService *TwitterUsersService) GetTwitterUsersPublic() {
	// 此方法为获取数据源定义的数据
	// 请自行实现
}
