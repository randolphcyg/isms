package domain

import "time"

// OverviewStats 概览统计数据
type OverviewStats struct {
	TotalSoftware    int64     `json:"total_software"`     // 软件总数
	TotalDevelopers  int64     `json:"total_developers"`   // 开发商总数
	TotalIndustries  int64     `json:"total_industries"`   // 行业总数
	TotalCountries   int64     `json:"total_countries"`    // 国家总数
	NewSoftwareCount int64     `json:"new_software_count"` // 新增软件数（最近30天）
	LastUpdated      time.Time `json:"last_updated"`       // 最后更新时间
}

// IndustryStatItem 行业统计项
type IndustryStatItem struct {
	IndustryID    int32   `json:"industry_id"`    // 行业ID
	IndustryName  string  `json:"industry_name"`  // 行业名称
	SoftwareCount int64   `json:"software_count"` // 软件数量
	Percentage    float64 `json:"percentage"`     // 占比（百分比）
}

// CountryStatItem 国家统计项
type CountryStatItem struct {
	CountryID     int32   `json:"country_id"`      // 国家ID
	CountryNameZh string  `json:"country_name_zh"` // 国家中文名称
	CountryNameEn string  `json:"country_name_en"` // 国家英文名称
	SoftwareCount int64   `json:"software_count"`  // 软件数量
	Percentage    float64 `json:"percentage"`      // 占比（百分比）
}

// DeveloperStatItem 开发商统计项
type DeveloperStatItem struct {
	DeveloperID     int32   `json:"developer_id"`      // 开发商ID
	DeveloperNameZh string  `json:"developer_name_zh"` // 开发商中文名称
	DeveloperNameEn string  `json:"developer_name_en"` // 开发商英文名称
	SoftwareCount   int64   `json:"software_count"`    // 软件数量
	Percentage      float64 `json:"percentage"`        // 占比（百分比）
}

// TrendStatItem 年份趋势统计项
type TrendStatItem struct {
	Year          int32 `json:"year"`           // 年份
	SoftwareCount int64 `json:"software_count"` // 软件数量
}

// StatusStatItem 状态统计项
type StatusStatItem struct {
	Status        string  `json:"status"`         // 状态值
	StatusLabel   string  `json:"status_label"`   // 状态标签
	SoftwareCount int64   `json:"software_count"` // 软件数量
	Percentage    float64 `json:"percentage"`     // 占比（百分比）
}

// 添加统计结果结构体定义
type IndustryStatResult struct {
	IndustryID    int32  `json:"industry_id"`
	IndustryName  string `json:"industry_name"`
	SoftwareCount int64  `json:"software_count"`
}

type CountryStatResult struct {
	CountryID     int32  `json:"country_id"`
	CountryNameZh string `json:"country_name_zh"`
	CountryNameEn string `json:"country_name_en"`
	SoftwareCount int64  `json:"software_count"`
}

type DeveloperStatResult struct {
	DeveloperID     int32  `json:"developer_id"`
	DeveloperNameZh string `json:"developer_name_zh"`
	DeveloperNameEn string `json:"developer_name_en"`
	SoftwareCount   int64  `json:"software_count"`
}

type TrendStatResult struct {
	Year          int32 `json:"year"`
	SoftwareCount int64 `json:"software_count"`
}

type StatusStatResult struct {
	Status        string `json:"status"`
	StatusLabel   string `json:"status_label"`
	SoftwareCount int64  `json:"software_count"`
}
