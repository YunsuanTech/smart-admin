
// 自动生成模板Tokens
package Binance
import (
)

// tokens表 结构体  Tokens
type Tokens struct {
    Id  *int `json:"id" form:"id" gorm:"primarykey;column:id;comment:;"`  //ID 
    Address  *string `json:"address" form:"address" gorm:"column:address;comment:;"`  //地址 
    ChainName  *string `json:"chainName" form:"chainName" gorm:"column:chain_name;comment:;"`  //标识 
    MarketCap  *int `json:"marketCap" form:"marketCap" gorm:"column:market_cap;comment:;"`  //市值 
    IsMeme  *int `json:"isMeme" form:"isMeme" gorm:"column:is_meme;comment:;"`  //isMeme 
    Okx  *int `json:"okx" form:"okx" gorm:"column:okx;comment:;"`  //ok上线 
    Gate  *int `json:"gate" form:"gate" gorm:"column:gate;comment:;"`  //gate上线 
    Bybit  *int `json:"bybit" form:"bybit" gorm:"column:bybit;comment:;"`  //bybit上线 
    Kucoin  *int `json:"kucoin" form:"kucoin" gorm:"column:kucoin;comment:;"`  //kucoin上线 
}


// TableName tokens表 Tokens自定义表名 tokens
func (Tokens) TableName() string {
    return "tokens"
}



