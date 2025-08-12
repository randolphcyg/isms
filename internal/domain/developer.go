package domain

import (
	"context"
	"fmt"
	"time"
)

// Developer 开发商领域模型
type Developer struct {
	ID            int32     `json:"id"`
	NameZh        string    `json:"name_zh"`         // 中文名称
	NameEn        string    `json:"name_en"`         // 英文名称
	CountryID     int32     `json:"country_id"`      // 所属国家ID
	CountryNameZh string    `json:"country_name_zh"` // 所属国家中文名称
	Website       *string   `json:"website"`         // 官网
	Description   *string   `json:"description"`     // 简介
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

// Validate 业务校验方法
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

// DeveloperRepo 开发商仓库接口
type DeveloperRepo interface {
	// Create 新增开发商
	Create(ctx context.Context, dev *Developer) (*Developer, error)
	// GetByID 根据ID查询开发商
	GetByID(ctx context.Context, id int32) (*Developer, error)
	// List 分页查询开发商
	List(ctx context.Context, page, pageSize, countryID uint32, keyword string) ([]*Developer, int64, error)
	// ExistByName 校验名称是否已存在
	ExistByName(ctx context.Context, nameZh string) (bool, error)

	ExistByNameExcludeID(ctx context.Context, nameZh string, excludeID int32) (bool, error)
	// Update 更新开发商
	Update(ctx context.Context, dev *Developer) (*Developer, error)
}
