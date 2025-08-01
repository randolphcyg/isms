package domain

import (
	"context"
	"fmt"
	"time"

	v1 "isms/api/isms/v1"
)

// IsmsSoftware 工业软件领域实体
type IsmsSoftware struct {
	ID                   int32              `json:"id"`
	NameZh               string             `json:"name_zh"`                         // 软件中文名称
	NameEn               string             `json:"name_en"`                         // 软件英文名称
	Version              string             `json:"version"`                         // 版本号
	DeveloperID          int32              `json:"developer_id"`                    // 开发商ID
	CountryID            int32              `json:"country_id"`                      // 产地国家ID
	ReleaseDate          *time.Time         `json:"release_date,omitempty"`          // 发布日期
	SysReq               *string            `json:"sys_req,omitempty"`               // 系统要求
	Description          *string            `json:"description,omitempty"`           // 软件描述
	SizeGb               *float64           `json:"size_gb,omitempty"`               // 软件大小（GB）
	DeploymentMethod     *string            `json:"deployment_method,omitempty"`     // 部署方式
	ComplianceInfo       *string            `json:"compliance_info,omitempty"`       // 合规性信息
	SecurityInfo         *string            `json:"security_info,omitempty"`         // 安全信息
	IntellectualProperty *string            `json:"intellectual_property,omitempty"` // 知识产权信息
	Status               int32              `json:"status"`                          // 状态（1：有效，0：下架）
	IndustryIDs          []int32            `json:"industry_ids"`                    // 适用行业小类ID列表
	OsIDs                []int32            `json:"os_ids"`                          // 支持的操作系统ID列表
	IndustryDetails      []*v1.IsmsIndustry `json:"industry_details,omitempty"`      // 适用行业详情
	CreatedAt            time.Time          `json:"created_at"`
	UpdatedAt            time.Time          `json:"updated_at"`
}

// Validate 业务校验方法
func (s *IsmsSoftware) Validate() error {
	if len(s.NameZh) == 0 || len(s.NameZh) > 500 {
		return fmt.Errorf("软件中文名称长度必须为1-500字符")
	}
	if len(s.NameEn) == 0 || len(s.NameEn) > 500 {
		return fmt.Errorf("软件英文名称长度必须为1-500字符")
	}
	if len(s.Version) == 0 || len(s.Version) > 100 {
		return fmt.Errorf("版本号长度必须为1-100字符")
	}
	if s.DeveloperID == 0 {
		return fmt.Errorf("开发商ID不能为空")
	}
	if s.CountryID == 0 {
		return fmt.Errorf("国家ID不能为空")
	}
	if len(s.IndustryIDs) == 0 {
		return fmt.Errorf("至少需要选择一个适用行业")
	}
	if len(s.OsIDs) == 0 {
		return fmt.Errorf("至少需要选择一个支持的操作系统")
	}
	return nil
}

// ListSoftwareOptions 分页查询参数
type ListSoftwareOptions struct {
	Page         int32   `json:"page"`
	PageSize     int32   `json:"page_size"`
	Keyword      string  `json:"keyword"`
	CountryID    int32   `json:"country_id"`
	Status       int32   `json:"status"`
	DeveloperID  int32   `json:"developer_id,omitempty"`
	IndustryIDs  []int32 `json:"industry_ids,omitempty"`
	CategoryCode string  `json:"category_code,omitempty"`
}

type ListSoftwareOptions2 struct {
	Page        uint32
	PageSize    uint32
	DeveloperID uint32
	Category    string
	Status      uint32
	Keyword     string
}

// SoftwareRepo 软件仓储接口
type SoftwareRepo interface {
	// Create 保存软件实体
	Create(ctx context.Context, software *IsmsSoftware) (*IsmsSoftware, error)

	// Update 更新软件实体
	Update(ctx context.Context, software *IsmsSoftware) error

	// FindByID 查询软件
	FindByID(ctx context.Context, id uint32) (*IsmsSoftware, error)

	ExistByNameAndVersion(ctx context.Context, name, version string) (bool, error)
	ExistByID(ctx context.Context, id uint32) (bool, error)
	List(ctx context.Context, opts ListSoftwareOptions) ([]*IsmsSoftware, int64, error)
}
