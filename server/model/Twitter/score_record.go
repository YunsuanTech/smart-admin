// 自动生成模板ScoreRecord
package Twitter

import (
	"time"
)

// scoreRecord表 结构体  ScoreRecord
type ScoreRecord struct {
	Id         *int       `json:"id" form:"id" gorm:"primarykey;column:id;comment:;"`               //ID
	Query      string     `json:"query" form:"query" gorm:"column:query;comment:;"`                 //条件
	ScoreAdded *float64   `json:"scoreAdded" form:"scoreAdded" gorm:"column:score_added;comment:;"` //总分
	Timestamp  *time.Time `json:"timestamp" form:"timestamp" gorm:"column:timestamp;comment:;"`     //日期
}

// TableName scoreRecord表 ScoreRecord自定义表名 score_record
func (ScoreRecord) TableName() string {
	return "score_record"
}
