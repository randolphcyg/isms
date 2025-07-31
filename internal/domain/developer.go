package domain

import (
	"context"
	"fmt"
	"time"
)

// Developer 开发商领域模型（核心业务实体）
// 封装开发商的属性和业务行为（如校验名称合法性）
type Developer struct {
	ID          uint32    // 自增ID
	NameZh      string    // 中文名称（业务核心属性）
	NameEn      string    // 英文名称
	CountryID   int32     // 所属国家ID（关联国家领域模型）
	CountryName string    // 冗余国家名称（用于展示，非核心属性）
	Website     string    // 官网
	Description string    // 简介
	CreatedAt   time.Time // 创建时间
	UpdatedAt   time.Time // 更新时间
}

// Validate 业务校验方法（封装开发商的核心业务规则）
func (d *Developer) Validate() error {
	if len(d.NameZh) == 0 || len(d.NameZh) > 200 {
		return fmt.Errorf("中文名称长度必须为1-200字符")
	}
	if len(d.NameEn) == 0 || len(d.NameEn) > 200 {
		return fmt.Errorf("英文名称长度必须为1-200字符")
	}
	if d.CountryID == 0 {
		return fmt.Errorf("国家ID不能为空")
	}
	return nil
}

// DeveloperRepo 开发商仓库接口（定义数据访问契约）
// 方法名直接对应业务操作，而非数据库CRUD
type DeveloperRepo interface {
	// Create 新增开发商（业务导向的方法名）
	Create(ctx context.Context, dev *Developer) (*Developer, error)
	// GetByID 根据ID查询开发商
	GetByID(ctx context.Context, id uint32) (*Developer, error)
	// List 分页查询开发商（支持按国家和关键词筛选）
	List(ctx context.Context, page, pageSize uint32, countryID uint32, keyword string) ([]*Developer, int64, error)
	// ExistByName 校验名称是否已存在（业务场景：避免重复创建）
	ExistByName(ctx context.Context, nameZh string) (bool, error)
}
