
// 自动生成模板ScoringRules
package Twitter
import (
)

// scoringRules表 结构体  ScoringRules
type ScoringRules struct {
    Id  *int `json:"id" form:"id" gorm:"primarykey;column:id;comment:;"`  //ID 
    FactorName  *string `json:"factorName" form:"factorName" gorm:"column:factor_name;comment:;" binding:"required"`  //加分因子 
    Weight  *float64 `json:"weight" form:"weight" gorm:"column:weight;comment:;" binding:"required"`  //权重 
    Condition  *string `json:"condition" form:"condition" gorm:"column:condition;comment:;" binding:"required"`  //条件 
    Description  *string `json:"description" form:"description" gorm:"column:description;comment:;"`  //描述 
}


// TableName scoringRules表 ScoringRules自定义表名 scoring_rules
func (ScoringRules) TableName() string {
    return "scoring_rules"
}



