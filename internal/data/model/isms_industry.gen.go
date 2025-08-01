// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameIsmsIndustry = "isms_industry"

// IsmsIndustry 第二产业行业分类表
type IsmsIndustry struct {
	ID              int32     `gorm:"column:id;type:smallint(5) unsigned;primaryKey;autoIncrement:true;comment:自增ID" json:"id"`                                                                            // 自增ID
	CategoryCode    string    `gorm:"column:category_code;type:char(1);not null;uniqueIndex:uk_category_subcategory,priority:1;index:idx_category,priority:1;comment:大类代码(B/C/D/E等)" json:"category_code"` // 大类代码(B/C/D/E等)
	CategoryName    string    `gorm:"column:category_name;type:varchar(50);not null;comment:大类名称(采矿业/制造业等)" json:"category_name"`                                                                          // 大类名称(采矿业/制造业等)
	SubcategoryCode string    `gorm:"column:subcategory_code;type:char(2);not null;uniqueIndex:uk_category_subcategory,priority:2;comment:小类代码(06/07/13等)" json:"subcategory_code"`                        // 小类代码(06/07/13等)
	SubcategoryName string    `gorm:"column:subcategory_name;type:varchar(100);not null;comment:小类名称(煤炭开采和洗选业等)" json:"subcategory_name"`                                                                  // 小类名称(煤炭开采和洗选业等)
	CreatedAt       time.Time `gorm:"column:created_at;type:datetime;not null;default:CURRENT_TIMESTAMP;comment:创建时间" json:"created_at"`                                                                   // 创建时间
	UpdatedAt       time.Time `gorm:"column:updated_at;type:datetime;not null;default:CURRENT_TIMESTAMP;comment:更新时间" json:"updated_at"`                                                                   // 更新时间
}

// TableName IsmsIndustry's table name
func (*IsmsIndustry) TableName() string {
	return TableNameIsmsIndustry
}
