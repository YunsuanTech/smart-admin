package Binance

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/Binance"
	BinanceReq "github.com/flipped-aurora/gin-vue-admin/server/model/Binance/request"
)

type TokensService struct{}

// CreateTokens 创建tokens表记录
// Author [yourname](https://github.com/yourname)
func (tokensService *TokensService) CreateTokens(tokens *Binance.Tokens) (err error) {
	err = global.GVA_DB.Create(tokens).Error
	return err
}

// DeleteTokens 删除tokens表记录
// Author [yourname](https://github.com/yourname)
func (tokensService *TokensService) DeleteTokens(id string) (err error) {
	err = global.GVA_DB.Delete(&Binance.Tokens{}, "id = ?", id).Error
	return err
}

// DeleteTokensByIds 批量删除tokens表记录
// Author [yourname](https://github.com/yourname)
func (tokensService *TokensService) DeleteTokensByIds(ids []string) (err error) {
	err = global.GVA_DB.Delete(&[]Binance.Tokens{}, "id in ?", ids).Error
	return err
}

// UpdateTokens 更新tokens表记录
// Author [yourname](https://github.com/yourname)
func (tokensService *TokensService) UpdateTokens(tokens Binance.Tokens) (err error) {
	err = global.GVA_DB.Model(&Binance.Tokens{}).Where("id = ?", tokens.Id).Updates(&tokens).Error
	return err
}

// GetTokens 根据id获取tokens表记录
// Author [yourname](https://github.com/yourname)
func (tokensService *TokensService) GetTokens(id string) (tokens Binance.Tokens, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&tokens).Error
	return
}

// GetTokensInfoList 分页获取tokens表记录
// Author [yourname](https://github.com/yourname)
func (tokensService *TokensService) GetTokensInfoList(info BinanceReq.TokensSearch) (list []Binance.Tokens, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&Binance.Tokens{})
	var tokenss []Binance.Tokens
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

	err = db.Find(&tokenss).Error
	return tokenss, total, err
}
func (tokensService *TokensService) GetTokensPublic() {
	// 此方法为获取数据源定义的数据
	// 请自行实现
}
