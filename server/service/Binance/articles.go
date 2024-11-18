package Binance

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/Binance"
	BinanceReq "github.com/flipped-aurora/gin-vue-admin/server/model/Binance/request"
)

type ArticlesService struct{}

// CreateArticles 创建articles表记录
// Author [yourname](https://github.com/yourname)
func (articlesService *ArticlesService) CreateArticles(articles *Binance.Articles) (err error) {
	err = global.GVA_DB.Create(articles).Error
	return err
}

// DeleteArticles 删除articles表记录
// Author [yourname](https://github.com/yourname)
func (articlesService *ArticlesService) DeleteArticles(id string) (err error) {
	err = global.GVA_DB.Delete(&Binance.Articles{}, "id = ?", id).Error
	return err
}

// DeleteArticlesByIds 批量删除articles表记录
// Author [yourname](https://github.com/yourname)
func (articlesService *ArticlesService) DeleteArticlesByIds(ids []string) (err error) {
	err = global.GVA_DB.Delete(&[]Binance.Articles{}, "id in ?", ids).Error
	return err
}

// UpdateArticles 更新articles表记录
// Author [yourname](https://github.com/yourname)
func (articlesService *ArticlesService) UpdateArticles(articles Binance.Articles) (err error) {
	err = global.GVA_DB.Model(&Binance.Articles{}).Where("id = ?", articles.Id).Updates(&articles).Error
	return err
}

// GetArticles 根据id获取articles表记录
// Author [yourname](https://github.com/yourname)
func (articlesService *ArticlesService) GetArticles(id string) (articles Binance.Articles, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&articles).Error
	return
}

// GetArticlesInfoList 分页获取articles表记录
// Author [yourname](https://github.com/yourname)
func (articlesService *ArticlesService) GetArticlesInfoList(info BinanceReq.ArticlesSearch) (list []Binance.Articles, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&Binance.Articles{})
	var articless []Binance.Articles
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

	err = db.Find(&articless).Error
	return articless, total, err
}
func (articlesService *ArticlesService) GetArticlesPublic() {
	// 此方法为获取数据源定义的数据
	// 请自行实现
}
