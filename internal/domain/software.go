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
  CountryName          string             `json:"country_name,omitempty"`          // 国家名称
  DeveloperName        string             `json:"developer_name,omitempty"`        // 开发商名称
  ReleaseYear          *int32             `json:"release_year,omitempty"`          // 发布年份
  ReleaseMonth         *int32             `json:"release_month,omitempty"`         // 发布月份
  ReleaseDay           *int32             `json:"release_day,omitempty"`           // 发布日
  CPUReq               *string            `json:"cpu_req,omitempty"`               // 处理器要求
  MemoryMinGb          *float64           `json:"memory_min_gb,omitempty"`         // 最小内存要求(GB)
  DiskMinGb            *float64           `json:"disk_min_gb,omitempty"`           // 最小磁盘空间要求(GB)
  SysReqOther          *string            `json:"sys_req_other,omitempty"`         // 其他系统要求
  Description          *string            `json:"description,omitempty"`           // 软件描述
  SizeBytes            *int64             `json:"size_bytes,omitempty"`            // 软件实际大小（字节，用于计算和存储，1KB=1024字节）
  DeploymentMethod     *string            `json:"deployment_method,omitempty"`     // 部署方式
  ComplianceInfo       *string            `json:"compliance_info,omitempty"`       // 合规性信息
  SecurityInfo         *string            `json:"security_info,omitempty"`         // 安全信息
  IntellectualProperty *string            `json:"intellectual_property,omitempty"` // 知识产权信息
  Status               string             `json:"status"`                          // 状态（"active"：有效；"inactive"：下架；"testing"：测试中；"discontinued"：停止维护）
  IndustryIDs          []int32            `json:"industry_ids"`                    // 适用行业小类ID列表
  OsIDs                []int32            `json:"os_ids"`                          // 支持的操作系统ID列表
  IndustryDetails      []*v1.IsmsIndustry `json:"industry_details,omitempty"`      // 适用行业详情
  BitWidths            []string           `json:"bit_widths,omitempty"`            // 支持的位宽列表
  CreatedAt            time.Time          `json:"created_at"`
  UpdatedAt            time.Time          `json:"updated_at"`
}

// Validate 业务校验方法
func (s *IsmsSoftware) Validate() error {
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
	//if len(s.OsIDs) == 0 {
	//	return fmt.Errorf("至少需要选择一个支持的操作系统")
	//}
	// Status字段校验
	switch s.Status {
	case "active", "inactive", "testing", "discontinued":
		// 有效值，无需处理
	case "":
		// 空值，设置默认值为"active"
		s.Status = "active"
	default:
		return fmt.Errorf("状态值无效，必须是active、inactive、testing或discontinued之一")
	}
	return nil
}

// ListSoftwareOptions 分页查询参数
type ListSoftwareOptions struct {
	Page         int32   `json:"page"`
	PageSize     int32   `json:"page_size"`
	Keyword      string  `json:"keyword"`
	CountryID    int32   `json:"country_id"`
	Status       string  `json:"status"`
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
// 修改SoftwareRepo接口，确保List方法签名正确
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
