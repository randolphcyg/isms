package domain

import (
	"context"
	"time"
)

// IndustryCategory 行业大类领域模型
type IndustryCategory struct {
	CategoryCode string // 大类编码（如"B"、"C"）
	CategoryName string // 大类名称（如"采矿业"、"制造业"）
}

// IsmsIndustry 行业分类（包含大类和小类）领域模型
type IsmsIndustry struct {
	ID              int32     `json:"id"`
	CategoryCode    string    `json:"category_code"`    // 大类代码（如B、C）
	CategoryName    string    `json:"category_name"`    // 大类名称（如采矿业）
	SubcategoryCode string    `json:"subcategory_code"` // 小类代码（如06、07）
	SubcategoryName string    `json:"subcategory_name"` // 小类名称
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

// IndustryRepo 仓库接口
type IndustryRepo interface {
	// GetAllCategories 查询所有行业大类（去重）
	GetAllCategories(ctx context.Context) ([]*IndustryCategory, error)
	// GetSubcategoriesByCode 根据大类编码查询小类
	GetSubcategoriesByCode(ctx context.Context, categoryCode string) ([]*IsmsIndustry, error)
	// GetCategoryByCode 根据编码查询单个大类
	GetCategoryByCode(ctx context.Context, categoryCode string) (*IndustryCategory, error)

	GetIndustriesByIDs(ctx context.Context, ids []int32) ([]*IsmsIndustry, error)

	// 统计方法
	Count(ctx context.Context) (int64, error)
}
