package initialize

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/Binance"
	"github.com/flipped-aurora/gin-vue-admin/server/model/Twitter"
)

func bizModel() error {
	db := global.GVA_DB
	err := db.AutoMigrate(Binance.Articles{}, Binance.Tokens{}, Twitter.TwitterUsers{}, Twitter.ScoringRules{}, Twitter.ScoreRecord{})
	if err != nil {
		return err
	}
	return nil
}
