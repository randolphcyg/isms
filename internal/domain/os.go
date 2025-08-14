package domain

import (
	"context"
	"fmt"
	"time"
)

// OS 操作系统领域模型
type OS struct {
	ID           int32     `json:"id"`
	Name         string    `json:"name"`         // 系统名称（如：Microsoft Windows、Ubuntu）
	Version      string    `json:"version"`      // 系统版本（如：7、10、20.04 LTS）
	Architecture string    `json:"architecture"` // 硬件架构
	Manufacturer *string   `json:"manufacturer"` // 系统开发商（如：Microsoft、Canonical）
	ReleaseYear  *int32    `json:"release_year"` // 发布年份
	Description  *string   `json:"description"`  // 系统说明（如包含的细分版本）
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

// Validate 业务规则校验
func (o *OS) Validate() error {
	if len(o.Name) == 0 || len(o.Name) > 200 {
		return fmt.Errorf("系统名称长度必须为1-200字符")
	}
	if len(o.Version) == 0 || len(o.Version) > 50 {
		return fmt.Errorf("系统版本长度必须为1-50字符")
	}
	if len(o.Architecture) == 0 {
		return fmt.Errorf("硬件架构不能为空")
	}
	if o.Manufacturer != nil && len(*o.Manufacturer) > 200 {
		return fmt.Errorf("系统开发商长度不能超过200字符")
	}
	if o.ReleaseYear != nil && *o.ReleaseYear > 2030 {
		return fmt.Errorf("发布年份不能超过2030年")
	}
	if o.Description != nil && len(*o.Description) > 65535 {
		return fmt.Errorf("系统说明长度不能超过65535字符")
	}
	return nil
}

// OSRepo 操作系统仓库接口
type OSRepo interface {
	Create(ctx context.Context, os *OS) (*OS, error)
	Update(ctx context.Context, os *OS) (*OS, error)
	Delete(ctx context.Context, id int32) error
	GetByID(ctx context.Context, id int32) (*OS, error)
	List(ctx context.Context, page, pageSize uint32, keyword, manufacturer string) ([]*OS, int64, error)
	ExistByNameVersionArch(ctx context.Context, name, version, architecture string, excludeID int32) (bool, error)
}
