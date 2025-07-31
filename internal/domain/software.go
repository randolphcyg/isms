package domain

import (
	"context"
	"time"

	v1 "isms/api/isms/v1"
)

// IsmsSoftware 工业软件领域实体（关联行业分类）
type IsmsSoftware struct {
	ID              int64              // 自增ID
	Name            string             // 软件名称（必填）
	NameEn          string             // 英文名称（可选）
	Developer       string             // 开发商（必填）
	Version         string             // 软件版本（必填）
	Category        string             // 软件类别（如CAD、CAE，必填）
	OsIDs           []int64            // 支持的操作系统ID列表（至少一个）
	IndustryIDs     []int32            // 适用行业小类ID列表（关联IsmsIndustry.ID）
	IndustryDetails []*v1.IsmsIndustry // 适用行业详情（冗余，便于展示）
	Description     string             // 描述（可选）
	CountryID       int64              // 所属国家ID（必填）
	Website         string             // 官网地址（可选）
	Status          int32              // 状态（0=停用，1=正常）
	CreatedAt       time.Time          // 创建时间
	UpdatedAt       time.Time          // 更新时间
}

// ListSoftwareOptions 分页查询参数（新增行业筛选）
type ListSoftwareOptions struct {
	Page         int32   // 页码
	PageSize     int32   // 每页条数
	Keyword      string  // 名称/开发商关键词
	Category     string  // 软件类别筛选
	CountryID    int64   // 国家ID筛选
	Status       int32   // 状态筛选
	IndustryIDs  []int32 // 行业小类ID筛选（可选）
	CategoryCode string  // 行业大类编码筛选（如"B"采矿业，可选）
}

// SoftwareRepo 软件仓储接口（扩展行业关联方法）
type SoftwareRepo interface {
	// Create 保存软件实体（含行业关联）
	Create(ctx context.Context, software *IsmsSoftware) (*IsmsSoftware, error)

	// Update 更新软件实体（含行业关联）
	Update(ctx context.Context, software *IsmsSoftware) error

	// FindByID 查询软件（关联行业详情）
	FindByID(ctx context.Context, id int64) (*IsmsSoftware, error)

	// List 分页查询软件（支持按行业筛选）
	List(ctx context.Context, opts ListSoftwareOptions) ([]*IsmsSoftware, int64, error)
}
