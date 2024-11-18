
// 自动生成模板Articles
package Binance
import (
	"time"
)

// articles表 结构体  Articles
type Articles struct {
    Id  *int `json:"id" form:"id" gorm:"primarykey;column:id;comment:;"`  //ID 
    Code  *string `json:"code" form:"code" gorm:"column:code;comment:;"`  //公告表示 
    Title  *string `json:"title" form:"title" gorm:"column:title;comment:;"`  //标题 
    Type  *string `json:"type" form:"type" gorm:"column:type;comment:;"`  //类型 
    ListingDate  *string `json:"listingDate" form:"listingDate" gorm:"column:listing_date;comment:;"`  //上市时间 
    ListinFees  *string `json:"ListinFees" form:"ListinFees" gorm:"column:Listin_Fees;comment:;"`  //上市费用 
    Address  *string `json:"address" form:"address" gorm:"column:address;comment:;"`  //地址 
    AnnouncementDate  *time.Time `json:"announcementDate" form:"announcementDate" gorm:"column:announcement_date;comment:;"`  //公告时间 
}


// TableName articles表 Articles自定义表名 articles
func (Articles) TableName() string {
    return "articles"
}



