// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameIsmsDeveloper = "isms_developer"

// IsmsDeveloper 软件开发商表（国家关联通过代码逻辑维护，非数据库外键）
type IsmsDeveloper struct {
	ID          int32     `gorm:"column:id;type:int(10) unsigned;primaryKey;autoIncrement:true;comment:自增ID" json:"id"`                                                             // 自增ID
	NameZh      string    `gorm:"column:name_zh;type:varchar(200);not null;uniqueIndex:uk_name_zh,priority:1;comment:开发商中文名称（如：施耐德电气）" json:"name_zh"`                              // 开发商中文名称（如：施耐德电气）
	NameEn      string    `gorm:"column:name_en;type:varchar(200);not null;comment:开发商英文名称（如：Schneider Electric）" json:"name_en"`                                                   // 开发商英文名称（如：Schneider Electric）
	CountryID   int32     `gorm:"column:country_id;type:smallint(5) unsigned;not null;index:idx_country,priority:1;comment:所属国家ID（关联isms_country.id，通过代码逻辑维护关联）" json:"country_id"` // 所属国家ID（关联isms_country.id，通过代码逻辑维护关联）
	Website     *string   `gorm:"column:website;type:varchar(500);comment:官方网站URL" json:"website"`                                                                                  // 官方网站URL
	Description *string   `gorm:"column:description;type:text;comment:开发商简介" json:"description"`                                                                                    // 开发商简介
	CreatedAt   time.Time `gorm:"column:created_at;type:datetime;not null;default:CURRENT_TIMESTAMP;comment:创建时间" json:"created_at"`                                                // 创建时间
	UpdatedAt   time.Time `gorm:"column:updated_at;type:datetime;not null;default:CURRENT_TIMESTAMP;comment:更新时间" json:"updated_at"`                                                // 更新时间
}

// TableName IsmsDeveloper's table name
func (*IsmsDeveloper) TableName() string {
	return TableNameIsmsDeveloper
}
